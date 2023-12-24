package aws

import (
	"aws-sdk-go-v2-iface-sample/aws/cognito/cognitoiface"
	mock_iface "aws-sdk-go-v2-iface-sample/mock/aws/cognito"
	"go.uber.org/mock/gomock"
	"testing"
)

type mockCognitoIDProvider struct {
	cognitoiface.CognitoIDProvider
}

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCognitoIDProvider := mock_iface.NewMockCognitoIDProvider(ctrl)
	_ = New(mockCognitoIDProvider)
}
