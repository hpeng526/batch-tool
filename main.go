package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"path/filepath"
	"fmt"
)

var config *Config

func main() {
	parent, _ := os.Getwd()
	if len(os.Args) < 2 {
		fmt.Printf("Usage: \n\tbatch-tool [cmd] [args]\n")
		return
	}
	log.Printf("exec in path: %s\n", parent)
	log.Printf("name is: %s\n", os.Args[1])
	log.Printf("args is: %s\n", os.Args[2:])
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	dirs, err := ioutil.ReadDir(parent)
	handleErr(err)
	for _, dir := range dirs {
		if dir.IsDir() {
			if !isIgnore(dir.Name()) {
				log.Printf("working dir: %s\n", dir.Name())
				cmd := exec.Command(cmdName, cmdArgs...)
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
	f, _ := exec.LookPath(os.Args[0])
	cp, _ := filepath.Abs(f)
	configFile := filepath.Join(filepath.Dir(cp), "config.json")
	c, err := os.Open(configFile)
	defer c.Close()
	if os.IsNotExist(err) {
		log.Fatalf("%s is not exist, please set up first", configFile)
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
