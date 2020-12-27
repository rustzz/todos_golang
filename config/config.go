package config

import (
	"fmt"
	"os/user"

	"github.com/joho/godotenv"
)

func Load() {
	// setting environment
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	godotenv.Load(fmt.Sprintf("%s/.cenv", user.HomeDir))
	return
}
