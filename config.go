package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

var CONFIG_FILE = func() string {
	dir, _ := os.UserConfigDir()
	return path.Join(dir, ".translate-env")
}()

type Config struct {
	Url  string `json:"url"`
	Key  string `json:"key"`
	Lang string `json:"lang"`
}

var CONFIG Config

func init() {
	if !fileExists(CONFIG_FILE) {
		os.MkdirAll(path.Dir(CONFIG_FILE), 0755)
		os.WriteFile(CONFIG_FILE, []byte("{}"), 0644)
	} else {
		data, _ := os.ReadFile(CONFIG_FILE)
		json.Unmarshal(data, &CONFIG)
	}

	if CONFIG.Url == "" {
		CONFIG.Url = "https://libretranslate.de"
	}
	if CONFIG.Lang == "" {
		CONFIG.Lang = "en"
	}
}

func SetConfig(args []string) {
	for _, rawArg := range args {
		arg := strings.Split(rawArg, "=")
		if len(arg) != 2 {
			fmt.Printf("Invalid argument: %s\n", rawArg)
			return
		}
		key, value := arg[0], arg[1]
		switch key {
		case "url":
			CONFIG.Url = value
		case "key":
			CONFIG.Key = value
		case "lang":
			CONFIG.Lang = value
		default:
			fmt.Printf("Invalid argument: %s\n", rawArg)
			return
		}
	}
	data, _ := json.Marshal(CONFIG)
	os.WriteFile(CONFIG_FILE, data, 0644)
}
