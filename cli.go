package main

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

type Product interface {
}

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
		{
			Name:   "new",
			Usage:  "Creates a new product",
			Action: newCommand,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "url",
					Usage:    "URL of the product",
					Required: true,
				},
				cli.IntFlag{
					Name:  "delay",
					Usage: "Delay (in minutes) to ping the product",
					Value: 15,
				},
			},
		},
		{
			Name:   "list",
			Usage:  "List all registred products",
			Action: listCommand,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "vendor",
					Usage: "Filter the products by vendor",
				},
			},
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

func newCommand(c *cli.Context) error {
	u, err := url.ParseRequestURI(c.String("url"))
	if err != nil {
		return err
	}

	product, err := getProductDetail(u)
	if err != nil {
		return err
	}

	if err := saveProduct(product); err != nil {
		return err
	}

	log.Println("Product created.")

	return nil
}

func listCommand(c *cli.Context) error {
	return nil
}

func getProductDetail(u *url.URL) (Product, error)
func saveProduct(p Product) error
