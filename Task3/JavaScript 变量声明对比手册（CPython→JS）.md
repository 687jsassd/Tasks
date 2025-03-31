# JavaScript 变量声明对比手册（C/Python→JS）

> 适用于有 C 或 Python 基础的 JS 初学者，重点对比变量声明方式与特性

---

## 1. 声明方式

### C 语言
```c
// 必须显式指定类型
int count = 10;            // 整型变量
const double PI = 3.14;    // 常量
static int global = 100;   // 静态变量

// 可先声明后赋值
int temp;
temp = 20;
```

### Python
```python
# 无需关键字直接赋值
score = 90                # 动态类型变量
MAX_SIZE = 100            # "常量"（约定全大写，实际可变）
global counter            # 声明全局变量（需在函数内使用）

# 必须立即初始化
# age  # ❌ NameError
```

### JavaScript
```javascript
// 三种声明方式
var name = "Alice";       // 函数作用域（旧方式）
let age = 25;             // 块级作用域变量
const PI = 3.14;          // 块级作用域常量

// 声明与赋值分离
let title;                // 合法（值为 undefined）
title = "Engineer";
```

---

## 2. 作用域规则

### C 语言
```c
// 块级作用域（由 {} 定义）
void demo() {
    int x = 10;
    if (1) {
        int y = 20;    // y 只在 if 块内有效
        x = 15;        // 可访问外层变量
    }
    printf("%d", y);   // ❌ 编译错误
}
```

### Python
```python
# 函数作用域（无块级作用域）
def test():
    if True:
        value = 30     # 实际属于函数作用域
    print(value)       # ✅ 输出 30

# 局部变量覆盖全局变量时需用 global
total = 0
def add():
    global total
    total += 1
```

### JavaScript
```javascript
function example() {
    var a = 1;        // 函数作用域
    if (true) {
        let b = 2;    // 块级作用域
        var c = 3;    // 提升到函数顶部
    }
    console.log(a);   // ✅ 1
    console.log(c);   // ✅ 3
    console.log(b);   // ❌ ReferenceError
}

// 全局变量差异
var globalVar = 1;    // 挂载到 window 对象
let moduleVar = 2;    // 不属于 window 对象
```

---

## 3. 变量提升（Hoisting）

### C 语言
```c
// 无变量提升
printf("%d", num);  // ❌ 编译错误（变量未声明）
int num = 5;
```

### Python
```python
# 无变量提升
print(num)  # ❌ NameError
num = 10
```

### JavaScript
```javascript
// var 声明会被提升
console.log(version);    // undefined（不会报错）
var version = "1.0";

// let/const 存在临时死区（TDZ）
console.log(status);     // ❌ ReferenceError
let status = "active";
```

---

## 4. 重复声明

### C 语言
```c
int id = 100;
int id = 200;  // ❌ 编译错误（重复定义）
```

### Python
```python
user_id = 100
user_id = 200  # ✅ 重新赋值（动态语言特性）
del user_id     # 删除变量
```

### JavaScript
```javascript
var count = 1;
var count = 2;   // ✅ 允许重复声明（可能引发 bug）

let score = 90;
let score = 95;  // ❌ SyntaxError

const MAX = 100;
const MAX = 200; // ❌ SyntaxError
```

---

## 5. 初始化要求

### C 语言
```c
int value;        // 未初始化（值为内存垃圾）
const int CODE;   // ❌ 编译错误（常量必须初始化）
```

### Python
```python
# 变量必须立即初始化
# value  # ❌ NameError
```

### JavaScript
```javascript
var name;        // ✅ 合法（值为 undefined）
let age;         // ✅ 合法

const PI;        // ❌ SyntaxError（必须初始化）
const RATE = 0.8;
```

---

## 关键差异总结表

| 特性       | C 语言                | Python          | JavaScript                          |
| ---------- | --------------------- | --------------- | ----------------------------------- |
| 声明方式   | 类型前缀（int/float） | 直接赋值        | `var`/`let`/`const`                 |
| 作用域     | 块级作用域            | 函数作用域      | 函数作用域 (var) / 块级 (let/const) |
| 变量提升   | 无                    | 无              | var 声明提升，let/const 有 TDZ      |
| 重复声明   | 禁止                  | 覆盖变量        | var 允许，let/const 禁止            |
| 常量       | `const`               | 约定全大写      | `const`（不可重新赋值）             |
| 全局变量   | `static`              | `global` 关键字 | var 声明挂载到 window               |
| 初始化要求 | 局部变量可不初始化    | 必须初始化      | const 必须初始化                    |

---

## 最佳实践建议

1. **声明选择**：优先使用 `const`，其次 `let`，避免使用 `var`
2. **命名规范**：常量用全大写（`API_KEY`），变量用驼峰命名（`userName`）
3. **作用域控制**：在代码块开始处集中声明变量（避免提升导致的意外）
4. **对象常量**：`const` 对象属性仍可修改（需用 `Object.freeze()` 完全锁定）
5. **严格模式**：使用 `'use strict';` 规避静默错误（如未声明变量赋值）
6. **模块化编程**：使用 `let` 替代 `var` 避免污染全局命名空间

