//go:build solution

package mutexes

import (
	"maps"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func NewContainer() *Container {
	return &Container{counters: map[string]int{}}
}

func (c *Container) Inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func (c *Container) Get(name string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counters[name]
}

func (c *Container) Snapshot() map[string]int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return maps.Clone(c.counters)
}

func IncMany(c *Container, name string, goroutines, n int) {
	var wg sync.WaitGroup
	for range goroutines {
		wg.Go(func() {
			for range n {
				c.Inc(name)
			}
		})
	}
	wg.Wait()
}

type Account struct {
	mu      sync.Mutex
	balance int
}

func NewAccount(balance int) *Account {
	return &Account{balance: balance}
}

func (a *Account) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

func (a *Account) Withdraw(amount int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if amount > a.balance {
		return false
	}
	a.balance -= amount
	return true
}

func (a *Account) Balance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

type Cache struct {
	mu     sync.Mutex
	values map[string]int
}

func NewCache() *Cache {
	return &Cache{values: map[string]int{}}
}

func (c *Cache) GetOrCompute(key string, compute func() int) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v, ok := c.values[key]; ok {
		return v
	}
	v := compute()
	c.values[key] = v
	return v
}
