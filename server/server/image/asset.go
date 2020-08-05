package image

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
	RemovesHands      bool             // Replaces hands
	RemovesHandsLeft  bool             // Replaces handsl eft
	RemovesHandsRight bool             // Replaces hands right
	AboveShirtPants   bool             // Should be assembled above SHirtPants
	AboveHands        bool             // Should be assembled above hands
	Weight            float64          // The weight of this accessory, determines how often it appears
}
