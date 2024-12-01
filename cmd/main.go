package main

import (
	"fmt"
	"log"
	"os"

	auth "ai-powered-study-planner-backend/cmd/cli"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "ai-study-planner-cli",
		Usage: "CLI for AI Study Planner",
		Action: func(*cli.Context) error {
			fmt.Println("Welcome to CLI")
			return nil
		},
		Commands: []*cli.Command{
			&auth.GetJWTTokenCommand,
		},
		Before: func(ctx *cli.Context) error {
			err := godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file")
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
