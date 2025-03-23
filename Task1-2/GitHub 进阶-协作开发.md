# GitHub 进阶-协作开发

## 目录
1. [协作开发模式概述](#协作开发模式概述)
2. [Fork 工作流详解](#fork-工作流详解)
3. [Pull Request 全流程](#pull-request-全流程)
4. [高级协作技巧](#高级协作技巧)
5. [常见问题处理](#常见问题处理)

---

## 协作开发模式概述

### 两种主要协作模式
1. **共享仓库模式**  
   - 特点：所有协作者拥有直接推送权限
   - 适用场景：小型团队/私有项目
   - 权限管理：通过仓库设置控制

2. **Fork & Pull 模式**  
   - 特点：开发者独立维护自己的仓库副本
   - 适用场景：开源项目/大型协作
   - 流程：  
     ```mermaid
     graph LR
         A[主仓库] --> B(Fork)
         B --> C[个人仓库]
         C --> D[本地开发]
         D --> E(Pull Request)
         E --> A
     ```

---

## Fork 工作流详解

### Fork 操作的本质
- 创建主仓库的独立副本（包含所有分支和提交历史）
- 完全独立的开发环境（不影响原始仓库）

### 标准工作流程
1. **Fork 主仓库**
   - 在 GitHub 页面点击右上角 `Fork` 按钮
   - 选择个人账号作为目标位置

2. **克隆本地副本**
   
   ```bash
   git clone https://github.com/你的账号/仓库名.git
   cd 仓库名
   ```
   
   3. **添加上游仓库**
   
   ```bash
   git remote add upstream https://github.com/原始作者/仓库名.git
   # 验证配置
   git remote -v
   ```
   
   4.**保持仓库同步**
   
   ```bash
   # 获取上游更新
   git fetch upstream
   # 合并到本地分支
   git checkout main
   git merge upstream/main
   ```
   
   ------
   
   ## Pull Request 全流程
   
   ### PR 生命周期
   
   ```mermaid
   sequenceDiagram
       参与者->>个人仓库: 开发提交
       个人仓库->>主仓库: 发起 PR
       主仓库维护者-->>PR: 代码审查
       PR-->>参与者: 请求修改
       参与者->>PR: 追加提交
       主仓库维护者->>主仓库: 合并 PR
   ```
   
   ### 创建优质 PR 的要点
   
   1. **分支规范**
   
      ```bash
      # 推荐分支命名格式
      feat/add-login-module
      fix/issue-123
      ```
   
   2. **提交信息**
   
      ```markdown
      ### 修改类型（必填）
      - feat: 新增功能
      - fix: 错误修复
      - docs: 文档更新
      - style: 代码格式调整
      
      ### 修改描述（必填）
      添加用户登录模块，包含手机号验证功能
      
      ### 关联Issue（可选）
      Close #123
      ```
   
   3. **代码审查准备**
   
      - 确保通过所有测试
      - 更新相关文档
      - 保持代码风格统一
   
   ------
   
   ## 高级协作技巧
   
   ### 1. 多分支协同开发
   
   ```bash
   # 基于上游仓库创建特性分支
   git checkout -b feature/new-module upstream/main
   
   # 定期同步上游更新
   git pull --rebase upstream main
   ```
   
   ### 2. PR 优化策略
   
   - **原子化提交**：每个 PR 只解决一个问题
   - **可视化辅助**：添加截图/屏幕录制
   - **模板应用**：使用 `.github/PULL_REQUEST_TEMPLATE.md`
   
   ### 3. 冲突解决方案
   
   ```bash
   # 当PR出现冲突时
   git fetch upstream
   git checkout feature-branch
   git rebase upstream/main
   # 手动解决冲突后
   git push --force-with-lease
   ```
   
   ------
   
   ## 常见问题处理
   
   ### Q1: Fork 后如何同步主仓库新分支？
   
   ```bash
   git fetch upstream
   git checkout -b new-branch upstream/new-branch
   ```
   
   ### Q2: PR 提交错误分支怎么办？
   
   1. 在本地使用 `git cherry-pick` 转移提交
   
   2. 强制推送到正确分支
   
      ```bash
      git push origin feature-branch --force-with-lease
      ```
   
   ### Q3: 如何撤回已合并的PR？
   
   1. 使用 revert 生成反向提交
   
      ```bash
      git revert -m 1 <合并提交的哈希>
      ```
   
   2. 创建新的 PR 进行撤回
   
   ------
   
   ## 最佳实践清单
   
   1. 每次开发新功能都创建独立分支
   2. 保持 `main` 分支与上游仓库同步
   3. 使用 `--force-with-lease` 代替强制推送
   4. PR 描述遵循项目规范模板
   5. 合并前执行 `git pull --rebase`