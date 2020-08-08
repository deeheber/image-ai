package cmd

import (
	"fmt"

	"image-ai/utils"
	"github.com/spf13/cobra"
)

var detectObjectsCmd = &cobra.Command{
	Use:   "detect-objects",
	Short: "Detects objects in an image",
	Long: `Detects objects in an image`,
	Run: func(cmd *cobra.Command, args []string) {
		imagePath, _ := cmd.Flags().GetString("image-path")

		fmt.Println("In detect objects command")
		getFileResponse := utils.GetFile(imagePath)

		fmt.Printf("Get file response is %s", getFileResponse)
	},
}

func init() {
	RootCmd.AddCommand(detectObjectsCmd)
	detectObjectsCmd.Flags().StringP("image-path", "p", "", "Relative file path to image")
	detectObjectsCmd.MarkFlagRequired("image-path")
}
