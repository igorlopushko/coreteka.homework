package service

import (
	"fmt"

	"github.com/igorlopushko/coreteka.homework/app/model"
	"github.com/igorlopushko/coreteka.homework/app/pkg/color"
	"github.com/sirupsen/logrus"
)

type GameService struct {
	Board      *model.Board
	IsGameOver bool
	boardSvc   BoardService
	isInit     bool
}

func NewGameService(b *model.Board) *GameService {
	svc := NewBoardService(b)
	return &GameService{
		Board:    b,
		boardSvc: svc,
	}
}

func (g *GameService) Start() error {
	err := g.boardSvc.Init()
	if err != nil {
		return err
	}

	g.isInit = true
	return nil
}

func (g *GameService) MakeStep(x, y int) (isAlreadyOpened bool, err error) {
	// check if the game has been started
	if !g.isInit {
		err = fmt.Errorf("Start() method of the GameService has to be called before make a step")
		logrus.Error(err)
		return false, err
	}

	// check if the cell is already opened
	if g.Board.Cells[y][x].Visibility == model.Opened {
		return true, nil
	}

	// open the cell
	trapped, err := g.boardSvc.OpenCell(x, y)
	if err != nil {
		g.IsGameOver = trapped
		return false, err
	}

	return false, nil
}

func (g *GameService) CheckIfWin() bool {
	opened := 0
	for _, r := range g.Board.Cells {
		for _, v := range r {
			if v.Visibility == model.Opened && v.Type != model.Hole {
				opened++
			}
		}
	}

	return g.Board.BlackHolesCount+opened == g.Board.Width*g.Board.Height
}

func (g *GameService) PrintBoard() {
	fmt.Print("  ")

	for i := 0; i < g.Board.Width; i++ {
		fmt.Print(i)
	}

	fmt.Println()
	fmt.Print("  ")
	for i := 0; i < g.Board.Width; i++ {
		fmt.Print("_")
	}

	fmt.Println()

	for y, r := range g.Board.Cells {
		fmt.Print(fmt.Sprint(y) + "|")
		for x := 0; x < len(r); x++ {
			if g.Board.Cells[y][x].Visibility == model.Closed {
				fmt.Print(".")
			} else {
				if g.Board.Cells[y][x].Type == model.Hole {
					fmt.Print(color.Red + "x" + color.Reset)
				} else {
					fmt.Print(g.Board.Cells[y][x].Value)
				}
			}
		}
		fmt.Println()
	}
}
