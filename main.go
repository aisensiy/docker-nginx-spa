package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func get_default_config(filepath string) map[string]interface{} {
	config := make(map[string]interface{})
	dat, err := ioutil.ReadFile(filepath)
	check(err)

	jsondata := string(dat)
	start := strings.IndexAny(jsondata, "{")

	if start == -1 {
		panic("can not find { in file")
	}

	if err := json.Unmarshal([]byte(jsondata[start:]), &config); err != nil {
		return make(map[string]interface{})
	} else {
		return config
	}
}

func main() {
	prefix := os.Getenv("CONFIG_VARS")
	if prefix == "" {
		fmt.Println("{\n}")
		return
	}

	var items map[string]interface{}

	if len(os.Args) > 1 {
		items = get_default_config(os.Args[1])
	} else {
		items = make(map[string]interface{})
	}

	for _, pair := range os.Environ() {
		kv := strings.Split(pair, "=")
		if !strings.HasPrefix(kv[0], prefix) {
			continue
		}
		items[kv[0][len(prefix):]] = kv[1]
	}
	result, _ := json.MarshalIndent(items, "", "  ")
	fmt.Println(string(result))
}
