结构体名首字母是大写的可以在其他包使用这个结构体，如果**希望小写/私有的也能在其他包使用**，此时就需要工厂模式来解决。

Factory **工厂模式**可以实现**跨包**初始化结构体实例。

```go
factory
└── main
    └── main.go
└── model
    ├── student_pub.go
    └── student_pri.go
```

```go
// /factory/main/main.go
package main

import {
    "fmt"
    "/factory/model"
}

func main() {
    stu_pub := model.Student {
        Name: "Tom",
        Score: "89.0",
    }
    fmt.Println(stu_pub)
    
    stu_pri := model.newStudent("Tom", 89.0)
    fmt.Println(stu_pri)
    fmt.Println(stu_pri.GetScore())
}

```

```go
// /factory/model/student_pub.go
package model

type Student struct {
    Name string
    Score float64
}

```

```go
// /factory/model/student_pri.go
package model

type student struct {
    name string
    score float64
}

func newStudent(n string, s float64) *student {
    return &student {
        Name: n,
        Score: s,
    }
}

func (s *student) GetScore() float64{
    return (s.score) 
}
```



