package main

import "fmt"

type User struct {
	name string
	age int
}

func (u User) printDetail(){
	fmt.Println("Name :", u.name);
	fmt.Println("Age", u.age);
}

func (u *User) changeName(new_name string) {
	u.name = new_name;
}

func (u *User) birthday() {
	u.age++;
}

func main() {
	user := User {
		name: "Asep",
		age: 20,
	}

	user.printDetail();

	printUserDetail := user.printDetail;

	user.changeName("Dimas");
	user.birthday();

	printUserDetail();

	printUserDetail = func() {
		user.printDetail()
	}
	printUserDetail();
}
