package main

import "regexp"

type Config struct {
	BatchCmd    string   `json:"batch_cmd"`
	IgnorePaths []string `json:"ignore_paths,omitempty"`
	RegExp      string   `json:"reg_exp,omitempty"`
	IsReg       bool
	Exp         *regexp.Regexp
}
