package bank

var deposits = make(chan int)
var balances = make(chan int)
var withdraw = make(chan int)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	b := Balance()
	withdraw <- amount
	if b-amount < 0 {
		return false
	}
	return true
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraw:
			if balance >= amount {
				balance -= amount
			}
		}
	}
}

func init() {
	go teller()
}
