package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func main() {
	// Session.NewSession() automatically uses the default credential chain:
	// 1. Environment variables
	// 2. Shared credentials file
	// 3. IAM role (EC2, ECS, etc.)

	sessOptions := session.Options{
		SharedConfigState: session.SharedConfigEnable, // Enable ~/.aws/config parsing
	}

	// Only set profile if AWS_PROFILE is not already set
	if profile := os.Getenv("AWS_PROFILE"); profile == "" && len(os.Args) > 1 {
		sessOptions.Profile = os.Args[1]
	}

	// Create session with automatic credential resolution
	sess, err := session.NewSessionWithOptions(sessOptions)
	if err != nil {
		log.Fatalf("unable to create session: %v", err)
	}

	// Create STS client
	svc := sts.New(sess)

	// Get caller identity
	result, err := svc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalf("unable to get caller identity: %v", err)
	}

	fmt.Printf("Account: %s\n", *result.Account)
	fmt.Printf("UserId: %s\n", *result.UserId)
	fmt.Printf("Arn: %s\n", *result.Arn)
}
