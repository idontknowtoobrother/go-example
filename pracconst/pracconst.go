package pracconst

import (
	"fmt"
	"log"
)

func Example() {

	log.SetPrefix("pracconst: ")
	const cantEditVarable int = 13
	infoMsg := fmt.Sprintf("cantEditVarable=%d (can't change, edit and reassing variables. it's usage such as config etc..)", cantEditVarable)
	log.Println()
	log.Println(infoMsg)
}
