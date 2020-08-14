package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/appditto/monKey/server/image"
	"github.com/golang/glog"
)

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

func getAccessoryAsset(fname string, path string) string {
	var err error
	asset := image.Asset{}
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
	asset.ShadowEye = false
	asset.ColorableRandom = hasTag(asset.FileName, "[colorable-random]")
	asset.RemovesEyes = hasTag(asset.FileName, "[removes-eyes]")
	asset.RemovesLegs = hasTag(asset.FileName, "[removes-legs]")
	asset.RemovesFeet = hasTag(asset.FileName, "[removes-feet]")
	asset.RemovesHandsLeft = hasTag(asset.FileName, "[removes-hands-left]") || hasTag(asset.FileName, "[removes-hands]")
	asset.RemovesHandsRight = hasTag(asset.FileName, "[removes-hands-right]") || hasTag(asset.FileName, "[removes-hands]")
	asset.AboveShirtPants = hasTag(asset.FileName, "[above-shirts-pants]")
	asset.AboveHands = hasTag(asset.FileName, "[above-hands]")
	asset.Weight = getWeight(asset.FileName)
	encoded, _ := json.Marshal(asset)
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprint(encoded), "[", "{"), "]", "}"), " ", ", ") + ","
}

func LoadAssetsToArray() {
	wd, err := os.Getwd()
	if err != nil {
		panic("Can't get working directory")
	}

	ret := "package image\n\n"

	var bodyAsset image.Asset
	ret += "var BodyPartsIllustrations = [][]byte{\n"
	fPath := path.Join(wd, "assets", "illustrations", string(image.BodyPart))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			bodyAsset = image.Asset{}
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
			bodyAsset.ShadowEye = hasTag(bodyAsset.FileName, "[shadow-eye]")
			bodyAsset.ColorableRandom = false
			bodyAsset.RemovesEyes = false
			bodyAsset.RemovesLegs = false
			bodyAsset.RemovesHandsLeft = false
			bodyAsset.RemovesHandsRight = false
			bodyAsset.AboveShirtPants = false
			bodyAsset.AboveHands = false
			bodyAsset.Weight = getWeight(bodyAsset.FileName)
			encoded, _ := json.Marshal(bodyAsset)
			ret += strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprint(encoded), "[", "{"), "]", "}"), " ", ", ") + ","
		}
		return nil
	})
	ret += "}\n"

	ret += "var HatIllustrations = [][]byte{\n"
	fPath = path.Join(wd, "assets", "illustrations", "accessories", string(image.Hats))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			ret += getAccessoryAsset(info.Name(), path)
		}
		return nil
	})
	ret += "}\n"

	ret += "var GlassesIllustrations = [][]byte{\n"
	fPath = path.Join(wd, "assets", "illustrations", "accessories", string(image.Glasses))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			ret += getAccessoryAsset(info.Name(), path)
		}
		return nil
	})
	ret += "}\n"

	ret += "var MiscIllustrations = [][]byte{\n"
	fPath = path.Join(wd, "assets", "illustrations", "accessories", string(image.Misc))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			ret += getAccessoryAsset(info.Name(), path)
		}
		return nil
	})
	ret += "}\n"

	ret += "var MouthsIllustrations = [][]byte{\n"
	fPath = path.Join(wd, "assets", "illustrations", "accessories", string(image.Mouths))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			ret += getAccessoryAsset(info.Name(), path)
		}
		return nil
	})
	ret += "}\n"

	ret += "var ShirtPantsIllustrations = [][]byte{\n"
	fPath = path.Join(wd, "assets", "illustrations", "accessories", string(image.ShirtPants))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			ret += getAccessoryAsset(info.Name(), path)
		}
		return nil
	})
	ret += "}\n"

	ret += "var ShoesIllustrations = [][]byte{\n"
	fPath = path.Join(wd, "assets", "illustrations", "accessories", string(image.Shoes))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			ret += getAccessoryAsset(info.Name(), path)
		}
		return nil
	})
	ret += "}\n"

	ret += "var TailsIllustrations = [][]byte{\n"
	fPath = path.Join(wd, "assets", "illustrations", "accessories", string(image.Tails))
	err = filepath.Walk(fPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".svg") {
			ret += getAccessoryAsset(info.Name(), path)
		}
		return nil
	})
	ret += "}"

	output := path.Join(wd, "image", "illustrations.go")
	outputF, err := os.Create(output)
	if err != nil {
		fmt.Printf("Failed to open file for writing %s", output)
	}
	defer outputF.Close()
	outputF.WriteString(ret)
}
