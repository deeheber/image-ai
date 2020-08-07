package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "image-ai",
	Short: "A CLI tool for analyzing the contents of an image",
	Long: `A CLI tool for analyzing the contents of an image`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
