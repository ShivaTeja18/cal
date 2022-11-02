package main

import (
	"cal/jMU"
	"cal/mqsetup"
)

func main() {
	mqsetup.MqSetup()
	//	go Consumer()
	jMU.Marshalling()
	mqsetup.Read()

}
