package opcodes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadOpcodes() {

	fileContent, _ := os.Open("etc/opcodes.json")

	/*if err != nil {
		log.Fatal(err)
		return a
	}*/

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var prefixed Prefixed

	json.Unmarshal(byteResult, &prefixed)

	for i := 0; i < len(prefixed.Instructions); i++ {
		fmt.Println("Mnemonic: " + prefixed.Instructions[i].mnemonic)
	}
}
