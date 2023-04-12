package ascendex

import (
	"fmt"
	"strconv"
)

func ParseOrderBookData(bbo *BestOrderBook, v interface{}) error {
	var err error

	bbo.Ask.Price, err = strconv.ParseFloat(v.(map[string]interface{})["ask"].([]interface{})[0].(string), 64)
	if err != nil {
		return fmt.Errorf("can't parse ask price: %w", err)
	}

	bbo.Ask.Amount, err = strconv.ParseFloat(v.(map[string]interface{})["ask"].([]interface{})[1].(string), 64)
	if err != nil {
		return fmt.Errorf("can't parse ask amount: %w", err)
	}

	bbo.Bid.Price, err = strconv.ParseFloat(v.(map[string]interface{})["bid"].([]interface{})[0].(string), 64)
	if err != nil {
		return fmt.Errorf("can't parse bid price: %w", err)
	}

	bbo.Bid.Amount, err = strconv.ParseFloat(v.(map[string]interface{})["bid"].([]interface{})[1].(string), 64)
	if err != nil {
		return fmt.Errorf("can't parse bid amount: %w", err)
	}

	return nil
}
