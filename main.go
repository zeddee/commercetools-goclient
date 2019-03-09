package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	types "github.com/zeddee/commercetoolsGoClient/types"
)

func main() {

	var sunglasses types.ProductProjectionsResults
	sunglassesJSONbytes, err := loadFile("sunglass.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(sunglassesJSONbytes, &sunglasses); err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Results:\n%+v\n", sunglasses)

	fmt.Printf("Sunglasses: %+v\n", sunglasses.Results[0].Name["de"])

}

// loadJSON loads a json file into a jsonTarget, a target struct, and returns it
func loadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
