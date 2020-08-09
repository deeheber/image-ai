package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var describeFaceCmd = &cobra.Command{
	Use:   "describe-face",
	Short: "Describes the largest fact in the image",
	Long:  `Detects the largest face in the image and returns a description of the face`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Describe face command")
	},
}

func init() {
	RootCmd.AddCommand(describeFaceCmd)
}
