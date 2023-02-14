package deploy

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
	"math/big"
	"os"
	"testing"
)

func GoContractDeployment(t *testing.T) {
	/*	bins, abis := compileFile()
		auth := generateSignature()
		// 部署合约
		_, tx, _, err := bind.DeployContract(auth, abis, common.FromHex(bins), client, value, gasLimit)
		if err != nil {
			// 处理错误
		}

		// 等待交易确认
		_, err = bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			// 处理错误
		}*/
	bins, abis := compileFile()
	auth := generateSignature()
	t.Log(bins, abis)
	t.Log(auth)
}

// @dev 合约文件读取
func compileFile() (string, abi.ABI) {
	// 读取编译后的二进制代码
	oldBin, err := os.ReadFile("build/MyContract.bin")
	if err != nil {
		panic("bin文件异常")
	}
	// 读取编译后的ABI
	oldAbi, err := os.ReadFile("build/MyContract.abi")
	if err != nil {
		panic("abi文件异常")
	}
	json, err := abi.JSON(io.Reader(bytes.NewReader(oldAbi)))
	if err != nil {
		panic("转换Abi格式异常")
	}
	return string(oldBin), json
}

// @dev 生成签名信息
func generateSignature() *bind.TransactOpts {
	// 创建账户密钥
	privateKey, err := crypto.HexToECDSA("YOUR_PRIVATE_KEY")
	if err != nil {
		panic("解析私钥异常")
	}
	// 创建签名器
	signer, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(0))
	if err != nil {
		panic("创建签名器异常")
	}
	return signer
}
