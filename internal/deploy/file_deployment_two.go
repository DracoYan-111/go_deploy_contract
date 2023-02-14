package deploy

//
//import (
//	"context"
//	"crypto/ecdsa"
//	"fmt"
//	"log"
//	"math/big"
//
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	"github.com/ethereum/go-ethereum/common/hexutil"
//	"github.com/ethereum/go-ethereum/crypto"
//	"github.com/ethereum/go-ethereum/ethclient"
//)
//
//func main() {
//	client, err := ethclient.Dial("https://mainnet.infura.io")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	privateKey, err := crypto.HexToECDSA("your_private_key")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	publicKey := privateKey.Public()
//	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//	if !ok {
//		log.Fatal("error casting public key to ECDSA")
//	}
//
//	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
//	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	gasPrice, err := client.SuggestGasPrice(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	auth := bind.NewKeyedTransactor(privateKey)
//	auth.Nonce = big.NewInt(int64(nonce))
//	auth.Value = big.NewInt(0)
//	auth.GasLimit = uint64(300000)
//	auth.GasPrice = gasPrice
//
//	input, err := hexutil.Decode("<bytecode>")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	address, tx, contract, err := bind.DeployContract(auth, abi, input, client)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("Contract deployed to address:", address.Hex())
//	fmt.Println("Transaction hash:", tx.Hash().Hex())
//
//	_ = contract
//}
