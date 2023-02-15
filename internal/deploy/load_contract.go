package deploy

import (
	box721 "GoContractDeployment/pkg"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

// 交互合约
func GoInteractiveContract(contract *box721.Box721, t *testing.T) {
	auth, _ := GoCreateConnection(t)
	tx, err := contract.Erc721Mint(auth, big.NewInt(0), common.HexToAddress("0x0000000000000000000000000000000000000001"), "")
	if err != nil {
		t.Log("发起交易异常", err)
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

// 查询合约
func GoQueryContract(contract *box721.Box721, t *testing.T) {
	name, err := contract.Name(nil)
	if err != nil {
		t.Log("查询失败", err)
	}
	t.Log(name)

}

// 创建合约并通过地址生成合约实例
func GoCreateAndGenerate(structure Structure, t *testing.T) *box721.Box721 {
	//contractAddr := GoContractDeployment(structure)
	example := GoLoadWithAddress(GoContractDeployment(structure, t), t)
	GoQueryContract(example, t)
	GoInteractiveContract(example, t)
	return example
}

// 通过地址生成合约实例
func GoLoadWithAddress(contractAddr string, t *testing.T) *box721.Box721 {
	_, client := GoCreateConnection(t)

	instance, err := box721.NewBox721(common.HexToAddress(contractAddr), client)
	if err != nil {
		panic(err)
	}
	t.Log("合约已加载", 6)

	return instance
}
