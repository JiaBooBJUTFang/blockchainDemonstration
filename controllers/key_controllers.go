package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type KeysControllers struct {
	beego.Controller
}

func (this *KeysControllers) Keyget() {
	this.TplName = "keys.html"
}

func (this *KeysControllers) MyPost() {
	readdata := this.Ctx.Input.RequestBody
	buf := make(map[string]interface{})
	err := json.Unmarshal(readdata, &buf)
	if err != nil {
		fmt.Println("json unmarshal err...", err)
		return
	}
	if req := buf["Reqnum"].(float64); req == 0 {
		if len(publicKey) == 0 || len(privateKey) == 0 {
			_, gk = Initkeys()
		}
		pri := ByteToString(gk.GetPrivKey())
		pub := ByteToStringpub(gk.GetPubKey())
		this.Data["json"] = map[string]interface{}{"prikey": pri, "pubkey": pub}
		this.ServeJSON()
		return
	}
	if req := buf["Reqnum"].(float64); req == 1 {
		_, gk = Initkeys()
		pri := ByteToString(gk.GetPrivKey())
		pub := ByteToStringpub(gk.GetPubKey())
		this.Data["json"] = map[string]interface{}{"prikey": pri, "pubkey": pub}
		this.ServeJSON()
		return
	}

}
