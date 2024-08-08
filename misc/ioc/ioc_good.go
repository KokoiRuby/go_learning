package main

import "errors"

/*
Good: Undo as control logic
*/

type Undo []func()

func (undo *Undo) Add(function func()) {
	*undo = append(*undo, function)
}

func (undo *Undo) Undo() error {
	functions := *undo
	if len(functions) == 0 {
		return errors.New("no functions to undo")
	}
	index := len(functions) - 1
	if function := functions[index]; function != nil {
		function()
		functions[index] = nil // For garbage collection
	}
	*undo = functions[:index]
	return nil
}

/*
Buz
*/

type IntSetNew struct {
	data map[int]bool
	undo Undo // dep inject
}

func NewIntSetNew() IntSetNew {
	return IntSetNew{data: make(map[int]bool)}
}

func (set *IntSetNew) Undo() error {
	return set.undo.Undo()
}

func (set *IntSetNew) Contains(x int) bool {
	return set.data[x]
}

func (set *IntSetNew) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.undo.Add(func() { set.Delete(x) })
	} else {
		set.undo.Add(nil)
	}
}

func (set *IntSetNew) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.undo.Add(func() { set.Add(x) })
	} else {
		set.undo.Add(nil)
	}
}
