# HTML 常用标签详解与示例

## 一、基础容器标签

### 1. `<div>` 标签
- **作用**：块级容器，用于页面布局与内容分组（无语义含义）
- **特点**：
  - 默认占满父容器宽度
  - 常用于 CSS 样式布局或 JavaScript 操作
- **示例**：
```html
<div class="container">
  <h2>产品列表</h2>
  <div class="product-item">...</div>
</div>
```

---

## 二、语义化结构标签（HTML5）

### 2. `<header>`
- **作用**：定义文档或区块的头部区域
- **示例**：
```html
<header>
  <h1>网站标题</h1>
  <nav>...</nav>
</header>
```

### 3. `<nav>`
- **作用**：定义导航链接集合
- **示例**：
```html
<nav>
  <ul>
    <li><a href="/">首页</a></li>
    <li><a href="/about">关于我们</a></li>
  </ul>
</nav>
```

### 4. `<section>`
- **作用**：定义文档中的独立内容区块
- **示例**：
```html
<section>
  <h2>最新公告</h2>
  <p>2023年度报告已发布...</p>
</section>
```

### 5. `<article>`
- **作用**：定义独立的内容单元（如博客文章、新闻稿）
- **示例**：
```html
<article>
  <h3>人工智能新突破</h3>
  <p>最新研究显示...</p>
</article>
```

### 6. `<aside>`
- **作用**：定义与主内容相关的侧边内容
- **示例**：
```html
<aside>
  <h4>相关链接</h4>
  <ul>...</ul>
</aside>
```

### 7. `<footer>`
- **作用**：定义文档或区块的底部区域
- **示例**：
```html
<footer>
  <p>© 2023 公司名称</p>
  <address>联系方式：contact@example.com</address>
</footer>
```

---

## 三、文本内容标签

### 8. `<p>`
- **作用**：段落文本
- **示例**：
```html
<p>这是第一个段落。</p>
<p>这是另一个独立的段落。</p>
```

### 9. `<span>`
- **作用**：行内容器，用于文本样式控制
- **示例**：
```html
<p>价格：<span class="price">¥299</span></p>
```

### 10. 标题标签 `<h1>`-`<h6>`
- **作用**：定义标题层级（h1最高级）
- **示例**：
```html
<h1>主标题</h1>
<h2>二级标题</h2>
<h3>三级标题</h3>
```

---

## 四、列表标签

### 11. 无序列表 `<ul>` + `<li>`
- **作用**：创建无顺序项目列表
- **示例**：
```html
<ul>
  <li>咖啡</li>
  <li>茶</li>
  <li>牛奶</li>
</ul>
```

### 12. 有序列表 `<ol>` + `<li>`
- **作用**：创建有顺序编号的列表
- **示例**：
```html
<ol>
  <li>切碎蔬菜</li>
  <li>热锅放油</li>
  <li>翻炒食材</li>
</ol>
```

---

## 五、媒体元素

### 13. `<img>`
- **作用**：嵌入图像
- **重要属性**：
  - `src`：图片路径（必需）
  - `alt`：替代文本（必需）
- **示例**：
```html
<img src="logo.png" alt="公司Logo" width="200">
```

### 14. `<video>`
- **作用**：嵌入视频
- **示例**：
```html
<video controls width="600">
  <source src="movie.mp4" type="video/mp4">
  您的浏览器不支持视频标签
</video>
```

---

## 六、表格标签

### 15. `<table>` 系列
```html
<table>
  <thead>
    <tr>
      <th>商品</th>
      <th>价格</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>鼠标</td>
      <td>¥99</td>
    </tr>
    <tr>
      <td>键盘</td>
      <td>¥199</td>
    </tr>
  </tbody>
</table>
```

---

## 七、交互元素

### 16. `<a>`
- **作用**：创建超链接
- **示例**：
```html
<a href="https://example.com" target="_blank">访问示例网站</a>
```

### 17. `<button>`
- **作用**：创建可点击按钮
- **示例**：
```html
<button type="button" onclick="alert('点击成功!')">立即购买</button>
```

---

## 八、使用原则

1. **语义优先**：
```html
<!-- 错误示例 -->
<div class="header">...</div>

<!-- 正确示例 -->
<header>...</header>
```

2. **正确嵌套**：
```html
<!-- 错误示例 -->
<p>文字 <div>错误嵌套</div></p>

<!-- 正确示例 -->
<div>
  <p>正确的嵌套结构</p>
</div>
```

3. **属性规范**：
```html
<!-- 正确使用布尔属性 -->
<input type="checkbox" checked>

<!-- 正确使用自定义属性 -->
<div data-user-id="123"></div>
```

4. **可访问性**：
```html
<!-- 为图标按钮添加说明 -->
<button aria-label="关闭弹窗">×</button>
```

---

## 完整页面结构示例
```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <meta charset="UTF-8">
  <title>示例页面</title>
</head>
<body>
  <header>
    <h1>电子商务平台</h1>
    <nav>
      <ul>
        <li><a href="/">首页</a></li>
        <li><a href="/products">商品分类</a></li>
      </ul>
    </nav>
  </header>

  <main>
    <section>
      <h2>热销商品</h2>
      <article class="product">
        <img src="product1.jpg" alt="无线耳机">
        <h3>降噪无线耳机</h3>
        <p>¥599</p>
        <button>加入购物车</button>
      </article>
    </section>
  </main>

  <footer>
    <p>© 2023 电商平台</p>
  </footer>
</body>
</html>
```

## 总结要点
1. `<div>` 和 `<span>` 作为通用容器，应根据内容选择更语义化的标签
2. 合理使用 HTML5 语义标签提升 SEO 和可访问性
3. 所有图片必须包含 `alt` 属性
4. 保持标签的合理嵌套关系
5. 交互元素应包含必要的 ARIA 属性

