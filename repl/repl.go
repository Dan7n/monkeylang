package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkeylang/lexer"
	"monkeylang/token"
	"os"
	"os/signal"
)

const PROMPT = ">>> "

func Init(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			fmt.Print("\nThank you for using our REPL\n")
			os.Exit(1)
		}
	}()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lxr := lexer.New(line)

		for tok := lxr.NextToken(); tok.Type != token.EOF; tok = lxr.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}

}
