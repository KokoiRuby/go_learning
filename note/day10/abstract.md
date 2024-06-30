抽象 Abstraction：把一类事物的**共有属性(字段)和方法**提取出来，**形成一个模型(模板)**。

```go
type Account struct {
	AccountNo string
	Password  string
	Balance   float64
}

func (account *Account) Deposite(money float64, pwd string) {}
func (account *Account) WithDraw(money float64, pwd string) {}
func (account *Account) Query(pwd string) {}
```

