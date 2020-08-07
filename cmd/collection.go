package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var collectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "Commands to manage collections",
	Long: `Commands to manage collections`,
}

var collectionListCmd = &cobra.Command{
	Use:   "list",
	Short: "List collections",
	Long: `List collections`,
	Run: func(cmd *cobra.Command, args []string) {
		collecitonList()
	},
}

func collecitonList() {
	fmt.Println("Collection list command")
}

var collectionCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a collection",
	Long: `Create a collection`,
	Run: func(cmd *cobra.Command, args []string) {
		collecitonCreate()
	},
}

func collecitonCreate() {
	fmt.Println("Collection create command")
}

var collectionDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a collection",
	Long: `Delete a collection`,
	Run: func(cmd *cobra.Command, args []string) {
		collectionDelete()
	},
}

func collectionDelete() {
	fmt.Println("Collection delete command")
}

var collectionAddImageCmd = &cobra.Command{
	Use:   "add-image",
	Short: "Add an image to a collection",
	Long: `Add an image to a collection`,
	Run: func(cmd *cobra.Command, args []string) {
		collectionAddImage()
	},
}

func collectionAddImage() {
	fmt.Println("Collection add image command")
}

var collectionDeleteImageCmd = &cobra.Command{
	Use:   "delete-image",
	Short: "Delete an image from a collection",
	Long: `Delete an image from a collection`,
	Run: func(cmd *cobra.Command, args []string) {
		collectionDeleteImage()
	},
}

func collectionDeleteImage() {
	fmt.Println("Collection delete image command")
}

func init() {
	RootCmd.AddCommand(collectionCmd)
	collectionCmd.AddCommand(collectionListCmd)
	collectionCmd.AddCommand(collectionCreateCmd)
	collectionCmd.AddCommand(collectionDeleteCmd)
	collectionCmd.AddCommand(collectionAddImageCmd)
	collectionCmd.AddCommand(collectionDeleteImageCmd)
}
