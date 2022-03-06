package utils

import "math/big"

// ToAIRE number of AIRE to Wei
func ToAIRE(aire uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(aire), big.NewInt(1e18))
}
