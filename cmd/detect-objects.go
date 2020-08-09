package cmd

import (
	"fmt"

	"image-ai/utils"

	"github.com/spf13/cobra"
)

var detectObjectsCmd = &cobra.Command{
	Use:   "detect-objects",
	Short: "Detects objects in an image",
	Long:  `Detects objects in an image`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get AWS Creds

		// Get image
		imagePath, _ := cmd.Flags().GetString("image-path")
		getFileResponse := utils.GetFile(imagePath)
		fmt.Printf("Get file response is %s", getFileResponse)

		// Call detect labels
		// Parse response
	},
}

func init() {
	RootCmd.AddCommand(detectObjectsCmd)
	detectObjectsCmd.Flags().StringP("image-path", "p", "", "Relative file path to image")
	detectObjectsCmd.MarkFlagRequired("image-path")
}
