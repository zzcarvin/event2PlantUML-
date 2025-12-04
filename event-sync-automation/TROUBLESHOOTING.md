# 故障排除指南

## 常见错误

### Error: GetFileAttributesEx ./events: The system cannot find the file specified.

**问题**：工作流找不到 `events` 目录。

**可能的原因**：
1. common-events 仓库中没有 `events` 目录
2. 目录名称不同（例如：`pkg/events`、`internal/events`）
3. 仓库结构不同

**解决方案**：

#### 方案 1：检查仓库结构

1. 在 common-events 仓库中确认事件结构体的实际位置
2. 如果目录名称不同，可以：
   - 重命名目录为 `events`
   - 或者使用环境变量指定路径

#### 方案 2：使用自定义路径

在 GitHub 仓库设置中添加变量：
- 名称：`EVENTS_DIR_PATH`
- 值：实际的事件目录路径（例如：`pkg/events` 或 `internal/events`）

然后修改工作流文件，使用这个变量。

#### 方案 3：修改工作流文件

在工作流文件中，可以手动指定路径：

```yaml
- name: Generate PlantUML from Go structs
  run: |
    cd common-events
    go run scripts/generate-plantuml.go \
      --input ./your-actual-events-directory \
      --output ../current-repo/output/event-structures.plantuml
```

### 工作流没有触发

**检查清单**：

1. **Webhook 配置**
   - 确认 common-events 仓库中有 `.github/workflows/notify-docs-update.yml`
   - 确认文件中的 `repository` 字段指向正确的文档仓库

2. **Secrets 配置**
   - 在 common-events 仓库中：
     - `DOCS_REPO_TOKEN` 是否存在且有效
     - `DOCS_REPO` 是否正确设置
   - 在文档仓库中：
     - `COMMON_EVENTS_TOKEN` 是否存在且有效
     - `COMMON_EVENTS_REPO` 变量是否正确设置

3. **分支和路径过滤**
   - 确认推送到了 `main`、`master` 或 `develop` 分支
   - 确认修改的文件路径匹配 `events/**/*.go`

4. **查看日志**
   - common-events 仓库的 Actions 标签页
   - 文档仓库的 Actions 标签页

### Token 权限错误

**错误信息**：`Permission denied` 或 `Bad credentials`

**解决方案**：
1. 确认 token 有正确的权限：
   - 对于 `COMMON_EVENTS_TOKEN`：需要读取 common-events 仓库的权限
   - 对于 `DOCS_REPO_TOKEN`：需要触发文档仓库工作流的权限

2. 重新生成 token 并更新 secrets

### 找不到 Go 文件

**问题**：工作流找不到 `.go` 文件

**检查**：
1. 确认事件结构体文件在指定目录中
2. 确认文件扩展名是 `.go`
3. 查看工作流日志中的目录列表

**调试步骤**：
```bash
# 在本地测试
cd common-events
find . -name "*.go" -type f
ls -la events/
```

### PlantUML 文件生成失败

**可能原因**：
1. Go 语法错误
2. 脚本解析失败
3. 输出目录权限问题

**解决方案**：
1. 检查 common-events 仓库中的 Go 文件是否有语法错误
2. 本地运行生成器测试：
   ```bash
   cd event-sync-automation
   bash test-generator.sh
   ```
3. 查看工作流日志中的详细错误信息

## 调试技巧

### 启用调试日志

在工作流中添加调试步骤：

```yaml
- name: Debug - List files
  run: |
    cd common-events
    echo "Repository structure:"
    find . -type f -name "*.go" | head -20
    echo ""
    echo "Current directory contents:"
    ls -la
```

### 本地测试

在本地测试生成器：

```bash
# 克隆 common-events 仓库到本地
git clone <common-events-repo-url> test-common-events

# 运行生成器
cd event-sync-automation
go run scripts/generate-plantuml.go \
  --input ../test-common-events/events \
  --output ../output/test.plantuml
```

### 检查工作流状态

1. 进入 GitHub 仓库的 Actions 标签页
2. 查看最近的工作流运行记录
3. 点击失败的工作流查看详细日志
4. 查看每个步骤的输出

## 获取帮助

如果问题仍然存在：

1. 检查工作流日志的完整输出
2. 确认所有 secrets 和 variables 都正确配置
3. 验证仓库结构和文件路径
4. 尝试手动触发一次工作流（如果启用了 workflow_dispatch）

