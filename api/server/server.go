package server

import (
	"bufio"
	"chess-console/di"
	"chess-console/pkg/shared/utils"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

var (
	port string
)

var ServerCmd = &cobra.Command{
	Use:   "api",
	Short: "Start the API server",
	Long:  "Start the chess-console API server with Echo framework",
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func init() {
	ServerCmd.Flags().StringVarP(&port, "port", "p", "", "Port to run the server on (overrides config)")
}

func startServer() {
	// Initialize di
	container := di.SetUp()

	// Initialize Echo (if we were running a web server, but here we run a console loop)
	e := echo.New()
	e.Validator = container.Validator

	fmt.Println("Chess Console Started!")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		container.Games.Print()
		fmt.Println("Turn:", container.Games.GetTurn())
		fmt.Print("Enter move (e.g., a2 a3): ")

		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		if input == "exit" || input == "quit" {
			break
		}

		sr, sc, er, ec, err := utils.ParseInput(input)
		if err != nil {
			container.Logger.Error("Invalid input: " + err.Error())
			continue
		}

		err = container.Games.Move(sr, sc, er, ec, container.Games.GetTurn())
		if err != nil {
			container.Logger.Error("Invalid move: " + err.Error())
			continue
		}

		if container.Games.IsGameOver() {
			container.Games.Print()
			container.Logger.Info("Game Over! King captured.")
			break
		}

		container.Games.SwitchTurn()
	}
}
