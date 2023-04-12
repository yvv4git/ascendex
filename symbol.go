package ascendex

import (
	"regexp"
	"strings"
)

func SymbolPrepare(symbol string) (string, error) {
	re := regexp.MustCompile(`^[A-Z]+\_[A-Z]+$`)
	if !re.Match([]byte(symbol)) {
		return "", ErrInvalidSymbolValue
	}

	symbol = strings.Replace(symbol, "_", "/", -1)

	return symbol, nil
}
