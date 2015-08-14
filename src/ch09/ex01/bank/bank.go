package bank

type withdrawReq struct {
	amount  int
	resultc chan<- bool
}

var (
	depc      = make(chan int)          // send amount to deposit
	balc      = make(chan int)          // receive balance
	withdrawc = make(chan *withdrawReq) // withdraw
)

func Deposit(amount int) { depc <- amount }
func Balance() int       { return <-balc }

func Withdraw(amount int) bool {
	resultc := make(chan bool)
	var req = withdrawReq{amount, resultc}
	withdrawc <- &req
	return <-resultc
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-depc:
			balance += amount
		case balc <- balance:
		case req := <-withdrawc:
			if balance >= req.amount {
				balance -= req.amount
				req.resultc <- true
			} else {
				req.resultc <- false
			}
		}
	}
}

func init() { go teller() }
