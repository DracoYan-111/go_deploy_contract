package internal

import (
	deploy "GoContractDeployment/internal/deploy"
	"GoContractDeployment/pkg/pancake"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math"
	"math/big"
)

// testnet
var network = "https://data-seed-prebsc-1-s3.binance.org:8545/"
var PancakeRouter = common.HexToAddress("0xD99D1c33F9fC3444f8101754aBC46c52416550D1")
var Wbnb = common.HexToAddress("0xae13d989dac2f0debff460ac112a837c89baa7cd")
var Cake = common.HexToAddress("0xFa60D973F7642B748046464e165A65B7323b0DEE")
var Usdt = common.HexToAddress("0xaB1a4d4f1D656d2450692D237fdD6C7f9146e814")

// mainnet
// var network = "https://bsc-dataseed1.binance.org/"
// var PancakeRouter = common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E")
// var Wbnb = common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
// var Usdt = common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56")

func GetBnbToUsdt(amountIn *big.Int) string {

	amountIn.Mul(amountIn, big.NewInt(10000005000))
	log.Println(amountIn)
	var path = []common.Address{Wbnb, Cake, Usdt}
	instance := GoLoadWithAddress()
	out, err := instance.GetAmountsOut(nil, amountIn, path)
	if err != nil {
		log.Fatal("价格查询失败", err)
	}
	convert := new(big.Float)
	convert.SetString(out[len(out)-1].String())
	value := new(big.Float).Quo(convert, big.NewFloat(math.Pow10(18)))

	return value.Text('f', 10)

	//instance := GoLoadWithAddress()
	//weth, _ := instance.WETH(nil)
	//
	//return weth.String()
}

// GoLoadWithAddress 通过地址生成合约实例
func GoLoadWithAddress() *pancake.PancakeRouter {
	_, client := deploy.GoCreateConnection(network)

	instance, err := pancake.NewPancakeRouter(PancakeRouter, client)
	if err != nil {
		panic(err)
	}
	log.Println("合约已加载", 6)

	return instance
}
