package service

import (
	"fmt"

	"github.com/igorlopushko/coreteka.homework/app/model"
	"github.com/igorlopushko/coreteka.homework/app/pkg/rand"
	"github.com/sirupsen/logrus"
)

// A BoardService defines a domain service to work with the game board.
type BoardService struct {
	Board  *model.Board
	isInit bool
}

// Creates a new instance of the BoardService.
func NewBoardService(b *model.Board) BoardService {
	return BoardService{Board: b}
}

// Initializes game board state. Has to be called before opening any cell.
func (s *BoardService) Init() error {
	err := generateBlackHoles(s.Board)
	if err != nil {
		return err
	}

	err = generateNumbers(s.Board)
	if err != nil {
		return err
	}

	s.isInit = true
	return nil
}

// Opens a cell with the specified coordinates.
func (s *BoardService) OpenCell(x, y int) (trapped bool, err error) {
	// check if the board is initialized before the usage
	if !s.isInit {
		err = fmt.Errorf("Init() method of the BoardService has to be called before open any cell")
		logrus.Error(err)
		return false, err
	}

	// check if the coordinates fit the board
	if x < 0 || y < 0 || x > s.Board.Width-1 || y > s.Board.Height-1 {
		err = fmt.Errorf("the cell is out of range of board coordinates (x:'%d' y:'%d')", x, y)
		logrus.Error(err)
		return false, err
	}

	// open the cell
	err = s.Board.Cells[y][x].Open()
	if err != nil {
		return true, err
	}

	// if the cell is empty apply the flood fill algorithm
	if s.Board.Cells[y][x].Type == model.Empty {
		err := floodFill(s.Board, x, y)
		if err != nil {
			return true, err
		}
	}

	return false, nil
}

func generateBlackHoles(b *model.Board) error {
	for i := 0; i < b.BlackHolesCount; i++ {
		for {
			x, y, err := rand.GetRandomCoordinates(b.Width, b.Height)
			if err != nil {
				return err
			}

			// try again if cell is already occupied
			if b.Cells[y][x] != nil {
				continue
			}

			c, err := model.NewCell(model.Hole)
			if err != nil {
				return err
			}
			b.Cells[y][x] = c
			break
		}
	}

	return nil
}

func generateNumbers(b *model.Board) error {
	// set empty values for not used cells
	for _, r := range b.Cells {
		for i := 0; i < len(r); i++ {
			if r[i] == nil {
				c, err := model.NewCell(model.Empty)
				if err != nil {
					return err
				}
				r[i] = c
			}
		}
	}

	// update numbers according to holes location
	for y, r := range b.Cells {
		for x, v := range r {
			if v.Type == model.Hole {
				innerX, innerY, xLength, yLength := getInnerSquare(x, y, b.Width, b.Height)
				yCount := 0

				for iY := innerY; yCount < yLength; iY++ {
					var xCount = 0

					for iX := innerX; xCount < xLength; iX++ {
						xCount++
						if b.Cells[iY][iX].Type == model.Hole {
							continue
						}
						b.Cells[iY][iX].Type = model.Number
						b.Cells[iY][iX].Value++
					}
					yCount++
				}
			}
		}
	}

	return nil
}

func floodFill(b *model.Board, x, y int) error {
	var stack []model.Point
	stack = append(stack, model.Point{X: x, Y: y})

	for len(stack) > 0 {
		n := len(stack) - 1
		p := stack[n]
		stack = stack[:n]

		innerX, innerY, xLength, yLength := getInnerSquare(p.X, p.Y, b.Width, b.Height)
		yCount := 0

		for iY := innerY; yCount < yLength; iY++ {
			xCount := 0

			for iX := innerX; xCount < xLength; iX++ {
				xCount++

				// skip cell if it is a hole or the same cell as the caller
				if b.Cells[iY][iX].Type == model.Hole ||
					iX == p.X && iY == p.Y {
					continue
				}

				// append cell to stack to process it further if it is an empty cell
				if b.Cells[iY][iX].Type == model.Empty &&
					b.Cells[iY][iX].Visibility == model.Closed {
					stack = append(stack, model.Point{X: iX, Y: iY})
				}

				// open cell if it is empty or number
				if (b.Cells[iY][iX].Type == model.Number || b.Cells[iY][iX].Type == model.Empty) &&
					b.Cells[iY][iX].Visibility == model.Closed {
					err := b.Cells[iY][iX].Open()
					if err != nil {
						return err
					}
				}
			}

			yCount++
		}
	}

	return nil
}

func getInnerSquare(x, y, width, height int) (innerX, innerY, xLength, yLength int) {
	if x == 0 {
		innerX = 0
	} else {
		innerX = x - 1
	}

	if y == 0 {
		innerY = 0
	} else {
		innerY = y - 1
	}

	xLength = 3
	yLength = 3

	if x-1 < 0 || x+xLength-1 > width {
		xLength = 2
	}

	if y-1 < 0 || y+yLength-1 > height {
		yLength = 2
	}

	return innerX, innerY, xLength, yLength
}
