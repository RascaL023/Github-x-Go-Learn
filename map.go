package main

import "fmt"

func print_detail_map(m map[string]string){
	// fmt.Println(m)
	fmt.Println("Username   :", m["username"])
	fmt.Println("Email      :", m["email"])
	fmt.Print("Password   : ")
	for i := 0; i < len(m["password"]); i++ {
		fmt.Print("*")
	}
	fmt.Println("\nTotal data :", len(m))
}

func maps(){
	accounts := map[string]string {
		"username": "rascal",
		"email": "rascal@gmail.com",
		"password": "qwerty123",
		"temps": "temp",
	}
	delete(accounts, "temps")

	print_detail_map(accounts)
}

func main(){
	maps()
}
