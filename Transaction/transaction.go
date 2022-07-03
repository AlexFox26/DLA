package transaction


import (
	"blockchain/account"
	"blockchain/key"
	"blockchain/signature"
	"time"
)

type Transaction struct{
	sender account.Account
	receiver account.Account
	signature []byte
}

//верификация и создание транзакции
func (objOper * Transaction)Vote(sen, res account.Account, r,s uint32, signHach [20]byte) *Transaction{
	objOper.sender = sen //записываем в транзакцию получателя и отправителя
	objOper.receiver = res
	currentdata:=new(time.Time)	//проверяем отправителя на способность вообще голосовать
	if(sen.dataLastVote != nil){//проверяем прошел ли год с момента прошлого голосования (было решено дать возможность голосовать раз в год)
		if( currentdata.Year() - sen.dataLastVote.Year() < 1 ){
			return nil	//если год не прошел, транзакция не подтвержается
		}
	}	
	//проверяем действительно ли операция голосования правильно подписана отправителем
	if (!signature.VeryfySignature(r,s, signHach, objOper.sender.keys.PubQ,objOper.sender.keys.PubP, objOper.sender.keys.PubG)){
		return nil	//если подписана неправильно, то транзакция не подтверждается
	}

	objOper.signature=signHach[:] //запоминаем подписанное письмо
	objOper.receiver.GetVote()	//добавляем голос получателю
	return objOper	
}

/*
	неободимости создавать посредника Operation нет, так как операция может быть только одна - это голосование
	нет смысла создавать 1000операций в одну транзакцию и их сохранять в блоке, если можно сохранить 1000 
	транзакций в блоке и сэкономить память из-за отсутсвия посредника. А данные все и так есть - это голосующий
	и за кого голосует, а так же подписанное письмо для верификации.
*/