package lexer

import (
	"strings"
)


func tokenize(expr string) []string {
	expr = strings.Replace(expr, "(", " ( ", -1)
	expr = strings.Replace(expr, ")", " ) ", -1)
	return strings.Split(expr, "")
}


func read_tokens(tokens []string) []interface{} {
	if len(tokens) == 0 {
		return nil
	}
	var block []interface{}
	var word string
	for len(tokens) > 0 {
		token := tokens[0]
		tokens = append(tokens[:0], tokens[1:]...)

		if token == "(" {
			block = append(block, read_tokens(tokens))
		} else if token == " " {
			if len(word) > 0 {
				block = append(block, word)
				word = ""
			} else {
				continue
			}
		} else if token == "\\'" {
			var str string
			str += token
			for tokens[0] != "\\'" {
				str += tokens[0]
				tokens = append(tokens[:0], tokens[1:]...)
			}
			block = append(block, str)
		} else if token == ")" {
			return block
		} else {
			word += token
		}
	}
	return block
}

func lex(input string) []interface{} {
	return read_tokens(tokenize(input))
}