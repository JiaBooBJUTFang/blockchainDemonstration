package controllers

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

const N = 4

type BlockchainControler struct {
	beego.Controller
}

func (this *BlockchainControler) Blockchaiget() {
	this.TplName = "blockchain.html"
}
func (c *BlockchainControler) Blockchainpost() {
	buf := make(map[string]interface{})
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &buf)
	if err != nil {
		fmt.Println("json unmarshal err", err)
		return
	}
	blockhigh1 := buf["Bl"].(string)
	data := buf["Da"].(string)
	fmt.Println("---", blockhigh1)
	fmt.Println("---", data)
	blockhigh, _ := strconv.ParseInt(blockhigh1, 10, 64)
	//chainnum, _ := strconv.ParseInt(chainnum1, 10, 64)

	block := NewBlock_simple(data, blockhigh)
	//hash_data := block.Hash //测试
	pw := NewProofofWork(block)
	nonce, hashdata1 := pw.run()
	hashdata := hex.EncodeToString(hashdata1)
	c.Data["json"] = map[string]interface{}{"hashdata": hashdata, "nonce": nonce}
	fmt.Println(nonce)
	fmt.Println(hashdata)
	c.ServeJSON()
	return
}
