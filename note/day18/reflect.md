反射 Reflection 是指在程序**运行期**对程序本身进行**访问和修改**的能力。

正常情况下，程序在编译时，源码变量被转换为内存地址，变量名不会被编译器写入到可执行部分，运行中程序无法获取自身的信息。

支持反射的语言可以在程序编译期将变量的反射信息，如**字段名称、类型信息、结构体信息**等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们。



![img](https://img.halfrost.com/Blog/ArticleImage/148_6_0.png)



:confused: 空接口可以存储任何类型的变量，**如何知道这个空接口保存的数据的类型？** 断言 or 反射

```go
var v interface{}

t := reflect.TypeOf(v)     // interface{} → reflect.Type
rVal := reflect.ValueOf(v) // interface{} → reflect.Value
rVal.Kind()                // reflect.Value.Kind()
rVal.Type()                // reflect.Value.Type()

iVal := rVal.interface()   // reflect.Value → interface{}
orig := iVal.(type)        // interface{} → type
```

如果要通过反射来修改变量，注意当使用 SetXxx 方法来设置需要**通过对应的指针类型**来完成，这样才能改变传入的变量的值，同时需要使用到 `redlect.Value.Elem()` 方法

```go
str := "Tom"
rfv := reflect.ValueOf(&str)
rfv.Elem().SetString("Jack")
fmt.Println(str)
```

遍历结构体字段，调用结构体方法 (ordered by ASCII dict order)，获取结构体标签。

- Type：表示一个值的**动态**类型，提供详细的类型信息；对于结构体，还能获取其字段和方法信息。
  - `NumField`：返回结构体成员的数量。
  - `Field`：返回结构体成员的信息。
  - `FieldByIndex`：返回结构体成员的信息。
  - `FieldByName`：返回结构体成员的信息。
  - `FieldByNameFunc`：返回结构体成员的信息。
- Kind：表示一个值的**静态**类型。无法提供详细信息。
  - `String`：反射值的静态类型的名称

```go
t := reflect.TypeOf(i)
t.Name()            // dynamic type
t.Kind().String()   // staic  type
if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
    // TODO
}

val := reflect.ValueOf(i)
val         // dynamic value
val.Elem()  // static value

// field & method count
fieldNum := t.Elem().NumField()
methodNum := t.Elem().NumMethod()

// iter filed & tag
for i := 0; i < fieldNum; i++ {
	fmt.Println(t.Elem().Field(i))
	tagVal := t.Field(i).Tag.Get("json")
	if tagVal != "" {
		fmt.Println(tagVal)
	}
}

// param & return count of method
t.Elem().Method(0).Type.NumIn()
t.Elem().Method(0).Type.NumOut()

// if had method
_, ok := t.Elem().MethodByName("Name")
if ok {
	fmt.Println("Found")
} 

// call method with/without arg
val.Method(1).Call(nil)

var params []reflect.Value
params = append(params, reflect.ValueOf(10))
params = append(params, reflect.ValueOf(20))
res := val.Method(0).Call(params)
fmt.Println(res[0].Int())
```

> 如何理解动态/静态类型 & 动态/静态值？

- **静态类型** 是在编译时已知的类型，用于编译时类型检查。
- **动态类型** 是在运行时确定的类型，用于接口变量及反射。
- **静态值** 是在编译时已知的值。
- **动态值** 是在运行时确定的值。

