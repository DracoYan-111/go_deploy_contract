package deploy

import (
	"GoContractDeployment/pkg/box721"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

// GoInteractiveContract interactive contract
func GoInteractiveContract(contract *box721.Box721, t *testing.T) {
	auth, _ := GoCreateConnection()
	tx, err := contract.Erc721Mint(auth, big.NewInt(0), common.HexToAddress("0x0000000000000000000000000000000000000001"), "")
	if err != nil {
		t.Log("<==== LoadContract:Initiate transaction exception ====>", err)
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

// GoQueryContract Query contract
func GoQueryContract(contract *box721.Box721, t *testing.T) {
	name, err := contract.Name(nil)
	if err != nil {
		t.Log("<==== LoadContract:Query failed ====>", err)
	}
	t.Log(name)
}

// GoCreateAndGenerate Create a contract and generate a contract instance through the address
func GoCreateAndGenerate(structure Structure, t *testing.T) *box721.Box721 {
	//contractAddr := GoContractDeployment(structure)
	_, address, _, _ := GoContractDeployment(structure)
	example := GoLoadWithAddress(address, t)
	GoQueryContract(example, t)
	GoInteractiveContract(example, t)
	return example
}

// GoLoadWithAddress Generate contract instance by address
func GoLoadWithAddress(contractAddr string, t *testing.T) *box721.Box721 {
	_, client := GoCreateConnection()

	instance, err := box721.NewBox721(common.HexToAddress(contractAddr), client)
	if err != nil {
		panic(err)
	}
	t.Log("LoadContract:Contract loaded", 6)

	return instance
}
