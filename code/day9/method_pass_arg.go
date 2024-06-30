package main

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

func main() {
	rect := Rectangle{width: 1, height: 1}
	area := rect.Area()
	println(area)
	rect.Scale(2)
	area = rect.Area()
	println(area)

}
