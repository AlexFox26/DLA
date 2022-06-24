package keypair

import (
	"math"
	"math/rand"
	"fmt"
	"time"
)
//Основано на алгоритме DSA
type Key struct{
	privateKey uint32
	publicKey uint32
   
	pubQ uint32
	pubP uint32
	pubG uint32
   }
   
   func (obj * Key) PrintKeys(){
	   fmt.Printf("Private key: %d\nPublic key: %d\n", obj.privateKey, obj.publicKey)
   }
   
   
   func (obj * Key) GenerateKeys() {
	obj.pubQ = 11
	obj.pubP = 23
	obj.pubG = 4 //2^(22/11)mod23
	   
	   rand.Seed(time.Now().UnixNano())
	   obj.privateKey = rand.Uint32() %obj.pubQ
	   obj.publicKey = uint32(math.Pow(float64(obj.pubG), float64(obj.privateKey))) % obj.pubP

   }
   
    func (obj * Key) GetpubQ() uint32{
		return obj.pubQ
	}
	func (obj * Key) GetpubP() uint32{
		return obj.pubP
	}
	func (obj * Key) GetpubG() uint32{
		return obj.pubG
	}
   func (obj * Key) GetPubKey() uint32 {
	   return obj.publicKey
   }
   func (obj * Key) GetPrivKey() uint32 {
	return obj.privateKey
}
