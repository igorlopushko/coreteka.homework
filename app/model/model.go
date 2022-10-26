// Package model is implemented to represent domain models.
package model

import (
	"errors"

	"github.com/sirupsen/logrus"
)

// A Point is a struct to represent point coordinates.
type Point struct {
	X int
	Y int
}

// A CellVisibilityType represents the visibility of the cell on the board.
type CellVisibilityType int32

// A CellType represents cell type.
type CellType int32

const (
	Closed CellVisibilityType = 0
	Opened CellVisibilityType = 1

	Empty  CellType = 0
	Number CellType = 1
	Hole   CellType = 2
)

// A Cell represents a cell object on the board.
type Cell struct {
	Visibility CellVisibilityType
	Type       CellType
	Value      int
}

// Creates a new instance of the Cell object.
func NewCell(t CellType) (*Cell, error) {
	if t == Number {
		err := errors.New("number cell type could not be set on initialization")
		logrus.Error(err)
		return nil, err
	}
	return &Cell{Visibility: Closed, Type: t}, nil
}

// Opens the specified cell on the board.
func (c *Cell) Open() error {
	c.Visibility = Opened

	if c.Type == Hole {
		err := errors.New("you trapped to the black hole")
		logrus.Info(err)
		return err
	}

	return nil
}

// A Board represents a gaming board with cells.
type Board struct {
	Cells           [][]*Cell
	Width           int
	Height          int
	BlackHolesCount int
}

// Creates a new instance of the Board object.
func NewBoard(width, height, blackHolesCount int) *Board {
	// create 2d slice of cells
	cells := make([][]*Cell, height)
	for i := 0; i < height; i++ {
		cells[i] = make([]*Cell, width)
	}

	return &Board{
		Cells:           cells,
		Height:          height,
		Width:           width,
		BlackHolesCount: blackHolesCount,
	}
}
