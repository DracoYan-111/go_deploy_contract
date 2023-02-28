package models

import (
	"time"
)

// ReceivePost received message structure
type ReceivePost struct {
	Opcode       string `json:"id"`
	ContractName string `json:"name"`
	ChainId      int64  `json:"chainId"`
}

// ReturnPost Returned information structure
type ReturnPost struct {
	Opcode         int64   `json:"productContractId"`
	ChainId        int64   `json:"chainId"`
	GasUST         float64 `json:"gasFee"`
	ContractAddr   string  `json:"contract"`
	ContractHash   string  `json:"txHash"`
	ContractOwner  string  `json:"owner"`
	ContractMinter string  `json:"minter"`
}

// DataPost database information structure
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
