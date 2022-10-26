// Package cmd is implemented to represent the command line tool to execute the program logic.
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/igorlopushko/coreteka.homework/app/config"
	"github.com/igorlopushko/coreteka.homework/app/model"
	"github.com/igorlopushko/coreteka.homework/app/pkg/color"
	"github.com/igorlopushko/coreteka.homework/app/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var fenvfile string

var Reset = "\033[0m"
var Red = "\033[31m"

var rootCmd = &cobra.Command{
	Use:   "go run main.go --env [path]",
	Short: "Coreteka: Proxx game",
	Long:  "Coreteka: Proxx game just to demonstrate Golang skills",
	RunE:  run,
	PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
		if fenvfile != "" {
			if err := godotenv.Load(fenvfile); err != nil {
				logrus.Error(fmt.Printf("failed to load envfile [%s]", fenvfile))
				return err
			}
		}
		return nil
	},
}

// Executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&fenvfile,
		"env",
		"e",
		"",
		"Path to env file")

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if fenvfile != "" {
		if err := godotenv.Load(fenvfile); err != nil {
			logrus.Warn(fmt.Printf("failed to load envfile [%s]", fenvfile))
		}
	}

	config.App.Load()

	logrus.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% - %msg%\n",
	})

	logLvl, err := logrus.ParseLevel(config.App.LogLevel)
	if err != nil {
		logrus.Warn("could not parse log level, using debug default")
		logLvl = logrus.DebugLevel
	}
	logrus.SetLevel(logLvl)
}

func run(cmd *cobra.Command, _ []string) error {
	defaultConfig := true
	w, h, c := config.App.Board.Width, config.App.Board.Height, config.App.Board.BlackHolesCount

	for {
		fmt.Println("Do you want to use default settings?")
		fmt.Printf("Height: '%d'; Width: '%d'; Black holes count: '%d'. (y/n):",
			config.App.Board.Width,
			config.App.Board.Height,
			config.App.Board.BlackHolesCount)

		r := bufio.NewReader(os.Stdin)
		txt, _ := r.ReadString('\n')
		txt = strings.Replace(txt, "\n", "", -1)
		if !strings.EqualFold(txt, "y") && !strings.EqualFold(txt, "n") {
			continue
		}

		switch strings.ToLower(txt) {
		case "y":
			defaultConfig = true
		case "n":
			defaultConfig = false
		}

		break
	}

	if !defaultConfig {
		w = getInputIntValue("\nPlease enter the board width", config.App.Board.WidthMin, config.App.Board.WidthMax)
		h = getInputIntValue("Please enter the board height", config.App.Board.HeightMin, config.App.Board.HeightMax)
		c = getInputIntValue("Please enter the black holes count", 1, w*h-1)

		fmt.Println("\nYour settings:")
		fmt.Printf("Height: '%d'; Width: '%d'; Black holes count: '%d'.\n", w, h, c)
	}

	b := model.NewBoard(w, h, c)
	g := service.NewGameService(b)

	err := g.Start()
	if err != nil {
		return err
	}

	g.PrintBoard()

	for {
		x := getInputIntValue("\nPlease enter X", 0, b.Width-1)
		y := getInputIntValue("\nPlease enter Y", 0, b.Height-1)

		isAlreadyOpened, err := g.MakeStep(x, y)

		if isAlreadyOpened {
			fmt.Println(color.Red + "The cell with specified coordinates is already opened." + color.Reset)
		}

		if err != nil {
			g.PrintBoard()
			if g.IsGameOver {
				fmt.Println(color.Red + "Game over!" + color.Reset)
			} else {
				fmt.Println(color.Red + "The game has been terminated! See logs for the details." + color.Reset)
			}
			break
		}

		g.PrintBoard()
		if g.CheckIfWin() {
			fmt.Println("You won!")
			return nil
		}
	}

	return nil
}

func getInputIntValue(m string, min, max int) int {
	for {
		// get the input from the user
		fmt.Print(m + ": ")
		r := bufio.NewReader(os.Stdin)
		txt, _ := r.ReadString('\n')
		txt = strings.Replace(txt, "\n", "", -1)

		// try to convert to int
		v, err := strconv.Atoi(txt)
		if err != nil {
			fmt.Println("Wrong input. Try again.")
			continue
		}

		// check if the number is in the range
		if v < min || v > max {
			fmt.Printf("Wrong input. Value is less than min value: '%d' or above its max value: '%d'.\n", min, max)
			continue
		}

		return v
	}
}
