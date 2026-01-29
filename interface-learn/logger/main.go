package main

import "fmt"

type Logger interface {
	Log(msg string)
}

type ConsoleLogger struct {}
type FileLogger struct {}

func (c ConsoleLogger) Log(msg string) {
	fmt.Println(msg);
}

func (f FileLogger) Log(msg string) {
	fmt.Println(msg);
}

func main(){
	
}
