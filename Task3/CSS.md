# CSS 入门
请配合[MDN CSS](https://developer.mozilla.org/zh-CN/docs/Web/CSS)进行学习。
## 一、CSS 的作用
**核心功能**：控制网页的样式与布局  
**优势**：

- 实现内容(HTML)与表现分离
- 统一管理网站风格
- 支持响应式设计
- 创建动画效果

```html
<!-- 使用前 -->
<h1><font color="red">标题</font></h1>

<!-- 使用后 -->
<style>
h1 { color: red; }
</style>
<h1>标题</h1>
```

---

## 二、基础选择器
| 选择器     | 示例     | 描述                    |
| ---------- | -------- | ----------------------- |
| 通用选择器 | `*`      | 选择所有元素            |
| 元素选择器 | `div`    | 选择所有指定标签元素    |
| 类选择器   | `.class` | 选择class属性匹配的元素 |
| ID选择器   | `#id`    | 选择id属性匹配的元素    |

```css
* { margin: 0; }               /* 清除默认边距 */
h2 { font-size: 20px; }        /* 所有h2元素 */
.btn { background: #f00; }      /* class="btn"的元素 */
#header { height: 60px; }       /* id="header"的元素 */
```

---

## 三、基础样式设置
### 1. 尺寸与颜色
```css
.box {
  width: 300px;         /* 宽度 */
  height: 150px;        /* 高度 */
  background: #e0e0e0;  /* 背景色 */
  color: #333;          /* 文字颜色 */
}
```

### 2. 文字样式
```css
.text {
  font-family: Arial, sans-serif;  /* 字体栈 */
  font-size: 16px;                /* 字号 */
  line-height: 1.5;                /* 行高 */
  text-align: center;             /* 对齐方式 */
}
```

---

## 四、盒模型（Box Model）【重点】
### 1. 组成结构
```css
.element {
  width: 200px;         /* 内容宽度 */
  padding: 20px;        /* 内边距 */
  border: 2px solid #000; /* 边框 */
  margin: 30px;         /* 外边距 */
}
```

### 2. 计算总尺寸
```
总宽度 = width + padding左右 + border左右 + margin左右
总高度 = height + padding上下 + border上下 + margin上下
```

### 3. 盒模型模式
```css
/* 标准盒模型（默认） */
box-sizing: content-box;

/* 边框盒模型（更易控制尺寸） */
box-sizing: border-box;
```

---

## 五、定位（Position）【重点】
| 定位方式 | 属性值   | 特点                     |
| -------- | -------- | ------------------------ |
| 静态定位 | static   | 默认值，不触发定位       |
| 相对定位 | relative | 相对于元素原位置偏移     |
| 绝对定位 | absolute | 相对于最近的定位祖先元素 |
| 固定定位 | fixed    | 相对于浏览器窗口固定     |
| 粘性定位 | sticky   | 在阈值范围内保持固定     |

```css
.item {
  position: relative;
  top: 10px;
  left: 20px;
}

.modal {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
```

---

## 六、布局方式
### 1. 常规流（Normal Flow）
默认布局方式，元素按HTML顺序排列

### 2. 浮动布局（Float）
```css
.left-col {
  float: left;
  width: 30%;
}

.right-col {
  float: right;
  width: 70%;
}
```

### 3. Flex布局（推荐）
```css
.container {
  display: flex;
  justify-content: space-between;
}
```

### 4. Grid布局（现代方案）
```css
.grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
}
```

---

## 七、常用伪类
| 伪类          | 描述                 |
| ------------- | -------------------- |
| :hover        | 鼠标悬停状态         |
| :active       | 激活状态（点击瞬间） |
| :nth-child(n) | 第n个子元素          |
| :first-child  | 第一个子元素         |
| :last-child   | 最后一个子元素       |

```css
/* 按钮交互效果 */
.btn:hover {
  background: #c00;
  transform: translateY(-2px);
}

/* 表格条纹效果 */
tr:nth-child(even) {
  background: #f5f5f5;
}
```

---

## 综合示例
```html
<style>
  .card {
    width: 300px;
    padding: 20px;
    border: 1px solid #ddd;
    border-radius: 8px;
    margin: 20px;
    position: relative;
  }
  
  .card:hover {
    box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  }
  
  .badge {
    position: absolute;
    top: -10px;
    right: -10px;
    background: #f00;
    color: white;
    padding: 5px 10px;
    border-radius: 15px;
  }
</style>

<div class="card">
  <span class="badge">New</span>
  <h3>产品标题</h3>
  <p>产品描述内容...</p>
</div>
```

---

## 学习建议
1. 使用Chrome开发者工具调试样式
2. 通过[CodePen](https://codepen.io/)实践练习
3. 掌握CSS调试技巧：
   ```css
   * { outline: 1px solid red; } /* 可视化所有元素边界 */
   ```
4. 阅读MDN官方文档：[MDN CSS](https://developer.mozilla.org/zh-CN/docs/Web/CSS)

