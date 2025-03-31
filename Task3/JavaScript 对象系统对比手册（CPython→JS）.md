# JavaScript 对象系统对比手册（C/Python→JS）

> 适用于有 C 或 Python 基础的 JS 初学者，重点对比面向对象特性与设计差异

---

## 1. 基本对象声明

### C 语言
```c
// 使用结构体模拟对象
struct Person {
    char name[20];
    int age;
};

// 需定义独立函数操作
void printPerson(struct Person p) {
    printf("%s, %d", p.name, p.age);
}

struct Person alice = {"Alice", 25};
```

### Python
```python
# 类定义
class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age
    
    def print_info(self):
        print(f"{self.name}, {self.age}")

alice = Person("Alice", 25)
```

### JavaScript
```javascript
// 对象字面量
const alice = {
    name: "Alice",
    age: 25,
    printInfo() {
        console.log(`${this.name}, ${this.age}`);
    }
};

// 构造函数模式
function Person(name, age) {
    this.name = name;
    this.age = age;
    this.printInfo = function() {
        console.log(`${this.name}, ${this.age}`);
    }
}
const bob = new Person("Bob", 30);
```

---

## 2. 继承机制

### C 语言
```c
// 无原生继承支持，需手动模拟
struct Employee {
    struct Person person;
    char position[20];
};
```

### Python
```python
# 类继承
class Employee(Person):
    def __init__(self, name, age, position):
        super().__init__(name, age)
        self.position = position

e = Employee("Carol", 28, "Engineer")
```

### JavaScript
```javascript
// 原型链继承
function Employee(name, age, position) {
    Person.call(this, name, age);
    this.position = position;
}
Employee.prototype = Object.create(Person.prototype);

// ES6 类继承
class Employee extends Person {
    constructor(name, age, position) {
        super(name, age);
        this.position = position;
    }
}

const carol = new Employee("Carol", 28, "Engineer");
```

---

## 3. 属性操作

### C 语言
```c
// 直接访问结构体成员
strcpy(alice.name, "Alice Smith");
alice.age += 1;
```

### Python
```python
# 动态属性操作
alice.email = "alice@example.com"  # 运行时添加属性
del alice.age                     # 删除属性
hasattr(alice, 'name')            # 检查属性存在
```

### JavaScript
```javascript
// 动态属性管理
alice.email = "alice@example.com";  // 添加属性
delete alice.age;                   // 删除属性
"name" in alice;                   // 检查属性存在

// 属性描述符
Object.defineProperty(alice, 'id', {
    value: 1001,
    writable: false
});

// 计算属性名
const prop = 'score';
const obj = {
    [prop + '_calc']: 95  // score_calc: 95
};
```

---

## 4. 特殊对象类型

### C 语言
```c
// 无内置高级对象，需自行实现
typedef struct {
    int year;
    int month;
    int day;
} Date;
```

### Python
```python
# 内置对象类型
from datetime import date
d = date(2023, 9, 15)

import json
data = json.loads('{"key": "value"}')
```

### JavaScript
```javascript
// 内置对象
const now = new Date();               // 日期对象
const pattern = /js/gi;               // 正则对象
const jsonData = JSON.parse('{"key": "value"}');

// 特殊对象
const map = new Map();                // 键值对集合
map.set('name', 'David');

const proxy = new Proxy(target, handler); // 代理对象
```

---

## 5. 原型系统

### Python
```python
# 方法解析顺序（MRO）
print(Employee.__mro__)  # 显示继承链
```

### JavaScript
```javascript
// 原型链访问
console.log(alice.__proto__ === Person.prototype); // true
console.log(Person.prototype.__proto__ === Object.prototype); // true

// 原型方法
Object.getPrototypeOf(alice);       // 获取原型
Object.setPrototypeOf(obj, proto);  // 修改原型（性能差）

// 原型继承关系
console.log(carol instanceof Employee);  // true
console.log(carol instanceof Person);    // true
```

---

## 关键差异总结表

| 特性         | C 语言     | Python                 | JavaScript              |
| ------------ | ---------- | ---------------------- | ----------------------- |
| 对象基础     | 结构体     | class 关键字           | 对象字面量/构造函数     |
| 继承机制     | 组合模拟   | 类继承                 | 原型链继承              |
| 方法定义     | 独立函数   | 类方法                 | 函数属性/原型方法       |
| 动态属性     | 不支持     | 支持                   | 支持                    |
| 访问控制     | 无         | 命名约定（_protected） | Symbol 私有属性（ES6+） |
| 多态实现     | 函数指针   | 鸭子类型               | 原型覆盖                |
| 实例检查     | 无         | isinstance()           | instanceof              |
| 内置对象类型 | 需手动实现 | datetime/json 等       | Date/RegExp/Map/Set 等  |
| 元编程能力   | 有限       | 装饰器/元类            | Proxy/Reflect           |
| 序列化支持   | 需手动实现 | pickle 模块            | JSON 方法               |

---

## 最佳实践建议

1. **创建对象**：优先使用类语法（ES6+）替代原型操作
2. **封装原则**：使用 `#` 私有字段（ES2022+）替代 `_` 命名约定
3. **继承策略**：组合优于继承，优先使用 Object.assign() 混入
4. **对象操作**：使用 `Object.hasOwn()` 替代 `in` 运算符检查自有属性
5. **不可变模式**：使用 `Object.freeze()` 防止意外修改重要对象
6. **原型污染**：避免修改 `Object.prototype` 等内置原型
7. **现代特性**：优先使用 Map/Set 替代普通对象存储键值对集合
8. **类型检查**：结合 `typeof` 与 `Object.prototype.toString` 精确判断类型
9. **拷贝策略**：深拷贝使用 `structuredClone()`（浏览器环境）
10. **模式匹配**：利用解构赋值提取对象属性 `const {name} = user;`

