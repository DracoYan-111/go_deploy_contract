package models

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

// ReturnPost 返回的信息结构
type ReturnPost struct {
	Opcode       int64          `json:"id"`
	ContractAddr common.Address `json:"contract address"`
	ContractHash string         `json:"contract hash"`
	ChainId      big.Int        `json:"chain id"`
	GasUsed      big.Int        `json:"gas price"`
	GasUST       big.Int        `json:"gas UST"`
}

// Post 数据库的信息结构
type Post struct {
	ID            int64     `json:"id"`
	Opcode        int64     `json:"opcode"`
	ContractName  string    `json:"contract name"`
	ContractAddr  string    `json:"contract address"`
	ContractHash  string    `json:"contract hash"`
	GasUsed       int64     `json:"gas price"`
	GasUST        int64     `json:"gas UST"`
	ChainId       int64     `json:"chain id"`
	CreatedAt     time.Time `json:"created_at"`
	CurrentStatus int64     `json:"current status"`
}
