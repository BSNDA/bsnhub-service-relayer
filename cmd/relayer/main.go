package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the entry point
var (
	rootCmd = &cobra.Command{
		Use:   "relayer",
		Short: "Relayer command line interface",
	}
)

func main() {
	cobra.EnableCommandSorting = false

	rootCmd.AddCommand(StartCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
