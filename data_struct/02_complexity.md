:confused: **Why?** 事后统计(运行/监控/统计)局限：依赖测试环境 (硬件影响)，受数据规模影响

#### Big O Notation

假设每行代码执行的时间都一样，为 unit_time。所有代码的执行时间 T(n) 与每行代码的执行次数 n 成正比。**行数越多，每行执行次数越多**，T(n) 越大。

公式：$$T(n) = O(f(n))$$ `T(n)` 表示代码执行的时间，`n` 表示数据规模的大小，`f(n)` 表示每行代码执行的次数总和。`O` 表示成正比。

Big O 并不具体表示代码真正的执行时间，而是**表示代码执行时间随数据规模增长的变化趋势** → 时间复杂度 Time Complexity

```python
# (2+2n) * unit_time  → O(2n+2) → O(n)
int cal(int n) {
	int sum = 0;            # 1
	int i = 1;              # 1
	for (; i <= n; ++i) {   # n
		sum = sum + i;      # n
	}
	return sum;
}
```

```python
# (3+2n*2n^2) * unit_time → O(2n^2+2n+3) → O(n^2)
int cal(int n) {
    int sum = 0;   # 1
    int i = 1;     # 1
    int j = 1;     # 1
    for (; i <= n; ++i) {       # n
    	j = 1;                  # n
    	for (; j <= n; ++j) {   # n * n
    		sum = sum + i * j;  # n * n
    	}
    }
}
```

:confused: **How to analyze?**

1. 只关注循环执行次数最多的一段代码
2. 加法：总复杂度等于量级最大的那段代码的复杂度  $$T(n) = O(f1(n)+f2(n)+f3(n)+...)$$
   - 常量时间和 n 规模无关可以忽略，∵ 它对增长趋势没有任何影响。
3. 乘法：**嵌套**代码的复杂度等于嵌套内外代码复杂度的乘积   $$T(n) = O(f1(n)*f2(n)*f3(n)+...)$$

| Big O     | Desc               |
| --------- | ------------------ |
| O(1)      | 常量阶             |
| O(logn)   | 对数阶             |
| O(2^n)    | 指数阶（非多项式） |
| O(n!)     | 阶乘阶（非多项式） |
| O(n)      | 线性阶             |
| O(n*logn) | 线性对数阶         |
| O(n^2)    | 平方阶             |
| O(n^3)    | 立方阶             |
| O(n^k)    | K次方阶            |

```python
# Geometric progression 2^0, 1, 2 ... i = n
# 2^x = n
# x = log2n
i=1;
while (i <= n) {
	i = i * 2;
}
```

:bookmark_tabs: 只要算法中不存在循环语句、递归语句，无论多少行代码，都是 O(1)。

:bookmark_tabs: 系数可以忽略。



<img src="https://miro.medium.com/v2/resize:fit:1050/1*5ZLci3SuR0zM_QlZOADv8Q.jpeg" alt="img" style="zoom: 80%;" />

#### Category

Best/Worst case：在最理想/最糟糕的情况下，执行代码的时间复杂度。

```python
# best:  O(1)
# worst: O(n)
int find(int[] array, int n, int x) {
    int i = 0;
    int pos = -1;
    for (; i < n; ++i) {
        if (array[i] == x) {
        	pos = i;
        	break;
        }
    }
    return pos;
}
```



Average case 平均复杂度

- ~~数组查找有 n+1 种情况 = 0~n-1 + 不在数组中~~

- ~~将遍历元素个数累加 1+2+3...+n+n = n(n+3)~~

- ~~$$n(n+3)/2(n+1)$$ 忽略常数和系数 → O(n)~~

- 要查找的变量 x，要么在数组里，要么就不在数组里，概率 1/2

- 要查找的数据出现在 0～n-1 这 n 个位置的概率也是一样的，1/n

- 根据概率乘法法则，要查找的数据出现在 0～n-1 中任意位置的概率就是 1/(2n)。

- 加权平均：$$1*1/2n+2*1/2n+3*1/2n+...+n*1/n+n*1/2=(3n+1)/4$$ → 忽略常数和系数 → O(n)

  

Amortized 均摊时间复杂度：一种特殊的平均时间复杂度

- 假设数组的长度是 n，根据数据插入的位置的不同，我们可以分为 n 种情况，每种情况的时间复杂度是 O(1)。
- 除此之外，还有一种“额外”的情况，就是在数组没有空闲空间时插入一个数据，这个时候的时间复杂度是 O(n)。
- n+1 种情况发生的概率一样，$$1*1/(n+1) + 1*1/(n+1) + ... n*1/(n+1)$$ → O(1)
- **摊还分析法：**每一次 O(n) 的插入操作，都会跟着 n-1 次 O(1) 的插入操作，所以把耗时多的那次操作均摊到接下来的 n-1 次耗时少的操作上，均摊下来，这一组连续的操作的均摊时间复杂度就是 O(1)。

```python
int[] array = new int[n];
int count = 0;

void insert(int val) {
    if (count == array.length) {
        int sum = 0;
        for (int i = 0; i < array.length; ++i) {
            sum = sum + array[i];
        }
        array[0] = sum;
        count = 1;
    }
    array[count] = val;
    ++count;
}
```

