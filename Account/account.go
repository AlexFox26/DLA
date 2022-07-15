package account

import (
	"blockchain/key"
	"blockchain/transaction"

	"crypto/sha1"
	"strconv"
	"time"
)
var IDAutoIncrementation uint64 = 0//в качестве уникального ID

type Account struct{
	ID uint32
	keys keypair.Key
	votesNumber uint64
	dateLastVote time.Time
}

//создание ключей
func (objAcc *Account)GetAccount(){
	objAcc.ID = uint32(IDAutoIncrementation)
	IDAutoIncrementation++
	objAcc.keys.GenerateKeys()
	objAcc.votesNumber=0
}

//обновление ключа пользователя
func (objAcc *Account)AddKeyPair(){
	objAcc.keys.GenerateKeys()
}

//узнать количество голосов за пользователя
func (objAcc *Account)GetBalance( ) uint64{
	return objAcc.votesNumber
}

//функция для получения голоса
func (objAcc *Account)GetVote(){
	objAcc.votesNumber++
}

//функция для голосования за другого пользователя
func (objAcc *Account)Vote(reseiver *Account) *transaction.Transaction {
	var message string //формируем стоку, где просто говорим о том, что я хочу проголосовать за такого-то пользователя (необходяма просто для подтверждения транзакции)
	message = "Person " + strconv.FormatUint(uint64(objAcc.ID), 10) + " vote for person "+ strconv.FormatUint(uint64(reseiver.ID), 10)
	sha1.New()
	resHach := sha1.Sum([]byte(message)) //хешируем наше сообщение
	r,s:=objAcc.keys.SingData(resHach)	//подписываем
	resTrans:=new(transaction.Transaction)
	return resTrans.Vote(objAcc, reseiver,r,s, resHach)	//отправляем на проверку и подписание транзакции
}