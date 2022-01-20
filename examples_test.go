package main

import (
	"fmt"
	"testing"
	"time"
)

func TestExample1(t *testing.T) {
	transactionString := "2021-01-01,buy,10000.00,1.00000000\n2021-02-01,sell,20000.00,0.50000000"
	lots := processTransactions(transactionString, "fifo")

	t1, _ := time.Parse("2006-01-02", "2021-01-01")

	if len(lots) != 1 {
		t.Errorf("got %q, wanted %q", len(lots), 1)
	}

	if lots[0].id != 1 {
		t.Errorf("got %q, wanted %q", lots[0].id, 1)
	}

	if lots[0].avgCost != 10000 {
		t.Errorf("got %q, wanted %q", fmt.Sprint(lots[0].avgCost), fmt.Sprint(10000))
	}

	if lots[0].quantity != 0.5 {
		t.Errorf("got %q, wanted %q", fmt.Sprint(lots[0].quantity), fmt.Sprint(0.5))
	}

	if !lots[0].tradeDate.Equal(t1) {
		t.Errorf(" 1 got %q, wanted %q", lots[0].tradeDate, t1)
	}

}

func TestExample2(t *testing.T) {
	transactionString := "2021-01-01,buy,10000.00,1.00000000\n2021-01-02,buy,20000.00,1.00000000\n2021-02-01,sell,20000.00,1.50000000"
	lots := processTransactions(transactionString, "fifo")

	t1, _ := time.Parse("2006-01-02", "2021-01-02")

	if len(lots) != 1 {
		t.Errorf("got %q, wanted %q", len(lots), 1)
	}

	if lots[0].id != 2 {
		t.Errorf("got %q, wanted %q", lots[0].id, 2)
	}

	if lots[0].avgCost != 20000 {
		t.Errorf("got %q, wanted %q", fmt.Sprint(lots[0].avgCost), fmt.Sprint(20000))
	}

	if lots[0].quantity != 0.5 {
		t.Errorf("got %q, wanted %q", fmt.Sprint(lots[0].quantity), fmt.Sprint(0.5))
	}

	if !lots[0].tradeDate.Equal(t1) {
		t.Errorf(" 1 got %q, wanted %q", lots[0].tradeDate, t1)
	}
}

func TestExample3(t *testing.T) {
	transactionString := "2021-01-01,buy,10000.00,1.00000000\n2021-01-02,buy,20000.00,1.00000000\n2021-02-01,sell,20000.00,1.50000000"
	lots := processTransactions(transactionString, "hifo")

	t1, _ := time.Parse("2006-01-02", "2021-01-01")

	if len(lots) != 1 {
		t.Errorf("got %q, wanted %q", len(lots), 1)
	}

	if lots[0].id != 1 {
		t.Errorf("got %q, wanted %q", lots[0].id, 1)
	}

	if lots[0].avgCost != 10000 {
		t.Errorf("got %q, wanted %q", fmt.Sprint(lots[0].avgCost), fmt.Sprint(10000))
	}

	if lots[0].quantity != 0.5 {
		t.Errorf("got %q, wanted %q", fmt.Sprint(lots[0].quantity), fmt.Sprint(0.5))
	}

	if !lots[0].tradeDate.Equal(t1) {
		t.Errorf(" 1 got %q, wanted %q", lots[0].tradeDate, t1)
	}

}
