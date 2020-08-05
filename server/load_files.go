package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/appditto/monKey/server/image"
	"github.com/golang/glog"
)

func LoadAssetsToArray() {
	wd, err := os.Getwd()
	if err != nil {
		panic("Can't get working directory")
	}

	ret := "package image\n\n"

	var bodyAsset image.Asset
	ret += "var BodyParts = [][]byte{\n"
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
			encoded, _ := json.Marshal(bodyAsset)
			ret += strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprint(encoded), "[", "{"), "]", "}"), " ", ", ") + ","
		}
		return nil
	})
	ret += "}\n"

	output := path.Join(wd, "image", "illustrations.go")
	outputF, err := os.Create(output)
	if err != nil {
		fmt.Printf("Failed to open file for writing %s", output)
	}
	defer outputF.Close()
	outputF.WriteString(ret)
}
