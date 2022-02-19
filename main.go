package main

import (
	"fmt"
	"main/config"
	"main/scanner"
	"main/token"
)

func TravelSubTokens(tokens []token.Token, r int) {
	if len(tokens) == 0 {
		return
	}

	for _, i := range tokens {
		for t := 0; t < r; t++ {
			fmt.Printf("\t")
		}

		fmt.Printf(
			"Type: %s\tValue:%s\r\n", i.ID, i.Value)

		TravelSubTokens(i.SubTokens, r)
	}

	r += 1
}

func main() {
	data := `[[123, false],[123, "test string", null, true]]`
	lex := scanner.NewLexer([]byte(data), config.Config{
		PoolSize: 10,
	})

	res := lex.Lex()
	fmt.Printf(
		"Type: %s\r\nValue:%s\r\n", res.ID, res.Value)

	TravelSubTokens(res.SubTokens, 1)
}
