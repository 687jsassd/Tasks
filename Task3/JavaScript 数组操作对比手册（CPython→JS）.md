# JavaScript 数组操作对比手册（C/Python→JS）

> 适用于有 C 或 Python 基础的 JS 初学者，重点对比数组特性与操作方法差异

---

## 1. 基本声明

### C 语言
```c
// 固定类型/长度
int nums[3] = {1, 2, 3};          // 显式声明类型和大小
char *names[] = {"Alice", "Bob"}; // 字符串数组

// 动态数组需手动管理内存
int *arr = malloc(5 * sizeof(int));
free(arr);  // 必须手动释放
```

### Python
```python
# 动态类型列表
numbers = [1, 2, 3]                 # 混合类型合法
names = ["Alice", 30, True]         # 元素类型自由

# 列表推导式
squares = [x**2 for x in range(10)]
```

### JavaScript
```javascript
// 动态类型数组
const nums = [1, 2, 3];            // 字面量声明
const mixed = ["Text", 25, true];  // 混合类型合法

// 构造函数方式
const emptyArr = new Array(3);     // 创建长度为3的空数组 [empty × 3]

// 类型化数组（特殊场景）
const buffer = new Int32Array(5);  // 固定类型数组
```

---

## 2. 动态特性

### C 语言
```c
// 固定长度数组（无法直接扩展）
int arr[3] = {1,2,3};
arr[3] = 4;  // ❌ 数组越界（可能引发内存错误）
```

### Python
```python
lst = [1,2,3]
lst.append(4)       # ✅ [1,2,3,4]
lst.insert(0, 0)    # ✅ [0,1,2,3,4]
lst.pop()           # ✅ 移除末尾元素
```

### JavaScript
```javascript
const arr = [1,2,3];

// 修改方法
arr.push(4);        // ✅ [1,2,3,4]（尾部添加）
arr.unshift(0);     // ✅ [0,1,2,3,4]（头部添加）
arr.pop();          // ✅ 移除最后一个元素
arr.splice(1, 0, 9) // 在索引1插入9 → [0,9,1,2,3]

// 长度可变
arr.length = 2;     // ✅ 截断数组 → [0,9]
```

---

## 3. 常用方法对比

### C 语言
```c
// 需要手动实现基本操作
int arr[5] = {1,2,3,4,5};

// 遍历数组
for(int i=0; i<5; i++) {
    printf("%d", arr[i]);
}

// 查找元素（需自己实现）
int index = -1;
for(int i=0; i<5; i++){
    if(arr[i] == 3) index = i;
}
```

### Python
```python
lst = [1,2,3,4,5]

# 常用方法
lst.reverse()           # 倒序 → [5,4,3,2,1]
lst2 = lst.copy()       # 浅拷贝
sum(lst)               # 求和 → 15

# 切片操作
sub = lst[1:4]         # [2,3,4]
```

### JavaScript
```javascript
const arr = [1,2,3,4,5];

// 关键方法
arr.reverse();          // [5,4,3,2,1]
arr.slice(1,4);         // [2,3,4]（类似Python切片）
arr.includes(3);       // true（存在检查）
arr.find(x => x > 3);  // 4（查找首个满足条件的元素）

// 高阶函数
arr.map(x => x*2);     // [2,4,6,8,10]
arr.filter(x => x%2);  // [1,3,5]
arr.reduce((a,b) => a+b, 0); // 求和 → 15
```

---

## 4. 特殊操作

### 数组合并
```python
# Python
a = [1,2]
b = [3,4]
c = a + b  # [1,2,3,4]
```

```javascript
// JavaScript
const a = [1,2];
const b = [3,4];
const c = [...a, ...b];    // 扩展运算符 → [1,2,3,4]
const d = a.concat(b);     // 方法调用 → 同上
```

### 解构赋值
```python
# Python
a, b, *rest = [1,2,3,4]  # a=1, b=2, rest=[3,4]
```

```javascript
// JavaScript
const [x, y, ...others] = [1,2,3,4];  // x=1, y=2, others=[3,4]
```

### 稀疏数组
```javascript
// JavaScript 特有
const sparseArr = [1,,3];  // 中间空位 → [1, empty, 3]
sparseArr[1];             // undefined
```

---

## 关键差异总结表

| 特性     | C 语言        | Python              | JavaScript               |
| -------- | ------------- | ------------------- | ------------------------ |
| 数组类型 | 固定类型数组  | list 动态集合       | Array 对象               |
| 元素类型 | 必须统一      | 允许混合类型        | 允许混合类型             |
| 内存管理 | 手动分配/释放 | 自动回收            | 自动回收                 |
| 动态扩容 | 不可          | 支持                | 支持                     |
| 多维数组 | 支持          | 嵌套列表            | 嵌套数组                 |
| 索引越界 | 导致内存错误  | 引发 IndexError     | 返回 undefined           |
| 切片操作 | 需手动实现    | 原生支持            | slice() 方法             |
| 迭代方式 | for 循环      | for-in / 列表推导式 | for循环/forEach/for-of   |
| 包含检查 | 需遍历检查    | in 运算符           | includes() 方法          |
| 空值处理 | 无            | None 占位           | 允许空位（sparse array） |

---

## 最佳实践建议

1. **声明方式**：优先使用 `const arr = []` 而非 `new Array()`
2. **不可变操作**：使用 `slice()`/`map()` 替代直接修改原始数组
3. **空值处理**：避免创建稀疏数组（用 `undefined` 显式填充空位）
4. **类型一致**：尽量保持数组元素类型一致（增强代码可读性）
5. **性能优化**：大规模数据操作使用 `TypedArray`
6. **拷贝技巧**：深拷贝使用 `JSON.parse(JSON.stringify(arr))`（注意函数丢失）
7. **现代语法**：优先使用扩展运算符 `...` 和箭头函数简化操作
8. **安全访问**：使用 `arr?.length` 可选链避免空值错误

