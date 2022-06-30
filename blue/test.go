package main

import (
	"encoding/json"
	"github.com/chrisish71/go-bluesub/blue/model"
	"github.com/chrisish71/go-bluesub/blue/util"
	"log"
	"os"
)

func main() {

	var blueSubtitles []model.Subtitle
	format := "stl"
	data, _ := os.ReadFile("subtitles.json")
	err := json.Unmarshal(data, &blueSubtitles)
	if err != nil {
		log.Panicln(err)
		return
	}
	asticodeSubtitle := util.BlueModelToAsticodeModel(blueSubtitles, format)
	filename := "/Users/chris/Desktop/subtitles." + format
	if err := asticodeSubtitle.Write(filename); err != nil {
		log.Panicln(err)
		return
	}
}
