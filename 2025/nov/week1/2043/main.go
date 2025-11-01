package main

import (
	"fmt"
)

type Bank struct {
	BalancePerAccount map[int]int64
}

func Constructor(balance []int64) Bank {
	balancePerAccount := make(map[int]int64)
	for account, money := range balance {
		balancePerAccount[account + 1] = money
	}
	return Bank{BalancePerAccount: balancePerAccount}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if _, exists := this.BalancePerAccount[account2]; !exists {
		return false
	}

	if _, exists := this.BalancePerAccount[account2]; !exists {
		return false
	}

	if this.BalancePerAccount[account1] < money {
		return false
	}

	this.BalancePerAccount[account1] -= money
	this.BalancePerAccount[account2] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if _, exists := this.BalancePerAccount[account]; !exists {
		return false
	}

	this.BalancePerAccount[account] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if _, exists := this.BalancePerAccount[account]; !exists {
		return false
	}

	if this.BalancePerAccount[account] < money {
		return false
	}

	this.BalancePerAccount[account] -= money
	return true
}

func main() {
	// result: []
	// balance := []int64{}
	// operations := [][]int{}

	// result: [null, true, true, true, false, false]
	balance := []int64{10, 100, 20, 50, 30}
	operations := [][]int{
		{2, 3, 10},
		{0, 5, 1, 20},
		{1, 5, 20},
		{0, 3, 4, 15},
		{2, 10, 50},
	}

	// result: []
	// balance := []int64{}
	// operations := [][]int{}

	obj := Constructor(balance)

	for _, ope := range operations {
		if ope[0] == 0 {
			account1 := ope[1]
			account2 := ope[2]
			money := ope[3]
			fmt.Printf("obj.Transfer(%d, %d, %d) = %t\n", account1, account2, money, obj.Transfer(account1, account2, int64(money)))
		} else if ope[0] == 1 {
			account := ope[1]
			money := ope[2]
			fmt.Printf("obj.Deposit(%d, %d) = %t\n", account, money, obj.Deposit(account, int64(money)))
		} else if ope[0] == 2 {
			account := ope[1]
			money := ope[2]
			fmt.Printf("obj.Withdraw(%d, %d) = %t\n", account, money, obj.Withdraw(account, int64(money)))
		}
	}
}

