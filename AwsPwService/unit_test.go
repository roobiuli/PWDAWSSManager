package AwsPwService

import (
	"os"
	"testing"

	"github.com/roobiuli/PWDKMREP/Common"
	"github.com/roobiuli/PWDKMREP/Password"
	"github.com/stretchr/testify/assert"
)

func TestPasswdGather(t *testing.T) {
	asserts := assert.New(t)

	Sm := &PWService{
		svc: &AwsSecretsManagerMock{},
	}

	PW := Password.Passwd{
		Name: "TestPass",
	}

	data, err := Sm.GetPassword(PW)

	asserts.NoError(err, "No returned on Mock")

	//asserts.IsType(secretsmanager.GetSecretValueOutput{}, data, "Returned data should be GetSecretValueOutput")

	asserts.Equal([]byte("MockedPasswd"), data, "ShouldBe the mocked password")

}

func TestWithRealService(t *testing.T) {
	asserts := assert.New(t)

	// Seting Rand Password

	os.Setenv("PastauTest", "replaceme")

	// Test password

	EVars := Common.FindVars("replaceme")

	PasServ := NewPWService()

	PassP := &Common.PassProcessor{}

	PassP.WithPassColletion(EVars).WithSecretManager(PasServ)

	if !PassP.IsNext() {
		PassP.ProcessPassword()
	}

}
