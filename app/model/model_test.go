package model

import (
	"testing"
)

func TestCellOpen_TryToOpenHoleCell_ReturnsError(t *testing.T) {
	c := Cell{Type: Hole, Value: 0, Visibility: Closed}

	err := c.Open()

	if err == nil {
		t.Errorf("Open() method has to return an error when opening hole cell")
	}
}

func TestCellOpen_TryToOpenEmptyCell_OpensCell(t *testing.T) {
	c := Cell{Type: Empty, Value: 0, Visibility: Closed}

	err := c.Open()

	if err != nil {
		t.Errorf("Open() method does not have to return an error when opening empty cell")
	}

	if c.Visibility != Opened {
		t.Errorf("Open() method has to set Visibility to Opened value")
	}
}
