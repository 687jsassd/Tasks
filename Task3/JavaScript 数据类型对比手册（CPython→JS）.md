# JavaScript 数据类型对比手册（C/Python→JS）

> 适用于有 C 或 Python 基础的 JS 初学者，重点对比数据类型特性与特殊行为

---

## 1. 基本数据类型

### C 语言
```c
// 显式类型声明
int age = 25;            // 整型
float price = 9.99;      // 单精度浮点
double pi = 3.1415926;   // 双精度浮点
char grade = 'A';        // 单个字符
_Bool isOpen = 1;        // 布尔值（C99 引入）
```

### Python
```python
# 动态类型推断
age = 25               # int
price = 9.99           # float
name = "Alice"         # str（不可变）
is_active = True       # bool
nothing = None         # 空值
# 没有 char 类型，单字符也是字符串
```

### JavaScript
```javascript
// 动态类型（类型随值改变）
let age = 25;            // number
let price = 9.99;        // number（无 int/float 区分）
let name = "Bob";        // string（不可变）
let isAdmin = true;      // boolean
let empty = null;        // 空值（需显式赋值）
let notDefined;          // undefined（默认值）

// 特殊类型
let id = Symbol("id");   // Symbol（唯一标识符）
let bigNum = 123n;       // BigInt（ES2020+）
```

---

## 2. 复杂数据类型

### C 语言
```c
// 数组（固定大小/类型）
int scores[3] = {90, 85, 95};  

// 结构体（自定义复合类型）
struct Person {
    char name[20];
    int age;
};
struct Person user = {"Tom", 30};

// 指针（内存地址）
int* ptr = &age;
```

### Python
```python
# 列表（动态可变）
scores = [90, 85, 95]  
scores.append(100)  # 添加元素

# 元组（不可变）
colors = ("red", "green")

# 字典（键值对）
user = {"name": "Tom", "age": 30}

# 集合（唯一元素）
unique_nums = {1, 2, 3}
```

### JavaScript
```javascript
// 数组（动态类型/长度可调）
let scores = [90, 85, 95];
scores.push(100);  // 添加元素

// 对象（类似 Python 字典）
let user = { 
    name: "Tom", 
    age: 30,
    sayHi: function() { console.log("Hi!") }  // 可包含函数
};

// 特殊对象类型
let date = new Date();          // 日期对象
let regex = /pattern/gi;        // 正则表达式对象
```

---

## 3. 类型转换

### C 语言
```c
// 显式强制转换
int total = (int)price + age;    // 浮点转整型（直接截断）
char numStr[10];
sprintf(numStr, "%d", age);      // 数字转字符串
```

### Python
```python
# 显式构造函数转换
total = int(price) + age        # 浮点转整型（舍去小数）
num_str = str(age)              # 数字转字符串
is_valid = bool("Hello")        # 非空字符串转 True
```

### JavaScript
```javascript
// 隐式自动转换
let result = "5" + 2;     // "52"（数字转字符串）
let sum = "5" - 2;        // 3（字符串转数字）

// 显式转换
let num = Number("123");  // 123（失败返回 NaN）
let str = String(123);    // "123"
let bool = Boolean("");   // false

// 特殊转换案例
Boolean(null)         // false
Boolean({})           // true
Number(undefined)     // NaN
```

---

## 4. 类型检查

### C 语言
```c
// 编译时类型检查（无运行时类型判断）
printf("%ld", sizeof(int));  // 检查类型大小
```

### Python
```python
print(type(age) == int)      # 检查精确类型
print(isinstance(age, int))  # 包括继承关系检查
```

### JavaScript
```javascript
// typeof 运算符（返回类型字符串）
typeof 42            // "number"
typeof "hello"       // "string"
typeof undefined     // "undefined"

// 特殊行为
typeof null          // "object"（历史遗留问题）
typeof [1,2,3]       // "object"

// 精确检查方法
Array.isArray([1,2,3])              // true
user instanceof Object              // true
Object.prototype.toString.call([])  // "[object Array]"
```

---

## 关键差异总结表

| 特性      | C 语言             | Python              | JavaScript               |
| --------- | ------------------ | ------------------- | ------------------------ |
| 类型系统  | 静态类型           | 动态类型            | 动态类型                 |
| 整数/浮点 | 明确区分 int/float | 自动处理 int/float  | 统一为 number 类型       |
| 布尔值    | _Bool（0/1）       | True/False          | true/false（小写）       |
| 空值      | NULL（指针）       | None                | null 和 undefined 双空值 |
| 字符串    | char[]（可变）     | str（不可变）       | string（不可变）         |
| 数组      | 固定大小/类型      | list（动态可变）    | Array（动态类型/长度）   |
| 复合结构  | 结构体             | dict/class          | Object（键值对）         |
| 类型转换  | 必须显式强制转换   | 显式构造函数转换    | 允许隐式自动转换         |
| 类型检查  | 编译时检查         | type()/isinstance() | typeof/instanceof        |

---

## 最佳实践建议

1. **类型安全**：优先使用 `===` 而非 `==`（避免隐式转换）
2. **空值处理**：主动初始化变量，用 `null` 表示"无值"，`undefined` 留给未赋值变量
3. **类型检查**：使用 `Array.isArray()` 检查数组，`Object.prototype.toString.call()` 精确判断类型
4. **大数处理**：超过 `2^53` 时使用 `BigInt`（后缀加 n）
5. **对象拷贝**：使用 `JSON.parse(JSON.stringify(obj))` 实现简单深拷贝（注意函数丢失问题）
6. **防御性编程**：始终验证外部数据（如 `Number(input)` 后检查 `NaN`）

