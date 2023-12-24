package main

import (
	"aws-sdk-go-v2-iface-sample/aws"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func main() {
	awsClient := aws.New(cognitoidentityprovider.New(cognitoidentityprovider.Options{}))
	_, _ = awsClient.CognitoIDProvider.GetUser(context.Background(), &cognitoidentityprovider.GetUserInput{})
}
