package image

import (
	"encoding/json"
	"sync"
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
	ShadowEye         bool             // Replace shadow with calculated shadow
	ColorableRandom   bool             // Replace with random color
	RemovesEyes       bool             // Replaces eyes
	RemovesHandsLeft  bool             // Replaces hands left
	RemovesHandsRight bool             // Replaces hands right
	RemovesLegs       bool             // Replaces feet
	AboveShirtPants   bool             // Should be assembled above SHirtPants
	AboveHands        bool             // Should be assembled above hands
	Weight            float64          // The weight of this accessory, determines how often it appears
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
}

var singleton *assetManager
var once sync.Once

func GetAssets() *assetManager {
	once.Do(func() {
		var err error
		// Deserialize all assets and keep in-mem
		var bodyPartAssets []Asset
		for _, ba := range BodyPartsIllustrations {
			var a Asset
			err = json.Unmarshal(ba, &a)
			bodyPartAssets = append(bodyPartAssets, a)
		}
		var glassesAssets []Asset
		for _, ga := range GlassesIllustrations {
			var a Asset
			err = json.Unmarshal(ga, &a)
			glassesAssets = append(glassesAssets, a)
		}
		var hatsAssets []Asset
		for _, ha := range HatIllustrations {
			var a Asset
			err = json.Unmarshal(ha, &a)
			hatsAssets = append(hatsAssets, a)
		}
		var miscAssets []Asset
		for _, ma := range MiscIllustrations {
			var a Asset
			err = json.Unmarshal(ma, &a)
			miscAssets = append(miscAssets, a)
		}
		var mouthsAssets []Asset
		for _, ma := range MouthsIllustrations {
			var a Asset
			err = json.Unmarshal(ma, &a)
			mouthsAssets = append(mouthsAssets, a)
		}
		var shirtPantsAssets []Asset
		for _, sa := range ShirtPantsIllustrations {
			var a Asset
			err = json.Unmarshal(sa, &a)
			shirtPantsAssets = append(shirtPantsAssets, a)
		}
		var shoesAssets []Asset
		for _, sa := range ShoesIllustrations {
			var a Asset
			err = json.Unmarshal(sa, &a)
			shoesAssets = append(shoesAssets, a)
		}
		var tailsAssets []Asset
		for _, ta := range TailsIllustrations {
			var a Asset
			err = json.Unmarshal(ta, &a)
			tailsAssets = append(tailsAssets, a)
		}
		if err != nil {
			panic("Failed to decode assets")
		}
		// Create object
		singleton = &assetManager{
			bodyPartAssets:   bodyPartAssets,
			glassesAssets:    glassesAssets,
			hatsAssets:       hatsAssets,
			miscAssets:       miscAssets,
			mouthsAssets:     mouthsAssets,
			shirtPantsAssets: shirtPantsAssets,
			shoesAssets:      shoesAssets,
			tailsAssets:      tailsAssets,
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
