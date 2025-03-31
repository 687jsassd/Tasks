# JavaScript 与 HTML/CSS 关系学习笔记

> 了解 JavaScript 如何与网页结构（HTML）和样式（CSS）协同工作

---

## 1. 三者的角色分工

| 技术     | 角色           | 类比         |
| -------- | -------------- | ------------ |
| **HTML** | 网页内容与结构 | 建筑骨架     |
| **CSS**  | 网页样式与布局 | 装修风格     |
| **JS**   | 网页行为与交互 | 电器控制系统 |

---

## 2. DOM（文档对象模型）

### 基本概念
- **DOM** = Document Object Model
- 浏览器将 HTML 文档解析为树状结构
- JavaScript 通过 DOM API 操作页面元素

```html
<!-- 示例HTML结构 -->
<div id="app">
    <button class="btn">点击我</button>
    <p>当前计数：<span id="counter">0</span></p>
</div>
```

---

## 3. JavaScript 操作 HTML

### 获取元素
```javascript
// 通过 ID 获取（返回单个元素）
const counterEl = document.getElementById('counter');

// 通过选择器获取（返回第一个匹配元素）
const btn = document.querySelector('.btn');

// 获取多个元素（返回NodeList）
const allButtons = document.querySelectorAll('button');
```

### 修改内容
```javascript
// 修改文本内容
counterEl.textContent = '5'; 

// 修改HTML内容（注意XSS风险）
app.innerHTML = '<strong>更新内容</strong>';

// 修改属性
btn.setAttribute('disabled', true);
```

### 事件处理
```javascript
// 添加点击事件监听
btn.addEventListener('click', function() {
    // 事件处理逻辑
    counterEl.textContent = parseInt(counterEl.textContent) + 1;
});
```

---

## 4. JavaScript 操作 CSS

### 直接修改样式
```javascript
// 修改单个样式属性
btn.style.backgroundColor = 'blue';

// 修改多个样式（推荐class切换）
counterEl.style.cssText = 'color: red; font-size: 20px;';
```

### 类名操作
```css
/* CSS 样式 */
.active {
    transform: scale(1.1);
    box-shadow: 0 2px 5px rgba(0,0,0,0.2);
}
```

```javascript
// 添加类
btn.classList.add('active');

// 移除类
btn.classList.remove('active');

// 切换类
btn.classList.toggle('active');
```

---

## 5. 动态创建元素

### 创建新元素
```javascript
// 创建新段落
const newParagraph = document.createElement('p');
newParagraph.textContent = '动态创建的内容';

// 添加到DOM中
document.body.appendChild(newParagraph);
```

### 克隆元素
```javascript
const template = document.querySelector('#list-item');
const newItem = template.cloneNode(true);
newItem.textContent = '新项目';
document.querySelector('ul').appendChild(newItem);
```

---

## 6. 数据绑定基础

### 简单数据绑定示例
```javascript
const data = { count: 0 };

function updateDisplay() {
    counterEl.textContent = data.count;
    btn.style.opacity = data.count > 5 ? 0.5 : 1;
}

btn.addEventListener('click', () => {
    data.count++;
    updateDisplay();
});
```

---

## 7. 性能注意事项

1. **减少DOM操作**：批量修改时使用`document.createDocumentFragment()`
2. **避免强制同步布局**：
   ```javascript
   // ❌ 不良写法（触发多次重排）
   element.style.width = '100px';
   const width = element.offsetWidth;
   element.style.height = width + 'px';
   
   // ✅ 优化写法（使用变量暂存）
   const newWidth = 100;
   element.style.width = newWidth + 'px';
   element.style.height = newWidth + 'px';
   ```
3. **使用事件委托**：
   ```javascript
   // ✅ 优于为每个按钮单独添加监听
   document.querySelector('.container').addEventListener('click', (e) => {
       if(e.target.matches('button')) {
           // 处理按钮点击
       }
   });
   ```

---

## 最佳实践

1. **关注点分离**：
   - 保持 HTML 结构清晰
   - 样式写在 CSS 中
   - 交互逻辑写在 JS 中

2. **渐进增强原则**：
   - 确保基础功能在不支持JS时仍可用
   ```html
   <form id="search-form">
       <input type="text" name="query">
       <!-- 无JS时的备用提交方式 -->
       <noscript><input type="submit" value="搜索"></noscript>
   </form>
   ```

3. **可访问性**：
   ```javascript
   // 使用 ARIA 属性增强可访问性
   errorMessage.setAttribute('aria-live', 'polite');
   ```

4. **模块化开发**：
   ```javascript
   // 使用现代模块化语法
   import { initCounter } from './counter.js';
   document.addEventListener('DOMContentLoaded', initCounter);
   ```

---

## 调试技巧

1. 使用浏览器开发者工具：
   - Elements 面板查看 DOM
   - Console 面板测试 JS 代码
   - Sources 面板调试 JS

2. 常用调试方法：
   ```javascript
   console.log(element);         // 输出元素信息
   console.dir(element);         // 查看对象属性
   debugger;                     // 设置断点
   element.style.outline = 'red solid 1px'; // 可视化元素边界
   ```

