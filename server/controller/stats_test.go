package controller

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/appditto/MonKey/server/database"
	"github.com/appditto/MonKey/server/models"
	"github.com/appditto/MonKey/server/utils"
	"gorm.io/gorm"
)

func TestServicesResponse(t *testing.T) {
	tests := []struct {
		name     string
		svcStats []ServiceStats
		want     string
	}{
		{
			name: "empty is an array",
			want: `[]`,
		},
		{
			name: "keeps order and names the empty service unspecified",
			svcStats: []ServiceStats{
				{StatsNumbers: StatsNumbers{Total: 10, Unique: 3}, Service: "kalium"},
				{StatsNumbers: StatsNumbers{Total: 4, Unique: 2}, Service: ""},
			},
			want: `[{"kalium":{"total":10,"unique":3}},{"unspecified":{"total":4,"unique":2}}]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(servicesResponse(tt.svcStats))
			if err != nil {
				t.Fatal(err)
			}
			if string(got) != tt.want {
				t.Errorf("Expected %s got %s", tt.want, got)
			}
		})
	}
}

func TestAllTimeTotalsMatchFullScan(t *testing.T) {
	sc := &StatsController{DB: testDB(t)}
	today := startOfDay(time.Now())
	rolloutDay := today.AddDate(0, 0, -2)
	oldSvc := "kalium"
	retiredSvc := "retired"

	seedOldWorkerRows(t, sc.DB, []models.Stats{
		{IPAddress: "1.1.1.1", BanAddress: "ban_a", Count: 5, CreatedAt: rolloutDay.AddDate(0, 0, -3)},
		{IPAddress: "1.1.1.1", BanAddress: "ban_a", Service: &oldSvc, Count: 2, CreatedAt: rolloutDay.AddDate(0, 0, -3)},
		{IPAddress: "2.2.2.2", BanAddress: "ban_b", Service: &retiredSvc, Count: 7, CreatedAt: rolloutDay.AddDate(0, 0, -1)},
		{IPAddress: "3.3.3.3", BanAddress: "ban_a", Service: &oldSvc, Count: 1, CreatedAt: rolloutDay.Add(time.Hour)},
		{IPAddress: "4.4.4.4", BanAddress: "ban_c", Count: 3, CreatedAt: rolloutDay.Add(2 * time.Hour)},
	})
	recordAll(t, sc, rolloutDay.Add(13*time.Hour), []StatsMessage{
		{IP: "3.3.3.3", Address: "ban_a", Svc: "kalium"},
		{IP: "4.4.4.4", Address: "ban_c"},
		{IP: "4.4.4.4", Address: "ban_c", Svc: "not-registered"},
		{IP: "5.5.5.5", Address: "ban_d", Svc: "creeper"},
		{IP: "1.1.1.1", Address: "ban_b", Svc: "creeper"},
	})
	runBackfill(t, sc.DB)
	assertMatchesFullScan(t, sc)

	recordAll(t, sc, time.Now(), []StatsMessage{
		{IP: "1.1.1.1", Address: "ban_a"},
		{IP: "1.1.1.1", Address: "ban_a"},
		{IP: "6.6.6.6", Address: "ban_e", Svc: "kalium"},
	})
	assertMatchesFullScan(t, sc)

	var wg sync.WaitGroup
	for worker := 0; worker < 2; worker++ {
		wg.Add(1)
		go func(seed int64) {
			defer wg.Done()
			random := rand.New(rand.NewSource(seed))
			services := []string{"", "kalium", "creeper", "wban"}
			for i := 0; i < 300; i++ {
				msg := StatsMessage{
					IP:      fmt.Sprintf("10.0.0.%d", random.Intn(8)),
					Address: fmt.Sprintf("ban_%d", random.Intn(20)),
					Svc:     services[random.Intn(len(services))],
				}
				if err := sc.record(msg, time.Now()); err != nil {
					t.Error(err)
				}
			}
		}(int64(worker))
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		runBackfill(t, sc.DB)
	}()
	wg.Wait()
	assertMatchesFullScan(t, sc)

	runBackfill(t, sc.DB)
	assertMatchesFullScan(t, sc)
}

func testDB(t *testing.T) *gorm.DB {
	host := os.Getenv("DB_MOCK_HOST")
	if host == "" {
		t.Skip("DB_MOCK_HOST not set")
	}
	db, err := database.NewConnection(&database.Config{
		Host:     host,
		Port:     utils.GetEnv("DB_MOCK_PORT", "5432"),
		User:     os.Getenv("DB_MOCK_USER"),
		Password: os.Getenv("DB_MOCK_PASS"),
		DBName:   utils.GetEnv("DB_MOCK_NAME", "testing"),
		SSLMode:  "disable",
	})
	if err != nil {
		t.Fatal(err)
	}
	err = db.Migrator().DropTable(&models.Stats{}, &models.StatsAddress{}, &models.StatsClient{}, &models.StatsServiceAddress{}, &models.StatsServiceDay{}, &models.StatsDay{})
	if err != nil {
		t.Fatal(err)
	}
	if err := database.Migrate(db); err != nil {
		t.Fatal(err)
	}
	return db
}

func seedOldWorkerRows(t *testing.T, db *gorm.DB, rows []models.Stats) {
	if err := db.Create(&rows).Error; err != nil {
		t.Fatal(err)
	}
}

func recordAll(t *testing.T, sc *StatsController, now time.Time, msgs []StatsMessage) {
	for _, msg := range msgs {
		if err := sc.record(msg, now); err != nil {
			t.Fatal(err)
		}
	}
}

func runBackfill(t *testing.T, db *gorm.DB) {
	backfill, err := os.ReadFile("../scripts/backfill_stats_totals.sql")
	if err != nil {
		t.Error(err)
		return
	}
	if err := db.Exec(string(backfill)).Error; err != nil {
		t.Error(err)
	}
}

func assertMatchesFullScan(t *testing.T, sc *StatsController) {
	t.Helper()
	want, err := sc.totals(sc.DB.Model(&models.Stats{}))
	if err != nil {
		t.Fatal(err)
	}
	got, err := sc.allTimeTotals()
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Expected totals %+v got %+v", want, got)
	}

	wantServices := serviceNumbers(t, sc, sc.DB.Model(&models.Stats{}), windowServiceColumns)
	gotServices := serviceNumbers(t, sc, sc.DB.Model(&models.StatsServiceDay{}), allTimeServiceColumns)
	if len(gotServices) != len(wantServices) {
		t.Errorf("Expected services %+v got %+v", wantServices, gotServices)
	}
	for svc, numbers := range wantServices {
		if gotServices[svc] != numbers {
			t.Errorf("Expected %q %+v got %+v", svc, numbers, gotServices[svc])
		}
	}
}

func serviceNumbers(t *testing.T, sc *StatsController, query *gorm.DB, columns string) map[string]StatsNumbers {
	t.Helper()
	svcStats, err := sc.serviceStats(query, columns)
	if err != nil {
		t.Fatal(err)
	}
	numbers := map[string]StatsNumbers{}
	for _, s := range svcStats {
		numbers[s.Service] = s.StatsNumbers
	}
	return numbers
}
