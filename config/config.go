package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type IConfiguration interface {
	Get(name string) string
	Init()
}

type config struct {
	loaded bool
}

var Config *config = func() *config {
	c := &config{}
	c.Init()
	fmt.Println("Initializing config")
	return c
}()

func (c *config) Get(name string) string {

	c.Init()
	val, exists := os.LookupEnv(name)
	if exists != true {
		fmt.Println("Error loading config value", name)
		return val
	}
	return val
}

func (c *config) Init() {
	//load the .env file

	//NOTE : IF TESTING THIS, MAKE SURE TO ADD THE .env file in the test folder
	godotenv.Load()
	c.loaded = true
}

func InitConfig() IConfiguration {
	c := &config{}
	c.Init()
	return c

}

// func (c *config) GetInt(name string) int {
// 	return 0
// }
