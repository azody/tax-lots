package main

import (
	"fmt"
	"testing"
	"time"
)

func TestFIFOSort(t *testing.T) {

	t1, _ := time.Parse("2006-01-02", "2021-01-01")
	t2, _ := time.Parse("2006-01-02", "2021-01-03")
	t3, _ := time.Parse("2006-01-02", "2021-01-03")
	t4, _ := time.Parse("2006-01-02", "2021-01-04")

	lots := []Lot{{tradeDate: t2}, {tradeDate: t3}, {tradeDate: t4}, {tradeDate: t1}}

	sortLots(lots, "fifo")

	if !lots[3].tradeDate.Equal(t1) {
		t.Errorf(" 1 got %q, wanted %q", lots[3].tradeDate, t1)
	}
	if !lots[0].tradeDate.Equal(t4) {
		t.Errorf(" 3 got %q, wanted %q", lots[3].tradeDate, t4)
	}
	if !lots[1].tradeDate.Equal(t2) {
		t.Errorf(" 2 got %q, wanted %q", lots[2].tradeDate.String(), t2.String())
	}

}

func TestHIFOSort(t *testing.T) {

	a1 := 100.0
	a2 := 200.0
	a3 := 200.0
	a4 := 400.0

	lots := []Lot{{avgCost: a2}, {avgCost: a3}, {avgCost: a4}, {avgCost: a1}}

	sortLots(lots, "hifo")

	if lots[0].avgCost != a1 {
		t.Errorf("got %q, wanted %q", fmt.Sprint(lots[0].avgCost), fmt.Sprint(a1))
	}
	if lots[2].avgCost != a3 {
		t.Errorf("got %q, wanted %q", fmt.Sprint(lots[2].avgCost), fmt.Sprint(a3))
	}
	if lots[3].avgCost != a4 {
		t.Errorf("got %q, wanted %q", fmt.Sprint(lots[3].avgCost), fmt.Sprint(a4))
	}
}
