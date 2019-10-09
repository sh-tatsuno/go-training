package bank

import "testing"

func TestWithDraw(t *testing.T) {
	cases := []struct {
		name    string
		amount  int
		balance int
		result  int
		canDraw bool
	}{
		{
			name:    "通常",
			balance: 300,
			amount:  100,
			result:  200,
			canDraw: true,
		},
		{
			name:    "ぴったり",
			balance: 100,
			amount:  100,
			result:  0,
			canDraw: true,
		},
		{
			name:    "下ろせない",
			balance: 100,
			amount:  200,
			result:  100,
			canDraw: false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			Deposit(c.balance)
			ok := Withdraw(c.amount)
			if ok != c.canDraw {
				t.Errorf("got %v, expected %v", ok, c.canDraw)
			}
			b := Balance()
			if b != c.result {
				t.Errorf("rest balance %d, expected %d", b, c.result)
			}
			Withdraw(Balance())
		})
	}
}
