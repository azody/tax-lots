package main

import (
	"fmt"
	"testing"
)

func TestOversold(t *testing.T) {

	transactionString := "2021-01-01,buy,100.00,2.0\n2021-01-02,buy,200.00,0.5"
	lots := processTransactions(transactionString, "fifo")

	if len(lots) != 2 {
		t.Errorf(" 1 got %q, wanted %q", len(lots), 2)
	}
	if lots[0].id != 2 {
		t.Errorf(" 2 got %q, wanted %q", lots[0].id, 2)

	}
	if lots[1].id != 1 {
		t.Errorf(" 3 got %q, wanted %q", lots[1].id, 1)

	}

	if lots[0].avgCost != 200 {
		t.Errorf(" 2 got %q, wanted %q", fmt.Sprint(lots[0].avgCost), fmt.Sprint(200))
	}

	if lots[1].quantity != 2 {
		t.Errorf(" 2 got %q, wanted %q", fmt.Sprint(lots[0].quantity), fmt.Sprint(100))
	}
}
