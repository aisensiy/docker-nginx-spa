package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	prefix := os.Getenv("CONFIG_VARS")
	if prefix == "" {
		fmt.Println("{\n}")
		return
	}
	var items []string

	for _, pair := range os.Environ() {
		kv := strings.Split(pair, "=")
		if !strings.HasPrefix(kv[0], prefix) {
			continue
		}
		items = append(items, fmt.Sprintf("  \"%s\": \"%s\"", kv[0][len(prefix):], kv[1]))
	}
	fmt.Println("{")
	fmt.Println(strings.Join(items, ",\n"))
	fmt.Println("}")
}
