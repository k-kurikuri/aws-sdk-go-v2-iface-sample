package aws

import "aws-sdk-go-v2-iface-sample/aws/cognito/cognitoiface"

type AWS struct {
	CognitoIDProvider cognitoiface.CognitoIDProvider
}

func New(cognitoIDProvider cognitoiface.CognitoIDProvider) *AWS {
	return &AWS{
		CognitoIDProvider: cognitoIDProvider,
	}
}
