package main

import (
	"fmt"
	"testing"
)

/**
* Basic Sanity Checks
 */
func Test2BuysFIFO(t *testing.T) {

	transactionString := "2021-01-01,buy,100.00,2.0\\n2021-01-02,buy,200.00,0.5"
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

func Test2BuysHIFO(t *testing.T) {

	transactionString := "2021-01-01,buy,100.00,2.0\\n2021-01-02,buy,200.00,0.5"
	lots := processTransactions(transactionString, "fifo")

	if len(lots) != 2 {
		t.Errorf(" 1 got %q, wanted %q", len(lots), 2)
	}
	if lots[0].id != 2 {
		t.Errorf(" 2 got %q, wanted %q", lots[0].id, 2)

	}
	if lots[1].id != 1 {
		t.Errorf("got %q, wanted %q", lots[1].id, 1)

	}

	if lots[0].avgCost != 200 {
		t.Errorf("got %q, wanted %q", fmt.Sprint(lots[0].avgCost), fmt.Sprint(200))
	}

	if lots[1].quantity != 2 {
		t.Errorf("got %q, wanted %q", fmt.Sprint(lots[0].quantity), fmt.Sprint(100))
	}
}

/**
* Test Lot Accumulation on same day Buy
 */
func Test2BuysSameDay(t *testing.T) {
	transactionString := "2021-01-01,buy,100.00,2.0\\n2021-01-01,buy,400.00,1.0"
	lots := processTransactions(transactionString, "fifo")

	if len(lots) != 1 {
		t.Errorf("got %q, wanted %q", len(lots), 1)
	}

	if lots[0].avgCost != 200 {
		t.Errorf("got %q, wanted %q", fmt.Sprint(lots[0].avgCost), fmt.Sprint(125))
	}

	if lots[0].quantity != 3 {
		t.Errorf("got %q, wanted %q", fmt.Sprint(lots[0].quantity), fmt.Sprint(3))
	}
}
