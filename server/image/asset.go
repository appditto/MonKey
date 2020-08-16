package image

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/golang/glog"
)

type IllustrationType string

const (
	BodyPart   IllustrationType = "body-parts"
	Glasses    IllustrationType = "glasses"
	Hats       IllustrationType = "hats"
	Misc       IllustrationType = "misc"
	Mouths     IllustrationType = "mouths"
	ShirtPants IllustrationType = "shirt-pants"
	Shoes      IllustrationType = "shoes"
	Tails      IllustrationType = "tails"
	Vanity     IllustrationType = "vanities"
)

type Asset struct {
	FileName          string           // File name of asset
	IllustrationPath  string           // Full path of illustration on the file system
	Type              IllustrationType // Type of illustration (body, hair, mouth, eye)
	SVGContents       []byte           // Full contents of SVG asset
	FurColored        bool             // whether this asset is fur colored
	EyeColored        bool             // whether this asset should be eye colored
	ShadowFur         bool             // Replace shadow with calculated shadow
	ShadowFurDark     bool             // Replace shadow with calculated shadow
	ShadowIris        bool             // Replace shadow with calculated shadow
	ColorableRandom   bool             // Replace with random color
	RemovesEyes       bool             // Replaces eyes
	RemovesHandsLeft  bool             // Replaces hands left
	RemovesHandsRight bool             // Replaces hands right
	RemovesLegs       bool             // Replaces legs
	RemovesFeet       bool             // Replaces feet
	AboveShirtPants   bool             // Should be assembled above SHirtPants
	AboveHands        bool             // Should be assembled above hands
	Weight            float64          // The weight of this accessory, determines how often it appears
	Address           string           // Fixed address this vanity belongs to
}

// Singleton to keep assets loaded in memory
type assetManager struct {
	bodyPartAssets   []Asset
	glassesAssets    []Asset
	hatsAssets       []Asset
	miscAssets       []Asset
	mouthsAssets     []Asset
	shirtPantsAssets []Asset
	shoesAssets      []Asset
	tailsAssets      []Asset
	vanityAssets     map[string]Asset
}

// Helpers for loading assets
func hasTag(fname string, tag string) bool {
	return strings.Contains(fname, tag)
}

func getWeight(fname string) float64 {
	reWeightSubstr := regexp.MustCompile(`\[w-(.*)`)
	weightStr := reWeightSubstr.FindString(fname)
	reWeightFloat := regexp.MustCompile(`[0-9]?([0-9]*[.])?[0-9]+`)
	weightNumStr := reWeightFloat.FindString(weightStr)

	asF, err := strconv.ParseFloat(weightNumStr, 64)
	if err != nil {
		return -1
	}
	return asF
}

func getAccessoryAsset(fname string, path string) Asset {
	var err error
	asset := Asset{}
	asset.FileName = fname
	asset.IllustrationPath = path
	asset.SVGContents, err = ioutil.ReadFile(path)
	if err != nil {
		glog.Fatalf("Couldn't load file %s", path)
		panic(err.Error())
	}
	asset.FurColored = false
	asset.EyeColored = false
	asset.ShadowFur = false
	asset.ShadowFurDark = false
	asset.ShadowIris = false
	asset.ColorableRandom = hasTag(asset.FileName, "[colorable-random]")
	asset.RemovesEyes = hasTag(asset.FileName, "[removes-eyes]")
	asset.RemovesLegs = hasTag(asset.FileName, "[removes-legs]")
	asset.RemovesFeet = hasTag(asset.FileName, "[removes-feet]")
	asset.RemovesHandsLeft = hasTag(asset.FileName, "[removes-hands-left]") || hasTag(asset.FileName, "[removes-hands]")
	asset.RemovesHandsRight = hasTag(asset.FileName, "[removes-hands-right]") || hasTag(asset.FileName, "[removes-hands]")
	asset.AboveShirtPants = hasTag(asset.FileName, "[above-shirts-pants]")
	asset.AboveHands = hasTag(asset.FileName, "[above-hands]")
	asset.Weight = getWeight(asset.FileName)
	return asset
}

var singleton *assetManager
var once sync.Once

