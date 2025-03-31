# JavaScript 流程控制对比手册（C/Python→JS）

> 适用于有 C 或 Python 基础的 JS 初学者，重点对比条件判断与循环结构特性差异

---

## 1. 条件语句

### C 语言
```c
// if-else 结构
int x = 10;
if (x > 5) {
    printf("大于5");
} else if (x == 5) {
    printf("等于5");
} else {
    printf("小于5");
}

// 三元表达式
int y = (x > 5) ? 1 : 0;
```

### Python
```python
# 缩进控制代码块
score = 85
if score >= 90:
    print("A")
elif 80 <= score < 90:
    print("B")  # 输出 B
else:
    print("C")

# 无 switch 语句
```

### JavaScript
```javascript
// if-else 结构（同 C 语法）
let age = 20;
if (age >= 18) {
    console.log("成年");
} else {
    console.log("未成年");
}

// 三元运算符（支持表达式）
const status = age >= 18 ? 'adult' : 'minor';

// switch 严格比较（===）
let fruit = 'apple';
switch (fruit) {
    case 'apple':
        console.log('苹果');
        break;
    default:
        console.log('未知水果');
}
```

---

## 2. 循环结构

### C 语言
```c
// for 循环
for(int i=0; i<5; i++){
    printf("%d", i);  // 0-4
}

// while 循环
int j = 5;
while(j > 0){
    j--;
}

// do-while 循环
do {
    // 至少执行一次
} while(j > 0);
```

### Python
```python
# for-in 迭代
for num in [1,2,3]:
    print(num)

# while 循环
count = 3
while count > 0:
    count -= 1

# 推导式
squares = [x**2 for x in range(5)]
```

### JavaScript
```javascript
// C 风格 for 循环
for (let i=0; i<5; i++) {
    console.log(i);
}

// for-of 迭代（ES6+）
const arr = [1,2,3];
for (const num of arr) {
    console.log(num);
}

// for-in 遍历对象键
const obj = {a:1, b:2};
for (const key in obj) {
    console.log(key);  // 输出 a, b
}

// while 循环（同 C）
let k = 3;
while (k-- > 0) {
    // 循环体
}
```

---

## 3. 流程控制语句

### C 语言
```c
// break/continue
for(int i=0; i<10; i++){
    if(i == 5) break;     // 终止循环
    if(i % 2) continue;   // 跳过本次迭代
}

// goto（慎用）
goto label;
// ...
label: 
```

### Python
```python
# break/continue
for n in range(10):
    if n == 5:
        break
    if n % 2 == 0:
        continue

# 无 goto
```

### JavaScript
```javascript
// break/continue（同 C）
for (let i=0; i<10; i++) {
    if (i === 5) break;
    if (i % 2) continue;
}

// 标签跳转（类似 goto）
outerLoop: 
for (let i=0; i<3; i++) {
    for (let j=0; j<3; j++) {
        if (i === 1) break outerLoop;
    }
}
```

---

## 4. 异常处理

### C 语言
```c
// 无原生异常处理
int result = someFunction();
if (result == ERROR_CODE) {
    handleError();
}
```

### Python
```python
try:
    x = 1 / 0
except ZeroDivisionError as e:
    print("错误:", e)
finally:
    print("清理资源")
```

### JavaScript
```javascript
// try-catch-finally（类似 Python）
try {
    JSON.parse('invalid json');
} catch (err) {
    console.log('解析错误:', err.message);
} finally {
    console.log('执行清理');
}

// throw 任意类型
throw '错误描述';
throw new Error('错误对象');
```

---

## 5. 真值判断（Truthy/Falsy）

| 值类型      | C 语言   | Python | JavaScript   |
| ----------- | -------- | ------ | ------------ |
| 0           | 假       | 假     | 假（falsy）  |
| 空字符串 "" | 不可比较 | 假     | 假（falsy）  |
| null        | 无       | 无     | 假（falsy）  |
| undefined   | 无       | 无     | 假（falsy）  |
| NaN         | 无       | 假     | 假（falsy）  |
| 空数组 []   | 无       | 真     | 真（truthy） |
| 空对象 {}   | 无       | 真     | 真（truthy） |
| "0"         | 无       | 真     | 真（truthy） |

---

## 关键差异总结表

| 特性           | C 语言          | Python            | JavaScript                      |
| -------------- | --------------- | ----------------- | ------------------------------- |
| 代码块界定     | 大括号 {}       | 缩进              | 大括号 {}                       |
| switch 语句    | 支持（值匹配）  | 无（用字典模拟）  | 支持（严格比较 ===）            |
| 循环变量作用域 | 块级作用域      | 函数作用域        | let 块级作用域 / var 函数作用域 |
| 迭代协议       | 无              | __iter__/__next__ | Symbol.iterator                 |
| 真值判断       | 0 为假，非0为真 | 明确布尔转换      | falsy/truthy 隐式转换           |
| 错误处理       | 返回错误码      | try/except        | try/catch                       |
| 流程跳转       | goto            | 无                | 标签跳转                        |
| 空值处理       | NULL 指针       | None              | null 和 undefined               |

---

## 最佳实践建议

1. **条件判断**：优先使用 `===` 严格相等比较（避免类型转换错误）
2. **循环变量**：始终使用 `let` 声明循环变量（避免 var 的变量提升问题）
3. **现代迭代**：优先使用 `for...of` 替代传统 for 循环遍历数组
4. **异常处理**：始终抛出 `Error` 对象而非原始值（保留调用栈信息）
5. **真值检查**：显式转换 `if (value !== null && value !== undefined)`
6. **switch 优化**：使用对象映射替代复杂 switch-case 结构
7. **循环控制**：避免在循环内声明函数（防止闭包陷阱）
8. **错误边界**：异步操作使用 `Promise.catch()` 替代 try-catch
9. **空值合并**：使用 `??` 运算符处理默认值 `const val = input ?? 'default'`
10. **可选链**：使用 `?.` 避免深层属性访问错误 `obj?.prop?.method?.()`

