package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)



type Config struct{
	ConnectEmail ConnectEmail
}


type ConnectEmail struct {
	Email string
	Password string
	Address string
}




func NewConfig() *Config{

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Не удалось загрузить конфиг")
	}

	return &Config{
		ConnectEmail: ConnectEmail{
			Email: os.Getenv("EMAIL"),
			Password: os.Getenv("PASSWORD"),
			Address: os.Getenv("ADDRESS"),
		},
	}
}