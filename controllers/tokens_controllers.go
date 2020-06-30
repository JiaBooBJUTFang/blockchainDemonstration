package controllers

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/astaxie/beego"
)

type TokensControler struct {
	beego.Controller
}

func (this *TokensControler) Tokensget() {
	this.TplName = "tokens.html"
}
func (c *TokensControler) Tokenspost() {
	buf := make(map[string]interface{})
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &buf)
	if err != nil {
		fmt.Println("json unmarshal err", err)
		os.Exit(0)
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
