package controller

import (
	"fmt"
	"os"
	"path"

	"github.com/appditto/monKey/server/image"
	"github.com/appditto/monKey/server/utils"
)

// For generating CSV documents for algorithm analysis
func TestAccessoryDistribution(seed string) {
	wd, _ := os.Getwd()
	output := path.Join(wd, "accessory_distribution.csv")
	outputF, err := os.Create(output)
	defer outputF.Close()
	if err != nil {
		fmt.Printf("Failed to open file for writing %s", output)
	}
	var address string
	var sha256 string
	var accessories image.Accessories
	iterations := 100000
	ret := "glas,hats,misc,mout,shpt,shoe,tail\n"
	glas := 0
	hats := 0
	misc := 0
	shpt := 0
	shoe := 0
	tail := 0
	countMap := make(map[string]int)
	countMap["one"] = 0
	countMap["two"] = 0
	countMap["three"] = 0
	countMap["four"] = 0
	countMap["five"] = 0
	countMap["six"] = 0
	for i := 0; i < iterations; i++ {
		address = utils.GenerateAddress()
		sha256 = utils.Sha256(address, seed)
		accessories, _ = image.GetAccessoriesForHash(sha256)
		glassesName := "none"
		localCount := 0
		if accessories.GlassesAsset != nil {
			glassesName = accessories.GlassesAsset.FileName
			glas += 1
			localCount += 1
		}
		hatName := "none"
		if accessories.HatAsset != nil {
			hatName = accessories.HatAsset.FileName
			hats += 1
			localCount += 1
		}
		miscName := "none"
		if accessories.MiscAsset != nil {
			miscName = accessories.MiscAsset.FileName
			misc += 1
			localCount += 1
		}
		shptName := "none"
		if accessories.ShirtPantsAsset != nil {
			shptName = accessories.ShirtPantsAsset.FileName
			shpt += 1
			localCount += 1
		}
		shoeName := "none"
		if accessories.ShoeAsset != nil {
			shoeName = accessories.ShoeAsset.FileName
			shoe += 1
			localCount += 1
		}
		tailName := "none"
		if accessories.TailAccessory != nil {
			tailName = accessories.TailAccessory.FileName
			tail += 1
			localCount += 1
		}
		mouthName := "none"
		if accessories.MouthAsset != nil {
			mouthName = accessories.MouthAsset.FileName
		}
		ret += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s\n", glassesName, hatName, miscName, shptName, shoeName, tailName, mouthName)
		switch localCount {
		case 1:
			countMap["one"] += 1
		case 2:
			countMap["one"] += 1
			countMap["two"] += 1
		case 3:
			countMap["one"] += 1
			countMap["two"] += 1
			countMap["three"] += 1
		case 4:
			countMap["one"] += 1
			countMap["two"] += 1
			countMap["three"] += 1
			countMap["four"] += 1
		case 5:
			countMap["one"] += 1
			countMap["two"] += 1
			countMap["three"] += 1
			countMap["four"] += 1
			countMap["five"] += 1
		case 6:
			countMap["one"] += 1
			countMap["two"] += 1
			countMap["three"] += 1
			countMap["four"] += 1
			countMap["five"] += 1
			countMap["six"] += 1
		}
	}
	outputF.WriteString(ret)
	fmt.Printf("Total iterations: %d\n-----\n", iterations)
	fmt.Printf("Glasses: %d, Percent of Total %f%%\n", glas, (float64(glas)/float64(iterations))*100.0)
	fmt.Printf("Hats: %d, Percent of Total %f%%\n", hats, (float64(hats)/float64(iterations))*100.0)
	fmt.Printf("Misc: %d, Percent of Total %f%%\n", misc, (float64(misc)/float64(iterations))*100.0)
	fmt.Printf("Shirt pants: %d, Percent of Total %f%%\n", shpt, (float64(shpt)/float64(iterations))*100.0)
	fmt.Printf("Shoes: %d, Percent of Total %f%%\n", shoe, (float64(shoe)/float64(iterations))*100.0)
	fmt.Printf("Tail A: %d, Percent of Total %f%%\n-----\n", tail, (float64(tail)/float64(iterations))*100.0)
	fmt.Printf("Percent with at least 1 accessory %f%%\n", (float64(countMap["one"])/float64(iterations))*100.0)
	fmt.Printf("Percent with at least 2 accessories %f%%\n", (float64(countMap["two"])/float64(iterations))*100.0)
	fmt.Printf("Percent with at least 3 accessories %f%%\n", (float64(countMap["three"])/float64(iterations))*100.0)
	fmt.Printf("Percent with at least 4 accessories %f%%\n", (float64(countMap["four"])/float64(iterations))*100.0)
	fmt.Printf("Percent with at least 5 accessories %f%%\n", (float64(countMap["five"])/float64(iterations))*100.0)
	fmt.Printf("Percent with at least 6 accessories %f%%\n", (float64(countMap["six"])/float64(iterations))*100.0)
}
