package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var rangeOfTypicalInt = map[string]uint64{
	"MaxInt8":  255,
	"MaxInt16": 65535,
	"MaxInt32": 4294967295,
	"MaxInt64": 18446744073709551615,
}
var rangeOfBigInt = map[string]float64{
	"Max128":  math.Pow(2, 128),
	"Max256":  math.Pow(2, 256),
	"Max512":  math.Pow(2, 512),
	"Max1024": math.Pow(2, 1024),
	"Max2048": math.Pow(2, 2048),
	"Max4096": math.Pow(2, 4096),
}

//проверка соответствия вероятного ключа реальному
func IsKeyFind(gotkey string) bool {
	return strings.Compare(realValueOfKey, gotkey) == 0
}

//поиск ключа длинной 8 байт
func Key8() string {
	var foundKey string
	var tempString string
	var i uint64 = 0
	for ; i < rangeOfTypicalInt["MaxInt8"]; i++ {
		tempString = strconv.FormatUint(i, 16)
		foundKey = "0" + tempString
		if IsKeyFind(foundKey) {
			foundKey = strconv.FormatUint(i, 16)
			break
		}
		foundKey = tempString
		if IsKeyFind(foundKey) {
			foundKey = strconv.FormatUint(i, 16)
			break
		}
	}
	return foundKey
}

//поиск ключа длинной 16 байт
func Key16() string {
	var foundKey string
	var tempString string
	var i uint64 = 0
	for ; i < rangeOfTypicalInt["MaxInt16"]; i++ {
		tempString = strconv.FormatUint(i, 16)
		foundKey = "0" + tempString
		if IsKeyFind(foundKey) {
			foundKey = strconv.FormatUint(i, 16)
			break
		}
		foundKey = tempString
		if IsKeyFind(foundKey) {
			foundKey = strconv.FormatUint(i, 16)
			break
		}
	}
	return foundKey
}

//поиск ключа длинной 32 байта
func Key32() string {
	var foundKey string
	var tempString string
	var i uint64 = 0
	for ; i < rangeOfTypicalInt["MaxInt32"]; i++ {
		tempString = strconv.FormatUint(i, 16)
		foundKey = "0" + tempString
		if IsKeyFind(foundKey) {
			foundKey = strconv.FormatUint(i, 16)
			break
		}
		foundKey = tempString
		if IsKeyFind(foundKey) {
			foundKey = strconv.FormatUint(i, 16)
			break
		}
	}
	return foundKey
}

//поиск ключа длинной 64 байта
func Key64() string {
	var foundKey string
	var tempString string
	var i uint64 = 0
	for ; i < rangeOfTypicalInt["MaxInt64"]; i++ {
		tempString = strconv.FormatUint(i, 16)
		foundKey = "0" + tempString
		if IsKeyFind(foundKey) {
			foundKey = strconv.FormatUint(i, 16)
			break
		}
		foundKey = tempString
		if IsKeyFind(foundKey) {
			foundKey = strconv.FormatUint(i, 16)
			break
		}
	}
	return foundKey
}

//поиск ключа длинной более 64 байт
func KeyBig() string {
	chFind := make(chan int, 1)  //канал для поиска последних 64 байт ключа
	chBiger := make(chan int, 1) //канал для увеличения байт до последних 64
	var foundKey string          //найденное слово
	var arr []uint64             //значение вероятного ключа в каждом порядке из последовательности 64 байт
	var activeOrder int = 0      //для увеличения значения вероятного ключа в каждом порядке из последовательности 64 байт
	var i uint64 = 0             //для поиска ключа в последни 64 байтах
	var tempString string        //временный кусочек из последних 64 байт
	chFind <- 0                  //начало поиска
	for {
		select {
		case <-chBiger:

			if arr[activeOrder]+1 <= rangeOfTypicalInt["MaxInt64"] { //если число все еще можно увеличить
				arr[activeOrder] = arr[activeOrder] + 1

				foundKey = "" //перезаписываем кусок ключа до last 64b
				for k := activeOrder; k >= 0; k-- {
					foundKey = foundKey + strconv.FormatUint(uint64(arr[k]), 16)
				}

				chFind <- 0
			} else { //если необходимо переходит на следующий порядок
				arr[activeOrder] = 0
				activeOrder++
				chBiger <- 0
			}

		case <-chFind:
			activeOrder = 0 //чтобы при необходимости увеличить ключ до крайних 64б увеличивались все значения по очереди

			for ; i < rangeOfTypicalInt["MaxInt64"]; i++ { //перебор ключей
				tempString = strconv.FormatUint(i, 16)

				if IsKeyFind(foundKey + tempString) {
					foundKey = foundKey + tempString //если ключ найден, закрыть все каналы и завершить поиск
					close(chBiger)
					close(chFind)
					return foundKey
				}
			}

			i = 0
			chBiger <- 0 //если последние 64б не подобрали ключ, нужно увеличить значение ключа до last 64b
		}
	}
}

