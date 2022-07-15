package signature

import (
	"math/big"
	"strconv"
)

func VeryfySignature(r, s uint32, mes [20]byte, pubKey, pubQ, pubP, pubG uint32) bool {
	var message [20]byte = mes
	var messString string
	for _, b := range message {
		messString = messString + strconv.FormatInt(int64(b), 10)
	}
	tempW := new(big.Int)
	tempU1 := new(big.Int)
	tempU2 := new(big.Int)
	tempW.Exp(big.NewInt(int64(s)), big.NewInt(int64(pubQ-2)), big.NewInt(int64(pubQ)))
	tempU1.SetString(messString, 10)
	tempU1.Mul(tempU1, tempW)
	tempU1.Mod(tempU1, big.NewInt(int64(pubQ)))

	tempU2.Mul(big.NewInt(int64(r)), tempW)
	tempU2.Mod(tempU2, big.NewInt(int64(pubQ)))

	tempW.Exp(big.NewInt(int64(pubG)), tempU1, nil)
	tempU1.Exp(big.NewInt(int64(pubKey)), tempU2, nil)
	tempU2.Mul(tempW, tempU1)
	tempU2.Mod(tempU2, big.NewInt(int64(pubP)))
	tempU2.Mod(tempU2, big.NewInt(int64(pubQ)))
	res := tempU2.Cmp(big.NewInt(int64(r)))
	if res == 0 {
		return true
	} else {
		return false
	}
}
