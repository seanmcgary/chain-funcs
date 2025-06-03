package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	// Load embedded configuration
	config, err := LoadEmbeddedConfig()
	if err != nil {
		log.Fatalf("Failed to load embedded config: %v", err)
	}

	app := &cli.App{
		Name:  "faas-cli",
		Usage: "CLI for interacting with the FaaS AVS contract",
		Commands: []*cli.Command{
			{
				Name:  "deploy-function",
				Usage: "Deploy a JavaScript function to the FaaS contract",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "rpc-url",
						Usage:    "Ethereum RPC URL",
						Value:    config.GetDefaultRPCURL(),
						EnvVars:  []string{"RPC_URL"},
					},
					&cli.StringFlag{
						Name:     "private-key",
						Usage:    "Private key for signing transactions",
						Required: true,
						EnvVars:  []string{"PRIVATE_KEY"},
					},
					&cli.StringFlag{
						Name:     "faas-address",
						Usage:    "FaaS contract address",
						Value:    config.GetFaaSAddress(),
						EnvVars:  []string{"FAAS_ADDRESS"},
					},
				},
				ArgsUsage: "<function-directory>",
				Action:    deployFunction,
			},
			{
				Name:  "call-function",
				Usage: "Call a deployed JavaScript function",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "rpc-url",
						Usage:    "Ethereum RPC URL",
						Value:    config.GetDefaultRPCURL(),
						EnvVars:  []string{"RPC_URL"},
					},
					&cli.StringFlag{
						Name:     "private-key",
						Usage:    "Private key for signing transactions",
						Required: true,
						EnvVars:  []string{"PRIVATE_KEY"},
					},
					&cli.StringFlag{
						Name:     "faas-address",
						Usage:    "FaaS contract address",
						Value:    config.GetFaaSAddress(),
						EnvVars:  []string{"FAAS_ADDRESS"},
					},
					&cli.StringFlag{
						Name:     "taskMailbox-address",
						Usage:    "TaskMailbox contract address",
						Value:    config.GetTaskMailboxAddress(),
						EnvVars:  []string{"TASKMAILBOX_ADDRESS"},
					},
					&cli.StringFlag{
						Name:     "function-id",
						Usage:    "Function ID (32-byte hex string)",
						Value:    config.GetFaaSAddress(),
					},
				},
				ArgsUsage: "<input-string>",
				Action:    callFunction,
			},
			{
				Name:  "config",
				Usage: "Show embedded contract configuration",
				Action: func(c *cli.Context) error {
					fmt.Printf("Embedded Contract Configuration:\n\n")
					fmt.Printf("RPC URL: %s\n", config.GetDefaultRPCURL())
					fmt.Printf("TaskMailbox Address: %s\n", config.GetTaskMailboxAddress())
					fmt.Printf("FaaS Address: %s\n", config.GetFaaSAddress())
					fmt.Printf("\nNote: These addresses are embedded at compile time from devnet deployment.\n")
					fmt.Printf("Use flags or environment variables to override these defaults.\n")
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}