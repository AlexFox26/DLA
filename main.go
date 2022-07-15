package main

import (
	blockchane "blockchain/block"
	"blockchain/transaction"
)

func main() {

	mainchain := new(blockchane.Blockchain)
	mainchain.InitBlockchane()
	mainchain.AddUser()
	mainchain.AddUser()
	mainchain.AddUser()

	var transArr []transaction.Transaction

	currentUser := mainchain.GetObjUser(0)
	transArr = append(transArr, *currentUser.Vote(mainchain.GetObjUser(1)))
	currentUser = mainchain.GetObjUser(2)
	transArr = append(transArr, *currentUser.Vote(mainchain.GetObjUser(1)))

	block := new(blockchane.Block)
	block.CreateBlock(transArr, mainchain.GetLastBlock())

	mainchain.ValidateBlock(block) //бавить блок
	mainchain.PrevBalance()
}
