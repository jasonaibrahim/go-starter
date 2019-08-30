package app

import (
	"encoding/gob"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Store *sessions.CookieStore
)

func Init() error {
	err := godotenv.Load()

	if err != nil {
		log.Print(err.Error())
		return err
	}

	Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	gob.Register(map[string]interface{}{})

	return nil
}
