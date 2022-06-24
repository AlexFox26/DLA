package signature

import (
	"math/rand"
	"math/big"
	"strconv"
	"blockchain/key"
)
//Основано на алгоритме DSA

func SingData(objKey *keypair.Key, mes [20]byte) (uint32, uint32) {
	var message [20]byte = mes
	var messString string
	for _, b := range message {
		messString = messString + strconv.FormatInt(int64(b), 10)
	} 
	var r uint32 = objKey.GetpubQ() //используем как временную переменную в первый раз
	var s uint32
	temp1 := objKey.GetpubG() //для читабельности
	temp3 := new(big.Int)
	temp4 := new(big.Int)
	for {
		temp2 := rand.Int()%(int(r)-1) + 1 //случайное число 0 < temp2 < objKey.pubQ
		temp3.Exp(big.NewInt(int64(temp1)), big.NewInt(int64(temp2)), big.NewInt(int64(objKey.GetpubP())))
		r = uint32(temp3.Uint64()) % objKey.GetpubQ()

		temp4.Exp(big.NewInt(int64(temp2)), big.NewInt(int64(objKey.GetpubQ()-2)), big.NewInt(int64(objKey.GetpubQ())))
		temp3.SetString(messString, 10)
		temp3.Add(temp3, (big.NewInt(int64(objKey.GetPrivKey()) * int64(r))))
		temp3.Mul(temp3, temp4)
		temp3.Mod(temp3, big.NewInt(int64(objKey.GetpubQ())))
		messString = temp3.Text(10)
		temp5, err2 := strconv.ParseInt(messString, 10, 64)
		if err2 != nil { return 0, 0	}
		s = uint32(temp5)
		if r != 0 && s != 0 {
			break
		}
	}
	return r, s
}

func VeryfySignature(r, s uint32, mes [20]byte, objKey *keypair.Key) bool {

	var message [20]byte = mes
	var messString string
	for _, b := range message {
		messString = messString + strconv.FormatInt(int64(b), 10)
	}

	tempW := new(big.Int)
	tempU1 := new(big.Int)
	tempU2 := new(big.Int)
	tempW.Exp(big.NewInt(int64(s)), big.NewInt(int64(objKey.GetpubQ()-2)), big.NewInt(int64(objKey.GetpubQ())))
	tempU1.SetString(messString, 10)
	tempU1.Mul(tempU1, tempW)
	tempU1.Mod(tempU1, big.NewInt(int64(objKey.GetpubQ())))

	tempU2.Mul(big.NewInt(int64(r)), tempW)
	tempU2.Mod(tempU2, big.NewInt(int64(objKey.GetpubQ())))

	tempW.Exp(big.NewInt(int64(objKey.GetpubG())), tempU1, nil)
	tempU1.Exp(big.NewInt(int64(objKey.GetPubKey())), tempU2, nil)
	tempU2.Mul(tempW, tempU1)
	tempU2.Mod(tempU2, big.NewInt(int64(objKey.GetpubP())))
	tempU2.Mod(tempU2, big.NewInt(int64(objKey.GetpubQ())))
	res := tempU2.Cmp(big.NewInt(int64(r)))
	if res == 0 {
		return true
	} else {
		return false
	}
}