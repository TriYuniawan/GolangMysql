package main

import (
	"fmt"
	"go-crud/controllers/pasiencontroller"
	"net/http"
)

func main() {
	fmt.Println("Server berjalan!")
	http.HandleFunc("/", pasiencontroller.Index)
	http.HandleFunc("/pasien", pasiencontroller.Index)
	http.HandleFunc("/pasien/index", pasiencontroller.Index)
	http.HandleFunc("/pasien/add", pasiencontroller.Add)
	http.HandleFunc("/pasien/edit", pasiencontroller.Edit)
	http.HandleFunc("/pasien/delete", pasiencontroller.Delete)

	http.ListenAndServe(":8080", nil)

}
