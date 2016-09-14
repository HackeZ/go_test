package main

import (
	"container/list"
	"errors"
	"log"
	"sync"
)

var (
	ErrNoMoreObject = errors.New("No More Object in Pool")
	ErrNoMoreSpace  = errors.New("No More Space in Pool")
)

// Object ..
type Object struct {
	ID int
}

// Pool a pool
type Pool interface {
	Borrow() (*Object, error)
	Return(*Object) error
}

// Allocate return index to Object.
type Allocate func(int) (*Object, error)

type objectPool struct {
	Size     int
	Allocate Allocate
	FreeList *list.List
	mu       sync.Mutex
}

func (p *objectPool) Borrow() (*Object, error) {
	// Lock It.
	p.mu.Lock()
	elem := p.FreeList.Front()
	if elem == nil {
		return nil, ErrNoMoreObject
	}

	obj := p.FreeList.Remove(elem)

	// Unlock It.
	p.mu.Unlock()

	o := obj.(*Object)
	log.Println("A Object Borrow in Pool")

	return o, nil
}

func (p *objectPool) Return(obj *Object) error {
	if p.Size == p.FreeList.Len() {
		return ErrNoMoreSpace
	}

	p.mu.Lock()
	p.FreeList.PushBack(obj)
	p.mu.Unlock()

	log.Println("A Object Return in Pool")
	return nil
}

// New return a available Pool
func New(initPoolSize int, alloc Allocate) (Pool, error) {
	op := &objectPool{
		Size:     initPoolSize,
		Allocate: alloc,
		FreeList: list.New(),
	}

	var obj *Object
	var err error
	for i := 0; i < initPoolSize; i++ {
		obj, err = op.Allocate(i)
		if err != nil {
			return nil, err
		}

		op.FreeList.PushFront(obj)
	}
	return op, nil
}

func main() {
	const poolSize = 3

	p, _ := New(poolSize, func(i int) (*Object, error) {
		return &Object{ID: i}, nil
	})

	// Borrow a Object to Test.
	ob, err := p.Borrow()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("borrow a object from pool: %#v\n", *ob)

	// Borrow All Object in Pool
	for i := 0; i < poolSize-1; i++ {
		o, err := p.Borrow()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("borrow a object from pool: %#v\n", *o)
	}

	p.Return(ob)
	ob1, err := p.Borrow()
	if err != nil {
		log.Fatal(err)
	}

	if ob != ob1 {
		log.Fatal("expect the same object")
	}
}
