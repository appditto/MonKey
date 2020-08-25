package image

import (
	"strconv"
	"strings"

	"github.com/appditto/monKey/server/color"
	"github.com/golang/glog"
)

// Chance an accessory category has of occuring
const glassesChance = 0.25
const hatChance = 0.3
const miscChance = 0.3
const shirtPantChance = 0.2
const shoeChance = 0.2
const tailChance = 0.2

// Accessories - represents accessories for monKey
type Accessories struct {
	TailAsset       *Asset
	TailAccessory   *Asset
	LegAsset        *Asset
	ArmsAsset       *Asset
	BodyUpperAsset  *Asset
	ShirtPantsAsset *Asset
	MiscAsset       *Asset
	EarAsset        *Asset
	FaceAsset       *Asset
	EyeAsset        *Asset
	GlassesAsset    *Asset
	NoseAsset       *Asset
	MouthAsset      *Asset
	HatAsset        *Asset
	FootLeftAsset   *Asset
	FootRightAsset  *Asset
	ShoeAsset       *Asset
	HandLeftAsset   *Asset
	HandRightAsset  *Asset
	AccessoryColors map[string]color.RGB
}

func GetAccessoriesForHash(hash string) (Accessories, error) {
	var accessories = Accessories{}
	accessories.AccessoryColors = make(map[string]color.RGB)

	// A monKey can use at most 69 characters of a hex string, worst case
	// Assuming it has every type of accessory and every colorable-random accessory
	// We need to add 5 characters to the end of our hash, we'll just take a random sample
	// And add it to the end
	hash = hash + string(hash[4]) + string(hash[2]) + string(hash[0]) + string(hash[19]) + string(hash[23])

	// Keep track of our index on the hash (cannot exceed 69)
	workingIdx := 0

	// Pick fur-color and eye-color
	furColor, _ := GetColor(hash[workingIdx:workingIdx+2], hash[workingIdx+2:workingIdx+4], hash[workingIdx+4:workingIdx+6])
	workingIdx += 6
	eyeColor, _ := GetColor(hash[workingIdx:workingIdx+2], hash[workingIdx+2:workingIdx+4], hash[workingIdx+4:workingIdx+6])
	workingIdx += 6

	// Get base body parts
	for _, bodyPart := range GetAssets().GetBodyParts() {
		localPart := bodyPart
		if strings.Contains(bodyPart.FileName, "arms") {
			if bodyPart.FurColored {
				accessories.AccessoryColors[bodyPart.FileName] = furColor
			} else if bodyPart.EyeColored {
				accessories.AccessoryColors[bodyPart.FileName] = eyeColor
			}
			accessories.ArmsAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "body-upper") {
			if bodyPart.FurColored {
				accessories.AccessoryColors[bodyPart.FileName] = furColor
			} else if bodyPart.EyeColored {
				accessories.AccessoryColors[bodyPart.FileName] = eyeColor
			}
			accessories.BodyUpperAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "ears") {
			if bodyPart.FurColored {
				accessories.AccessoryColors[bodyPart.FileName] = furColor
			} else if bodyPart.EyeColored {
				accessories.AccessoryColors[bodyPart.FileName] = eyeColor
			}
			accessories.EarAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "eyes") {
			if bodyPart.FurColored {
				accessories.AccessoryColors[bodyPart.FileName] = furColor
			} else if bodyPart.EyeColored {
				accessories.AccessoryColors[bodyPart.FileName] = eyeColor
			}
			accessories.EyeAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "face") {
			if bodyPart.FurColored {
				accessories.AccessoryColors[bodyPart.FileName] = furColor
			} else if bodyPart.EyeColored {
				accessories.AccessoryColors[bodyPart.FileName] = eyeColor
			}
			accessories.FaceAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "foot-left") {
			if bodyPart.FurColored {
				accessories.AccessoryColors[bodyPart.FileName] = furColor
			} else if bodyPart.EyeColored {
				accessories.AccessoryColors[bodyPart.FileName] = eyeColor
			}
			accessories.FootLeftAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "foot-right") {
			if bodyPart.FurColored {
				accessories.AccessoryColors[bodyPart.FileName] = furColor
			} else if bodyPart.EyeColored {
				accessories.AccessoryColors[bodyPart.FileName] = eyeColor
			}
			accessories.FootRightAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "hand-left") {
			if bodyPart.FurColored {
				accessories.AccessoryColors[bodyPart.FileName] = furColor
			} else if bodyPart.EyeColored {
				accessories.AccessoryColors[bodyPart.FileName] = eyeColor
			}
			accessories.HandLeftAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "hand-right") {
			if bodyPart.FurColored {
				accessories.AccessoryColors[bodyPart.FileName] = furColor
			} else if bodyPart.EyeColored {
				accessories.AccessoryColors[bodyPart.FileName] = eyeColor
			}
			accessories.HandRightAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "legs") {
			if bodyPart.FurColored {
				accessories.AccessoryColors[bodyPart.FileName] = furColor
			} else if bodyPart.EyeColored {
				accessories.AccessoryColors[bodyPart.FileName] = eyeColor
			}
			accessories.LegAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "nose") {
			if bodyPart.FurColored {
				accessories.AccessoryColors[bodyPart.FileName] = furColor
			} else if bodyPart.EyeColored {
				accessories.AccessoryColors[bodyPart.FileName] = eyeColor
			}
			accessories.NoseAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "tail") {
			if bodyPart.FurColored {
				accessories.AccessoryColors[bodyPart.FileName] = furColor
			} else if bodyPart.EyeColored {
				accessories.AccessoryColors[bodyPart.FileName] = eyeColor
			}
			accessories.TailAsset = &localPart
		}
	}

	// Pick accessories based on accessoryChance

	// Get threshold for existing
	glassesChanceLocal := glassesChance
	maxHasAccessoryValue := int64(255 * glassesChanceLocal)
	hasGlasses := false
	hasGlassesWorkingVal, _ := strconv.ParseInt(hash[workingIdx:2+workingIdx], 16, 64)
	if hasGlassesWorkingVal <= maxHasAccessoryValue {
		hasGlasses = true
	}
	workingIdx += 2

	hatChanceLocal := hatChance
	maxHasAccessoryValue = int64(255 * hatChanceLocal)
	hasHats := false
	hasHatsWorkingVal, _ := strconv.ParseInt(hash[workingIdx:2+workingIdx], 16, 64)
	if hasHatsWorkingVal <= maxHasAccessoryValue {
		hasHats = true
	}
	workingIdx += 2

	miscChanceLocal := miscChance
	maxHasAccessoryValue = int64(255 * miscChanceLocal)
	hasMisc := false
	hasMiscWorkingVal, _ := strconv.ParseInt(hash[workingIdx:2+workingIdx], 16, 64)
	if hasMiscWorkingVal <= maxHasAccessoryValue {
		hasMisc = true
	}
	workingIdx += 2

	shirtPantChanceLocal := shirtPantChance
	maxHasAccessoryValue = int64(255 * shirtPantChanceLocal)
	hasShirtPants := false
	hasShirtPantsWorkingVal, _ := strconv.ParseInt(hash[workingIdx:2+workingIdx], 16, 64)
	if hasShirtPantsWorkingVal <= maxHasAccessoryValue {
		hasShirtPants = true
	}
	workingIdx += 2

	shoeChanceLocal := shoeChance
	maxHasAccessoryValue = int64(255 * shoeChanceLocal)
	hasShoes := false
	hasShoesWorkingVal, _ := strconv.ParseInt(hash[workingIdx:2+workingIdx], 16, 64)
	if hasShoesWorkingVal <= maxHasAccessoryValue {
		hasShoes = true
	}
	workingIdx += 2

	tailChanceLocal := tailChance
	maxHasAccessoryValue = int64(255 * tailChanceLocal)
	hasTails := false
	hasTailsWorkingVal, _ := strconv.ParseInt(hash[workingIdx:2+workingIdx], 16, 64)
	if hasTailsWorkingVal <= maxHasAccessoryValue {
		hasTails = true
	}
	workingIdx += 2 // Up to 24 at this point

	// Pick accessories if we have them
	if hasGlasses {
		accessories.GlassesAsset = GetAccessoryFromHexWithWeight(hash[workingIdx:3+workingIdx], GetAssets().GetGlasses())
		workingIdx += 3
	}

	if hasHats {
		accessories.HatAsset = GetAccessoryFromHexWithWeight(hash[workingIdx:3+workingIdx], GetAssets().GetHats())
		if accessories.HatAsset.ColorableRandom {
			randColor, _ := GetColor(hash[workingIdx:workingIdx+2], hash[workingIdx+2:workingIdx+4], hash[workingIdx+4:workingIdx+6])
			workingIdx += 6
			accessories.AccessoryColors[accessories.HatAsset.FileName] = randColor
		}
		workingIdx += 3
	}

	if hasMisc {
		accessories.MiscAsset = GetAccessoryFromHexWithWeight(hash[workingIdx:3+workingIdx], GetAssets().GetMisc())
		workingIdx += 3
	}

	if hasShirtPants {
		accessories.ShirtPantsAsset = GetAccessoryFromHexWithWeight(hash[workingIdx:3+workingIdx], GetAssets().GetShirtPantsAssets())
		if accessories.ShirtPantsAsset.ColorableRandom {
			randColor, _ := GetColor(hash[workingIdx:workingIdx+2], hash[workingIdx+2:workingIdx+4], hash[workingIdx+4:workingIdx+6])
			workingIdx += 6
			accessories.AccessoryColors[accessories.ShirtPantsAsset.FileName] = randColor
		}
		workingIdx += 3
	}

	if hasShoes {
		accessories.ShoeAsset = GetAccessoryFromHexWithWeight(hash[workingIdx:3+workingIdx], GetAssets().GetShoeAssets())
		if accessories.ShoeAsset.ColorableRandom {
			randColor, _ := GetColor(hash[workingIdx:workingIdx+2], hash[workingIdx+2:workingIdx+4], hash[workingIdx+4:workingIdx+6])
			workingIdx += 6
			accessories.AccessoryColors[accessories.ShoeAsset.FileName] = randColor
		}
		workingIdx += 3
	}

	if hasTails {
		accessories.TailAccessory = GetAccessoryFromHexWithWeight(hash[workingIdx:3+workingIdx], GetAssets().GetTailAssets())
		if accessories.TailAccessory.ColorableRandom {
			randColor, _ := GetColor(hash[workingIdx:workingIdx+2], hash[workingIdx+2:workingIdx+4], hash[workingIdx+4:workingIdx+6])
			workingIdx += 6
			accessories.AccessoryColors[accessories.TailAccessory.FileName] = randColor
		}
		workingIdx += 3
	}

	// Mouth always exists
	accessories.MouthAsset = GetAccessoryFromHexWithWeight(hash[workingIdx:3+workingIdx], GetAssets().GetMouthAssets())
	workingIdx += 3 // Up to 69 at this point

	return accessories, nil
}

// Get weighted accessory given 2-digit hex string
func GetAccessoryFromHexWithWeight(hexStr string, assets []Asset) *Asset {
	maxPossible := 4095.0
	target, _ := strconv.ParseInt(hexStr, 16, 64)

	// Interval starting point
	curStart := 0

	// Find combined weight of each asset
	combinedWeight := 0.0
	for _, asset := range assets {
		combinedWeight += asset.Weight
	}

	// Find target of each asset
	for idx, asset := range assets {
		targetVal := int((asset.Weight / combinedWeight) * maxPossible)
		if targetVal < 1 {
			targetVal = 1
		}
		curEnd := curStart + targetVal
		if idx == len(assets)-1 {
			curEnd = int(maxPossible)
		}
		if int(target) >= curStart && int(target) < curEnd {
			return &asset
		}
		curStart += targetVal
	}

	// Fallback in case nothing happened
	glog.Warningf("Couldn't find asset with appropriate weight, defaulting to index-0 asset %s %s", hexStr, assets[0].FileName)
	return &assets[0]
}
