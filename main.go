package main

import (
	_ "blockchain/routers"
	"net/http"

	"github.com/astaxie/beego"
)
func main() {
	http.Handle("/static", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	beego.Run()
}
