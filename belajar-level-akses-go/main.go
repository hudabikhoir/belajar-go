package main

import (
	/* komoponen yang berada di package lain yang di import bisa dijadikan selevel dengan komponen package peng-import
	dengan menambahkan tandan titik (.), jika sebelumnya "belajar-level-akses-go/library" maka akan menjadi
	. "belajar-level-akses-go/library" */
	. "belajar-level-akses-go/library"
	"fmt"
	"reflect"
	"runtime"
)

//interface
func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}
func main() {
	/* pemanggilan function dari library jika sebelumnya library.SayHello(), setelah kita samakan levelnya dengan
	menambahkan tanda titik (.) pemanggilannya langsung SayHello */
	SayHello("Nur Huda")
	// pemanggilan struct
	var s1 = Student{"Nur Huda", 21}
	fmt.Println("Nama \t:", s1.Name)
	fmt.Println("Umur \t:", s1.Age)
	fmt.Println("----------------------------")
	/* sayHelloDude() dari partial.go bisa langsung dipanggil tanpa di export karena memiliki nama package yang sama */
	// sayHelloDude("Huda")
	var number = 3
	//reflection
	var reflectValue = reflect.ValueOf(number)
	// mengembalikan tipe reflect
	fmt.Println("Type \t:", reflectValue.Type())
	// mengembalikan value reflect
	fmt.Println("Value \t:", reflectValue.Interface())
	fmt.Println("----------------------------")
	// goroutine
	// digunakan untuk menentukan jumlah core yang diaktifkan untuk eksekusi program
	runtime.GOMAXPROCS(2)
	// setiap execute goroutine selalu ditandai dengan 'go'
	go print(5, "async")
	print(5, "sync")

	var input string
	fmt.Scanln(&input)
}
