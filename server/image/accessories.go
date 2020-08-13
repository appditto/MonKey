package image

import (
	"strconv"
	"strings"

	"github.com/golang/glog"
)

// Chance an accessory category has of occuring
const accessoryChance = 0.2

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
}

func GetAccessoriesForHash(hash string) (Accessories, error) {
	var accessories = Accessories{}

	// Get base body parts
	for _, bodyPart := range GetAssets().GetBodyParts() {
		localPart := bodyPart
		if strings.Contains(bodyPart.FileName, "arms") {
			accessories.ArmsAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "body-upper") {
			accessories.BodyUpperAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "ears") {
			accessories.EarAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "eyes") {
			accessories.EyeAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "face") {
			accessories.FaceAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "foot-left") {
			accessories.FootLeftAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "foot-right") {
			accessories.FootRightAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "hand-left") {
			accessories.HandLeftAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "hand-right") {
			accessories.HandRightAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "legs") {
			accessories.LegAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "nose") {
			accessories.NoseAsset = &localPart
		} else if strings.Contains(bodyPart.FileName, "tail") {
			accessories.TailAsset = &localPart
		}
	}

	// Pick accessories based on accessoryChance
	// Get threshold for existing
	maxHasAccessoryValue := int64(255 * accessoryChance)

	workingIdx := 0
	hasGlasses := false
	hasGlassesWorkingVal, _ := strconv.ParseInt(hash[workingIdx:2+workingIdx], 16, 64)
	if hasGlassesWorkingVal <= maxHasAccessoryValue {
		hasGlasses = true
	}
	workingIdx += 2

	hasHats := false
	hasHatsWorkingVal, _ := strconv.ParseInt(hash[workingIdx:2+workingIdx], 16, 64)
	if hasHatsWorkingVal <= maxHasAccessoryValue {
		hasHats = true
	}
	workingIdx += 2

	hasMisc := false
	hasMiscWorkingVal, _ := strconv.ParseInt(hash[workingIdx:2+workingIdx], 16, 64)
	if hasMiscWorkingVal <= maxHasAccessoryValue {
		hasMisc = true
	}
	workingIdx += 2

	hasMouth := false
	hasMouthWorkingVal, _ := strconv.ParseInt(hash[workingIdx:2+workingIdx], 16, 64)
	if hasMouthWorkingVal <= maxHasAccessoryValue {
		hasMouth = true
	}
	workingIdx += 2

	hasShirtPants := false
	hasShirtPantsWorkingVal, _ := strconv.ParseInt(hash[workingIdx:2+workingIdx], 16, 64)
	if hasShirtPantsWorkingVal <= maxHasAccessoryValue {
		hasShirtPants = true
	}
	workingIdx += 2

	hasShoes := false
	hasShoesWorkingVal, _ := strconv.ParseInt(hash[workingIdx:2+workingIdx], 16, 64)
	if hasShoesWorkingVal <= maxHasAccessoryValue {
		hasShoes = true
	}
	workingIdx += 2

	hasTails := false
	hasTailsWorkingVal, _ := strconv.ParseInt(hash[workingIdx:2+workingIdx], 16, 64)
	if hasTailsWorkingVal <= maxHasAccessoryValue {
		hasTails = true
	}
	workingIdx += 2 // 14

	// Pick accessories if we have them
	if hasGlasses {
		accessories.GlassesAsset = GetAccessoryFromHexWithWeight(hash[workingIdx:2+workingIdx], GetAssets().GetGlasses())
		workingIdx += 2
	}

	if hasHats {
		accessories.HatAsset = GetAccessoryFromHexWithWeight(hash[workingIdx:2+workingIdx], GetAssets().GetHats())
		workingIdx += 2
	}

	if hasMisc {
		accessories.MiscAsset = GetAccessoryFromHexWithWeight(hash[workingIdx:2+workingIdx], GetAssets().GetMisc())
		workingIdx += 2
	}

	if hasMouth {
		accessories.MouthAsset = GetAccessoryFromHexWithWeight(hash[workingIdx:2+workingIdx], GetAssets().GetMouthAssets())
		workingIdx += 2
	}

	if hasShirtPants {
		accessories.ShirtPantsAsset = GetAccessoryFromHexWithWeight(hash[workingIdx:2+workingIdx], GetAssets().GetShirtPantsAssets())
		workingIdx += 2
	}

	if hasShoes {
		accessories.ShoeAsset = GetAccessoryFromHexWithWeight(hash[workingIdx:2+workingIdx], GetAssets().GetShoeAssets())
		workingIdx += 2
	}

	if hasTails {
		accessories.TailAsset = GetAccessoryFromHexWithWeight(hash[workingIdx:2+workingIdx], GetAssets().GetTailAssets())
		workingIdx += 2
	}

	// Working idx could be up to 28 here, if we dont have enough left for color we can either re-use parts of the hash or re-hash
	// TODO - color

	return accessories, nil
}

// Get weighted accessory given 2-digit hex string
func GetAccessoryFromHexWithWeight(hexStr string, assets []Asset) *Asset {
	maxPossible := 255.0
	asInt, _ := strconv.ParseInt(hexStr, 16, 64)

	// Map start/end value for each asset
	// Represents an inclusive start, exclusive end
	assetValueMap := make(map[string][]int)
	curStart := 0

	// Find combined weight of each asset
	combinedWeight := 0.0
	for _, asset := range assets {
		combinedWeight += asset.Weight
	}

	// Average size of each
	averageSize := maxPossible / combinedWeight

	// Find range of assets < 1 weight
	for _, asset := range assets {
		if asset.Weight < 1 {
			endVal := int(averageSize * asset.Weight)
			if endVal <= 0 {
				endVal = 1
			}
			assetValueMap[asset.FileName] = []int{curStart, curStart + endVal}
			curStart = curStart + endVal
		}
	}

	// Re-calculate average given new weights
	maxPossible -= float64(curStart)
	combinedWeight = 0
	if maxPossible != 255.0 {
		for _, asset := range assets {
			if asset.Weight >= 1 {
				combinedWeight += 1
			}
		}
		averageSize = maxPossible / combinedWeight
		for _, asset := range assets {
			if asset.Weight >= 1 {
				endVal := int(averageSize)
				assetValueMap[asset.FileName] = []int{curStart, curStart + endVal}
				curStart = curStart + endVal
			}
		}
	}

	// Pick asset
	for fname, interval := range assetValueMap {
		if interval[0] <= int(asInt) && interval[1] > int(asInt) {
			for _, asset := range assets {
				if asset.FileName == fname {
					return &asset
				}
			}
		}
	}

	// Fallback in case nothing happened
	glog.Warningf("Couldn't find asset with appropriate weight, defaulting to index-0 asset %s %s", hexStr, assets[0].FileName)
	return &assets[0]
}
