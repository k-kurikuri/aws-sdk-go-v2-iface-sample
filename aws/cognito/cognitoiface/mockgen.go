//go:generate mockgen -source=cognitoiface.go -package=mock_$GOPACKAGE -destination=../../../mock/aws/cognito/mock_cognitoiface.go -self_package=iface-sample/aws/cognito/$GOPACKAGE

package cognitoiface
