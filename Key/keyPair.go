package keypair

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

//Основано на алгоритме DSA
const   PubQ uint32 = 11
const	PubP uint32  = 23
const	PubG uint32 = 4 //2^(22/11)mod23


type Key struct{
	privateKey uint32
   }
   
   func (obj * Key) PrintKeys(){
	   fmt.Printf("Private key: %d\nPublic key: %d\n", obj.privateKey, obj.GetPubKey())
   }
   
   
   func (obj * Key) GenerateKeys() {
   
	   rand.Seed(time.Now().UnixNano())
	   obj.privateKey = rand.Uint32() % PubQ
   }
   
    func (obj * Key) GetpubQ() uint32{
		return PubQ
	}
	func (obj * Key) GetpubP() uint32{
		return PubP
	}
	func (obj * Key) GetpubG() uint32{
		return PubG
	}
   func (obj * Key) GetPubKey() uint32 {
		var publicKey uint32
		publicKey = uint32(math.Pow(float64(PubG), float64(obj.privateKey))) % PubP
	    return publicKey
   }
   func (obj * Key) GetPrivKey() uint32 {
		return obj.privateKey
   }

   func (objKey * Key)SingData(mes [20]byte) (uint32, uint32) {
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