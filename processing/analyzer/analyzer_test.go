package analyzer

import (
	"disp_bot/utils"
	"reflect"
	"testing"
)

func TestAnalyzer_Analyze(t *testing.T) {
	type args struct {
		parsedData utils.ParsedData
	}
	tests := []struct {
		name string
		args args
		want utils.ProcData
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Analyzer{}
			if got := a.Analyze(tt.args.parsedData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Analyze() = %v, want %v", got, tt.want)
			}
		})
	}
}
