package main

import "fmt"

type Record struct{
	Id int
	Name string
	Email string
	Height float32
	IsActive bool
}

func main(){
	rec1 := Record{
		Id: 1,
		Name: "Ra$caL",
		Email: "rascal@gmail.com",
		Height: 168.5,
		IsActive: true,
	}

	rec2 := Record{
		Id: 2,
		Name: "Asep",
		Email: "asep@gmail.com",
		Height: 170,
		IsActive: false,
	}

	// recs := []Record{rec1,rec2}
	// recs := make([]Record, 2, 5) => akan terisi 2 value kosong!!
	recs := make([]Record, 0, 5)
	recs = append(recs, rec1)
	recs = append(recs, rec2)
	print_all(recs)
}

func print_all(recs []Record) {
	fmt.Printf("%-5s %-10s %-20s %-8s %-8s\n",
			"ID", "Nama", "Email", "Height", "Active")
	fmt.Println("-----------------------------------------------------------")

	for _, r := range recs {
		r.printDetailRecord()
	}
	fmt.Println("")
}

func (r Record) printDetailRecord(){
	fmt.Printf("%-5d %-10s %-20s %-8.2f %-8t\n",
	r.Id, r.Name, r.Email, r.Height, r.IsActive)
}

// func print_detail(rec Record){
// 	fmt.Println("ID     :", rec.Id)
// 	fmt.Println("Nama   :", rec.Name)
// 	fmt.Println("Email  :", rec.Email)
// 	fmt.Println("Height :", rec.Height)
// 	fmt.Println("Status :", rec.IsActive)
// }
