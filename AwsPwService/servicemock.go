package AwsPwService

import (
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
)

type AwsSecretsManagerMock struct {
	secretsmanageriface.SecretsManagerAPI
}

func (a *AwsSecretsManagerMock) GetSecretValue(i *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error)  {
	dumypass := "MockedPasswd"
	return &secretsmanager.GetSecretValueOutput{
		SecretString: &dumypass,
	}, nil
}