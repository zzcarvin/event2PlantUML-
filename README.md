# Event Sync Automation - Test Repository

这个仓库包含了事件结构同步自动化工具，用于从 common-events 包自动生成 PlantUML 文档。

## 📁 文件结构

```
test1111111/
├── .github/
│   └── workflows/
│       └── update-event-docs.yml          # GitHub Action 工作流（已配置）
├── event-sync-automation/
│   ├── scripts/
│   │   └── generate-plantuml.go          # PlantUML 生成器脚本
│   ├── common-events-example/            # 示例事件结构
│   └── ...                               # 其他文档文件
└── output/                               # 生成的 PlantUML 文件将在这里
```

## 🚀 快速开始

### 1. 配置 GitHub Secrets

在当前仓库（文档仓库）的设置中（Settings → Secrets and variables → Actions）：

**必需：**
- `COMMON_EVENTS_TOKEN` - 有读取 common-events 仓库权限的 GitHub Token

**可选：**
- `COMMON_EVENTS_REPO` - common-events 仓库名称（例如：`your-org/common-events`）

### 2. 设置自动触发（重要！）

要让结构体改动时自动触发文档更新，需要在 **common-events 仓库** 中配置 webhook。

**详细步骤请查看：** `event-sync-automation/SETUP-WEBHOOK.md`

简单来说：
1. 在 common-events 仓库创建 `.github/workflows/notify-docs-update.yml`
2. 复制 `event-sync-automation/common-events-webhook.yml` 的内容
3. 在 common-events 仓库配置 `DOCS_REPO_TOKEN` 和 `DOCS_REPO` secrets

### 3. 测试

1. 在 common-events 仓库中修改 `events/` 目录下的任何 `.go` 文件
2. 提交并推送
3. 文档仓库的 Actions 应该会自动触发并更新文档

### 3. 查看输出

生成的 PlantUML 文件将保存在 `output/event-structures.plantuml`

## 📝 本地测试

你也可以在本地测试生成器：

```bash
cd event-sync-automation
bash test-generator.sh
```

## 📚 详细文档

查看 `event-sync-automation/README.md` 获取更多信息。

