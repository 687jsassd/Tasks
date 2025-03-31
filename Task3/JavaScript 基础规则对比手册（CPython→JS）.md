# JavaScript 基础规则对比手册（C/Python→JS）

> 适用于有 C 或 Python 基础的 JS 初学者，重点对比变量作用域、分号、注释等基础规则

---

## 1. 变量作用域

### C 语言
```c
// 块级作用域（由 {} 定义）
int main() {
    if (1) {
        int x = 10; // x 只在 if 块内有效
    }
    printf("%d", x); // ❌ 编译错误：x 未定义
}

// 需显式声明变量类型
```

### Python
```python
# 函数作用域（无块级作用域）
def func():
    if True:
        y = 20
    print(y)  # ✅ 输出 20

# 全局变量需用 global 声明
global_var = 10
def foo():
    global global_var
    global_var += 1
```

### JavaScript
```javascript
// 函数作用域 (var) vs 块级作用域 (let/const)
function test() {
    if (true) {
        var a = 10;    // 提升到函数顶部
        let b = 20;    // 仅在 if 块内有效
    }
    console.log(a);  // ✅ 10
    console.log(b);  // ❌ ReferenceError
}

// 变量提升（仅声明被提升，赋值不提升）
console.log(c);     // undefined（不会报错）
var c = 30; 
console.log(c);     // 30
```

---

## 2. 分号使用

### C 语言
```c
// 必须使用分号
int x = 5;
printf("%d", x);
```

### Python
```python
# 禁止使用分号（除非同一行写多个语句）
x = 10
print(x)  # 正确
y = 20; print(y)  # 有效但不推荐
```

### JavaScript
```javascript
// 分号可选（但推荐始终手动添加）
let name = "Alice"
console.log(name)  // ✅ 正确

// 自动分号插入 (ASI) 的陷阱
function getData() {
    return 
    { id: 1 }  // ❌ 实际返回 undefined（ASR 在 return 后插入分号）
}

// 必须加分号的场景
const list = [1, 2, 3]
;[1, 2, 3].forEach(console.log)  // 行以 [ 开头时需前置分号
```

---

## 3. 注释语法

### C 语言
```c
// 单行注释

/* 
多行注释
可跨行 
*/
```

### Python
```python
# 单行注释

"""
三引号字符串常被用作
多行注释（但本质是字符串，不推荐）
'''
```

### JavaScript
```javascript
// 单行注释（同 C）

/* 
多行注释（同 C）
支持嵌套 /* 注释 */ （C 语言不支持嵌套）
*/
```

---

## 关键差异总结表

| 特性       | C 语言          | Python       | JavaScript                          |
| ---------- | --------------- | ------------ | ----------------------------------- |
| 变量作用域 | 块级作用域      | 函数作用域   | 函数作用域 (var) / 块级 (let/const) |
| 变量声明   | 必须声明类型    | 直接赋值     | `var`/`let`/`const`                 |
| 分号       | 强制使用        | 禁止使用     | 可选（建议始终手动添加）            |
| 注释       | `//` 和 `/* */` | `#` 和 `"""` | `//` 和 `/* */`（同 C）             |
| 变量提升   | 无              | 无           | `var` 声明会提升                    |

---

## 最佳实践建议

1. **变量声明**：始终使用 `let`/`const` 替代 `var`（避免意外提升和全局污染）
2. **分号策略**：统一代码风格（要么全加，要么全不加），建议手动添加分号
3. **作用域控制**：避免在块内用 `var` 声明变量（优先使用 `let`）
4. **严格模式**：在 JS 文件开头添加 `"use strict";` 以规避静默错误

