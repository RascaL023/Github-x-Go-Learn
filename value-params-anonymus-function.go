package main

import "fmt"

func filterAdmin(name string) string {
	if name == "RascaL" {
		return "Atmin"
	}

	return name
}

func helloTo(name string, filter func(string) string, blackListed func(string) bool) string {
	if(blackListed(name)) {
		return "You are banned"
	}
	return "Hello " + filter(name)
}

func main(){
	var name string
	fmt.Print("Input nama: ")
	fmt.Scan(&name)

	blackList := func(n string) bool {
		return n == "banned"
	}

	greet := helloTo
	fmt.Println(greet(name, filterAdmin, blackList))
}
