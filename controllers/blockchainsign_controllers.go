package controllers

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type BlockchainsignControllers struct {
	beego.Controller
}

func (this *BlockchainsignControllers) Get() {
	this.TplName = "Blockchainsign.html"
}
func (c *BlockchainsignControllers) Post() {
	buf := make(map[string]interface{})
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &buf)
	if err != nil {
		fmt.Println("json unmarshal err", err)
		return
	}
	blockhigh1 := buf["Bl"].(string)
	blockhigh, _ := strconv.ParseInt(blockhigh1, 10, 64)
	data := buf["Da"].(string)
	block := NewBlock_simple(data, blockhigh)
	pw := NewProofofWork(block)
	nonce, hashdata1 := pw.run()
	hashdata := hex.EncodeToString(hashdata1)
	fmt.Println(nonce)
	fmt.Println(hashdata)
	c.Data["json"] = map[string]interface{}{"hashdata": hashdata, "nonce": nonce}
	c.ServeJSON()
	return
}
