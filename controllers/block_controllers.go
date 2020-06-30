package controllers

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type BlockControler struct {
	beego.Controller
}
type Bb struct {
	Bl string `json:"bl"`
	Ts string `json:"ts"`
	Da string `json:"da"`
}

func (c *BlockControler) Blockget() {
	c.TplName = "block.html"
}

func (this *BlockControler) Blockpost() {
	bb := &Bb{}
	var err error
	redata := this.Ctx.Input.RequestBody
	if err = json.Unmarshal(redata, bb); err == nil {
		blockhigh1 := bb.Bl
		//timenow:=time.Now().Format("2006-01-02 15:04:05")
		blockhigh, _ := strconv.ParseInt(blockhigh1, 10, 64)
		data := bb.Da
		fmt.Println("success...", bb)
		block := NewBlock_simple(data, blockhigh)
		pw := NewProofofWork(block)
		nonce, hashdata := pw.run()
		hashreturn := hex.EncodeToString(hashdata)
		fmt.Println("nonce = ", nonce, "hash = ", hashreturn)
		this.Data["hashre"] = hashreturn
		this.Data["json"] = map[string]interface{}{"datas": hashreturn, "nonce": nonce}
	} else {
		beego.Info(err.Error())
	}
	this.ServeJSON()
}
