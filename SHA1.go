package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

var h0 uint64 = 0x67452301
var h1 uint64 = 0xefcdab89
var h2 uint64 = 0x98badcfe
var h3 uint64 = 0x10325476
var h4 uint64 = 0xc3d2e1f0

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

//перевод из 2 в 16
func ConvBitToHex(s string) string {
	message := s
	var result string
	var block string
	for a, b := range message {
		block = block + string(b)
		if (a+1)%4 == 0 {
			block = TwoToHex(block)
			result = result + block
			block = ""
		}
	}
	result = "0x" + result
	return result
}

//операция побитового "not"
func Not(value uint64) uint64 {
	a := strconv.FormatUint(value, 2)
	var result uint64
	b := []byte(a)
	for i, j := range b {
		if string(j) == "1" {
			b[i] = 0
		} else {
			b[i] = 1
		}
	}
	var i int = len(b) - 1
	for _, j := range b {
		if string(j) == "\x01" {
			result = result + uint64(math.Pow(2, float64(i)))
		}
		i--
	}
	return result
}

//заполняет исходное сообщение до кратность 512 битам
func Filling(s string) string {
	message := s
	var messInt, _ = new(big.Int).SetString(message, 16)             //исодное сообщение в Int формате
	lenMessageInBit := strconv.FormatInt(int64(messInt.BitLen()), 2) //длинна исходного сообщения в двоичной системе
	messageInBit := messInt.Text(2)                                  //исходное сообщение в двоичном формате
	if len(messageInBit)%4 != 0 {                                    //дополняем в начале необходимым количеством "0"
		for i := 0; len(messageInBit)%4 != 0; i++ {
			messageInBit = "0" + messageInBit
		}
	}
	lack := 512 - (len(messageInBit) % 512) //сколько бит не хватает до блока в 512 бит
	if lack >= (64 + 1) {                   //если нет смысла создавать еще один блок
		numberOfZeroNeed := lack - 1 - len(lenMessageInBit) //количество "0" необходимых для заполнения
		messageInBit = messageInBit + "1"                   //в исходное сообщение добавляем "1"
		for i := 0; i < numberOfZeroNeed; i++ {             //добавляем "0"-ли
			messageInBit = messageInBit + "0"
		}
		messageInBit = messageInBit + lenMessageInBit //в конец добавляем длинну сообщения
		messageInBit = ConvBitToHex(messageInBit)     //переводим все в Hex
		return messageInBit                           //возвращаем результат
	} else { //если нужно создавать еще один блок в 512 бит
		numberOfZeroNeed := lack - 1
		messageInBit = messageInBit + "1"
		for i := 0; i < numberOfZeroNeed+448+(64-len(lenMessageInBit)); i++ {
			messageInBit = messageInBit + "0"
		}
		messageInBit = messageInBit + lenMessageInBit
		messageInBit = ConvBitToHex(messageInBit)
		return messageInBit
	}
}

func SHA1(s string) string {
	message := s
	message = Filling(message) //заполнение исходного сообщения
	message = strings.Replace(message, "0x", "", 1)
	var chunks []string
	messageInByte := []byte(message)
	numberOfChunks := (len(message) / 128)
	for i := 0; i < numberOfChunks; i++ {
		chunks = append(chunks, string(messageInByte[i*128:(i+1)*128]))
	}

	for _, j := range chunks {
		wordsByte := []byte(j)
		var wordsString []string
		for i := 0; i < 16; i++ {
			wordsString = append(wordsString, string(wordsByte[i*8:i*8+8]))
		}
		for i := 16; i < 80; i++ {
			temp1, _ := strconv.ParseUint(wordsString[i-3], 16, 64)
			temp2, _ := strconv.ParseUint(wordsString[i-8], 16, 64)
			temp3, _ := strconv.ParseUint(wordsString[i-14], 16, 64)
			temp4, _ := strconv.ParseUint(wordsString[i-16], 16, 64)
			tempRes := (temp1 ^ temp2 ^ temp3 ^ temp4) << 5
			resString := strconv.FormatUint(tempRes, 16)
			wordsString = append(wordsString, resString)
		}

		a := h0
		b := h1
		c := h2
		d := h3
		e := h4
		var f uint64
		var k uint64
		for i := 0; i < 80; i++ {
			if 0 <= i && i <= 19 {
				f = (b & c) | ((Not(b)) & d)
				k = 0x5a827999
			}
			if 20 <= i && i <= 39 {
				f = b ^ c ^ d
				k = 0x6ed9eba1
			}
			if 40 <= i && i <= 59 {
				f = (b & c) | (b % d) | (c & d)
				k = 0x8f1bbcdc
			}
			if 60 <= i && i <= 79 {
				f = b ^ c ^ d
				k = 0xca62c1d6
			}
			tempWi, _ := strconv.ParseUint(wordsString[i], 16, 64)
			temp := (a << 5) + f + e + k + tempWi
			e = d
			d = c
			c = b << 30
			b = a
			a = temp
		}
		h0 = h0 + a
		h1 = h1 + b
		h2 = h2 + c
		h3 = h3 + d
		h4 = h4 + e
	}

	res := strconv.FormatUint((h0 << 128), 16)
	res += strconv.FormatUint((h1 << 96), 16)
	res += strconv.FormatUint((h2 << 64), 16)
	res += strconv.FormatUint((h3 << 32), 16)
	res += strconv.FormatUint(h0, 16)
	return res
}
func main() {
	fmt.Printf("Message: %s", SHA1("223442a778e997f7977c7999d8075a466e5"))
}
