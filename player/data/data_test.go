package data

import (
	"testing"
)

var d *Data

// Unit tests
// I would use testing tables so we can add many cases and range over them
// No time in 2 hours for it

// proccessInput and generateCSV functions are defined in main, should be in a helpers file
// and those functions should be tested too

// Data methods are dependendant, Decode will fail if the numeric base is not
// calculated previously, and so on

func TestCalcBase(t *testing.T) {
	d = &Data{
		PlayerName:  "P1",
		CodedBase:   "asdfghjklñ",
		CodedPoints: "ñ",
	}

	d.CalcBase()

	if d.NumericBase != len(d.CodedBase) {
		t.Errorf("Expected %v, got %v", len(d.CodedBase), d.NumericBase)
	}
}

func TestDecode(t *testing.T) {
	d = &Data{
		PlayerName:  "P2",
		CodedBase:   "ab",
		CodedPoints: "b",
	}

	d.CalcBase()
	d.Decode()

	if len(d.DecodedPoints) != len(d.CodedPoints) {
		t.Errorf("Expected number of decoded digits %v, got %v", len(d.DecodedPoints), len(d.CodedPoints))
	}
}

func TestToBase10(t *testing.T) {
	d = &Data{
		PlayerName:  "P1",
		CodedBase:   "asdfghjklñ",
		CodedPoints: "ñ",
	}

	d.CalcBase()
	d.Decode()
	d.ToBase10()

	// Harcodedd no time..
	if d.Points != 9 {
		t.Errorf("Expected points 9, got %v", d.Points)
	}
}
