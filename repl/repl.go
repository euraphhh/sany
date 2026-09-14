package repl

import (
	"bufio"
	"fmt"
	"io"
	"sany/lexer"
	"sany/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		
		// Criamos um novo Lexer com a linha digitada pelo usuário
		l := lexer.New(line)

		// Lemos todos os tokens e imprimimos na tela até o fim do texto
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
