# Git 基本操作指南

## 目录
1. [环境准备](#环境准备)
2. [仓库操作](#仓库操作)
3. [文件追踪](#文件追踪)
4. [提交管理](#提交管理)
5. [分支操作](#分支操作)
6. [远程协作](#远程协作)
7. [版本回退](#版本回退)
8. [实用技巧](#实用技巧)
9. [命令速查表](#命令速查表)

---

## 环境准备

### 安装与配置
```bash
# 安装Git后初始配置
git config --global user.name "Your Name"
git config --global user.email "your@email.com"
```

### 重要概念

| 概念     | 说明                     | 类比           |
| :------- | :----------------------- | :------------- |
| 工作区   | 直接看到的文件目录       | 办公桌         |
| 暂存区   | 已标记改动的文件列表     | 待处理的文件篮 |
| 本地仓库 | 存储完整历史记录的数据库 | 文件归档柜     |
| 远程仓库 | 云端共享的代码仓库       | 公司文件服务器 |

------

## 仓库操作

### 初始化仓库

```bash
# 创建新仓库
git init 
# 克隆现有仓库
git clone https://github.com/user/repo.git
```

### 仓库状态检查

```bash
git status        # 查看当前状态
git log           # 查看提交历史
git log --oneline # 简洁版历史
```

------

## 文件追踪

### 基础操作流程图

```mermaid
graph LR
    A[工作区修改] --> B[git add]
    B --> C[暂存区]
    C --> D[git commit]
    D --> E[本地仓库]
```

### 常用命令

```bash
git add .                     # 添加所有改动
git add filename              # 添加特定文件
git reset HEAD filename       # 取消暂存
git rm --cached filename     # 停止追踪文件
```

### .gitignore 文件

```bash
# 示例：忽略所有.class文件
*.class
# 忽略特定目录
/node_modules/
# 不忽略特殊文件
!important.class
```

------

## 提交管理

### 标准提交

```bash
git commit -m "描述性信息"    # 快速提交
git commit                   # 打开编辑器编写详细说明
```

### 提交规范建议

```bash
feat: 添加用户登录功能
^--^  ^------------^
|     |
|     ➔ 简要描述（英文小写开头，不加句号）
|
➔ 提交类型（feat/fix/docs/style/refactor/test等）
```

------

## 分支操作

### 基础分支管理

```bash
git branch                   # 查看本地分支
git branch new-feature       # 创建新分支
git checkout develop         # 切换分支
git checkout -b hotfix       # 创建并切换分支
git branch -d old-branch     # 删除分支
```

### 合并与冲突

```bash
git merge feature-branch     # 合并指定分支到当前分支
git mergetool                # 使用可视化工具解决冲突
```

------

## 远程协作

### 远程仓库配置

```bash
git remote add origin https://github.com/user/repo.git
git remote -v                # 查看远程仓库
git fetch origin             # 获取远程更新（不合并）
git pull origin main         # 拉取并合并远程更新
```

### 推送代码

```bash
git push -u origin main      # 首次推送建立关联
git push                     # 后续简化推送
```

------

## 版本回退

### 撤销操作

```bash
git reset HEAD~1             # 撤销最近一次提交（保留修改）
git reset --hard HEAD~1      # 彻底回退到前一个版本
git revert HEAD             # 创建反向提交来撤销改动
```

### 版本穿梭

```bash
git checkout 3a8b5d2        # 切换到指定提交
git checkout main           # 返回主分支
```

------

## 实用技巧

### 储藏临时修改

```bash
git stash        # 保存当前工作进度
git stash pop    # 恢复最近储藏的内容
```

### 比较差异

```bash
git diff                  # 工作区与暂存区差异
git diff --staged         # 暂存区与最新提交差异
git diff branch1..branch2 # 比较两个分支差异
```

------

## 命令速查表

| 操作场景         | 常用命令                  |
| :--------------- | :------------------------ |
| 初始化仓库       | `git init`                |
| 克隆仓库         | `git clone <url>`         |
| 添加文件到暂存区 | `git add <file>`          |
| 提交更改         | `git commit -m "message"` |
| 查看状态         | `git status`              |
| 查看提交历史     | `git log`                 |
| 创建分支         | `git branch <branch>`     |
| 切换分支         | `git checkout <branch>`   |
| 合并分支         | `git merge <branch>`      |
| 拉取远程更新     | `git pull`                |
| 推送本地提交     | `git push`                |
| 撤销工作区修改   | `git checkout -- <file>`  |
| 查看远程仓库     | `git remote -v`           |

------

> **最佳实践提示**：
>
> 1. 保持提交原子化（每次提交只做一件事）
> 2. 每天开始工作前先执行 `git pull`
> 3. 重要分支（如main）启用分支保护
> 4. 使用 `.gitignore` 管理不需要追踪的文件
> 5. 提交信息采用规范格式（如Conventional Commits）