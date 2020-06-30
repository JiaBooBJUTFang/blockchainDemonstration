package controllers

import (
"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/astaxie/beego"
)

type DistributedControler struct {
	beego.Controller
}

func (this *DistributedControler) Distributedget() {
	this.TplName = "distributed.html"

}
func (c *DistributedControler) Distributedpost() {
buf := make(map[string]interface{})
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &buf)
	if err != nil {
		fmt.Println("json unmarshal err", err)
		os.Exit(0)
	}
	chainnum := buf["Chain"].(float64)
	blockhigh1 := buf["Bl"].(string)
	blockhigh, _ := strconv.ParseInt(blockhigh1, 10, 64)
	data := buf["Da"].(string)
	block := NewBlock_simple(data, blockhigh)
	//hash_data := block.Hash //测试
	pw := NewProofofWork(block)
	nonce, hashdata1 := pw.run()
	hashdata := hex.EncodeToString(hashdata1)
	c.Data["json"] = map[string]interface{}{"hashdata": hashdata, "nonce": nonce, "blockhigh": blockhigh, "chainnum": chainnum}
	c.ServeJSON()
	c.StopRun()

}
