package main

import "fmt"

func main() {
	// variable _ digunakan untuk membuang var yang tidak terpakai
	_ = "belajar go"
	// penulisan multiple variable agar lebih ringkas
	first_name, _ := "Nur", "Huda"
	first_name = "Huda"
	// konstanta
	const last_name = "Huda"

	fmt.Println("hello", first_name, last_name)
	// fmt.Println(name)
	fmt.Println(last_name)

	//operator aritmatika
	var value = ((2 + 4) / 3)
	var isEqual = (value == 2)

	fmt.Println("nilai", value, isEqual)
	var hasil = ""
	if value == 10 {
		hasil = "sempurna"
	} else if value > 8 {
		hasil = "lulus"
	} else {
		hasil = "tidak lulus nilai"
	}
	fmt.Println(hasil, value)

	//perulangan
	for i := 0; i < 5; i++ {
		fmt.Println("angka", i)
	}

	//array
	var nama = [...]string{"huda", "nur"}
	//multidimensional array
	var multiArr = [2][2]string{
		[2]string{
			"huda",
			"nur",
		},
		[2]string{
			"bikhoir", "y",
		},
	}

	fmt.Println(nama)
	fmt.Println(multiArr)
	// _ untuk membuang key dari array
	for _, nm := range nama {
		fmt.Println("user =", nm)
	}
	//slice
	var namaSlc = []string{"nur", "huda", "bikhoir", "yoi"}
	// pemanggilam dimulai dari index ke-0, hingga elemen sebelum index ke-2
	fmt.Println(namaSlc[0:3])

	var aNamaSlc = namaSlc[0:3]
	var bNamaSlc = namaSlc[1:3]
	// function len() untuk menghitung jumlah slice yang ada
	fmt.Println("aNamaSlc len", len(aNamaSlc))
	fmt.Println("bNamaSlc len", len(bNamaSlc))
	// function cap()untuk menghitung lebar atau jumlah maksimum slice
	fmt.Println("aNamaSlc cap", cap(aNamaSlc))
	fmt.Println("bNamaSlc cap", cap(bNamaSlc))

	var chicken = map[string]int{
		"januari":  30,
		"februari": 80,
		"maret":    50,
		"april":    30,
		"mei":      80,
		"juni":     50,
	}
	for key, val := range chicken {
		fmt.Println(key, ":", val)
	}
	fmt.Println(chicken)
	// delete elemen chicken dengan index "juni"
	delete(chicken, "juni")
	fmt.Println(chicken)

	var chickenExist, isExist = chicken["mei"]

	if isExist {
		fmt.Println(chickenExist)
	} else {
		fmt.Println("Chicken does not exist")
	}
	// kombinasi slice dan
	var siswa = []map[string]string{
		{"nama": "Nur Huda", "jenis_kelamin": "L"},
		{"nama": "Bikhoir", "jenis_kelamin": "L"},
		{"nama": "Renata", "jenis_kelamin": "P"},
	}
	for _, murid := range siswa {
		fmt.Println("nama:", murid["nama"], "jenis kelamin:", murid["jenis_kelamin"])
	}
}
