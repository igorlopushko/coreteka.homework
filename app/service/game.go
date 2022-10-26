// Package service determines domain services behavior.
package service

import (
	"fmt"

	"github.com/igorlopushko/coreteka.homework/app/model"
	"github.com/igorlopushko/coreteka.homework/app/pkg/color"
	"github.com/sirupsen/logrus"
)

// A GameService defines a domain service to proceed with the game.
type GameService struct {
	Board      *model.Board
	IsGameOver bool
	boardSvc   BoardService
	isInit     bool
}

// Creates a new instance of the GameService.
func NewGameService(b *model.Board) *GameService {
	svc := NewBoardService(b)
	return &GameService{
		Board:    b,
		boardSvc: svc,
	}
}

// Starts the game. Has to be called before making a step in the game.
func (g *GameService) Start() error {
	err := g.boardSvc.Init()
	if err != nil {
		return err
	}

	g.isInit = true
	return nil
}

// Makes a step in the game by specifying coordinates.
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

// Checks if the game is finished and the player has won.
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

// Outputs a game board to the console.
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
