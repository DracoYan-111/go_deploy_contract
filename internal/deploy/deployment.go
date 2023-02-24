package deploy

import (
	"GoContractDeployment/pkg/box721"
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
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
func GoContractDeployment(structure Structure) (string, string, *big.Int, int64) {
	auth, client := GoCreateConnection("https://data-seed-prebsc-1-s1.binance.org:8545")

	balance, err := client.BalanceAt(context.Background(), auth.From, nil)

	if balance.Int64() < 5e16 {
		log.Println("用户余额不足", balance)

		return "", "", big.NewInt(0), 0
	}
	address, txData, _, err := box721.DeployBox721(
		auth,
		client,
		structure.Name,
		structure.Symbol,
		structure.Minter,
		structure.TokenURIPrefix,
	)

	if err != nil {
		log.Println("创建合约异常", err)

		return "", "", big.NewInt(0), 0
	}
	log.Println(structure.Name, "开始部署:", txData.Hash().Hex())

	gasUsed, err := goTransactionNews(client, txData.Hash().Hex())

	gas := gasUsed.Mul(gasUsed, txData.GasPrice())
	time.Sleep(3 * time.Second)

	return address.Hex(), txData.Hash().Hex(), gas.Add(gas, big.NewInt(5e10)), 1

}

// goTransactionNews 查询使用的gas
func goTransactionNews(client *ethclient.Client, hash string) (*big.Int, error) {
	time.Sleep(7 * time.Second)

	txHash := common.HexToHash(hash)

	// 获取交易
	_, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Println(err)
	}

	// 如果交易未被打包，等待它被打包
	if isPending {
		log.Println("交易正在打包")

		return new(big.Int).SetUint64(0), errors.New("交易进行中")
	} else {
		// 获取交易所使用的gas数量
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			log.Println(err, "获取交易所使用的gas数量")
		}
		return new(big.Int).SetUint64(receipt.GasUsed), nil
	}
}

// GoCreateConnection createConnection
func GoCreateConnection(url string) (*bind.TransactOpts, *ethclient.Client) {
	var client *ethclient.Client
	var err error
	if len(url) > 0 {
		// Connect to node
		client, err = ethclient.Dial(url)
		if err != nil {
			log.Println("<==== 连接到节点异常 ====>", err)
		} else {
			log.Println("<++++ 连接到节点成功 ++++>")
		}
	} else {
		// Connect to node
		client, err = ethclient.Dial(RpcUrl)
		if err != nil {
			log.Println("<==== 连接到节点异常 ====>", err)
		} else {
			log.Println("<++++ 连接到节点成功 ++++>")
		}
	}

	// Create private key instance
	privateKey, err := crypto.HexToECDSA(UserPrivateKey)
	if err != nil {
		log.Println("<==== 加载私钥异常 ====>", err)
	}

	//Get the current chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Println("<==== 获取链ID异常 ====>", err)
	}
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	// Get the latest random number of the current user
	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Println("<==== 最新nonce异常 ====>", err)
	}

	// Estimated gasPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println("<==== gasPrice获取异常 ====>", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(5100000) // in units
	auth.GasPrice = gasPrice

	return auth, client
}
