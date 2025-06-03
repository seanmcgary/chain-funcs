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
				Name:  "register-function",
				Usage: "Register a JavaScript or Python function with content stored on-chain",
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
					&cli.BoolFlag{
						Name:    "skip-deps",
						Usage:   "Skip dependency installation for Python functions",
						EnvVars: []string{"SKIP_DEPS"},
					},
				},
				ArgsUsage: "<function-directory>",
				Action:    registerFunction,
			},
			{
				Name:  "deploy-function",
				Usage: "Deploy a JavaScript or Python function with content stored remotely in S3",
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
						Name:     "s3-base-url",
						Usage:    "S3 base URL for remote deployment (e.g., s3://bucket-name/path/)",
						Required: true,
						EnvVars:  []string{"S3_BASE_URL"},
					},
					&cli.BoolFlag{
						Name:    "skip-deps",
						Usage:   "Skip dependency installation for Python functions",
						EnvVars: []string{"SKIP_DEPS"},
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
						Name:     "ws-rpc-url",
						Usage:    "Ethereum WebSocket RPC URL for event subscriptions",
						Value:    config.GetDefaultWSRPCURL(),
						EnvVars:  []string{"WS_RPC_URL"},
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
						Required: true,
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
					fmt.Printf("WebSocket RPC URL: %s\n", config.GetDefaultWSRPCURL())
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