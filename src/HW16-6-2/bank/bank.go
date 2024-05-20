package bank

import (
	"errors"
	"sync"
)

var m sync.RWMutex

type Client struct {
	amount int
}

type BankClient interface {
	// Deposit deposits given amount to clients account
	Deposit(amount int)

	// Withdrawal withdraws given amount from clients account.
	// return error if clients balance less the withdrawal amount
	Withdrawal(amount int) error

	// Balance returns clients balance
	Balance() int
}

func (c *Client) Deposit(a int) {
	m.Lock()
	c.amount += a
	m.Unlock()
}

func (c *Client) Withdrawal(a int) error {
	if c.amount < a {
		return errors.New("баланс недостаточен. операция не выполнена")
	}
	m.Lock()
	c.amount -= a
	m.Unlock()
	return nil
}

func (c *Client) Balance() int {
	m.RLock()
	defer m.RUnlock()
	return c.amount
}

func New() Client {
	return Client{}
}
