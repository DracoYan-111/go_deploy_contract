package internal

import (
	deploy "GoContractDeployment/internal/deploy"
	"GoContractDeployment/pkg/pancake"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math"
	"math/big"
)

// PancakeRouter main net
var PancakeRouter = common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E")
var Wbnb = common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
var Usdt = common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56")

func GetBnbToUsdt(amountIn *big.Int) float64 {

	log.Println(amountIn)
	var path = []common.Address{Wbnb, Usdt}
	instance := GoLoadWithAddress()
	out, err := instance.GetAmountsOut(nil, amountIn, path)
	if err != nil {
		log.Fatal("<==== GetPrice:Price query failed ====>", err)
	}

	convert := new(big.Float)
	convert.SetString(out[len(out)-1].String())
	value, _ := new(big.Float).Quo(convert, big.NewFloat(math.Pow10(18))).Float64()

	return value
}

// GoLoadWithAddress Generate contract instance by address
func GoLoadWithAddress() *pancake.PancakeRouter {
	_, client := deploy.GoCreateConnection()

	instance, err := pancake.NewPancakeRouter(PancakeRouter, client)
	if err != nil {
		panic(err)
	}
	log.Println("GetPrice:Pancake Swap contract loaded")

	return instance
}
