package deploy

import (
	"GoContractDeployment/pkg/box721"
	"GoContractDeployment/utils"
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

type Structure struct {
	Name           string
	Symbol         string
	Minter         common.Address
	TokenURIPrefix string
}

// GoContractDeployment Create a contract and return the contract address
func GoContractDeployment(structure Structure) (string, string, *big.Int, int64) {
	auth, client := GoCreateConnection()

	balance, err := client.BalanceAt(context.Background(), auth.From, nil)

	if balance.Int64() < 5e16 {
		log.Println("Deployment:Insufficient user balance", balance)

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
		log.Println("Deployment:Create contract exception", err)

		return "", "", big.NewInt(0), 0
	}
	log.Println(structure.Name, "Deployment:Start deployment:", txData.Hash().Hex())

	gasUsed, err := goTransactionNews(client, txData.Hash().Hex())

	gas := gasUsed.Mul(gasUsed, txData.GasPrice())
	time.Sleep(3 * time.Second)

	return address.Hex(), txData.Hash().Hex(), gas.Add(gas, big.NewInt(5e12)), 1

}

// goTransactionNews Query the gas used
func goTransactionNews(client *ethclient.Client, hash string) (*big.Int, error) {
	time.Sleep(7 * time.Second)

	txHash := common.HexToHash(hash)

	_, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Println(err)
	}

	if isPending {
		log.Println("Deployment:Transaction is being packaged")

		return new(big.Int).SetUint64(0), errors.New("Deployment:Transaction in progress")
	} else {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			log.Println(err, "Deployment:Get the amount of gas used by the transaction")
		}
		return new(big.Int).SetUint64(receipt.GasUsed), nil
	}
}

// GoCreateConnection CreateConnection
func GoCreateConnection() (*bind.TransactOpts, *ethclient.Client) {
	var client *ethclient.Client
	var err error
	// Connect to node
	loading, err := utils.ConfigurationLoading("web3", []string{"rpcUrl", "privateKey"})
	if err != nil {
		log.Panicln("ReturnStatus:", err)
	}
	client, err = ethclient.Dial(loading[0])
	if err != nil {
		log.Println("<==== Deployment:Connection to node exception ====>", err)
	} else {
		log.Println("<++++ Deployment:Connected to node successfully ++++>")
	}
	//}

	// Create private key instance
	privateKey, err := crypto.HexToECDSA(loading[1])
	if err != nil {
		log.Println("<==== Deployment:Exception loading private key ====>", err)
	}

	//Get the current chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Println("<==== Deployment:Obtaining the chain ID is abnormal ====>", err)
	}
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	// Get the latest random number of the current user
	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Println("<==== Deployment:Latest nonce exception ====>", err)
	}

	// Estimated gasPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println("<==== Deployment:Gas Price get exception ====>", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(5100000) // in units
	auth.GasPrice = gasPrice

	return auth, client
}
