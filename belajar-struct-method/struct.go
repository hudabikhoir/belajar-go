package main

import (
	"fmt"
	"strings"
)

// struct parent
type student struct {
	name    string
	grade   int
	address string
	parent
}

// struct child
type parent struct {
	address string
	father  string
	mother  string
}

/*
	pendeklarasian method hampir sama seperti function, hanya saja kita perlu
	mendeklarasikan struct. Struct yang digunakan akan menjadi pemilik method.
	pada code berikut method sayHello() milik struct student. parameter pertama
	sebagai aliasing.
*/
func (s student) sayHello() {
	fmt.Println("halo,", s.name)
}
func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}
func main() {
	// struct
	var p1 = parent{"Nganjuk", "Bapak", "Ibu"}
	var s1 = student{"Nur huda", 1, "Kertosono", p1}
	// anonymous struct
	var p2 = parent{"Nganjuk", "Bapak2", "Ibu2"}
	var s2 = struct {
		// deklarasi anonymous function
		student
		kelas string
	}{
		// inisialisasi anonymus function. jika embed dengan struct lain wajib menyertkan struct yang dituju
		student{"Asep", 2, "Kertosono", p2}, // embed struct student
		"XI-IPS2",
	}
	// kombinasi slice and struct
	var allParent = []parent{
		{"kertosono", "bapak1", "ibu1"},
		{"kertosono2", "bapak2", "ibu2"},
		{"kertosono3", "bapak3", "ibu3"},
	}
	fmt.Println("name s1 \t:", s1.name)
	fmt.Println("grade s1 \t:", s1.grade)
	fmt.Println("address s1 \t:", s1.address)
	// pemanggilan struct di dalam parent
	fmt.Println("father s1 \t:", s1.parent.father)
	fmt.Println("mother s1 \t:", s1.parent.mother)
	fmt.Println("address s1 \t:", s1.parent.address)
	fmt.Println("----------------------------------")
	fmt.Println("name s2 \t:", s2.name)
	fmt.Println("grade s2 \t:", s2.grade)
	fmt.Println("address s2 \t:", s2.address)
	fmt.Println("kelas s2 \t:", s2.kelas)
	// pemanggilan struct di dalam parent
	fmt.Println("father s2 \t:", s2.parent.father)
	fmt.Println("mother s2 \t:", s2.parent.mother)
	fmt.Println("address s2 \t:", s2.parent.address)
	fmt.Println("----------------------------------")
	// loop kombinasi struct dan slice
	for _, orang_tua := range allParent {
		fmt.Println("father \t:", orang_tua.father, "| mother \t:", orang_tua.mother, "| address \t:", orang_tua.address)
	}
	// penggunaan method
	s1.sayHello()
	// karena method getNameAt masih return. harus ditampung ke variable terlebih dahulu. untuk selanjutnya di print
	var name = s1.getNameAt(1)
	fmt.Println("nama panggilan \t:", name)
}
