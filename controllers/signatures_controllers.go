package controllers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
)

type SignaturesControllers struct {
	beego.Controller
}
type Msg struct {
	Message string `json:"msg"`
}

var msgsign string
var gk *GKey

func (this *SignaturesControllers) Get() {

	this.TplName = "signature.html"

}
func (this *SignaturesControllers) Post() {
	var err error
	redata := this.Ctx.Input.RequestBody
	buf := make(map[string]interface{})
	err = json.Unmarshal(redata, &buf)
	if err != nil {
		fmt.Println("gk.sign err ...")
		return
	}
	if reqnum := buf["Reqid"].(float64); reqnum == 0 {
		if len(publicKey) == 0 || len(privateKey) == 0 {
			err, gk = Initkeys()
			if err != nil {
				fmt.Println("err in Initkey...")
				return
			}
		}
		pri := ByteToString(gk.GetPrivKey())
		pub := ByteToStringpub(gk.GetPubKey())
		this.Data["json"] = map[string]interface{}{"prikey": pri, "pubkey": pub}
		this.ServeJSON()
	} else {
		if reqnum := buf["Reqid"].(float64); reqnum == 1 {
			msg := buf["Message"].(string)
			msgsign, err = gk.Sign([]byte(msg))
			if err != nil {
				fmt.Println("gk.sign err ...")
				return
			}
			fmt.Println("msg = ", msg)
			fmt.Println("sign =", msgsign)
			this.Data["json"] = map[string]interface{}{"megsign": msgsign}
			this.ServeJSON()
		}

		if reqnum := buf["Reqid"].(float64); reqnum == 2 {
			//Verify
			msg := buf["Message"].(string)
			puk := buf["Puk"].(string)
			sign := buf["Signature"].(string)
			//Todo puk==publicKey
			if strings.Compare(puk, ByteToStringpub(gk.GetPubKey())) != 0 {
				this.Data["json"] = map[string]interface{}{"result": false}
				this.ServeJSON()
			}
			if strings.Compare(sign, msgsign) != 0 {
				this.Data["json"] = map[string]interface{}{"result": false}
				this.ServeJSON()
			}		
			flag, err := Verify([]byte(msg), sign, &gk.PublicKey)
			if err != nil {
				fmt.Println("err in Verify...", err)
				this.Data["json"] = map[string]interface{}{"result": false}
				this.ServeJSON()
				return
			}
			if flag == true {
				fmt.Println("flag = ", flag)
				this.Data["json"] = map[string]interface{}{"result": true}
				this.ServeJSON()
			}
			if flag == false {
				fmt.Println("flag = ", flag)
				this.Data["json"] = map[string]interface{}{"result": false}
				this.ServeJSON()
			}
		}
	}

}
