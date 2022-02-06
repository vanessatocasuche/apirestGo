package service

import (
	"github.com/vanessatocasuche/apirestGo/repository"
)

/**
Package available to verify that the Volume file is static
*/

// Method that connects the controller with the service

func GetFile() interface{} {
	return repository.GetFile_()
}
