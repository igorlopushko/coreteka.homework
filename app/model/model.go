package model

import (
	"errors"

	"github.com/sirupsen/logrus"
)

type Point struct {
	X int
	Y int
}

type CellVisibilityType int32
type CellType int32

const (
	Closed CellVisibilityType = 0
	Opened CellVisibilityType = 1

	Empty  CellType = 0
	Number CellType = 1
	Hole   CellType = 2
)

type Cell struct {
	Visibility CellVisibilityType
	Type       CellType
	Value      int
}

func (c *Cell) Open() error {
	c.Visibility = Opened

	if c.Type == Hole {
		err := errors.New("the bomb has been exploded")
		logrus.Info(err)
		return err
	}

	return nil
}

func NewCell(t CellType) (*Cell, error) {
	if t == Number {
		err := errors.New("number cell type could not be set on initialization")
		logrus.Error(err)
		return nil, err
	}
	return &Cell{Visibility: Closed, Type: t}, nil
}

type Board struct {
	Cells           [][]*Cell
	Width           int
	Height          int
	BlackHolesCount int
}

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
