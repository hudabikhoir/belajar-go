package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "apa kabar")
}
func main() {
	var urlString = "http://depeloper.com:80/hello?name=john wick&age=27"
	var u, e = url.Parse(urlString)
	if e != nil {
		fmt.Println(e.Error)
		return
	}
	fmt.Printf("protocol: %s\n", u.Scheme)
	fmt.Printf("host \t: %s\n", u.Host)
	fmt.Printf("path \t: %s\n", u.Path)

	var nama = u.Query()["name"][0]
	var umur = u.Query()["age"][0]
	fmt.Printf("Nama \t: %s | umur \t: %s\n", nama, umur)
	// penentuan aksi ketka url tertentu diakses user
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hallo")
	})
	http.HandleFunc("/index", index)
	// templating
	http.HandleFunc("/template", func(w http.ResponseWriter, r *http.Request) {
		// data yang akan di parsing
		var data = map[string]string{
			"nama":  "Nur Huda Bikhoir",
			"pesan": "Selamat",
		}

		var t, err = template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		t.Execute(w, data)
	})
	fmt.Println("starting web server at http://localhost:8080/ ...")
	http.ListenAndServe(":8080", nil)
}
