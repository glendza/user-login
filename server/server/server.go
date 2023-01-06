package server

import (
	"fmt"

	"glendza/login-code-challenge/config"
)

func Init() {
	config.Init()
	c := *config.GetConfig()
	r := NewRouter()
	r.Run(fmt.Sprintf("0.0.0.0:%d", c.Port))
}
