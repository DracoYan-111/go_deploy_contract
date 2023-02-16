package deploy

import (
	box721 "GoContractDeployment/pkg"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"testing"
)

const RpcUrl = "http://127.0.0.1:8545/"
const UserPrivateKey = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"

type Structure struct {
	Name           string
	Symbol         string
	Minter         common.Address
	TokenURIPrefix string
}

// GoContractDeployment 创建合约并返回合约地址
func GoContractDeployment(structure Structure, t *testing.T) string {
	auth, client := GoCreateConnection(t)

	address, _, _, err := box721.DeployBox721(
		auth,
		client,
		structure.Name,
		structure.Symbol,
		structure.Minter,
		structure.TokenURIPrefix,
	)

	if err == nil {
		t.Log("最新nonce异常合约创建完成", address.Hex())
	}
	return address.Hex()
}

// GoCreateConnection createConnection
func GoCreateConnection(t *testing.T) (*bind.TransactOpts, *ethclient.Client) {

	// Connect to node
	client, err := ethclient.Dial(RpcUrl)
	if err != nil {
		t.Log("连接到节点异常", err)
	}

	// Create private key instance
	privateKey, err := crypto.HexToECDSA(UserPrivateKey)
	if err != nil {
		t.Log("加载私钥异常", err)
	}

	//Get the current chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		t.Log("获取链ID异常", err)
	}
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	// Get the latest random number of the current user
	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		t.Log("最新nonce异常", err)
	}

	// Estimated gasPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		t.Log("gasPrice获取异常", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(30000000) // in units
	auth.GasPrice = gasPrice

	return auth, client
}
