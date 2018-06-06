package sensors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

func loadInputFile(location string, input interface{}) error {
	raw, err := ioutil.ReadFile(location)
	if err != nil {
		return err
	}
	err = json.Unmarshal(raw, &input)
	if err != nil {
		return err
	}
	return err
}

func writeOutputFile(outputFile string, data []byte) error {
	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	_, err = f.Write(data)

	return err

}

func round2decimals(i float64) float64 {

	t := fmt.Sprintf("%.2f", math.Round(i*100)/100) // round up to nearest 2 decimals
	r, _ := strconv.ParseFloat(t, 64)
	return r
}
