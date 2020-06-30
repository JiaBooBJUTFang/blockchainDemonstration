package routers

import (
	"blockchain/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/login", &controllers.MainController{}, "get:Mainget")
	beego.Router("/blockinfo", &controllers.BlockControler{}, "get:Blockget;post:Blockpost")
	beego.Router("/hash", &controllers.HashControler{}, "get:Hashget")
	beego.Router("/distributed", &controllers.DistributedControler{}, "get:Distributedget;post:Distributedpost")
	beego.Router("/blockchain", &controllers.BlockchainControler{}, "get:Blockchaiget;post:Blockchainpost")
	beego.Router("/coinbase", &controllers.CoinbaseControler{}, "get:Coinbaseget;post:Coinbasepost")
	beego.Router("/tokens", &controllers.TokensControler{}, "get:Tokensget;post:Tokenspost")

	beego.Router("/signature", &controllers.SignaturesControllers{})
	beego.Router("/transaction", &controllers.TransactionControllers{})
	beego.Router("/Blockchainsign", &controllers.BlockchainsignControllers{})
	beego.Router("/keys", &controllers.KeysControllers{}, "get:Keyget;post:MyPost")

	beego.Router("/chaindemo", &controllers.IndexControllers{})
	beego.Router("/chaindemo2", &controllers.Index2Controllers{})
}
