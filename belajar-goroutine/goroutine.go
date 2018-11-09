package main

import (
	"fmt"
	"runtime"
	"time"
)

func print(mesaage string, limit int) {
	for i := 0; i < limit; i++ {
		fmt.Println((i + 1), mesaage)
	}
}
func main() {
	// digunakan untuk menentukan jumlah core yang diaktifkan untuk eksekusi program
	runtime.GOMAXPROCS(1)

	// setiap pemanggilan goroutine selalu diawali 'go'
	go print("yoi pa kabar?", 5)
	print("hallo", 5)

	// menampilkan waktu sekarang
	var time1 = time.Now()
	fmt.Println("sekarang jam :", time1)

	// merubah format waktu
	var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
	fmt.Println("formated :", time2)

	var input string
	fmt.Scanln(&input)
}
