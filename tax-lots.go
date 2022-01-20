package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Creating structs for object that would be seen in a db
type Transaction struct {
	tradeDate   time.Time
	tradeAction string
	price       float64
	quantity    float64
}

type Lot struct {
	id        int
	tradeDate time.Time
	avgCost   float64
	quantity  float64
}

func main() {
	transactionString := "2021-01-01,buy,10000.00,1.00000000\n2021-02-01,sell,20000.00,0.50000000"
	accountingMethod := "fifo"
	processTransactions(transactionString, accountingMethod)
	os.Exit(0)
}

func processTransactions(transactionString string, accountingMethod string) []Lot {

	/**
	* Validate and Parse Transactions
	* Optional step
	* Going straight to lot processing would be more efficient, but likely never happen in a real world application
	* In my experience, real world financial data is filled with errors
	 */
	transactions := parseTransactions(transactionString)

	/**
	* Sort Transactions so it goes in chronological order with buys before sells so buys are aggregated first
	* Examples seem to be this way however there could be cancels/corrections etc... in real data
	 */
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].tradeDate.Before(transactions[j].tradeDate)
	})

	fmt.Printf("%+v", transactions)

	/**
	* Process Lots
	* Add/Adjust Lots as needed
	* Sort Lots depending on accounting method
	 */
	lots := []Lot{}
	lotId := 1
	currentLot := Lot{ // Check if will modify after being added to array
		tradeDate: time.Time{},
	}

	for _, element := range transactions {

		// Update Lots
		if element.tradeAction == "buy" {
			if element.tradeDate.After(currentLot.tradeDate) {
				lots = append(lots, Lot{
					id:        lotId,
					tradeDate: element.tradeDate,
					avgCost:   element.price,
					quantity:  element.quantity,
				})
				lotId++
				currentLot = lots[len(lots)-1]
			} else {
				_, tempLot := updateLot(currentLot, element)

				// Can't pass by reference in go?
				for i, e := range lots {
					if e.id == currentLot.id {
						lots[i] = tempLot
						break
					}
				}
			}
			sortLots(lots, accountingMethod) // Sort lots to keep the next to be depleted lot in last index
		} else {
			for element.quantity > 0 {
				if len(lots) == 0 {
					fmt.Println("Oversold Condition, Sell Occured with Quantity Greater than Amount Held")
					os.Exit(1)
				}
				remainder, updatedLot := updateLot(lots[len(lots)-1], element)
				if remainder > 0 {
					lots = lots[:len(lots)-1] // Remove depleted lots
				} else {
					lots[len(lots)-1] = updatedLot
				}
				element.quantity = remainder // Repeat Transaction with new quantity
			}
			// Don't need to resort on sell as only change would be potential removal of final element
		}

	}
	fmt.Printf("\n----------------------------------------------------------\n")
	for _, element := range lots {
		fmt.Printf("%+v \n", element)
	}
	return lots
}

func parseTransactions(transactionString string) []Transaction {

	transactions := []Transaction{}
	transactionStrings := strings.Split(transactionString, "\n")

	// Formatting and Validation for Transactions
	for _, element := range transactionStrings {
		// index is the index where we are
		// element is the element from someSlice for where we are
		currentStringArray := strings.Split(element, ",")

		// Validate Time
		t, err := time.Parse("2006-01-02", currentStringArray[0])

		if err != nil {
			fmt.Println("Invalid Date Given " + currentStringArray[0])
			os.Exit(1)
		}

		// Validate Trade Action
		if currentStringArray[1] != "buy" && currentStringArray[1] != "sell" {
			fmt.Println("Transaction Type {} not supported", currentStringArray[1])
			os.Exit(1)
		}

		// Validate Price
		p, err := strconv.ParseFloat(currentStringArray[2], 64)
		if err != nil {
			fmt.Println("Invalid Price Given: " + currentStringArray[2])
			os.Exit(1)
		}

		if p < 0 {
			fmt.Println("Negative Price Given for a Transaction: " + currentStringArray[2])
			os.Exit(1)
		}

		// Validate Quantity
		q, err := strconv.ParseFloat(currentStringArray[3], 64)
		if err != nil {
			fmt.Println("Invalid Quantity Given: " + currentStringArray[3])
			os.Exit(1)
		}
		if q < 0 {
			fmt.Println("Negative Quantity Given for a Transaction: " + currentStringArray[3])
			os.Exit(1)
		}

		transactions = append(transactions,
			Transaction{
				tradeDate:   t,
				tradeAction: currentStringArray[1],
				price:       p,
				quantity:    q,
			})
	}
	return transactions
}

/**
* Sort lots based on accounting method
* Final element is the first to be taken down (First in, highest avg cost)
 */
func sortLots(lots []Lot, accountingMethod string) {

	if accountingMethod == "fifo" {
		sort.Slice(lots, func(i, j int) bool {
			return lots[i].tradeDate.After(lots[j].tradeDate)
		})
	} else if accountingMethod == "hifo" {
		sort.Slice(lots, func(i, j int) bool {
			return lots[i].avgCost < lots[j].avgCost
		})
	}
}

/**
* returns remaining quantity
 */
func updateLot(currentLot Lot, transaction Transaction) (float64, Lot) {

	if transaction.tradeAction == "buy" {
		quantity := currentLot.quantity + transaction.quantity
		avgCost := ((currentLot.quantity * currentLot.avgCost) + (transaction.quantity * transaction.price)) / quantity
		currentLot.quantity = quantity
		currentLot.avgCost = avgCost
		return 0, currentLot
	} else {
		quantity := currentLot.quantity - transaction.quantity
		if quantity > 0 {
			currentLot.quantity = quantity
			return 0, currentLot
		} else {
			currentLot.quantity = 0
			return -quantity, currentLot
		}

	}

}
