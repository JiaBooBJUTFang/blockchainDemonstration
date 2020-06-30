package controllers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
)
type TransactionControllers struct {
	beego.Controller
}

func (this *TransactionControllers) Get() {
	this.TplName = "transaction.html"
}
func (this *TransactionControllers) Post() {
	var err error
	buf := make(map[string]interface{})
	redata := this.Ctx.Input.RequestBody
	err = json.Unmarshal(redata, &buf)
	if err != nil {
		fmt.Println("Unmashal err  in *TransactionControllers ...")
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
		return
	}
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
		return
	}
	if reqnum := buf["Reqid"].(float64); reqnum == 2 {
		msg := buf["Message"].(string)
		sign := buf["Signature"].(string)
		fmt.Println("verify msg = ", msg)
		fmt.Println("verify sign =", sign)
		if strings.Compare(sign, msgsign) != 0 {
			this.Data["json"] = map[string]interface{}{"result": false}
			this.ServeJSON()
			return
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
