package main

import (
	//"fmt"
	"log"
	//"net/http"
	//"io/ioutil"
	//"encoding/json"
	//"time"  
	//"encoding/base64"
	"github.com/pelletier/go-toml"
)

func ErrorCheck(e error) {
    if e != nil {
		log.Fatalln(e)
        panic(e)
    }
}

func main() {
	config, _ := toml.LoadFile("config.toml")
	log.Println(config.Get("account.token").(string))
}