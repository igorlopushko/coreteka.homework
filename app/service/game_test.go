package service

import (
	"testing"

	"github.com/igorlopushko/coreteka.homework/app/model"
)

func TestStart_CorrectBoard_StartsTheGame(t *testing.T) {
	w, h, c := 10, 10, 5
	b := model.NewBoard(w, h, c)
	svc := NewGameService(b)

	err := svc.Start()

	if err != nil {
		t.Errorf("Start() method does not have to return an error")
	}
}

func TestMakeStep_GameIsNotStarted_ReturnsError(t *testing.T) {
	w, h, c := 10, 10, 5
	b := model.NewBoard(w, h, c)
	svc := NewGameService(b)

	err := svc.Start()
	if err != nil {
		t.Errorf("Start() method does not have to return an error")
	}

	_, err = svc.MakeStep(5, 5)

	if err != nil {
		t.Errorf("MakeStep() method has to return an error if the game is not started")
	}
}

func TestMakeStep_AlreadyOpenedCell_ReturnsError(t *testing.T) {
	w, h, c := 10, 10, 5
	b := model.NewBoard(w, h, c)
	svc := NewGameService(b)

	err := svc.Start()
	if err != nil {
		t.Errorf("Start() method does not have to return an error")
	}

	b.Cells[5][5].Visibility = model.Opened

	isAlreadyOpened, _ := svc.MakeStep(5, 5)

	if !isAlreadyOpened {
		t.Errorf("MakeStep() method has to return an error when try to open already opened cell")
	}
}

func TestMakeStep_OpenHoleCell_ReturnsError(t *testing.T) {
	w, h, c := 10, 10, 5
	b := model.NewBoard(w, h, c)
	svc := NewGameService(b)

	err := svc.Start()
	if err != nil {
		t.Errorf("Start() method does not have to return an error")
	}

	cellX, cellY, isFound := 0, 0, false
	for y, r := range svc.Board.Cells {
		for x, v := range r {
			if v.Type == model.Hole {
				cellX = x
				cellY = y
				isFound = true
				break
			}
			if isFound {
				break
			}
		}
	}

	_, err = svc.MakeStep(cellX, cellY)

	if err == nil {
		t.Errorf("MakeStep() method has to return an error when try to open hole cell")
	}
}

func TestCheckIfWin_NotAllCellsAreOpened_ReturnsNotWin(t *testing.T) {
	w, h, c := 10, 10, 5
	b := model.NewBoard(w, h, c)
	svc := NewGameService(b)

	err := svc.Start()
	if err != nil {
		t.Errorf("Start() method does not have to return an error")
	}

	result := svc.CheckIfWin()

	if result {
		t.Errorf("CheckIfWin() method has to return false if not all empty/number cells are opened")
	}
}

func TestCheckIfWin_AllCellsAreOpened_ReturnsWin(t *testing.T) {
	w, h, c := 10, 10, 5
	b := model.NewBoard(w, h, c)
	svc := NewGameService(b)

	err := svc.Start()
	if err != nil {
		t.Errorf("Start() method does not have to return an error")
	}

	for _, r := range svc.Board.Cells {
		for _, v := range r {
			if v.Type != model.Hole {
				v.Visibility = model.Opened
			}
		}
	}

	result := svc.CheckIfWin()

	if !result {
		t.Errorf("CheckIfWin() method has to return true if all empty/number cells are opened")
	}
}
