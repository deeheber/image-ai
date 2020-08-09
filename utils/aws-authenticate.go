package utils

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func GetAWSSession(profile string, region string) *session.Session {
	if profile == "" {
		profile = "default"
	}

	if region == "" {
		region = "us-west-2"
	}

	fmt.Printf("\n\nUsing AWS profile: %s and AWS region: %s \n\n", profile, region)

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewSharedCredentials("", profile),
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Verify creds were successfully retrieved
	// Might not be necessary???
	_, err2 := sess.Config.Credentials.Get()

	if err2 != nil {
		fmt.Printf("Unable to authenticate with AWS: %s \n", err2)
		os.Exit(1)
	}

	return sess
}
