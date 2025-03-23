# Git 文件追踪的注意事项

## 问题场景还原
假设你的仓库结构如下：

project/
├── .git/
├── tracked_file.txt
└── new_file.txt（通过资源管理器新建）

执行操作：

```bash
git commit -m "添加新文件"
```

## 实际结果分析

1. **新文件未被追踪**

   - `new_file.txt` 不会包含在本次提交中

   - 执行 `git status` 会显示：

     ```bash
     Untracked files:
       (use "git add <file>..." to include in what will be committed)
             new_file.txt
     ```

2. **潜在风险**

   - 文件可能永久丢失（如果误删且未提交）
   - 其他协作者无法获取该文件
   - 可能导致后续操作混乱（如误判文件存在状态）

## 正确操作流程

### 标准工作流

```mermaid
graph LR
    A[工作区新建文件] --> B[git add]
    B --> C[暂存区]
    C --> D[git commit]
    D --> E[版本库]
```

### 具体步骤

1. 创建新文件后执行：

   ```bash
   git add new_file.txt
   # 或添加所有新文件
   git add .
   ```

2. 确认暂存状态：

   ```bash
   git status
   # 应显示：
   # Changes to be committed:
   #   new file:   new_file.txt
   ```

3. 执行提交：

   ```bash
   git commit -m "正确添加新文件"
   ```

## 常见错误场景

### 场景1：忘记添加直接提交

```bash
# 补救措施（修改最后一次提交）
git add new_file.txt
git commit --amend
```

### 场景2：误添加大文件/敏感文件

```bash
# 从暂存区移除但不删除文件
git rm --cached huge_file.zip
# 添加到.gitignore
echo "huge_file.zip" >> .gitignore
```

## 最佳实践建议

1. **养成检查习惯**：每次提交前运行 `git status`
2. **分步操作**：按文件类型分批添加（如 `git add *.html`）
3. **使用.gitignore**：配置自动过滤不需要追踪的文件
4. **IDE整合**：使用VSCode等工具的Git可视化界面

> 重要原则：Git不会自动追踪任何新文件，必须显式告知需要追踪哪些文件！