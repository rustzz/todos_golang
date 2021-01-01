package config

import (
	"fmt"
	"os/user"

	"github.com/joho/godotenv"
)

// Load : ...
func Load() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	godotenv.Load(fmt.Sprintf("%s/.cenv", user.HomeDir))
	return
}
