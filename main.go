package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var config *Config

func main() {
	parent, _ := os.Getwd()
	log.Printf("exec in path: %s\n", parent)
	log.Printf("args is: %s\n", os.Args[1:])
	gitArgs := os.Args[1:]
	dirs, err := ioutil.ReadDir(parent)
	handleErr(err)
	for _, dir := range dirs {
		if dir.IsDir() {
			if !isIgnore(dir.Name()) {
				log.Printf("working dir: %s\n", dir.Name())
				cmd := exec.Command(config.BatchCmd, gitArgs...)
				cmd.Dir = parent + string(os.PathSeparator) + dir.Name()
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err = cmd.Run()
				if err != nil {
					log.Printf("cmd Error: %s\n", err)
					continue
				}
			}
		}
	}

}

func isIgnore(name string) (ignore bool) {
	ignore = false
	if config.IsReg {
		ignore = config.Exp.MatchString(name)
		return
	} else {
		for _, v := range config.IgnorePaths {
			if strings.EqualFold(v, name) {
				ignore = true
				return
			}
		}
	}
	return
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	c, err := os.Open("./config.json")
	defer c.Close()
	if os.IsNotExist(err) {
		log.Fatalf("config.json is not exist, please set up first")
	}
	handleErr(err)
	decoder := json.NewDecoder(c)
	err = decoder.Decode(&config)
	handleErr(err)
	if len(config.IgnorePaths) > 0 {
		config.IsReg = false
	}
	if config.RegExp != "" {
		config.IsReg = true
		config.Exp = regexp.MustCompile(config.RegExp)
	}
}
