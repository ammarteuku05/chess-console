package main

import (
	"chess-console/api/server"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "chess-console",
		Short: "chess-console Application",
		Long:  "A chess console application implemented with Go",
	}

	// Add subcommands
	rootCmd.AddCommand(server.ServerCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