func GetAssets() *assetManager {
	once.Do(func() {
		var err error

		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		var bodyPartAssets []Asset
		var bodyAsset Asset
		fPath := path.Join(wd, "assets", "illustrations", string(BodyPart))
		err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
			if strings.Contains(info.Name(), ".svg") {
				bodyAsset = Asset{}
				bodyAsset.FileName = info.Name()
				bodyAsset.IllustrationPath = path
				bodyAsset.SVGContents, err = ioutil.ReadFile(path)
				if err != nil {
					glog.Fatalf("Couldn't load file %s", path)
					panic(err.Error())
				}
				bodyAsset.FurColored = hasTag(bodyAsset.FileName, "[fur-color]")
				bodyAsset.EyeColored = hasTag(bodyAsset.FileName, "[eye-color]")
				bodyAsset.ShadowFur = hasTag(bodyAsset.FileName, "[shadow-fur]")
				bodyAsset.ShadowFurDark = hasTag(bodyAsset.FileName, "[shadow-fur-dark]")
				bodyAsset.ShadowIris = hasTag(bodyAsset.FileName, "[shadow-iris]")
				bodyAsset.ColorableRandom = false
				bodyAsset.RemovesEyes = false
				bodyAsset.RemovesLegs = false
				bodyAsset.RemovesHandsLeft = false
				bodyAsset.RemovesHandsRight = false
				bodyAsset.AboveShirtPants = false
				bodyAsset.AboveHands = false
				bodyAsset.Weight = getWeight(bodyAsset.FileName)
				bodyPartAssets = append(bodyPartAssets, bodyAsset)
			}
			return nil
		})

		var hatAssets []Asset
		fPath = path.Join(wd, "assets", "illustrations", "accessories", string(Hats))
		err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
			if strings.Contains(info.Name(), ".svg") {
				hatAssets = append(hatAssets, getAccessoryAsset(info.Name(), path))
			}
			return nil
		})

		var glassesAssets []Asset
		fPath = path.Join(wd, "assets", "illustrations", "accessories", string(Glasses))
		err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
			if strings.Contains(info.Name(), ".svg") {
				glassesAssets = append(glassesAssets, getAccessoryAsset(info.Name(), path))
			}
			return nil
		})

		var miscAssets []Asset
		fPath = path.Join(wd, "assets", "illustrations", "accessories", string(Misc))
		err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
			if strings.Contains(info.Name(), ".svg") {
				miscAssets = append(miscAssets, getAccessoryAsset(info.Name(), path))
			}
			return nil
		})

		var mouthAssets []Asset
		fPath = path.Join(wd, "assets", "illustrations", "accessories", string(Mouths))
		err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
			if strings.Contains(info.Name(), ".svg") {
				mouthAssets = append(mouthAssets, getAccessoryAsset(info.Name(), path))
			}
			return nil
		})

		var shirtPantsAssets []Asset
		fPath = path.Join(wd, "assets", "illustrations", "accessories", string(ShirtPants))
		err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
			if strings.Contains(info.Name(), ".svg") {
				shirtPantsAssets = append(shirtPantsAssets, getAccessoryAsset(info.Name(), path))
			}
			return nil
		})

		var shoeAssets []Asset
		fPath = path.Join(wd, "assets", "illustrations", "accessories", string(Shoes))
		err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
			if strings.Contains(info.Name(), ".svg") {
				shoeAssets = append(shoeAssets, getAccessoryAsset(info.Name(), path))
			}
			return nil
		})

		var tailAssets []Asset
		fPath = path.Join(wd, "assets", "illustrations", "accessories", string(Tails))
		err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
			if strings.Contains(info.Name(), ".svg") {
				tailAssets = append(tailAssets, getAccessoryAsset(info.Name(), path))
			}
			return nil
		})

		var vanityAsset Asset
		vanityAssets := make(map[string]Asset)
		fPath = path.Join(wd, "assets", "illustrations", "accessories", string(Tails))
		err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
			if strings.Contains(info.Name(), ".svg") {
				vanityAsset = Asset{}
				vanityAsset.FileName = info.Name()
				vanityAsset.IllustrationPath = path
				vanityAsset.SVGContents, err = ioutil.ReadFile(path)
				if err != nil {
					glog.Fatalf("Couldn't load file %s", path)
					panic(err.Error())
				}
				vanityAsset.Address = strings.Split(vanityAsset.FileName, ".svg")[0]
				vanityAssets[vanityAsset.Address] = vanityAsset
			}
			return nil
		})

		if err != nil {
			panic("Failed to decode assets")
		}
		// Create object
		singleton = &assetManager{
			bodyPartAssets:   bodyPartAssets,
			glassesAssets:    glassesAssets,
			hatsAssets:       hatAssets,
			miscAssets:       miscAssets,
			mouthsAssets:     mouthAssets,
			shirtPantsAssets: shirtPantsAssets,
			shoesAssets:      shoeAssets,
			tailsAssets:      tailAssets,
			vanityAssets:     vanityAssets,
		}
	})
	return singleton
}

// GetBodyParts - get complete list of body parts assets
func (sm *assetManager) GetBodyParts() []Asset {
	return sm.bodyPartAssets
}

// GetGlasses - get complete list of glasses parts assets
func (sm *assetManager) GetGlasses() []Asset {
	return sm.glassesAssets
}

// GetHats - get complete list of hat parts assets
func (sm *assetManager) GetHats() []Asset {
	return sm.hatsAssets
}

// GetMisc - get complete list of misc assets
func (sm *assetManager) GetMisc() []Asset {
	return sm.miscAssets
}

// GetMouthAssets - get complete list of mouth assets
func (sm *assetManager) GetMouthAssets() []Asset {
	return sm.mouthsAssets
}

// GetShirtPantsAssets - get complete list of shirt pants assets
func (sm *assetManager) GetShirtPantsAssets() []Asset {
	return sm.shirtPantsAssets
}

// GetShoeAssets - get complete list of shoe assets
func (sm *assetManager) GetShoeAssets() []Asset {
	return sm.shoesAssets
}

// GetTailAssets - get complete list of tail assets
func (sm *assetManager) GetTailAssets() []Asset {
	return sm.tailsAssets
}

// GetVanityAsset - get vanity asset
func (sm *assetManager) GetVanityAsset(address string) []byte {
	if asset, ok := sm.vanityAssets[strings.ToLower(address)]; ok {
		return asset.SVGContents
	}
	return nil
}
