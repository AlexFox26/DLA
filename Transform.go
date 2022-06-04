package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

//обрезает лишние ненужные "0" в начале или конце строки. Place = "Start" or "End"
func cutZero(s, place string) string {
	Pl := place
	letters := []rune(s)
	if Pl != "End" && Pl != "Start" {
		return "error: place not found"
	}
	var isZeroSequence bool = false
	var numberOfZero int

	switch place {
	case "End":

		for _, b := range letters {
			if b == 48 {
				isZeroSequence = true
				numberOfZero++
			} else {
				isZeroSequence = false
				numberOfZero = 0
			}
		}
		if isZeroSequence {
			if (len(letters)-numberOfZero)%2 == 1 {
				numberOfZero--
			}
			return string(letters[0 : len(letters)-numberOfZero])
		}
		return string(letters)

	case "Start":

		for _, b := range letters {
			if b == 48 {
				numberOfZero++
			} else {
				if (len(letters)-numberOfZero)%2 == 1 {
					numberOfZero--
				}
				if numberOfZero < 0 {
					numberOfZero = 0
				}
				return string(letters[numberOfZero:len(letters)])
			}
		}
	}
	return "error"
}

//меняет порядок байт в строке на симметричный
func Invert(s string) string {
	letters := []rune(s)
	i1 := 0
	i2 := 1
	j1 := len(letters) - 2
	j2 := len(letters) - 1
	for i2 < j1 {

		letters[i1], letters[j1] = letters[j1], letters[i1]
		letters[i2], letters[j2] = letters[j2], letters[i2]
		j1 -= 2
		j2 -= 2
		i1 += 2
		i2 += 2
	}
	return string(letters)
}

//соответствие двуричной системы шестнадцатиричной
func TwoToHex(s string) string {
	switch s {
	case "0000":
		return "0"
	case "0001":
		return "1"
	case "0010":
		return "2"
	case "0011":
		return "3"
	case "0100":
		return "4"
	case "0101":
		return "5"
	case "0110":
		return "6"
	case "0111":
		return "7"
	case "1000":
		return "8"
	case "1001":
		return "9"
	case "1010":
		return "a"
	case "1011":
		return "b"
	case "1100":
		return "c"
	case "1101":
		return "d"
	case "1110":
		return "e"
	case "1111":
		return "f"
	}
	return "error: not number"
}

//перевод из Hex в Little Endian или BIG Endian.  Endian = "Little" or "BIG"
func HexToEndian(message, Endian string) big.Int {
	End := Endian //хороший тон не трогать пришедшие переменные
	result := message
	if End != "BIG" && End != "Little" {
		return *big.NewInt(-1)
	} //выходим с ошибкой, если неправильно указан Endian
	result = strings.ToLower(result)              //переводим все в нижний регистр
	result = strings.Replace(result, "0x", "", 1) //убираем 0х в начале
	if End == "Little" {
		result = Invert(result)           //переставляем байты
		result = cutZero(result, "Start") //обрезаем мусор в виде "0" в начале
	}
	var res, _ = new(big.Int).SetString(result, 16) //переводим получившееся число в большой Int
	return *res
}

//перевод из Little Endian или BIG Endian в Hex. Endian = "Little" or "BIG"
func EndianToHex(message big.Int, Endian string) string {
	End := Endian
	number := message
	if End != "BIG" && End != "Little" {
		return "error: Endian not found"
	} //выходим с ошибкой, если неправильно указан Endian
	var result string           //результат перевода значения
	var numberInPartByte string //для сбора 4 последовательных бит, чтобы после перевести число в 16 ричную систему
	numOfBit := number.BitLen() //количество бит, которое образует число
	if numOfBit%4 != 0 {
		for i := 0; i < 4-(numOfBit%4); i++ {
			numberInPartByte = numberInPartByte + "0"
		}
	}
	/*добавляем в начало последовательности "0"-ли, если количество бит не кратно 4
	необходимо когда первое число в 2 системе начинается с 0 (например 0010)
	тогда первые 00 не включены в число "numberOfBit"*/
	numOfBit-- //количество != индекс
	for ; numOfBit >= 0; numOfBit-- {
		temp := number.Bit(numOfBit)                                            //запоминаем значение бита
		numberInPartByte = numberInPartByte + strconv.FormatInt(int64(temp), 2) //записываем в последовательность из 4 бит дл перевода в 16.с.
		if numOfBit%4 == 0 {                                                    //когда собрали 4 бита, переводим в 16 и записываем в результат
			numberInPartByte = TwoToHex(numberInPartByte)
			result = result + numberInPartByte
			numberInPartByte = "" //обнуляем для создания новой последовательности 4 бит
		}
	}
	if End == "Little" {
		result = Invert(result)
	} //меняем последовательность байт для Little Endian
	res := "0x" + result //не забываем про 0ч в начале числа в 16 системе
	return res
}

func main() {

	message := "0xff00000000000000000000000000000000000000000000000000000000000000"
	//message := "0xaaaa000000000000000000000000000000000000000000000000000000000000"
	//message := "0xFFFFFFFF"
	//message := "0xF000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	//message := "0x0bae00"

	fmt.Printf("Value Hex: %s\n", message)
	A := new(big.Int)
	*A = HexToEndian(message, "Little")
	fmt.Printf("Hex to LE: %d\n", A)

	B := new(big.Int)
	*B = HexToEndian(message, "BIG")
	fmt.Printf("Hex to BE: %d\n", B)

	resultLEtoHex := EndianToHex(*A, "Little")
	fmt.Printf("LE to Hex: %v\n", resultLEtoHex)

	resultBEtoHex := EndianToHex(*B, "BIG")
	fmt.Printf("BE to Hex: %v\n", resultBEtoHex)
}
