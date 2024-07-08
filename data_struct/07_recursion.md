:confused: **什么样的问题可以用“递归”来解决？**同时满足以下条件：

1. 一个问题的解可以**分解**为几个子问题的解 → Recursive Case
2. 这个问题与分解之后的子问题，求解思路完全一样
3. 存在递归**终止条件** → Base Case



:confused: **如何编写递归代码？**写出<u>递推公式</u>，找到<u>终止条件</u>。

e.g. N 个台阶，每次可以跨 1 or 2 个台阶。请文有几种走法？

- 第一步走了 1 个台阶 + (n-1) 个台阶走法；第一步走了 2 个台阶 + (n-2) 个台阶走法
- 中止条件：
  - 当只有 2 个台阶的时候，有 2 种走法。（跨 1 个台阶 + 跨 2 个台阶）
  - 当只有 1 个台阶的时候，只有 1 种走法。
- 递归公式：
  - $$f(n)=f(n-1)+f(n-2)$$
  - $$f(1)=1$$
  - $$f(2)=2$$

关键：把它**抽象**成一个递推公式，不用想一层层的调用关系，不要试图用人脑去分解递归的每个步骤。



**栈溢出 Stackoverflow**

如果递归求解的规模很大，调用层次很深，函数一直入栈，就会有堆栈溢出的风险。

解决：限制调用最大深度。注：Go 种通过代码显式追踪来限制；控制 `GOMAXPROCS`；控制栈大小 `OS ulimit -s`；控制最大线程数 SetMaxThreads 



**避免重复计算**

通过一个数据结构（比如散列表）来**保存已经求解过的值**。每次递归调用判断是否求解过。



:confused: **vs?**

:smile: 易理解，简洁

:cry: 空间复杂度高，存在栈溢出，重复计算，过多函数耗时等问题。



**:bookmark_tabs: “最终推荐人”** 追朔判断 actor_id 是否存在 referer_id？不存在则为“最终推荐人”。

```java
long findRootReferrerId(long actorId) {
    referrerId = select by actorId
    if referrerId == null {
        return actorId
    }
    return findRootReferrerId(referrerId)
}
```

