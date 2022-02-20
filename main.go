package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"main/config"
	"main/scanner"
	"main/token"
)

func TravelSubTokens(tokens []token.Token, r int) {
	if len(tokens) == 0 {
		r -= 1
		return
	} else {
		r += 1
	}

	for _, i := range tokens {
		for t := 1; t < r; t++ {
			fmt.Printf("\t")
		}

		fmt.Printf(
			"Type: %s\tValue:%s\r\n", i.ID, i.Value)

		TravelSubTokens(i.SubTokens, r)
	}


}

func main() {
	// data := "[\r\n 123," + `"hello"` + "\r\n, [123, \r\n, 789]]"
	data, err :=ioutil.ReadFile("test.json")
	if err != nil {
		log.Fatalln(err)
	}

	lex := scanner.NewLexer([]byte(data), config.Config{
		PoolSize: 10,
	})

	res := lex.Lex()
	fmt.Printf(
		"Type: %s\r\nValue:%s\r\n", res.ID, res.Value)

	TravelSubTokens(res.SubTokens, 1)
}
