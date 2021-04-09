package AwsPwService

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
	"github.com/roobiuli/PWDKMREP/Password"
)

type PWService struct {
	//svc *secretsmanager.SecretsManager
	svc secretsmanageriface.SecretsManagerAPI
}

func NewPWService() *PWService {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}
	s := secretsmanager.New(sess)
	return &PWService{svc: s}
}

func (p *PWService) GetPassword(pa *Password.Passwd) (*string, error) {
	input := secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(pa.Name),
	}
	ov , err := p.svc.GetSecretValue(&input)
	if err != nil {
		return nil, err
	}

	return ov.SecretString, nil
}