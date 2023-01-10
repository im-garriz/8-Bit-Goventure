package opcodes

type Operand struct {
	immediate bool   `json:immediate`
	name      string `json:operands`
	bytes     uint8  `json:operands`
	value     int16  `json:operands`
	adjust    string `json:operands`
}

type Instruction struct {
	opcode    uint8
	immediate bool      `json:immediate`
	operands  []Operand `json:operands`
	cycles    []uint8   `json:cycles`
	bytes     uint8     `json:bytes`
	mnemonic  string    `json:mnemonic`
	comment   string
}

type Operands struct {
	Operands []Operand `json:operands`
}

type Unprefixed struct {
	Instructions []Instruction `json:unprefixed`
}

type Prefixed struct {
	Instructions []Instruction `json:prefixed`
}

/*func ReadOpcodes() ([]Instruction, []Instruction) {

	fileContent, nil := os.Open("etc/opcodes.json")

	if err != nil {
		log.Fatal(err)
		return a
	}

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var prefixed Prefixed

	json.Unmarshal(byteResult, &prefixed)

	for i := 0; i < len(prefixed.Instructions); i++ {
		fmt.PrintLn("Mnemonic: " + prefixed.Instructions[i].mnemonic)
	}

	return prefixed, prefixed
}*/
