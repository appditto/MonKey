package image

import "strings"

// Chance an accessory category has of occuring
const accessoryChance = 0.2

// Accessories - represents accessories for monKey
type Accessories struct {
	TailAsset                *Asset
	TailAccessory            *Asset
	LegAsset                 *Asset
	ArmsAsset                *Asset
	BodyUpperAsset           *Asset
	ShirtPantsAsset          *Asset
	MiscAboveShirtPantsAsset *Asset
	EarAsset                 *Asset
	FaceAsset                *Asset
	EyeAsset                 *Asset
	GlassesAsset             *Asset
	NoseAsset                *Asset
	MouthAsset               *Asset
	HatAsset                 *Asset
	FootLeftAsset            *Asset
	FootRightAsset           *Asset
	HandLeftAsset            *Asset
	HandRightAsset           *Asset
	MiscAboveHandsAsset      *Asset
}

func GetAccessoriesForHash(hash string) (Accessories, error) {
	var accessories = Accessories{}

	// Get base body parts
	for _, bodyPart := range GetAssets().GetBodyParts() {
		if strings.Contains(bodyPart.FileName, "arms") {
			accessories.ArmsAsset = &bodyPart
		} else if strings.Contains(bodyPart.FileName, "body-upper") {
			accessories.BodyUpperAsset = &bodyPart
		} else if strings.Contains(bodyPart.FileName, "ears") {
			accessories.EarAsset = &bodyPart
		} else if strings.Contains(bodyPart.FileName, "eyes") {
			accessories.EyeAsset = &bodyPart
		} else if strings.Contains(bodyPart.FileName, "face") {
			accessories.FaceAsset = &bodyPart
		} else if strings.Contains(bodyPart.FileName, "foot-left") {
			accessories.FootLeftAsset = &bodyPart
		} else if strings.Contains(bodyPart.FileName, "foot-right") {
			accessories.FootRightAsset = &bodyPart
		} else if strings.Contains(bodyPart.FileName, "hand-left") {
			accessories.HandLeftAsset = &bodyPart
		} else if strings.Contains(bodyPart.FileName, "hand-right") {
			accessories.HandRightAsset = &bodyPart
		} else if strings.Contains(bodyPart.FileName, "legs") {
			accessories.LegAsset = &bodyPart
		} else if strings.Contains(bodyPart.FileName, "nose") {
			accessories.NoseAsset = &bodyPart
		} else if strings.Contains(bodyPart.FileName, "tail") {
			accessories.TailAsset = &bodyPart
		}
	}

	// Pick accessories based on accessoryChance

	return accessories, nil
}
