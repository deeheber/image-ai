package cmd

import (
	"fmt"
	"os"

	"image-ai/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/spf13/cobra"
)

var detectObjectsCmd = &cobra.Command{
	Use:   "detect-objects",
	Short: "Detects objects in an image",
	Long:  `Detects objects in an image`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// Get image
		imagePath, _ := cmd.Flags().GetString("image-path")
		fileBlob := utils.GetFile(imagePath)

		// Authenticate with AWS
		awsRegion, _ := cmd.Flags().GetString("aws-region")
		awsProfile, _ := cmd.Flags().GetString("aws-profile")
		sess := utils.GetAWSSession(awsProfile, awsRegion)

		// Call detect labels
		svc := rekognition.New(sess)
		input := &rekognition.DetectLabelsInput{
			Image: &rekognition.Image{
				Bytes: fileBlob,
			},
			MaxLabels:     aws.Int64(5),
			MinConfidence: aws.Float64(70.000000),
		}
		result, err := svc.DetectLabels(input)

		// Parse response
		// TODO handle errors better
		/*
			https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/handling-errors.html
			https://docs.aws.amazon.com/rekognition/latest/dg/error-handling.html#error-handling.MessagesAndCodes
			https://docs.aws.amazon.com/sdk-for-go/api/aws/awserr/#Error
		*/
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Labels found for image: %s\n\n", imagePath)

		for _, label := range result.Labels {
			fmt.Println("Name:", aws.StringValue(label.Name))
			fmt.Printf("Confidence:%f\n", *label.Confidence)
			fmt.Println("----------------------")
	}
	},
}

func init() {
	RootCmd.AddCommand(detectObjectsCmd)
	detectObjectsCmd.Flags().StringP("image-path", "p", "", "Relative file path to image")
	detectObjectsCmd.MarkFlagRequired("image-path")
}
