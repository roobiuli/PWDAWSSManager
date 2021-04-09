package AwsPwService

import (
	"github.com/roobiuli/PWDKMREP/Password"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestPasswdGather(t *testing.T)  {
	asserts := assert.New(t)

	Sm := &PWService{
		svc: &AwsSecretsManagerMock{},
	}

	PW := Password.Passwd{
		Name:  "TestPass",
	}

	data, err := Sm.GetPassword(PW)

	asserts.NoError(err, "No returned on Mock")

	//asserts.IsType(secretsmanager.GetSecretValueOutput{}, data, "Returned data should be GetSecretValueOutput")

	asserts.Equal("MockedPasswd", data, "ShouldBe the mocked password")

}