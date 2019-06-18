package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
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
	//log.Print(config)
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
					fmt.Println("started again")
				} else {
					fmt.Println("not work")
				}
			} else {
				fmt.Println("working")
			}

			time.Sleep(15 * time.Second)
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

	//gopre.Pre(con.Program)

	cmd := exec.Command(con.Program.PROGRAM)
	cmd.Dir = con.Path.PATH

	_ , err := os.Stat(cmd.Dir)

	if err != nil{

		fmt.Println(err)
		log.Fatal("path hatasi")

	}


	if err := cmd.Start(); err != nil {
		fmt.Println("cmd.Start() failed", err)
		return false
	}

	return true

}
