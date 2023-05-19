package parser

import (
	"reflect"
	"testing"
)

func TestHaul(t *testing.T) {
	tests := []struct {
		name             string
		args             []string
		wantCars         map[string]ServiceID
		wantNotProcessed []string
	}{
		{
			name: "1",
			args: []string{
				"1 \n 2 \n 3 \n 4 \n",
				"1 \n 2 \n 3 \n",
				"1 \n 2",
			},
			wantCars: map[string]ServiceID{
				"": "",
			},
			wantNotProcessed: []string{
				"",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCars, gotNotProcessed := Haul(tt.args)
			if !reflect.DeepEqual(gotCars, tt.wantCars) {
				t.Errorf("Haul() gotCars = %v, want %v", gotCars, tt.wantCars)
			}
			if !reflect.DeepEqual(gotNotProcessed, tt.wantNotProcessed) {
				t.Errorf("Haul() gotNotProcessed = %v, want %v", gotNotProcessed, tt.wantNotProcessed)
			}
		})
	}
}

func Test_getCarNumAndServ(t *testing.T) {
	tests := []struct {
		name       string
		mes        string
		wantCarNum string
		wantServ   string
		wantErr    bool
	}{
		{
			"",
			"url\ngrz\n47\nцветок\nмусор",
			"grz",
			"47",
			false,
		},
		{
			"",
			"url\ngrz\n47",
			"grz",
			"47",
			false,
		},
		{
			"",
			"url\ngrz47",
			"grz47",
			"grz47",
			false,
		},
		{
			"",
			"url",
			"",
			"",
			true,
		},
		{
			"",
			"",
			"",
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCarNum, gotServ, err := getCarNumAndServ(tt.mes)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCarNumAndServ() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCarNum != tt.wantCarNum {
				t.Errorf("getCarNumAndServ() gotCarNum = %v, want %v", gotCarNum, tt.wantCarNum)
			}
			if gotServ != tt.wantServ {
				t.Errorf("getCarNumAndServ() gotServ = %v, want %v", gotServ, tt.wantServ)
			}
		})
	}
}
