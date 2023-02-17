package deploy

import (
	"GoContractDeployment/pkg/box721"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"testing"
	"time"
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
	auth, client := GoCreateConnection("https://data-seed-prebsc-1-s1.binance.org:8545/", t)

	address, txData, _, err := box721.DeployBox721(
		auth,
		client,
		structure.Name,
		structure.Symbol,
		structure.Minter,
		structure.TokenURIPrefix,
	)

	if err != nil {
		t.Log("创建合约异常", address.Hex())
	}
	t.Log("开始等待", txData.Hash().Hex())
	time.Sleep(5 * time.Second)
	return address.Hex()
}

// GoTransactionNews 查询使用的gas
func GoTransactionNews(client *ethclient.Client, hash string, t *testing.T) uint64 {
	txHash := common.HexToHash(hash)

	// 获取交易
	_, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	// 如果交易未被打包，等待它被打包
	if isPending {
		log.Fatal("Transaction is pending")
	}

	// 获取交易所使用的gas数量
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	t.Log(receipt.GasUsed)
	return receipt.GasUsed
}

// GoCreateConnection createConnection
func GoCreateConnection(url string, t *testing.T) (*bind.TransactOpts, *ethclient.Client) {
	var client *ethclient.Client
	var err error
	if len(url) > 0 {
		// Connect to node
		fmt.Println("000000000000000000000000000000")
		client, err = ethclient.Dial(url)
		if err != nil {
			t.Log("连接到节点异常", err)
		}
	} else {
		fmt.Println("11111111111111111111111111111111")
		// Connect to node
		client, err = ethclient.Dial(RpcUrl)
		if err != nil {
			t.Log("连接到节点异常", err)
		}
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
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6000000) // in units
	auth.GasPrice = gasPrice

	return auth, client
}
