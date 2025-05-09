## 1. 井号的使用(标题Header)
1. md支持添加**6级标题**，分别用1-6个井号作为前缀以应用。渲染时，井号前缀省略。
2. \#并不直接控制字号，而是一种语法结构。
3. 以上打出井号是因为使用了**转义符**\\
4. 对于典型字号，见下面的表格↓  注意典型字号并不绝对，实际渲染依软件不同而变化。

| 标题级别 | 典型字号 | HTML等效标签 |
| ---- | ---- | -------- |
| 一级   | 32px | `<h1>`   |
| 二级   | 28px | `<h2>`   |
| 三级   | 24px | `<h3>`   |
| 四级   | 20px | `<h4>`   |
| 五级   | 16px | `<h5>`   |
| 六级   | 14px | `<h6>`   |
5. 注意：井号后加空格才可生效
6. 其他的标题定义方式：使用=或者-以标记一二级标题。

## 2.段落格式
1. 可以使用两个或以上的末尾空格加回车以重起一段，但是更常用的还是直接空一行。
2. md没有特别要求的段落格式，以上是实践建议而不是硬性限制。  

## 3.字体
- md本身（原生）不支持直接设置字体，但是通过一些语法可以间接控制文本样式。
1. 强调：
    - 强调分为粗体和斜体（以及斜粗体）
    - 粗体**shili**表示 -> \*\* xxx \*\*
    - 斜体*shili2*表示 -> \* xx \*
    - 粗斜体***shili***表示 -> \*\*\* x \*\*\*
    - 星号\*可以用下划线\_替换，全部替换时效果不变，但是不能部分替换。
2. 分割线：
    这些写法都可以制造分割线：
    \* \* \*  三个星号 空格分隔
    `*****`  五联星号
    `- - -` 三横线 空格分隔
3. 删除线：
    在文字始末各加入两个波浪号即可
    `~~你好~~` -> ~~你好~~
4. 脚注：
    使用`[^]`来创建脚注
    [^1脚注]

[^1]: 这是一个脚注qwq
	

## 4.列表
- md支持有序和无序列表两种，其中无序列表只需要用星号，加号，减号作为前缀即可创建。有序列表则用数字加点号以创建。
- 范例：`- 这是无序列表` `1. 这是有序列表`
- 有序列表的顺序会自动调整，如我输入：
    `1. 1`
    `2. 2`
    `5. 555`
    那么实际上`5. 555`会被自动调整为`3. 555`
- 列表嵌套只需在子列表中的选项前面添加两个或四个空格即可。

## 5.区块声明
> 区别于前面的格式，这里使用了区块声明，略微更加美观一些。
> 
> 我们使用`>`来声明一个区块
> > 区块声明可以嵌套，仅需叠加大于号字符即可。
> 
> 理论上，区块声明不影响列表使用，但是至少在obsidian编辑器中是无法渲染的。

## 6.代码块
> 使用\` 来包裹代码，当然一般情况下也可以直接当字符串的引号使用。
> 
> 使用三个反引号来包裹多行的代码块，见下面的例子：
>  > ```python
>  > print('hello world')
>  > ```

## 7. 链接
- 使用`[链接文本](https://example.com)`的格式来创建超链接，点击该链接即可跳转对应网站。
    这是一个示例： [带标题链接](https://example.com "点击这里")  
- 或者，如果不想要标题，也可以采用这个格式：`<链接地址>`。
    一个示例：<https://www.baidu.com >

## 8. 图片
- 使用这种格式插入图片：
    `![图片描述](https://example.com/image.jpg)`
    或者，如果你想要标题↓↓
    `![Alt文本](https://example.com/image.jpg "可选标题")`
- 注意，如果需要添加本地图片，应该把地址部分改为图片路径。
- 这是两个示例，分别表示本地图片和非本地图片：
    ![本地图片](1.jpg "本地图片")
    ![远程图片](https://img2.baidu.com/it/u=3498204231,2724330704&fm=253&fmt=auto&app=138&f=JPEG?w=800&h=1040 "远程标题")
    `注：github在线显示的情况下，远程图片可能不能正确获取？`

## 9.表格
- 使用下面的格式创建表格:
    | 表头1 | 表头2 |
    |-------|-------|
    | 单元格1 | 单元格2 |
    | 单元格3 | 单元格4 |
- 特别地，可以修改表头和单元格间的横线列，前加:则左对齐，尾加:则右对齐，两端都加则中心对齐。

