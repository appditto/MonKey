package image

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"sync"

	svg "github.com/ajstarks/svgo"
	"github.com/golang/glog"
	minify "github.com/tdewolff/minify/v2"
	minifysvg "github.com/tdewolff/minify/v2/svg"
)

const DefaultSize = 4000           // Default SVG width/height attribute
const lodBwReplacement = "#9CA2AF" // Replace white with this color on bw assets

type SVG struct {
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
	Doc    string `xml:",innerxml"`
}

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
	if err := xml.Unmarshal(accessories.FootLeftAsset.SVGContents, &footLeft); err != nil {
		glog.Fatalf("Unable to parse foot left SVG %v", err)
		return nil, err
	}
	if err := xml.Unmarshal(accessories.FootRightAsset.SVGContents, &footRight); err != nil {
		glog.Fatalf("Unable to parse foot right SVG %v", err)
		return nil, err
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
	io.WriteString(canvas.Writer, tail.Doc)
	canvas.Gend()

	// Add tail accessory if present
	if tailAccessory.Doc != "" {
		canvas.Group(fmt.Sprintf("id=\"%s\"", "tailAcc"), "fill=\"none\"")
		io.WriteString(canvas.Writer, tailAccessory.Doc)
		canvas.Gend()
	}

	// Add legs if not removed
	if legs.Doc != "" {
		canvas.Group(fmt.Sprintf("id=\"%s\"", "legs"), "fill=\"none\"")
		io.WriteString(canvas.Writer, legs.Doc)
		canvas.Gend()
	}

	// Add arms always
	canvas.Group(fmt.Sprintf("id=\"%s\"", "arms"), "fill=\"none\"")
	io.WriteString(canvas.Writer, arms.Doc)
	canvas.Gend()

	// Add bodyUpper always
	canvas.Group(fmt.Sprintf("id=\"%s\"", "bodyUpper"), "fill=\"none\"")
	io.WriteString(canvas.Writer, bodyUpper.Doc)
	canvas.Gend()

	// Add shirt pants
	if shirtPants.Doc != "" {
		canvas.Group(fmt.Sprintf("id=\"%s\"", "shirtPants"), "fill=\"none\"")
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
	io.WriteString(canvas.Writer, face.Doc)
	canvas.Gend()

	// Add Eyes if not removed
	if eyes.Doc != "" {
		canvas.Group(fmt.Sprintf("id=\"%s\"", "eyes"), "fill=\"none\"")
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
	/*
		var ret []byte
		ret, _ = getMinifier().minifier.Bytes("image/svg+xml", b.Bytes())
	*/
	// TODO - return minified version
	return b.Bytes(), nil
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
		minifier.AddFunc("image/svg+xml", minifysvg.Minify)
		mSingleton = &minifySingleton{
			minifier: minifier,
		}
	})
	return mSingleton
}
