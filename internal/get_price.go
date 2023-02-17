package internal

import (
	deploy "GoContractDeployment/internal/deploy"
	"GoContractDeployment/pkg/pancake"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math"
	"math/big"
	"testing"
)

var PancakeRouter = common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E")
var Bnb = common.HexToAddress("0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c")
var Usdt = common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56")

func GetBnbToUsdt(amountIn *big.Int, t *testing.T) string {

	amountIn.Mul(amountIn, big.NewInt(10000005000))
	var path = []common.Address{Bnb, Usdt}
	instance := GoLoadWithAddress(t)
	out, err := instance.GetAmountsOut(nil, amountIn, path)
	if err != nil {
		log.Fatal("价格查询失败", err)
	}
	convert := new(big.Float)
	convert.SetString(out[1].String())
	value := new(big.Float).Quo(convert, big.NewFloat(math.Pow10(18)))

	return value.Text('f', 10)
}

// GoLoadWithAddress 通过地址生成合约实例
func GoLoadWithAddress(t *testing.T) *pancake.PancakeRouter {
	_, client := deploy.GoCreateConnection("https://bsc-dataseed1.ninicoin.io/", t)

	instance, err := pancake.NewPancakeRouter(PancakeRouter, client)
	if err != nil {
		panic(err)
	}
	t.Log("合约已加载", 6)

	return instance
}
