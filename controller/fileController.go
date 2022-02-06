package controller

import (
	s "github.com/vanessatocasuche/apirestGo/service"
)

// Method that connects the route with the service

func GetFile() interface{} {
	return s.GetFile()
}
