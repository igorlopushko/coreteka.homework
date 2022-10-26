package service

import (
	"fmt"

	"github.com/igorlopushko/coreteka.homework/app/model"
	"github.com/igorlopushko/coreteka.homework/app/pkg/color"
)

type Game struct {
	Board      *model.Board
	IsGameOver bool
	boardSvc   BoardService
}

func NewGame(b *model.Board) *Game {
	return &Game{
		Board:    b,
		boardSvc: BoardService{Board: b},
	}
}

func (g *Game) Start() error {
	err := g.boardSvc.Init(g.Board)
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) MakeStep(x, y int) (isAlreadyOpened bool, err error) {
	// check if the cell is already opened
	if g.Board.Cells[y][x].Visibility == model.Opened {
		return true, nil
	}

	// open the cell
	exploded, err := g.boardSvc.OpenCell(x, y)
	if err != nil {
		g.IsGameOver = exploded
		return false, err
	}

	return false, nil
}

func (g *Game) CheckIfWin() bool {
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

func (g *Game) PrintBoard() {
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