//генератор ключей
func GeneratorKey(base int) string {
	var keyValue string
	switch base {
	case 8:
		keyValue = strconv.FormatUint(uint64(rand.Intn(int(rangeOfTypicalInt["MaxInt8"]))), 16)
	case 16:
		keyValue = strconv.FormatUint(uint64(rand.Intn(int(rangeOfTypicalInt["MaxInt16"]))), 16)
	case 32:
		keyValue = strconv.FormatUint(uint64(rand.Intn(int(rangeOfTypicalInt["MaxInt32"]))), 16)
	case 64:
		keyValue = strconv.FormatUint(uint64(rand.Int63()), 16)
	case 128:
		for i := 0; i < 2; i++ {
			keyValue = keyValue + strconv.FormatUint(uint64(rand.Int63()), 16)
		}
	case 256:
		for i := 0; i < 4; i++ {
			keyValue = keyValue + strconv.FormatUint(uint64(rand.Int63()), 16)
		}
	case 512:
		for i := 0; i < 8; i++ {
			keyValue = keyValue + strconv.FormatUint(uint64(rand.Int63()), 16)
		}
	case 1024:
		for i := 0; i < 16; i++ {
			keyValue = keyValue + strconv.FormatUint(uint64(rand.Int63()), 16)
		}
	case 2048:
		for i := 0; i < 32; i++ {
			keyValue = keyValue + strconv.FormatUint(uint64(rand.Int63()), 16)
		}
	case 4096:
		for i := 0; i < 64; i++ {
			keyValue = keyValue + strconv.FormatUint(uint64(rand.Int63()), 16)
		}
	}
	return keyValue
}

var realValueOfKey string //реальный ключ
func main() {
	rand.Seed(time.Now().UnixNano())
	var foundKey string
	var j uint64
	var start time.Time
	var duration time.Duration
	for ; j <= 9; j++ {

		realValueOfKey = GeneratorKey(int(math.Pow(2, float64(j+3))))
		fmt.Printf("Real key of %vbit order: %v\n", (math.Pow(2, float64(j+3))), realValueOfKey)

		if j <= 3 {

			start = time.Now()
			if j == 0 {
				fmt.Printf("Number of options: %v\n", rangeOfTypicalInt["MaxInt8"])
				foundKey = Key8()
			}
			if j == 1 {
				fmt.Printf("Number of options: %v\n", rangeOfTypicalInt["MaxInt16"])
				foundKey = Key16()
			}
			if j == 2 {
				fmt.Printf("Number of options: %v\n", rangeOfTypicalInt["MaxInt32"])
				foundKey = Key32()
			}
			if j == 3 {
				fmt.Printf("Number of options: %v\n", rangeOfTypicalInt["MaxInt64"])
				foundKey = Key64()
			}
			duration = time.Since(start)

		} else {
			if j == 4 {
				fmt.Printf("Number of options: %v\n", rangeOfBigInt["Max128"])
			}
			if j == 5 {
				fmt.Printf("Number of options: %v\n", rangeOfBigInt["Max256"])
			}
			if j == 6 {
				fmt.Printf("Number of options: %v\n", rangeOfBigInt["Max512"])
			}
			if j == 7 {
				fmt.Printf("Number of options: %v\n", rangeOfBigInt["Max1024"])
			}
			if j == 8 {
				fmt.Printf("Number of options: %v\n", rangeOfBigInt["Max2048"])
			}
			if j == 9 {
				fmt.Printf("Number of options: %v\n", rangeOfBigInt["Max4096"])
			}
			start = time.Now()
			foundKey = KeyBig()
			duration = time.Since(start)
		}

		fmt.Printf("Needed time: %v\n", duration.Seconds())
		fmt.Printf("Found key: %s \n\n", foundKey)

	}
	fmt.Println("Congratulations! You have incredible patience!")
}
