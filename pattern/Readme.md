### Polymorphism

**多态 Polymorphism** 是 OOP 中的一个重要概念。

多态 ← **继承 Inheritance** & **方法重写 Overriding** & **接口 Interface**

多态通过使用**父类**的引用/指针来引用**子类**对象（parent → child * N），从而实现对不同类型对象的**统一操作**。

具体执行由子类重写父类方法 or 实现接口的类型方法决定，**运行时根据实际对象的类型来确定调用哪个方法**。（呈现不同形态）

### [Overview](https://www.yuque.com/aceld/lfhu8y)

**设计模式**是在特定环境下人们**解决某类问题可复用的解决方案**。

**软件设计模式**是一套可重用、多数人知晓的、经过分类编目的、代码设计经验的总结。

1. 创建型 Creational：如何创建对象。

2. 结构型 Structural：如何实现类或对象的组合。

3. 行为型 Behavioral：类或对象怎样交互以及怎样分配职责。

**学习有助于：**

1. 如何将代码分散在几个不同的类中？
2. 为什么要有“接口”？
3. 何谓针对抽象编程？
4. 何时不应该使用继承？
5. 如果不修改源代码增加新功能？
6. 更好地阅读和理解现有类库与其他系统中的源代码。

### Principle

> 目的：高内聚 High Cohesion + 低耦合 Low Coupling.

**单一职责原则（Single Responsibility Principle，SRP）**：一个类应该只负责一项职责。

**开放封闭原则（Open-Closed Principle，OCP）**：类应该对扩展开放，对修改关闭（可能会增加内部耦合度）。

**里氏替换原则（Liskov Substitution Principle，LSP）**：任何抽象类/接口都可以用其实现类替换。

**依赖倒置原则（Dependency Inversion Principle，DIP）**：面向抽象/接口编程 = 依赖于抽象/接口，而不是具体实现类。

- 一般代码设计：接口层/实现层/逻辑层。

  - 逻辑层

    ↓ (Dep)

  - 抽象层/接口层

    ↑ (Dep)

  - 实现层

**接口隔离原则（Interface Segregation Principle，ISP）**：接口设计得精细小而专一，而不是一个大而笼统的接口。不强迫实现依赖它们不需要的接口方法，即接口设计也尽量符合单一职责原则。

**合成/聚合复用原则（Composition/Aggregation Reuse Principle，CARP）**：首选使用对象组合或聚合，而不是通过继承来实现代码复用（需要把父类所有的方法都拿过来，有些可能是我们不需要的）。轻量优化：对象不依赖，只有在调用的时候依赖耦合。

**迪米特原则（Law of Demeter，LoD）**： 一个对象应该尽可能少地了解其他对象的细节，从而降低之间的耦合性。通常抽象一个中间接口层 as proxy，各模块相互调用通过统一接口实现。

