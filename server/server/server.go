package server

import (
	"fmt"

	"glendza/login-code-challenge/config"
)

func Init() {
	config.Init()
	c := *config.GetConfig()
	r := NewRouter()
	r.Run(fmt.Sprintf("localhost:%d", c.Port))
}
