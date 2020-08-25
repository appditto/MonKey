package image

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
	"sync"

	svg "github.com/ajstarks/svgo"
	"github.com/golang/glog"
	minify "github.com/tdewolff/minify/v2"
	minifyxml "github.com/tdewolff/minify/v2/xml"
)

// DefaultSize SVG width/height attribute
const DefaultSize = 4000

// Replace white with this color on bw assets
const lodBwReplacement = "#9CA2AF"

// SVG struct
type SVG struct {
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
	Doc    string `xml:",innerxml"`
}

// PureSVG Return SVG minified, minus width/height/etc attributes
func PureSVG(svgData []byte) ([]byte, error) {
	var pureSVG SVG
	if err := xml.Unmarshal(svgData, &pureSVG); err != nil {
		glog.Fatal("Unable to parse SVG")
		return nil, err
	}
	var b bytes.Buffer
	canvas := svg.New(&b)
	canvas.Startraw(fmt.Sprintf("viewBox=\"0 0 %d %d\"", DefaultSize, DefaultSize), "fill=\"none\"")
	io.WriteString(canvas.Writer, pureSVG.Doc)
	// End document
	canvas.End()

	return b.Bytes(), nil
}

// CombineSVG is a function that combines svgs
func CombineSVG(accessories Accessories) ([]byte, error) {
	var (
		tail          SVG
		tailAccessory SVG
		legs          SVG
		arms          SVG
		bodyUpper     SVG
		shirtPants    SVG
		misc          SVG
		ears          SVG
		face          SVG
		eyes          SVG
		glasses       SVG
		nose          SVG
		mouth         SVG
		hat           SVG
		footLeft      SVG
		footRight     SVG
		shoes         SVG
		handLeft      SVG
		handRight     SVG
	)
	// Parse all SVG assets that are relevant to the final image
	if err := xml.Unmarshal(accessories.TailAsset.SVGContents, &tail); err != nil {
		glog.Fatalf("Unable to parse tail SVG %v", err)
		return nil, err
	}
	if accessories.TailAccessory != nil {
		if err := xml.Unmarshal(accessories.TailAccessory.SVGContents, &tailAccessory); err != nil {
			glog.Fatalf("Unable to parse tail Accessory SVG %v", err)
			return nil, err
		}
	}
	if accessories.ShirtPantsAsset == nil || !accessories.ShirtPantsAsset.RemovesLegs {
		if err := xml.Unmarshal(accessories.LegAsset.SVGContents, &legs); err != nil {
			glog.Fatalf("Unable to parse legs SVG %v", err)
			return nil, err
		}
	}
	if err := xml.Unmarshal(accessories.ArmsAsset.SVGContents, &arms); err != nil {
		glog.Fatalf("Unable to parse arms SVG %v", err)
		return nil, err
	}
	if err := xml.Unmarshal(accessories.BodyUpperAsset.SVGContents, &bodyUpper); err != nil {
		glog.Fatalf("Unable to parse body upper SVG %v", err)
		return nil, err
	}
	if accessories.ShirtPantsAsset != nil {
		if err := xml.Unmarshal(accessories.ShirtPantsAsset.SVGContents, &shirtPants); err != nil {
			glog.Fatalf("Unable to parse shirt pants SVG %v", err)
			return nil, err
		}
	}
	if accessories.MiscAsset != nil {
		if err := xml.Unmarshal(accessories.MiscAsset.SVGContents, &misc); err != nil {
			glog.Fatalf("Unable to parse misc SVG %v", err)
			return nil, err
		}
	}
	if err := xml.Unmarshal(accessories.EarAsset.SVGContents, &ears); err != nil {
		glog.Fatalf("Unable to parse ears SVG %v", err)
		return nil, err
	}
	if err := xml.Unmarshal(accessories.FaceAsset.SVGContents, &face); err != nil {
		glog.Fatalf("Unable to parse face SVG %v", err)
		return nil, err
	}
	if accessories.GlassesAsset == nil || !accessories.GlassesAsset.RemovesEyes {
		if err := xml.Unmarshal(accessories.EyeAsset.SVGContents, &eyes); err != nil {
			glog.Fatalf("Unable to parse eye SVG %v", err)
			return nil, err
		}
	}
	if accessories.GlassesAsset != nil {
		if err := xml.Unmarshal(accessories.GlassesAsset.SVGContents, &glasses); err != nil {
			glog.Fatalf("Unable to parse glasses SVG %v", err)
			return nil, err
		}
	}
	if err := xml.Unmarshal(accessories.NoseAsset.SVGContents, &nose); err != nil {
		glog.Fatalf("Unable to parse nose SVG %v", err)
		return nil, err
	}
	if accessories.MouthAsset != nil {
		if err := xml.Unmarshal(accessories.MouthAsset.SVGContents, &mouth); err != nil {
			glog.Fatalf("Unable to parse mouth SVG %v", err)
			return nil, err
		}
	}
	if accessories.HatAsset != nil {
		if err := xml.Unmarshal(accessories.HatAsset.SVGContents, &hat); err != nil {
			glog.Fatalf("Unable to parse hat SVG %v", err)
			return nil, err
		}
	}
	if accessories.ShoeAsset == nil || !accessories.ShoeAsset.RemovesFeet {
		if err := xml.Unmarshal(accessories.FootLeftAsset.SVGContents, &footLeft); err != nil {
			glog.Fatalf("Unable to parse foot left SVG %v", err)
			return nil, err
		}
	}
	if accessories.ShoeAsset == nil || !accessories.ShoeAsset.RemovesFeet {
		if err := xml.Unmarshal(accessories.FootRightAsset.SVGContents, &footRight); err != nil {
			glog.Fatalf("Unable to parse foot right SVG %v", err)
			return nil, err
		}
	}
	if accessories.ShoeAsset != nil {
		if err := xml.Unmarshal(accessories.ShoeAsset.SVGContents, &shoes); err != nil {
			glog.Fatalf("Unable to parse shoes SVG %v", err)
			return nil, err
		}
	}
	if accessories.MiscAsset == nil || !accessories.MiscAsset.RemovesHandsLeft {
		if err := xml.Unmarshal(accessories.HandLeftAsset.SVGContents, &handLeft); err != nil {
			glog.Fatalf("Unable to parse hand left SVG %v", err)
			return nil, err
		}
	}
	if accessories.MiscAsset == nil || !accessories.MiscAsset.RemovesHandsRight {
		if err := xml.Unmarshal(accessories.HandRightAsset.SVGContents, &handRight); err != nil {
			glog.Fatalf("Unable to parse hand right SVG %v", err)
			return nil, err
		}
	}

	// Create new SVG writer for final assembly
	var b bytes.Buffer
	canvas := svg.New(&b)
	canvas.Startraw(fmt.Sprintf("viewBox=\"0 0 %d %d\"", DefaultSize, DefaultSize))

	// Add tail
	canvas.Group(fmt.Sprintf("id=\"%s\"", "tail"), "fill=\"none\"")
	if clr, ok := accessories.AccessoryColors[accessories.TailAsset.FileName]; ok {
		furReplacer := strings.NewReplacer("#7f6145", clr.ToHTML(true), "#7F6145", clr.ToHTML(true))
		tail.Doc = furReplacer.Replace(tail.Doc)
		// Shadow
		shadowOpacity := GetShadowOpacityFur(clr)
		shadowReplacer := strings.NewReplacer("fill-opacity=\".24\"", fmt.Sprintf("fill-opacity=\"%f\"", shadowOpacity), "fill-opacity=\"0.24\"", fmt.Sprintf("fill-opacity=\"%f\"", shadowOpacity))
		tail.Doc = shadowReplacer.Replace(tail.Doc)
	}
	io.WriteString(canvas.Writer, tail.Doc)
	canvas.Gend()

	// Add tail accessory if present
	if tailAccessory.Doc != "" {
		canvas.Group(fmt.Sprintf("id=\"%s\"", "tailAcc"), "fill=\"none\"")
		if clr, ok := accessories.AccessoryColors[accessories.TailAccessory.FileName]; ok {
			randReplacer := strings.NewReplacer("#62ffaa", clr.ToHTML(true), "#62FFAA", clr.ToHTML(true))
			tailAccessory.Doc = randReplacer.Replace(tailAccessory.Doc)
		}
		io.WriteString(canvas.Writer, tailAccessory.Doc)
		canvas.Gend()
	}

	// Add legs if not removed
	if legs.Doc != "" {
		canvas.Group(fmt.Sprintf("id=\"%s\"", "legs"), "fill=\"none\"")
		if clr, ok := accessories.AccessoryColors[accessories.LegAsset.FileName]; ok {
			furReplacer := strings.NewReplacer("#7f6145", clr.ToHTML(true), "#7F6145", clr.ToHTML(true))
			legs.Doc = furReplacer.Replace(legs.Doc)
			// Shadow
			shadowOpacity := GetShadowOpacityFur(clr)
			shadowReplacer := strings.NewReplacer("fill-opacity=\".14\"", fmt.Sprintf("fill-opacity=\"%f\"", shadowOpacity), "fill-opacity=\"0.14\"", fmt.Sprintf("fill-opacity=\"%f\"", shadowOpacity))
			legs.Doc = shadowReplacer.Replace(legs.Doc)
		}
		io.WriteString(canvas.Writer, legs.Doc)
		canvas.Gend()
	}

	// Add arms always
	canvas.Group(fmt.Sprintf("id=\"%s\"", "arms"), "fill=\"none\"")
	if clr, ok := accessories.AccessoryColors[accessories.ArmsAsset.FileName]; ok {
		furReplacer := strings.NewReplacer("#7f6145", clr.ToHTML(true), "#7F6145", clr.ToHTML(true))
		arms.Doc = furReplacer.Replace(arms.Doc)
		// Shadow
		shadowOpacity := GetShadowOpacityFur(clr)
		shadowReplacer := strings.NewReplacer("fill-opacity=\".14\"", fmt.Sprintf("fill-opacity=\"%f\"", shadowOpacity), "fill-opacity=\"0.14\"", fmt.Sprintf("fill-opacity=\"%f\"", shadowOpacity))
		arms.Doc = shadowReplacer.Replace(arms.Doc)

	}
	io.WriteString(canvas.Writer, arms.Doc)
	canvas.Gend()

	// Add bodyUpper always
	canvas.Group(fmt.Sprintf("id=\"%s\"", "bodyUpper"), "fill=\"none\"")
	if clr, ok := accessories.AccessoryColors[accessories.BodyUpperAsset.FileName]; ok {
		furReplacer := strings.NewReplacer("#7f6145", clr.ToHTML(true), "#7F6145", clr.ToHTML(true))
		bodyUpper.Doc = furReplacer.Replace(bodyUpper.Doc)
		// Shadow
		shadowOpacity := GetShadowOpacityFur(clr)
		shadowReplacer := strings.NewReplacer("fill-opacity=\".14\"", fmt.Sprintf("fill-opacity=\"%f\"", shadowOpacity), "fill-opacity=\"0.14\"", fmt.Sprintf("fill-opacity=\"%f\"", shadowOpacity))
		bodyUpper.Doc = shadowReplacer.Replace(bodyUpper.Doc)
	}
	io.WriteString(canvas.Writer, bodyUpper.Doc)
	canvas.Gend()

	// Add shirt pants
	if shirtPants.Doc != "" {
		canvas.Group(fmt.Sprintf("id=\"%s\"", "shirtPants"), "fill=\"none\"")
		if clr, ok := accessories.AccessoryColors[accessories.ShirtPantsAsset.FileName]; ok {
			randReplacer := strings.NewReplacer("#62ffaa", clr.ToHTML(true), "#62FFAA", clr.ToHTML(true))
			shirtPants.Doc = randReplacer.Replace(shirtPants.Doc)
		}
		io.WriteString(canvas.Writer, shirtPants.Doc)
		canvas.Gend()
	}

	// Add misc [above-shirt-pants]
	if misc.Doc != "" && accessories.MiscAsset.AboveShirtPants {
		canvas.Group(fmt.Sprintf("id=\"%s\"", "misc"), "fill=\"none\"")
		io.WriteString(canvas.Writer, misc.Doc)
		canvas.Gend()
	}

	// Add Ears always
	canvas.Group(fmt.Sprintf("id=\"%s\"", "ears"), "fill=\"none\"")
	io.WriteString(canvas.Writer, ears.Doc)
	canvas.Gend()

	// Add face always
	canvas.Group(fmt.Sprintf("id=\"%s\"", "face"), "fill=\"none\"")
	if clr, ok := accessories.AccessoryColors[accessories.FaceAsset.FileName]; ok {
		furReplacer := strings.NewReplacer("#7f6145", clr.ToHTML(true), "#7F6145", clr.ToHTML(true))
		face.Doc = furReplacer.Replace(face.Doc)
		// Shadow
		shadowOpacity := GetShadowOpacityFur(clr)
		shadowReplacer := strings.NewReplacer("fill-opacity=\".14\"", fmt.Sprintf("fill-opacity=\"%f\"", shadowOpacity), "fill-opacity=\"0.14\"", fmt.Sprintf("fill-opacity=\"%f\"", shadowOpacity))
		face.Doc = shadowReplacer.Replace(face.Doc)
	}
	io.WriteString(canvas.Writer, face.Doc)
	canvas.Gend()

	// Add Eyes if not removed
	if eyes.Doc != "" {
		canvas.Group(fmt.Sprintf("id=\"%s\"", "eyes"), "fill=\"none\"")
		if clr, ok := accessories.AccessoryColors[accessories.EyeAsset.FileName]; ok {
			eyeReplacer := strings.NewReplacer("#313cc4", clr.ToHTML(true), "#313CC4", clr.ToHTML(true))
			eyes.Doc = eyeReplacer.Replace(eyes.Doc)
			// Shadow
			shadowOpacity := GetShadowOpacityIris(clr)
			shadowReplacer := strings.NewReplacer("fill-opacity=\".13\"", fmt.Sprintf("fill-opacity=\"%f\"", shadowOpacity), "fill-opacity=\"0.13\"", fmt.Sprintf("fill-opacity=\"%f\"", shadowOpacity))
			eyes.Doc = shadowReplacer.Replace(eyes.Doc)
		}
		io.WriteString(canvas.Writer, eyes.Doc)
		canvas.Gend()
	}

	// Glasses
	if glasses.Doc != "" {
		canvas.Group(fmt.Sprintf("id=\"%s\"", "glasses"), "fill=\"none\"")
		io.WriteString(canvas.Writer, glasses.Doc)
		canvas.Gend()
	}

	// Nose
	canvas.Group(fmt.Sprintf("id=\"%s\"", "nose"), "fill=\"none\"")
	io.WriteString(canvas.Writer, nose.Doc)
	canvas.Gend()

	// Mouth
	if mouth.Doc != "" {
		canvas.Group(fmt.Sprintf("id=\"%s\"", "mouth"), "fill=\"none\"")
		io.WriteString(canvas.Writer, mouth.Doc)
		canvas.Gend()
	}

	// Hat
	if hat.Doc != "" {
		canvas.Group(fmt.Sprintf("id=\"%s\"", "hat"), "fill=\"none\"")
		if clr, ok := accessories.AccessoryColors[accessories.HatAsset.FileName]; ok {
			randReplacer := strings.NewReplacer("#62ffaa", clr.ToHTML(true), "#62FFAA", clr.ToHTML(true))
			hat.Doc = randReplacer.Replace(hat.Doc)
		}
		io.WriteString(canvas.Writer, hat.Doc)
		canvas.Gend()
	}

	// Foot left
	if footLeft.Doc != "" {
		canvas.Group(fmt.Sprintf("id=\"%s\"", "footLeft"), "fill=\"none\"")
		io.WriteString(canvas.Writer, footLeft.Doc)
		canvas.Gend()
	}

	// Foot right
	if footRight.Doc != "" {
		canvas.Group(fmt.Sprintf("id=\"%s\"", "footRight"), "fill=\"none\"")
		io.WriteString(canvas.Writer, footRight.Doc)
		canvas.Gend()
	}

	// Shoes
	if shoes.Doc != "" {
		canvas.Group(fmt.Sprintf("id=\"%s\"", "shoes"), "fill=\"none\"")
		if clr, ok := accessories.AccessoryColors[accessories.ShoeAsset.FileName]; ok {
			randReplacer := strings.NewReplacer("#62ffaa", clr.ToHTML(true), "#62FFAA", clr.ToHTML(true))
			shoes.Doc = randReplacer.Replace(shoes.Doc)
		}
		io.WriteString(canvas.Writer, shoes.Doc)
		canvas.Gend()
	}

	// Hand left
	if handLeft.Doc != "" {
		canvas.Group(fmt.Sprintf("id=\"%s\"", "handLeft"), "fill=\"none\"")
		io.WriteString(canvas.Writer, handLeft.Doc)
		canvas.Gend()
	}

	// Hand right
	if handRight.Doc != "" {
		canvas.Group(fmt.Sprintf("id=\"%s\"", "handRight"), "fill=\"none\"")
		io.WriteString(canvas.Writer, handRight.Doc)
		canvas.Gend()
	}

	// Add misc [above-hands]
	if misc.Doc != "" && accessories.MiscAsset.AboveHands {
		canvas.Group(fmt.Sprintf("id=\"%s\"", "misc"), "fill=\"none\"")
		io.WriteString(canvas.Writer, misc.Doc)
		canvas.Gend()
	}

	// End document
	canvas.End()

	// Minify
	var ret []byte
	ret, _ = getMinifier().minifier.Bytes("image/svg+xml", b.Bytes())
	return ret, nil
}

// Singleton to get minifier
type minifySingleton struct {
	minifier *minify.M
}

var mSingleton *minifySingleton
var onceM sync.Once

func getMinifier() *minifySingleton {
	onceM.Do(func() {
		minifier := minify.New()
		minifier.AddFunc("image/svg+xml", minifyxml.Minify)
		mSingleton = &minifySingleton{
			minifier: minifier,
		}
	})
	return mSingleton
}
