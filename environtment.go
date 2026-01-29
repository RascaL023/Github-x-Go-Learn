package main

import (
	"fmt"
	"os"
)

func systemEnvironment() {
	user := os.Getenv("USER");
	fmt.Println("User:", user);
}

func try(path string) string {
	return os.ExpandEnv(path);
}

func main(){
	fmt.Println(try("$HOME/Documents"));
	fmt.Println(try("$MYENV"));
}
