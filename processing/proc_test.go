package processing

import (
	"disp_bot/processing/analyzer"
	"disp_bot/processing/parser"
	"disp_bot/utils"
	"reflect"
	"testing"
)

func TestProc_Processing(t *testing.T) {
	type fields struct {
		parser   *parser.Parser
		analyzer *analyzer.Analyzer
	}
	type args struct {
		data *utils.UnProcData
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *utils.ProcData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Proc{
				parser:   tt.fields.parser,
				analyzer: tt.fields.analyzer,
			}
			if got := p.Processing(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Processing() = %v, want %v", got, tt.want)
			}
		})
	}
}
