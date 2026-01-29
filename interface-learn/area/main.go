package main

import "fmt"

type Calculate interface {
	Area() float64
}

type Square struct {
	sisi int
}

type Circle struct {
	r float64
}

type Triangle struct {
	a float64
	t float64
}

func (s Square) Area() float64 {
	return float64(s.sisi * s.sisi);
}

func (c Circle) Area() float64 {
	return c.r * c.r * 3.14;
}

func (t Triangle) Area() float64 {
	return t.a * t.t * 0.5;
}

func main(){
	var inp int;

	fmt.Println("Area Calculator");
	fmt.Println("1. Square");
	fmt.Println("2. Circle");
	fmt.Println("3. Triangle");
	fmt.Scan(&inp);

	var calc Calculate;

	switch inp {
	case 1: 
		var sisi int;
		fmt.Print("Input sisi: ");
		fmt.Scan(&sisi);
		calc = Square{sisi};
	case 2:
		var r float64;
		fmt.Print("Input jari - jari: ");
		fmt.Scan(&r);
		calc = Circle{r};
	case 3:
		var a, t float64;
		fmt.Print("Input alas: ");
		fmt.Scan(&a);
		fmt.Print("Input tinggi: ");
		fmt.Scan(&t);
		calc = Triangle{a, t};
	default: 
		calc = nil;
		fmt.Print("Invalid");
	}

	fmt.Println(calc.Area());
}
