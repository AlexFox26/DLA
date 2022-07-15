package transaction

import (
	"blockchain/account"
	"blockchain/signature"
	"time"
)

type Transaction struct{
	sender account.Account
	receiver account.Account
	signature []byte
} 

//верификация и создание транзакции
func (objOper *Transaction)Vote(sen, res account.Account, r,s uint32, signHach [20]byte) *Transaction{
	objOper.sender = sen //записываем  транзакцию получателя и отправителя
	objOper.receiver = res   
	currentdata:=new(time.Time)	//проверяем отправителя на способность вообе голосовать
	if(sen.dataLastVote !=nil){//проверяем прошел ли год с момента прошлого голосования (было решено дать возможность голосовать раз в год)
		if( currentdata.Year() - sen.dataLastVote.Year() < 1 ){
			 return nil	//если год не п  рошел, транзакция не подтвержается
		}
	}	 
	//роверяем действительно ли операция голосования правильно подписана отправителем
	if(!signature.VeryfySignature(r,s, signHach, objOper.sender.keys.PubQ, objOper.sender.keys.PubP, objOper.sender.keys.PubG)){
		return nil	//если подписана неправильно, то транзакция не подтверждается
	}   
 
	objOper.signature=signHach[:] //запоминаем подписанное письмо
	objOper.receiver.GetVote()	//добавляем голос получателю
	return objOper	  
}      
