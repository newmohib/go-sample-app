package main

import "testing"

var tests = []struct {
	name string
	dividend float32
	divisor float32
	expected float32
	isErr bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0, true},
	{"expect-5", 50.0, 10.0, 5, false},
	{"expect-fraction", -1.0, -777.0, 0.0012870013, false},
}

func TestDevision(t *testing.T)  {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil{
				t.Errorf("Name: %s, Expected an error but did not get one", tt.name)
			}
		}else{
			if err != nil {
				// t.Error("did not expect an error but got one", err.Error())
				t.Errorf("Name: %s, did not expect an error but got one %v", tt.name, err.Error())
			}
		}

		if got != tt.expected{

			t.Errorf("Name: %s, Expected %f but got %f", tt.name, tt.expected, got)
		}
	}
	
}

// func TestDivide(t *testing.T)  {
// 	_, err := divide(10.0, 1)
// 	if err != nil {
// 		t.Error("Got an error when we should not have")
// 	}
	
// }

// func TestBadDivide(t *testing.T)  {
// 	_, err := divide(10.0, 0)
// 	if err == nil {
// 		t.Error("Did not get an error when we should have")
// 	}
	
// }