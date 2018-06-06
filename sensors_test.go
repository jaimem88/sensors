package sensors

import (
	"io/ioutil"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestProcessData(t *testing.T) {
	logrus.SetOutput(ioutil.Discard)
	type args struct {
		inputFile  string
		outputFile string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"read input.json file and dump response to output_test.json",
			args{
				inputFile:  "input.json",
				outputFile: "output_test.json",
			},
			false,
		},
		{
			"error on empty input file",
			args{
				inputFile:  "",
				outputFile: "output_test.json",
			},
			true,
		},
		{
			"error on empty output file",
			args{
				inputFile:  "input.json",
				outputFile: "",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ProcessData(tt.args.inputFile, tt.args.outputFile); (err != nil) != tt.wantErr {
				t.Errorf("ProcessData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
