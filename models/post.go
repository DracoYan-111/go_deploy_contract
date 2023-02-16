package models

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// Post type details
type Post struct {
	ID           int64          `json:"id"`
	ContractAddr common.Address `json:"contract address"`
	ChainId      big.Int        `json:"chain id"`
	// created_at time.Time `json:"created_at"`
	// updated_at time.Time `json:"updated_at"`
}
