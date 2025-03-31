# JavaScript 面向对象编程对比手册（C/Python→JS）

> 适用于有 C 或 Python 基础的 JS 初学者，重点对比 OOP 实现机制与设计模式差异

---

## 1. 基本对象创建

### C 语言
```c
// 结构体模拟对象
typedef struct {
    float x;
    float y;
} Point;

// 独立函数作为"方法"
void movePoint(Point *p, float dx, float dy) {
    p->x += dx;
    p->y += dy;
}

Point p1 = {1.0, 2.0};
movePoint(&p1, 0.5, -0.5);
```

### Python
```python
class Point:
    def __init__(self, x, y):
        self.x = x
        self.y = y
    
    def move(self, dx, dy):
        self.x += dx
        self.y += dy

p1 = Point(1.0, 2.0)
p1.move(0.5, -0.5)
```

### JavaScript
```javascript
// 构造函数模式
function Point(x, y) {
    this.x = x;
    this.y = y;
    
    this.move = function(dx, dy) {
        this.x += dx;
        this.y += dy;
    }
}
const p1 = new Point(1.0, 2.0);

// ES6 类语法
class Point {
    constructor(x, y) {
        this.x = x;
        this.y = y;
    }
    
    move(dx, dy) {
        this.x += dx;
        this.y += dy;
    }
}
```

---

## 2. 继承机制

### C 语言
```c
// 组合模拟继承
typedef struct {
    Point base;  // 基类组合
    float z;
} Point3D;

void move3D(Point3D *p, float dx, float dy, float dz) {
    movePoint(&(p->base), dx, dy);
    p->z += dz;
}
```

### Python
```python
class Point3D(Point):
    def __init__(self, x, y, z):
        super().__init__(x, y)
        self.z = z
    
    def move(self, dx, dy, dz):
        super().move(dx, dy)
        self.z += dz
```

### JavaScript
```javascript
// 原型链继承（旧方式）
function Point3D(x, y, z) {
    Point.call(this, x, y);
    this.z = z;
}
Point3D.prototype = Object.create(Point.prototype);

// ES6 类继承
class Point3D extends Point {
    constructor(x, y, z) {
        super(x, y);
        this.z = z;
    }
    
    move(dx, dy, dz) {
        super.move(dx, dy);
        this.z += dz;
    }
}
```

---

## 3. 多态实现

### C 语言
```c
// 函数指针实现多态
typedef struct {
    void (*draw)(void);  // 虚函数指针
} Shape;

void drawCircle() { /* 圆形绘制逻辑 */ }
void drawSquare() { /* 方形绘制逻辑 */ }

Shape circle = {drawCircle};
Shape square = {drawSquare};
```

### Python
```python
# 鸭子类型实现多态
class Circle:
    def draw(self):
        print("绘制圆形")

class Square:
    def draw(self):
        print("绘制方形")

def render(shapes):
    for shape in shapes:
        shape.draw()
```

### JavaScript
```javascript
// 原型链多态
class Shape {
    draw() {
        throw new Error("抽象方法需要实现");
    }
}

class Circle extends Shape {
    draw() { console.log("绘制圆形"); }
}

class Square extends Shape {
    draw() { console.log("绘制方形"); }
}

function render(shapes) {
    shapes.forEach(shape => shape.draw());
}
```

---

## 4. 静态成员

### C 语言
```c
// 全局变量模拟静态成员
int instanceCount = 0;

void createObject() {
    instanceCount++;
}
```

### Python
```python
class Logger:
    _instance = None  # 类变量作为静态成员
    
    @classmethod
    def get_instance(cls):
        if not cls._instance:
            cls._instance = Logger()
        return cls._instance
```

### JavaScript
```javascript
class Logger {
    static instanceCount = 0;  // 静态字段（ES2022+）
    
    constructor() {
        Logger.instanceCount++;
    }
    
    static getInstance() {    // 静态方法
        if (!this._instance) {
            this._instance = new Logger();
        }
        return this._instance;
    }
}
```

---

## 5. 私有成员

### C 语言
```c
// 无真正封装，通过命名约定
typedef struct {
    int public_data;
    int _private_data;  // 下划线约定私有
} Data;
```

### Python
```python
class SecureData:
    def __init__(self):
        self.__secret = 100  # 名称修饰为 _SecureData__secret
    
    def get_secret(self):
        return self.__secret
```

### JavaScript
```javascript
// 闭包私有（旧方式）
function SecureData() {
    let secret = 100;
    this.getSecret = () => secret;
}

// 类私有字段（ES2022+）
class SecureData {
    #secret = 100;
    
    getSecret() {
        return this.#secret;
    }
}
```

---

## 关键差异总结表

| 特性     | C 语言         | Python       | JavaScript             |
| -------- | -------------- | ------------ | ---------------------- |
| 对象基础 | 结构体 + 函数  | class 关键字 | 构造函数/ES6 类        |
| 继承机制 | 组合模拟       | 类继承       | 原型链继承/ES6 extends |
| 多态实现 | 函数指针       | 鸭子类型     | 原型方法覆盖           |
| 封装支持 | 命名约定       | 名称修饰     | 闭包/#私有字段         |
| 静态成员 | 全局变量       | 类变量       | static 关键字          |
| 实例检查 | 无             | isinstance() | instanceof             |
| 抽象类   | 无             | ABC 模块     | 无原生支持（可用 TS）  |
| 接口支持 | 无             | Protocol     | 无原生支持（可用 TS）  |
| 构造方法 | 独立初始化函数 | __init__     | constructor            |
| 内存管理 | 手动分配/释放  | 自动 GC      | 自动 GC                |

---

## 最佳实践建议

1. **类语法优先**：始终使用 ES6 类语法替代原型操作
2. **组合优于继承**：使用 Object.assign() 实现混入模式
3. **私有字段**：现代项目中使用 # 语法替代闭包私有
4. **类型检查**：结合 TS 实现接口和抽象类功能
5. **不可变模式**：使用 Object.freeze() 保护重要对象
6. **工厂函数**：替代复杂的构造函数逻辑
7. **原型污染**：避免修改内置原型（如 Array.prototype）
8. **内存优化**：使用 WeakMap 管理私有数据
9. **方法绑定**：类方法中使用箭头函数或 bind() 保持 this 上下文
10. **模块化设计**：拆分类功能到不同模块（ES Modules）

