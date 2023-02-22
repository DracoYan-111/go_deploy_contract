package models

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

const (
	UpdateTaskOne = "update one"
)

const (
	InsertIntoJob = "INSERT INTO go_test_db (opcode, contract_name, chain_id) VALUES (?, ?, ?)"
	SelectOperate = "SELECT * FROM go_test_db WHERE current_status=?"
	SelectGetOne  = "SELECT * FROM go_test_db WHERE current_status=0 LIMIT 1"
)

// ReceivePost 接收参数
type ReceivePost struct {
	Opcode       string `json:"id"`
	ContractName string `json:"name"`
	ChainId      int64  `json:"chainId"`
}

// ReturnPost 返回的信息结构
type ReturnPost struct {
	Opcode       int64          `json:"id"`
	ContractAddr common.Address `json:"contract address"`
	ContractHash string         `json:"contract hash"`
	ChainId      big.Int        `json:"chain id"`
	GasUsed      big.Int        `json:"gas price"`
	GasUST       big.Int        `json:"gas UST"`
}

// DataPost 数据库的信息结构
type DataPost struct {
	ID            int64     `json:"id"`
	Opcode        string    `json:"opcode"`
	ContractName  string    `json:"contract_name"`
	ContractAddr  string    `json:"contract_address"`
	ContractHash  string    `json:"contract_hash"`
	GasUsed       int64     `json:"gas_used"`
	GasUST        float64   `json:"gas_usdt"`
	ChainId       int64     `json:"chain_id"`
	CreatedAt     time.Time `json:"created_at"`
	CurrentStatus int64     `json:"current_status"`
}
