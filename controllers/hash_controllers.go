package controllers

import (
	"github.com/astaxie/beego"
)

type HashControler struct {
	beego.Controller
}

func (this *HashControler) Hashget() {
	this.TplName = "hash.html"
}
