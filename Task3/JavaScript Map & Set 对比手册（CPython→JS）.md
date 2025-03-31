# JavaScript Map & Set 对比手册（C/Python→JS）

> 适用于有 C 或 Python 基础的 JS 初学者，重点对比集合类型特性与操作差异

---

## 1. 基本声明

### C 语言
```c
// 无原生 Map/Set 支持，需手动实现
// 模拟 Map（键值对数组）
struct MapEntry {
    char* key;
    int value;
};
struct MapEntry map[10];
int mapSize = 0;

// 模拟 Set（数组去重）
int set[10];
int setSize = 0;
```

### Python
```python
# 字典（类似 Map）
person = {"name": "Alice", "age": 25}

# 集合
unique_nums = {1, 2, 3}

# 类型化字典（Python 3.9+）
from typing import Dict
scores: Dict[str, int] = {"math": 90}
```

### JavaScript
```javascript
// Map 对象
const person = new Map();
person.set("name", "Alice");
person.set("age", 25);

// Set 对象
const uniqueNums = new Set([1, 2, 3]);

// 对象字面量（非真正 Map）
const objMap = { 
    name: "Bob", // 键自动转为字符串
    1: "One"     // 数字键会转为字符串
};
```

---

## 2. 主要特性对比

| 特性           | C 模拟              | Python               | JavaScript                   |
| -------------- | ------------------- | -------------------- | ---------------------------- |
| **键类型**     | 固定类型（如char*） | 不可变类型（如str）  | 任意类型（对象/函数等）      |
| **顺序保留**   | 数组顺序            | 字典(Python3.7+有序) | Map/Set 保持插入顺序         |
| **重复检测**   | 手动实现            | 自动去重             | 自动去重（Set）              |
| **空值支持**   | 需特殊处理          | 允许 None            | 允许 null/undefined          |
| **直接序列化** | 需自定义格式        | 支持 JSON            | Map/Set 需转换后才可 JSON 化 |
| **内存管理**   | 手动释放            | 自动回收             | 自动回收                     |
| **内置方法**   | 无                  | 丰富方法             | 专用 API（set/get/has等）    |

---

## 3. 常用操作对比

### 添加元素
```c
// C 模拟 Map
map[mapSize++] = (struct MapEntry){"age", 25}; 

// C 模拟 Set（需遍历检查重复）
if(!contains(set, 3)) set[setSize++] = 3;
```

```python
# Python
person["gender"] = "female"  # 字典赋值
unique_nums.add(4)           # 集合添加
```

```javascript
// JavaScript
person.set("gender", "female");  // Map 添加
uniqueNums.add(4);               // Set 添加

// 对象特殊操作
objMap[Symbol.iterator] = function*() { /*...*/ };  // 添加 Symbol 键
```

### 删除元素
```python
# Python
del person["age"]          # 字典删除
unique_nums.discard(2)     # 集合删除
```

```javascript
// JavaScript
person.delete("age");        // Map 删除
uniqueNums.delete(2);        // Set 删除

// 对象删除
delete objMap.name;
```

### 查找判断
```c
// C 模拟查找
int findValue(char* key) {
    for(int i=0; i<mapSize; i++){
        if(strcmp(map[i].key, key) == 0) return map[i].value;
    }
    return -1;
}
```

```python
# Python
"name" in person           # True
2 in unique_nums           # True
```

```javascript
// JavaScript
person.has("name");         // true
uniqueNums.has(2);          // true

// 对象判断
"name" in objMap;           // true
```

### 获取大小
```python
# Python
len(person)       # 2
len(unique_nums) # 3
```

```javascript
// JavaScript
person.size;      // 2
uniqueNums.size;  // 3

// 对象需转换
Object.keys(objMap).length;  // 2
```

---

## 4. 转换与迭代

### Python
```python
# 字典 ↔ 列表
items = list(person.items())  # [('name', 'Alice'), ('age', 25)]

# 集合 ↔ 列表
nums_list = list(unique_nums)
```

### JavaScript
```javascript
// Map ↔ 数组
const entries = Array.from(person);  // [["name","Alice"], ["age",25]]

// Set ↔ 数组
const numsArray = [...uniqueNums];   // [1,3,4]

// 迭代方式
for(const [key, val] of person){ /*...*/ }
uniqueNums.forEach(v => console.log(v));
```

---

## 5. 特殊用例

### 对象键比较
```javascript
const map = new Map();
const keyObj = {id: 1};

map.set(keyObj, "对象键");
map.get(keyObj);          // ✅ "对象键"
map.get({id: 1});         // ❌ undefined（不同对象引用）

// 对象字面量键自动转字符串
const obj = {1: "One"};
obj[1] === obj["1"];      // ✅ true
```

### 链式操作
```javascript
const map = new Map()
    .set(1, "a")
    .set(2, "b")
    .set(3, "c");
```

### 数学运算
```javascript
// 集合操作（模拟数学集合）
const setA = new Set([1,2,3]);
const setB = new Set([3,4,5]);

// 并集
new Set([...setA, ...setB]);        // {1,2,3,4,5}

// 交集
new Set([...setA].filter(x => setB.has(x))); // {3}
```

---

## 最佳实践建议

1. **键类型选择**：需要非字符串键时强制使用 Map
2. **性能优化**：大数据量操作优先使用 Map/Set（比对象/数组高效）
3. **数据去重**：使用 Set 替代数组 + filter 去重
4. **内存泄漏**：使用 WeakMap 存储 DOM 节点等易销毁对象
5. **数据转换**：优先使用扩展运算符 `...` 替代 Array.from()
6. **对象陷阱**：避免将用户输入直接作为对象键（可能覆盖原型属性）
7. **类型安全**：使用 Map 存储明确类型的键值对（替代无约束的对象）
8. **迭代顺序**：依赖插入顺序时务必使用 Map/Set（普通对象不保证顺序）
9. **JSON 处理**：转换 Map 到 JSON 时使用 `JSON.stringify([...map])`
10. **浏览器兼容**：旧版 IE 不支持 Map/Set，需要 polyfill

