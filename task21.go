package main

import "fmt"

type vehicle interface {
	refillTank()
}

type lada struct {
}

func (l *lada) refillTank() {
	fmt.Println("Tank is empty. Needs refilling")
}

type tesla struct {
}

func (t *tesla) rechargeBattery() {
	fmt.Println("Battery low. Insert the charger.")
}

// Tesla don't have func "refillTank", but have similar "rechargeBattery"
// So we use an adapter pattern
type teslaAdapter struct {
	_tesla *tesla
}

func (tA *teslaAdapter) refillTank() {
	tA._tesla.rechargeBattery()
}

type client struct {
}

func (c *client) notifyOnEmptyTank(v vehicle) {
	v.refillTank()
}

func task21() {
	fmt.Println("------------Task 21----------------")
	client := &client{}
	var _lada lada
	var _tesla tesla

	client.notifyOnEmptyTank(&_lada)

	_tA := teslaAdapter{
		_tesla: &_tesla,
	}
	client.notifyOnEmptyTank(&_tA)
}
