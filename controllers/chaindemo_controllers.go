package controllers

import (
	"github.com/astaxie/beego"
)

//index
type IndexControllers struct {
	beego.Controller
}

func (this *IndexControllers) Get() {
	this.TplName = "chaindemo.html"
}

type Index2Controllers struct {
	beego.Controller
}

func (this *Index2Controllers) Get() {
	this.TplName = "chaindemo2.html"
}
