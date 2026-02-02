package model

import (
	"github.com/camargogg/golang.git/src/configuration/logger"
	resterr "github.com/camargogg/golang.git/src/configuration/rest_err"
)

func (*UserDomain) CreateUser(user UserDomain) *resterr.RestErr {
	logger.Info("Init create UserModel")
	return nil
}