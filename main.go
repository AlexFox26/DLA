package main

import (
	"blockchain/key"
	"blockchain/signature"
	"bufio"
	"crypto/sha1"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hello, enter the message: ")

	in:=bufio.NewReader(os.Stdin)
	message, err1:= in.ReadBytes( 10)
	if(err1 != nil){
		return
	}
	sha1.New()
	keyobj:=new(keypair.Key)
	keyobj.GenerateKeys()
	fmt.Printf("%v\n", message)
	message1 := sha1.Sum(message)
	fmt.Printf("%v\n", message1)


	r,s:=signature.SingData(keyobj, message1 )
	b:=signature.VeryfySignature(r,s, message1, keyobj)
	fmt.Printf("%v", b)
}