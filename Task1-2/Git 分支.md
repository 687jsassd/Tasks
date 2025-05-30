# Git 分支完全指南

## 目录
1. [什么是分支？](#什么是分支)
2. [何时创建/不创建分支](#何时创建不创建分支)
3. [分支管理实践](#分支管理实践)
4. [分支合并机制](#分支合并机制)
5. [常见问题与误区](#常见问题与误区)
6. [实用命令速查](#实用命令速查)

---

## 什么是分支？
### 核心概念
- **分支本质**：指向提交对象的可变指针
- **默认分支**：通常为 `main` 或 `master`
- **HEAD指针**：指向当前所在分支的最新提交

### 可视化理解
```bash
          main
           ↓
A ← B ← C ← D
     ↗
feature/login
```



## 何时创建/不创建分支

### 应该创建新分支的场景

| 场景        | 示例                 | 优点                     |
| :---------- | :------------------- | :----------------------- |
| 开发新功能  | `feat/user-profile`  | 隔离功能开发不影响主分支 |
| 修复紧急bug | `hotfix/payment-bug` | 快速修复生产环境问题     |
| 实验性尝试  | `experiment/new-ui`  | 避免污染稳定代码库       |
| 多人协作    | `dev/john/feature-x` | 防止代码冲突             |

### 无需创建分支的情况

1. 修改文档/配置文件等简单变更
2. 项目规模极小且单人开发
3. 立即要合并的微小改动

------

## 分支管理实践

### 分支类型

| 分支类型   | 生命周期 | 典型命名         | 用途             |
| :--------- | :------- | :--------------- | :--------------- |
| 主分支     | 永久     | `main`, `master` | 稳定可发布的代码 |
| 开发分支   | 长期     | `develop`        | 日常开发集成     |
| 功能分支   | 短期     | `feat/search`    | 特定功能开发     |
| 热修复分支 | 超短期   | `hotfix/xxx`     | 紧急生产问题修复 |
| 发布分支   | 中期     | `release/v1.2`   | 版本发布准备     |

### 最佳实践

1. **命名规范**：使用 `/` 分隔分类（如 `feat/docs/fix`）
2. **及时清理**：合并后删除已结束的分支
3. **定期同步**：`git pull origin main` 保持基础分支更新
4. **小步提交**：每个分支保持专注的修改范围

------

## 分支合并机制

### 合并类型

1. **快进合并（Fast-Forward）**

   - 条件：目标分支没有新提交
   - 结果：直接移动分支指针

   ```bash
   A ← B ← C (main)
        ↗
       D (feature)
   # 合并后：A ← B ← C ← D
   ```

2. **三方合并（3-Way Merge）**

   - 条件：分支出现分叉
   - 结果：生成新的合并提交

   ```bash
   A ← B ← C (main)
     ↖
       D ← E (feature)
   # 合并后生成新提交 F
   ```

### 合并冲突处理

1. 冲突标识：

   ```bash
   <<<<<<< HEAD
   main分支的内容
   =======
   feature分支的内容
   >>>>>>> feature
   ```

2. 解决步骤：

   - 运行 `git status` 查看冲突文件
   - 手动编辑文件保留需要的内容
   - 删除冲突标记（`<<<`, `===`, `>>>`）
   - 执行 `git add` 和 `git commit`

------

## 常见问题与误区

### Q1: 分支会占用额外存储空间吗？

- **不会**：Git分支只是指针，真正存储的是提交对象

### Q2: 可以同时工作在多个分支吗？

- **不能**：一次只能在一个分支工作（通过`checkout`切换）

### Q3: 误删分支如何恢复？

```bash
# 查看最近所有提交的哈希值
git reflog
# 根据哈希重建分支
git branch <分支名> <提交哈希>
```

### Q4: 远程分支和本地分支的区别？

- 远程分支格式：`remotes/origin/<分支名>`
- 需要显式推送：`git push -u origin <分支名>`

------

## 实用命令速查

| 操作                 | 命令                                |
| :------------------- | :---------------------------------- |
| 创建新分支           | `git branch <分支名>`               |
| 切换分支             | `git checkout <分支名>`             |
| 创建并切换分支       | `git checkout -b <分支名>`          |
| 查看所有分支         | `git branch -a`                     |
| 删除本地分支         | `git branch -d <分支名>`            |
| 强制删除未合并分支   | `git branch -D <分支名>`            |
| 删除远程分支         | `git push origin --delete <分支名>` |
| 合并分支到当前分支   | `git merge <分支名>`                |
| 变基操作（谨慎使用） | `git rebase <目标分支>`             |

------

> **重要提示**：在团队协作中，始终遵循项目的分支管理规范！当不确定时，先创建分支总是更安全的选择。