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
 	path := "$HOME/Documents/Kuliah";
	if strings.HasPrefix(path, "$") {
		trim := strings.TrimPrefix(path, "$");
		parts := strings.SplitN(trim, "/", 2);
		fmt.Println(parts);
	}

	// rep := "mechabox";
	// custom := "$waybarPreset";
}

func main(){
	try();
}
