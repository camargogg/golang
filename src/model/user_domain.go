package model

import resterr "github.com/camargogg/golang.git/src/configuration/rest_err"

type UserDomain struct {
	Email string
	Password string
	Name  string 
	Age int8
}

type UserDomainInterface interface {
	CreateUser(UserDomain) *resterr.RestErr
	UpdateUser(string, UserDomain) *resterr.RestErr
	FindUser(string) (*UserDomain, *resterr.RestErr)
	DeleteUser(userID string) *resterr.RestErr
}