package try_test

import (
	"GoContractDeployment/internal/deploy"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestDeploy(t *testing.T) {
	stream :=
		deploy.Structure{
			Name:           "TianYun",
			Symbol:         "TianYun",
			Minter:         common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"),
			TokenURIPrefix: "test",
		}
	fmt.Print(stream)

	deploy.GoCreateAndGenerate(stream, t)
}
