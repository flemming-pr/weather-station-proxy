package main

import (
	"testing"
)

func TestMphToKnots(t *testing.T) {
	t.Parallel()
    r := &RequestProbe{
        WindSpeed: 10.0,
        WindGust:  20.0,
    }

    r.mphToKnots()

    if r.WindSpeed != 10.0*0.868976 {
        t.Errorf("WindSpeed was incorrect, got: %f, want: %f.", r.WindSpeed, 10.0*0.868976)
    }

    if r.WindGust != 20.0*0.868976 {
        t.Errorf("WindGust was incorrect, got: %f, want: %f.", r.WindGust, 20.0*0.868976)
    }
}

func TestFahrenheitToCelcius(t *testing.T) {
	t.Parallel()
	r := &RequestProbe{
		Temperature: 10.0,
	}

	r.fahrenheitToCelcius()

	if r.Temperature != (10.0-32)*5/9 {
		t.Errorf("Temperature was incorrect, got: %f, want: %f.", r.Temperature, (10.0-32)*5/9)
	}
}
