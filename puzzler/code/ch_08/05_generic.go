package main

// generic
// interface
type List[T any] interface {
	Add(index int, val T)
	Append(val T)
	Delete(index T)
}

// struct
type Node1[T any] struct {
	val T
}

type LinkedList1[T any] struct {
	head *Node1[T]
}

func (l *LinkedList1[T]) Add(index int, val T) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList1[T]) Append(val T) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList1[T]) Delete(index T) {
	//TODO implement me
	panic("implement me")
}

// method
func Sum[T Number](vals ...T) T {
	var sum T
	for _, val := range vals {
		sum = sum + val
	}
	return sum
}

// generic constraint
type Number interface {
	int | uint
}

func main() {

}
