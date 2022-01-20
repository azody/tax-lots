package main

import (
	"fmt"
	"testing"
)

/**
* Basic Sanity Checks
 */

func Test2BuysAndSellFIFO(t *testing.T) {

	transactionString := "2021-01-01,buy,100.00,2.0\n2021-01-02,buy,200.00,0.5\n2021-01-03,sell,100.00,1.0"
	lots := processTransactions(transactionString, "fifo")

	if lots[1].quantity != 1 {
		t.Errorf("got %q, wanted %q", fmt.Sprint(lots[0].quantity), fmt.Sprint(1))
	}
}

func Test2BuysAndSellHIFO(t *testing.T) {

	transactionString := "2021-01-01,buy,500.00,2.0\n2021-01-02,buy,200.00,0.5\n2021-01-03,sell,100.00,1.0"
	lots := processTransactions(transactionString, "fifo")

	if lots[1].quantity != 1 {
		t.Errorf("got %q, wanted %q", fmt.Sprint(lots[0].quantity), fmt.Sprint(1))
	}
}
