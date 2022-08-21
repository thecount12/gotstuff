package main

import "testing"

// lastly coverage profile
// go test -coverprofile=coverage.out && go tool cover -html=coverage.out


// 6. table test
var tests = []struct {
	name string
	dividend float32
	divisor float32
	expected float32
	isErr bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"valid-data", 100.0, 0.0, 0, true},
	{"expected-5", 50.0, 10.0, 5.0, false},
	{"expected-fraction", -1.0, -777.0, 0.0012870013, false},
	
}

// 7. table func
func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but gone one", err.Error())
			}
		}
		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}


// 3. basic tests just require a function
// 4. now run 'go test' or 'go test -v'
func TestDivid(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

//5 . just make a copy and change name and make ==
func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Did not get an error when we should have")
	}
}
