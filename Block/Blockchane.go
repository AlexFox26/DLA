package blockchane

import (
	"blockchain/account"
	"blockchain/transaction"
	"crypto/sha1"
	"fmt"
)


type Blockchain struct{
	coinDatabase map[string] int
	blockHistory []Block
	users []account.Account
}

type Block struct {
	idBlock []byte
	prevHash []byte
	listTrans []transaction.Transaction 
}
func (objBlockchain *Blockchain)GetObjUser(userId int) *account.Account{
	return &objBlockchain.users[userId]
}
func (objBlockchain *Blockchain)InitBlockchane(){
	genesus:=new(Block)
	genesus.prevHash = []byte("0")
	genesus.listTrans=nil
	sha1.New()
	resHach := sha1.Sum([]byte("Genesis block"))
	genesus.idBlock= resHach[:]
	objBlockchain.ValidateBlock(genesus)
}
func (objBlockchain *Blockchain)AddUser(){
	ref:=new(account.Account)
	ref.GetAccount()
	objBlockchain.users = append(objBlockchain.users, *ref)
	//ref=nil
}
func (objBlockchain *Blockchain)PrevBalance(){
	for a,b:=range objBlockchain.users{
		fmt.Printf("User %d have %d votes", a, b.GetBalance())
	}
}

func (objBlockchain *Blockchain)ValidateBlock( ovjBlock *Block){
	objBlockchain.blockHistory = append(objBlockchain.blockHistory, *ovjBlock)
}


func (objBlockchain *Blockchain)GetLastBlock( ) *Block{
	
	var Index int
	for a,_:=range objBlockchain.blockHistory{
		Index=a
	}
	
	return &objBlockchain.blockHistory[Index]
}
func ( objBlock *Block)CreateBlock( objTrans []transaction.Transaction, prevBlock *Block) *Block{
	
	objBlock.prevHash = prevBlock.idBlock
	for _,b:= range objTrans{
		objBlock.listTrans = append(objBlock.listTrans, b)
	}
	var forID []byte
	for _,b:= range objTrans{
		forID = append(forID, b.signature)
	}
	sha1.New()
	resHach := sha1.Sum(forID) //хешируем наше сообщение
	objBlock.idBlock = resHach[:]
	return objBlock
}