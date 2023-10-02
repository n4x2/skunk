package app

import (
	"flag"
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/n4x2/zoo/pass"
	"github.com/n4x2/zoo/to"
)

func generateHandler(fs *flag.FlagSet, args []string) error {
	var (
		len                  int
		number, symbol, copy bool
	)

	fs.Parse(args)

	copyFlag := fs.Lookup("copy")
	lenFlag := fs.Lookup("len")
	numberFlag := fs.Lookup("numbers")
	symbolFlag := fs.Lookup("symbols")

	if lenFlag != nil {
		pLen, err := to.Int(lenFlag.Value.String())
		if err != nil {
			return err
		}

		len = pLen
	}

	if numberFlag != nil {
		pNumber, err := to.Bool(numberFlag.Value.String())
		if err != nil {
			return err
		}

		number = pNumber
	}

	if symbolFlag != nil {
		pSymbol, err := to.Bool(symbolFlag.Value.String())
		if err != nil {
			return err
		}

		symbol = pSymbol
	}

	password, err := pass.Generate(number, symbol, len)
	if err != nil {
		return err
	}

	if copyFlag != nil {
		pCopy, err := to.Bool(copyFlag.Value.String())
		if err != nil {
			return err
		}

		copy = pCopy
	}

	if copy {
		err := clipboard.WriteAll(password)
		if err != nil {
			return err
		}

		fmt.Println("Copied into clipboard")
		return nil
	}

	fmt.Println(password)

	return nil
}
