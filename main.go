package main

import (
	"net/http"
	"fmt"
	"time"
	"os/exec"
	"os"
	"github.com/BurntSushi/toml"
	"log"
)

type status int

const (
	statusOk status = 200
)

type httpOparation struct {
	status status
}

type url struct {
	URL string
}

type program struct {
	PROGRAM string
}

type path struct {
	PATH string
}

type Config struct {
	Url     url
	Program program
	Path    path
}

var con Config

func ReadConfig() Config {
	var configfile = "./config.toml"

	if _, err := os.Stat(configfile); os.IsNotExist(err) {
		// path/to/whatever does not exist
		currentPath, _ := os.Getwd()
		log.Fatal("config.toml not found in ", currentPath)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	log.Print(config)
	return config
}

func main() {

	con = ReadConfig()

	check()

}

func check() {

	var httpR = httpOparation{}
	httpR.status = statusOk
	for {

		{
			request := httpR.httpRequest(con.Url.URL)

			
			fmt.Println(request)

			if !request {
				if startExe() {
					fmt.Println("Yeniden baslasiiiiiin")
				} else {
					fmt.Println("olmadi yar")
				}
			} else {
				fmt.Println("sorun yok")
			}

			time.Sleep(222 * time.Millisecond)
		}

	}
}

func (httpR httpOparation) httpRequest(url string) bool {

	response, err := http.Get(url)

	if err != nil {
		return false
	}

	if response.StatusCode == 200 {
		return true

	} else {
		return false
	}

}

func startExe() bool {

	cmd := exec.Command(con.Program.PROGRAM, ">", " &", " 1 &")
	cmd.Dir = con.Path.PATH

	if err := cmd.Start(); err != nil {
		fmt.Println("cmd.Start() failed", err)
		return false
	}
	return true

}
