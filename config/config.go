package config

import (
	"fmt"
	"log"
	"os/user"

	"github.com/joho/godotenv"
)

// Load : ...
func Load() {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalln(err.Error())
	}
	godotenv.Load(fmt.Sprintf("%s/.cenv", currentUser.HomeDir))
	return
}
