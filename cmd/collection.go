package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var collectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "Commands to manage collections",
	Long:  `Commands to manage collections`,
}

var collectionListCmd = &cobra.Command{
	Use:   "list",
	Short: "List collections",
	Long:  `List collections`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Collection list command")
	},
}

var collectionCreateCmd = &cobra.Command{
	Use:   "create [collection name]",
	Short: "Create a collection",
	Long:  `Create a collection`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fmt.Printf("Collection create with collection name %s \n", name)
	},
}

var collectionDeleteCmd = &cobra.Command{
	Use:   "delete [collection name]",
	Short: "Delete a collection",
	Long:  `Delete a collection`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fmt.Printf("Collection delete with collection name %s \n", name)
	},
}

var collectionAddImageCmd = &cobra.Command{
	Use:   "add-image",
	Short: "Add an image to a collection",
	Long:  `Add an image to a collection`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Collection add image command")
	},
}

var collectionDeleteImageCmd = &cobra.Command{
	Use:   "delete-image",
	Short: "Delete an image from a collection",
	Long:  `Delete an image from a collection`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Collection delete image command")
	},
}

func init() {
	RootCmd.AddCommand(collectionCmd)
	collectionCmd.AddCommand(collectionListCmd)
	collectionCmd.AddCommand(collectionCreateCmd)
	collectionCmd.AddCommand(collectionDeleteCmd)
	collectionCmd.AddCommand(collectionAddImageCmd)
	collectionCmd.AddCommand(collectionDeleteImageCmd)
}
