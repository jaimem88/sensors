package sensors

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// ProcessData loads the inputFile into memory and does the processing of all records in the file
// Risk: if the input file is too big, this can cause fatal error depending on the machine where it runs
func ProcessData(inputFile, outputFile string) error {

	var inputData InputData
	err := loadInputFile(inputFile, &inputData)
	if err != nil {
		log.WithError(err).Errorln("sensors: failed to load input file")
		return err
	}
	log.Infof("Processing %d records", len(inputData))

	data := map[string][]float64{}
	for _, temp := range inputData {
		data[temp.ID] = append(data[temp.ID], temp.Temperature)
	}

	results := []*Result{}
	for id, temperatures := range data {
		result := &Result{
			ID:      id,
			Average: average(temperatures),
			Median:  median(temperatures),
			Mode:    mode(temperatures),
		}
		results = append(results, result)
	}

	jsonResults, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		log.WithError(err).Errorln("sensors: failed to convert result to JSON")
		return err
	}

	if err := writeOutputFile(outputFile, jsonResults); err != nil {
		log.WithError(err).Errorln("sensors: failed to write result to file")
		return err
	}
	log.Infof("Results: \n%s\n", jsonResults)
	return nil
}
