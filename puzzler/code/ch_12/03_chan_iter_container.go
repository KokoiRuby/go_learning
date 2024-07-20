package main

import (
	"fmt"
	"time"
)

type item struct {
	ID   int
	Name string
}

type container struct {
	Items []item
}

func (c *container) Len() int {
	return len(c.Items)
}

func (c *container) Iter() <-chan item {
	ch := make(chan item)
	go func() {
		for i := 0; i < c.Len(); i++ {
			ch <- c.Items[i]
			time.Sleep(time.Second)
		}
		close(ch)
	}()
	return ch
}

func main() {
	c := &container{
		Items: []item{
			{ID: 1, Name: "Item 1"},
			{ID: 2, Name: "Item 2"},
			{ID: 3, Name: "Item 3"},
		},
	}

	for x := range c.Iter() {
		fmt.Printf("Received: %+v\n", x)
	}

	fmt.Println("Done")
}
