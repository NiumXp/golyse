package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Golyse"
	app.Usage = "..."

	app.Commands = []cli.Command{
		{
			Name:   "start",
			Usage:  "Starts the server",
			Action: startCommand,
		},
		{
			Name:   "stop",
			Usage:  "Stops the server",
			Action: stopCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func startCommand(c *cli.Context) error {
	if _, err := http.Get("http://localhost:8080/running"); err == nil {
		return errors.New("A server is already running.")
	}

	// cmd := exec.Command("go", "run", "server.go")
	cmd := exec.Command("golyse-server")
	err := cmd.Start()

	if err == nil {
		log.Println("Server started.")
	}

	return err
}

func stopCommand(c *cli.Context) error {
	response, err := http.Get("http://localhost:8080/stop")

	if err != nil {
		log.Println("Server is not running.")
		return nil
	}

	if response.StatusCode != 200 {
		log.Fatal(err)
	}

	log.Println("Server stopped.")
	return nil
}
