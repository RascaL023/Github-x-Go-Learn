package main

import (
	"fmt"
	"os"

	"strings"
)

func systemEnvironment() {
	user := os.Getenv("USER");
	fmt.Println("User:", user);
}

func try() {
 	path := "data|output/$HOME/Documents/Kuliah";
	for i := 0; i < len(path); i++ {
		if path[i] != '$' {
			continue;
		}

		start := i + 1;
		end := start;

		for end < len(path) {

		}
	}
	// rep := "mechabox";
	// custom := "$waybarPreset";
}

func main(){
	try();
}
