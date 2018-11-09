package library

import (
	"fmt"
)

// pemberian level akses pada struct juga sama, ditandai dengan nama awal function huruf besar atau kecil
type Student struct {
	Name string
	Age  int
}

// level akses function SayHello() adalah public, ditandai dengan nama awal function huruf besar
func SayHello(name string) {
	fmt.Println("Halo")
	/* karena kita tidak bisa akses function introduce secara langsung, maka pemanggilan bisa dilakukan
	di function SayHello() yang bersifat public */
	introduce(name)
}

// sedangkan function introduce() level akses nya private, ditandai dengan nama awal function menggunakan huruf kecil
func introduce(name string) {
	fmt.Println("Perkenalkan nama saya", name)
}
