package image

import (
	"strconv"
	"strings"

	"github.com/golang/glog"
)

// Chance an accessory category has of occuring
const glassesChance = 0.2
const hatChance = 0.2
const miscChance = 0.2
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
	maxHasAccessoryValue := int64(4095 * glassesChance)
	workingIdx := 0
	hasGlasses := false
	hasGlassesWorkingVal, _ := strconv.ParseInt(hash[workingIdx:3+workingIdx], 16, 64)
	if hasGlassesWorkingVal <= maxHasAccessoryValue {
		hasGlasses = true
	}
	workingIdx += 3

	maxHasAccessoryValue = int64(4095 * hatChance)
	hasHats := false
	hasHatsWorkingVal, _ := strconv.ParseInt(hash[workingIdx:3+workingIdx], 16, 64)
	if hasHatsWorkingVal <= maxHasAccessoryValue {
		hasHats = true
	}
	workingIdx += 3

	maxHasAccessoryValue = int64(4095 * miscChance)
	hasMisc := false
	hasMiscWorkingVal, _ := strconv.ParseInt(hash[workingIdx:3+workingIdx], 16, 64)
	if hasMiscWorkingVal <= maxHasAccessoryValue {
		hasMisc = true
	}
	workingIdx += 3

	maxHasAccessoryValue = int64(4095 * shirtPantChance)
	hasShirtPants := false
	hasShirtPantsWorkingVal, _ := strconv.ParseInt(hash[workingIdx:3+workingIdx], 16, 64)
	if hasShirtPantsWorkingVal <= maxHasAccessoryValue {
		hasShirtPants = true
	}
	workingIdx += 3

	maxHasAccessoryValue = int64(4095 * shoeChance)
	hasShoes := false
	hasShoesWorkingVal, _ := strconv.ParseInt(hash[workingIdx:3+workingIdx], 16, 64)
	if hasShoesWorkingVal <= maxHasAccessoryValue {
		hasShoes = true
	}
	workingIdx += 3

	maxHasAccessoryValue = int64(4095 * tailChance)
	hasTails := false
	hasTailsWorkingVal, _ := strconv.ParseInt(hash[workingIdx:3+workingIdx], 16, 64)
	if hasTailsWorkingVal <= maxHasAccessoryValue {
		hasTails = true
	}
	workingIdx += 3 // 18

	// Pick accessories if we have them
	if hasGlasses {
		accessories.GlassesAsset = GetAccessoryFromHexWithWeight(hash[workingIdx:3+workingIdx], GetAssets().GetGlasses())
		workingIdx += 3
	}

	if hasHats {
		accessories.HatAsset = GetAccessoryFromHexWithWeight(hash[workingIdx:3+workingIdx], GetAssets().GetHats())
		workingIdx += 3
	}

	if hasMisc {
		accessories.MiscAsset = GetAccessoryFromHexWithWeight(hash[workingIdx:3+workingIdx], GetAssets().GetMisc())
		workingIdx += 3
	}

	if hasShirtPants {
		accessories.ShirtPantsAsset = GetAccessoryFromHexWithWeight(hash[workingIdx:3+workingIdx], GetAssets().GetShirtPantsAssets())
		workingIdx += 3
	}

	if hasShoes {
		accessories.ShoeAsset = GetAccessoryFromHexWithWeight(hash[workingIdx:3+workingIdx], GetAssets().GetShoeAssets())
		workingIdx += 3
	}

	if hasTails {
		accessories.TailAccessory = GetAccessoryFromHexWithWeight(hash[workingIdx:3+workingIdx], GetAssets().GetTailAssets())
		workingIdx += 3
	}

	// Mouth always exists
	accessories.MouthAsset = GetAccessoryFromHexWithWeight(hash[workingIdx:3+workingIdx], GetAssets().GetMouthAssets())
	workingIdx += 3

	// Working idx could be up to 39 here, if we dont have enough left for color we can either re-use parts of the hash or re-hash
	// TODO - color

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
