package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// memastikan angka random benar-benar unik
	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 8)
	fmt.Println(randomValue)

	var nama = []string{"Nur", "Huda"}
	printNama("Halo,", nama)

	pembagian(4, 2)
	pembagian(0, 2)

	var luas, keliling float64 = lingkaran(5)
	fmt.Printf("Luas lingkaran: %.2f\n", luas)
	fmt.Printf("Keliling lingkaran: %.2f\n", keliling)

	var rata_rata float64 = rataRata(3, 5, 6, 3, 2, 7, 8, 9, 7, 10, 5, 6)
	fmt.Printf("Rata-rata: %.2f\n", rata_rata)

	yourHobbies("Nur Huda", "Design", "code")

	var numbers = []int{3, 5, 6, 3, 1, 6, 7, 8, 9}

	//anonymous function
	var getMinMax = func(n []int) (min int, max int) {
		for i, e := range n {
			switch {
			case i == 0:
				min, max = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return
	}
	var min, max = getMinMax(numbers)
	fmt.Printf("data: %v\nmax: %v\nmin: %v\n", numbers, max, min)

	/*
		Immediately-Invoked Function Expression IIFE
		function di eksekusi langsung pada saat deklarasinya
	*/
	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(7)
	fmt.Println("original number: ", numbers)
	fmt.Println("filtered number: ", newNumbers)
	// function return closure
	var maxNum = 5
	var howMany, getNumber = findMax(numbers, maxNum)
	var theNum = getNumber()
	fmt.Println("number \t:", numbers)
	fmt.Println("find \t:", max)
	fmt.Println("found \t:", howMany)
	fmt.Printf("value \t: %v\n", theNum)

	// function sebagai parameter
	var data = []string{"Nur", "huda", "bikhoir", "james", "bond"}
	var dataContain = filter(data, func(each string) bool {
		return strings.Contains(each, "a")
	})
	var dataLen = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("filter ada huruf a \t:", dataContain)
	fmt.Println("filter ada 5 huruf \t:", dataLen)
}
func printNama(pesan string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(pesan, nameString)
}

// type return harus dideklariskan setelah nama function
func randomWithRange(min, max int) int {
	var acak = rand.Int()%(min-max+1) + min
	return acak
}

func pembagian(x, y int) {
	if x == 0 || y == 0 {
		fmt.Printf("invalid divider. %d cannot divided %d\n", x, y)
		return
	}
	var res = x / y
	fmt.Printf("%d / %d = %d\n", x, y, res)
}

func lingkaran(d float64) (luas float64, keliling float64) {
	//hitung luas
	luas = math.Pi * math.Pow(d/2, 2)
	// hitung keliing
	keliling = math.Pi * d

	//kembalikan 2 nilai
	return
}
func rataRata(angka ...int) (rata_rata float64) {
	var total int
	for _, ang := range angka {
		total += ang
	}

	rata_rata = float64(total) / float64(len(angka))

	return
}
func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Nama: %s\n", name)
	fmt.Printf("Hobi: %s\n", hobbiesAsString)
}

// function closure sebagai nilai kembalian
func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}
func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, dt := range data {
		if filtered := callback(dt); filtered {
			result = append(result, dt)
		}
	}
	return result
}
