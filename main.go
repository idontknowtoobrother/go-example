package main

import (
	"fmt"

	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
	"github.com/idontknowtoobrother/go-example/hexafu"
	"github.com/idontknowtoobrother/go-example/pracconst"
)

func main() {

	fullname := "Jakkrit"
	var lastname string = "Chaopron"

	id := uuid.New()

	fmt.Printf("FULLNAME=%s\nLASTNAME=%s\nMACADDRESS=%v\nUUID=%v", fullname, lastname, faker.MacAddress(), id)

	fmt.Printf("\nGetRandomInt=%d", hexafu.GetRandomInt(100))

	fullname = "Chaopron"
	fmt.Printf("\nNEW_FULLNAME=%s", fullname)

	pracconst.Example()
}
