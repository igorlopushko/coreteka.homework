package rand

import (
	"testing"
)

func TestGetRandomCoordinates_ProvideInputParams_ReturnsCorrectCoordinates(t *testing.T) {
	w, h := 10, 10
	x, y, err := GetRandomCoordinates(10, 10)

	if x < 0 || x > w || y < 0 || y > h {
		t.Errorf("GetRandomCoordinates() method has return values between 0 and width/height")
	}

	if err != nil {
		t.Errorf("GetRandomCoordinates() method does not have to return an error")
	}
}
