# 设置自动触发 Webhook

要让文档在结构体改动时自动更新，需要在 **common-events 仓库** 中设置一个 webhook 工作流。

## 步骤 1: 在 common-events 仓库中创建工作流

1. 在 common-events 仓库中创建文件：
   - 路径：`.github/workflows/notify-docs-update.yml`
   - 内容：复制 `common-events-webhook.yml` 文件的内容

## 步骤 2: 配置 GitHub Secrets

在 **common-events 仓库** 的 Settings → Secrets and variables → Actions 中添加：

### 必需的 Secrets:

1. **DOCS_REPO_TOKEN**
   - 值：一个有写入权限的 GitHub Personal Access Token
   - 这个 token 需要有权限在文档仓库中触发 `repository_dispatch` 事件
   - 建议创建一个 Fine-grained Personal Access Token，只给文档仓库的权限

2. **DOCS_REPO**
   - 值：文档仓库的完整名称（例如：`your-org/test1111111` 或 `your-username/test1111111`）

### 如何创建 Personal Access Token:

1. 登录 GitHub
2. 点击右上角头像 → Settings
3. 左侧菜单 → Developer settings
4. Personal access tokens → Tokens (classic) 或 Fine-grained tokens
5. 生成新 token，需要的权限：
   - 对于 Fine-grained token：
     - Repository access: 选择文档仓库
     - Repository permissions:
       - Actions: Read and write (用于触发 workflow)
       - Contents: Read (用于读取仓库)
   - 对于 Classic token：
     - `repo` scope (完整仓库访问权限)

## 步骤 3: 测试

1. 在 common-events 仓库中修改 `events/` 目录下的任意 `.go` 文件
2. 提交并推送到 main/master/develop 分支
3. 检查文档仓库的 Actions 标签页，应该会自动触发 "Update Event Documentation" 工作流

## 工作原理

```
common-events 仓库推送事件
    ↓
检测到 events/**/*.go 文件变更
    ↓
触发 notify-docs-update.yml 工作流
    ↓
发送 repository_dispatch 事件到文档仓库
    ↓
文档仓库的 update-event-docs.yml 工作流自动运行
    ↓
生成并提交更新的 PlantUML 文档
```

## 故障排除

### 工作流没有自动触发

1. **检查路径过滤**：
   - 确保修改的文件路径匹配 `events/**/*.go`
   - 只修改 `.go` 文件才会触发

2. **检查分支**：
   - 确保推送到了 `main`、`master` 或 `develop` 分支
   - 其他分支不会触发

3. **检查 Secrets**：
   - 验证 `DOCS_REPO_TOKEN` 和 `DOCS_REPO` 是否正确配置
   - 确认 token 有正确的权限

4. **查看日志**：
   - 在 common-events 仓库的 Actions 标签页查看 `notify-docs-update` 工作流的执行日志
   - 在文档仓库的 Actions 标签页查看 `update-event-docs` 工作流的执行日志

### Token 权限问题

如果遇到权限错误，确保 token 有：
- 对文档仓库的 `repository_dispatch` 事件触发权限
- 如果使用 Fine-grained token，确保选择了正确的仓库和权限范围

## 备用方案：定时检查

即使 webhook 失败，工作流仍然会每 6 小时运行一次（通过 `schedule`），确保文档定期更新。

## 安全提示

- 不要将 token 提交到代码仓库
- 定期轮换 token
- 使用 Fine-grained token 以最小权限原则
- 考虑使用 GitHub App 替代 Personal Access Token（更安全）

