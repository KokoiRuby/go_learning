封装：将数据和操作封装在一个类中，通过限制对内部数据和方法的访问，隐藏了实现的细节。封装提供了数据的安全性和保护性，外部只能通过公共接口访问数据和方法。

seter/getter-like

```go
// /encapsulation/model/person.go
package modedl

import {
    "fmt"
}

type person struct {
    Name string
    age int
    sal float64
}

// factory as constructor
func newPerson(name String) *person{
    return &person{
        Name: name,
    }
}

// seter/getter for pri fields
func (p *person) SetAge(age int) {
    p.age = aget
}

func (p *person) GetAge() int {
    return p.age
}

func (p *person) SetSal(sal float64) {
    p.sal = sal
}

func (p *person) GetSal() float64 {
    return p.sal
}
```

```go
// /encapsulation/main.go
package main
import (
	"fmt"
	"encapsulate/model"
)

func main() {
    p := model.NewPerson("Jack")
    p.SetAge(18)
    p.SetSal(5000)
    fmt.Println(p)
    fmt.Println(p.GetAge(), pGetSal())
}
```

