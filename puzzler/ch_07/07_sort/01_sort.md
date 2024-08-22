go 的排序操作需要该类型实现 `Interface` 接口

```go
type Interface interface {
    // return # of elem in the collection
    Len() int
    // if collection[i] < collection[i]
    Less(i, j int) bool
    // swap collection[i] & collection[j]
    Swap(i, j int)
}
```

已实现该接口的类型（重定义）。

```go
type IntSlice []int
type Float64Slice []float64
type StringSlice []string

// extra methods
Sort()
Search()
```

```go
// sort functions
func Ints(a []int)
func IntsAreSorted(a []int) bool
func SearchInts(a []int, x int) int

func Float64s(a []float64)
func Float64sAreSorted(a []float64) bool
func SearchFloat64s(a []float64, x float64) int
```

排序默认 Ascending。降序

```go
sort.Sort(sort.Reverse(sort.IntSlice(intList)))
sort.Sort(sort.Reverse(sort.Float64Slice(float64List)))
sort.Sort(sort.Reverse(sort.StringSlice(stringList)))
```

**结构体类型的排序**

1. 通过使用 `sort.Sort(sl)` 实现的，只要 sl 实现了 `sort.Interface` 方法即可。

:smile: 简单。

:cry: 只能根据单一字段比较。

```go
type Person struct {
	Name string
	Age  int
}

type Persons []Person

func (p Persons) Len() int {
	return len(p)
}

func (p Persons) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
```

2. 动态 less 方法

:smile: 支持不同字段排序。

:cry: 复杂。

```go
type Cat struct {
	Name string
	Age  int
}

type CatWrapper struct {
	Cats []Cat
	by   func(p1, p2 *Cat) bool
}

func (cw *CatWrapper) Len() int {
	return len(cw.Cats)
}

// Less is dynamic
func (cw *CatWrapper) Less(i, j int) bool {
	return cw.by(&cw.Cats[i], &cw.Cats[j])
}

func (cw *CatWrapper) Swap(i, j int) {
	cw.Cats[i], cw.Cats[j] = cw.Cats[j], cw.Cats[i]
}

// encap
type SortBy func(p, q *Cat) bool

func SortCats(cats []Cat, by SortBy) {
	sort.Sort(&CatWrapper{cats, by})
}

// call
SortCats(cats, func(p, q *Cat) bool {
	return p.Age < q.Age
})
```

3. 预定义

```go
type Fruit struct {
	Name   string
	Weight int
}

type Fruits []Fruit

func (f Fruits) Len() int {
	return len(f)
}

func (f Fruits) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// enCap into ByName

type ByName struct{ Fruits }

func (b ByName) Less(i, j int) bool {
	return b.Fruits[i].Name < b.Fruits[j].Name
}

// enCap to Weight

type ByWeight struct{ Fruits }

func (b ByWeight) Less(i, j int) bool {
	return b.Fruits[i].Weight < b.Fruits[j].Weight
}

// call
sort.Sort(ByName{fruits})
sort.Sort(ByWeight{fruits})
```

**复杂结构排序**

`[][]int`

```go
type matrix [][]int

func (m matrix) Len() int {
	return len(m)
}

func (m matrix) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m matrix) Less(i, j int) bool {
	return m[i][1] < m[j][1]   // by row
}
```

`[]map[string]int`

```go
type mapSl []map[string]int

func (msl mapSl) Len() int {
	return len(msl)
}
func (msl mapSl) Swap(i, j int) {
	msl[i], msl[j] = msl[j], msl[i]
}
func (msl mapSl) Less(i, j int) bool {
	return msl[i]["a"] < msl[j]["a"]
}
```

**稳定排序**，不改变元素的相对位置。

```go
sort.Stable
```

