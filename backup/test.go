package main

import (
    "fmt"
    "log"
    "os"

    "github.com/urfave/cli/v2"
)

func main() {
    tasks := []string{"cook", "clean", "laundry", "eat", "sleep", "code"}

    app := &cli.App{
        EnableBashCompletion: true,
        Commands: []*cli.Command{
            {
                Name:    "complete",
                Aliases: []string{"c"},
                Usage:   "complete a task on the list",
                Action: func(cCtx *cli.Context) error {
                    fmt.Println("completed task: ", cCtx.Args().First())
                    return nil
                },
                BashComplete: func(cCtx *cli.Context) {
                    // This will complete if no args are passed
                    if cCtx.NArg() > 0 {
                        return
                    }
                    for _, t := range tasks {
                        fmt.Println(t)
                    }
                },
            },
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
