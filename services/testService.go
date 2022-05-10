package services

import (
	"github.com/natecarson1995/go-microservice-template/repositories"
)

func GetTestData() string {
	return "Bar " + repositories.GetTestData()
}
