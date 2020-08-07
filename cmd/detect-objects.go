package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var detectObjectsCmd = &cobra.Command{
	Use:   "detect-objects",
	Short: "Detects objects in an image",
	Long: `Detects objects in an image`,
	Run: func(cmd *cobra.Command, args []string) {
		detectObjects()
	},
}

func detectObjects() {
	fmt.Println("Detect objects command")
}

func init() {
	RootCmd.AddCommand(detectObjectsCmd)
}
