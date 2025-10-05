package storage

import (
	"errors"
	"sync"
)

type Employee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

type Storage interface {
	Insert(e *Employee)
	Delete(id int)
	Update(id int, e Employee)
	Get(id int) (Employee, error)
}

type MemoryStorage struct {
	counter int
	hashmap map[int]Employee
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{counter: 1, hashmap: make(map[int]Employee)}
}

func (ms *MemoryStorage) Insert(e *Employee) {
	ms.Lock()
	e.ID = ms.counter
	ms.hashmap[e.ID] = *e
	ms.counter++
	ms.Unlock()
}

func (ms *MemoryStorage) Delete(id int) {
	ms.Lock()
	delete(ms.hashmap, id)
	ms.Unlock()
}

func (ms *MemoryStorage) Update(id int, e Employee) {
	ms.Lock()
	ms.hashmap[id] = e
	ms.Unlock()
}

func (ms *MemoryStorage) Get(id int) (Employee, error) {
	ms.Lock()
	defer ms.Unlock()
	e, err := ms.hashmap[id]
	if !err {
		return Employee{}, errors.New("this employee doesn't exist")
	}
	return e, nil
}
