package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//генерирует простое число
func PrimeNumberGen() uint8 {
	//напрямую сгенерировать простое число сложно. Поэтому можно создать любое рандомное число
	//и найти ближайшее старшее простое число. Для этого можно делить рандомное число подряд на все числа начиная с 3
	//и если рандомное число делиться без остатка, то увеличить рандомное число на 2 (во избежания лишних проверок парных чисел)
	var value uint8

	value = uint8(rand.Int31n(255)) //рандомное число
	value += 10                     //если n или m будут слишком маленькими, алгоритм не сработает. Поэтому сделаем заранее достаточно большие p и q
	if value%2 == 0 {               //избавляемся от всех парных чисел
		value += 1
	}
	var iterator uint8 = 3               //начинаем делить с 3
	for ; iterator < value; iterator++ { //если мы уже делим число само на себя, то делить не нужно и число простое
		if value%iterator == 0 { //если число поделилось на другое, но наше "рандомное" число нужо увеличить
			value += 2 //увеличить на 2 чтобы не перебирать парные числа
		}
		if value <= 2 {
			return value
		}
	}
	return value
}

//поиск взаимно простого числа
func RelativPrime(m uint16) uint16 {
	//m - число, для которого нужно найти взаимно простое
	//для этого смотрим на какие числа нацело делиться наше m
	//и придумываем рандомное число до тех пор, пока не найдем то число
	//которое не делиться на цело ни на одно из тех что мы нашли
	value := m
	var dividers []uint16
	var iterator uint16 = 2
	var result uint16
	for ; iterator < value; iterator++ {
		if value%iterator == 0 {
			dividers = append(dividers, iterator)
		}
	}
	for {
		result = uint16(rand.Int31n(int32(value)))
		for _, b := range dividers {
			if result%b == 0 {
				result = 0
				break
			}
		}
		if result != 0 {
			break
		}
	}
	return result
}

//функция ассиметричного шифрования RSA
func RSA(s string) {
	message := s
	var openE int64
	var openN uint16
	var privateD uint16
	var p uint8 = PrimeNumberGen()
	var q uint8 = PrimeNumberGen()
	var m uint16
	openN = uint16(p) * uint16(q)
	m = uint16(p-1) * uint16(q-1)
	privateD = RelativPrime(m)
	openE = 1
	for ; ; openE++ { //поиски числа e из уловия (e*d)mod(m)=1

		if (openE*int64(privateD))%int64(m) == 1 {
			break
		}
	}

	fmt.Printf("Open e: %d\nOpen n:%d\nPrivate d: %d\n", openE, openN, privateD)
	//---------------------шифрование-----------------------
	var encrypt []string
	messageByte := []byte(message)
	messageByte = messageByte[0 : len(messageByte)-2]
	var tempBI1 big.Int //иза того, что числа 81^20 или подобные выходят за рамки uint64 приодится использовать big.Int
	for _, b := range messageByte {
		tempBI1.Exp(big.NewInt(int64(b)), big.NewInt(openE), big.NewInt(int64(openN))) //удобно, сразу и степень находит и по модулю делает
		encrypt = append(encrypt, tempBI1.Text(10))
	}
	fmt.Print("Encrypt text: ")
	for _, b := range encrypt {
		fmt.Printf("%s ", b)
	}
	fmt.Println()
	//-------------------------------дешифрование--------------------
	var decrypt []rune
	var temp2 big.Int
	for _, b := range encrypt {
		tempInt64, err2 := strconv.ParseInt(b, 10, 64)
		if err2 != nil {
			log.Fatalln(err2)
		}
		temp2.SetInt64(tempInt64)

		tempBI1.Exp(big.NewInt(tempInt64), big.NewInt(int64(privateD)), big.NewInt(int64(openN)))
		decrypt = append(decrypt, rune(tempBI1.Int64()))
	}

	fmt.Print("Decrypt text: ")
	for _, b := range decrypt {
		fmt.Printf("%q", b)
	}
	fmt.Println()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Print("Enter the message: ")
	in := bufio.NewReader(os.Stdin)
	message, err1 := in.ReadString('\n')
	if err1 != nil {
		log.Fatalln(err1)
	}
	RSA(message)
}
