package service

import (
	"testing"

	"github.com/igorlopushko/coreteka.homework/app/model"
)

func TestInit_ProvideInputParams_ReturnsCorrectBoard(t *testing.T) {
	w, h, c := 10, 10, 5
	b := model.NewBoard(w, h, c)
	svc := NewBoardService(b)

	err := svc.Init()
	if err != nil {
		t.Errorf("Init() method does not have to return an error")
	}

	valueCellsCount, holeCellsCount := 0, 0
	for _, r := range svc.Board.Cells {
		for _, v := range r {
			if v.Type == model.Empty || v.Type == model.Number {
				valueCellsCount++
			} else if v.Type == model.Hole {
				holeCellsCount++
			}
		}
	}

	if holeCellsCount != c {
		t.Errorf("Init() method initializes board with wrong number of holes: '%d'", holeCellsCount)
	}

	if valueCellsCount != w*h-c {
		t.Errorf("Init() method initializes board with wrong number of value cells: '%d'", valueCellsCount)
	}
}

func TestOpenCell_BoardIsNotInit_ReturnsError(t *testing.T) {
	w, h, c := 10, 10, 5
	b := model.NewBoard(w, h, c)
	svc := NewBoardService(b)

	_, err := svc.OpenCell(1, 1)

	if err == nil {
		t.Errorf("OpenCell() method has to return an error with not initialized board")
	}
}

func TestOpenCell_IncorrectCoordinates_ReturnsError(t *testing.T) {
	w, h, c := 10, 10, 5
	b := model.NewBoard(w, h, c)
	svc := NewBoardService(b)

	err := svc.Init()
	if err != nil {
		t.Errorf("Init() method does not have to return an error")
	}

	_, err = svc.OpenCell(11, 11)

	if err == nil {
		t.Errorf("OpenCell() method has to return an error with incorrect coordinates")
	}
}

func TestOpenCell_OpenHoleCell_ReturnsError(t *testing.T) {
	w, h, c := 10, 10, 5
	b := model.NewBoard(w, h, c)

	svc := NewBoardService(b)
	err := svc.Init()
	if err != nil {
		t.Errorf("Init() method does not have to return an error")
	}

	b.Cells[5][5].Type = model.Hole

	ex, err := svc.OpenCell(5, 5)

	if !ex || err == nil {
		t.Errorf("OpenCell() method has to return an error when open hole cell")
	}
}

func TestOpenCell_OpenEmptyCell_OpensCell(t *testing.T) {
	w, h, c := 10, 10, 5
	b := model.NewBoard(w, h, c)

	svc := NewBoardService(b)
	err := svc.Init()
	if err != nil {
		t.Errorf("Init() method does not have to return an error")
	}

	var cellX, cellY int
	var isFound bool
	for y, r := range svc.Board.Cells {
		for x, v := range r {
			if v.Type == model.Empty && v.Visibility == model.Closed {
				cellX = x
				isFound = true
				break
			}
		}
		if isFound {
			cellY = y
			break
		}
	}

	ex, err := svc.OpenCell(cellX, cellY)

	if ex || err != nil {
		t.Errorf("OpenCell() method does not have to return an error when open empty cell")
	}
}
