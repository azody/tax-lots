Build Instructions
===========================================================
1) Have Latest Version of Go installed on your machine
   - https://go.dev/dl/
2) Navigate to Project Directory
   - Example: > cd ~/git/tax-lots
3) Build using go
   - go build tax-lots.go 

Run Instructions 
===========================================================
1) Run program: >./tax-lots {TransactionString} {AccountingMethod}
   Example: ./tax-lots 2021-01-01,buy,10000.00,1.00000000\n2021-01-02,sell,20000.00,0.50000000 fifo
2) Accounting Methods Supported: FIFO, HIFO