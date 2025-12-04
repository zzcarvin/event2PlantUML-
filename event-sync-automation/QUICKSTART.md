# Quick Start Guide

## 🚀 5-Minute Setup

### 1. 复制工作流文件到正确位置

将工作流文件从 `event-sync-automation/.github/workflows/update-event-docs.yml` 复制到仓库根目录：

```bash
# 在仓库根目录执行
mkdir -p .github/workflows
cp event-sync-automation/.github/workflows/update-event-docs.yml .github/workflows/
```

### 2. 配置 GitHub Secrets

在 GitHub 仓库设置中（Settings → Secrets and variables → Actions）：

- **添加 Secret**: `COMMON_EVENTS_TOKEN`
  - 值：一个有读取 common-events 仓库权限的 GitHub Personal Access Token

- **添加 Variable** (可选): `COMMON_EVENTS_REPO`
  - 值：`your-org/common-events` (替换为你的实际仓库名)

### 3. 测试

1. 在 GitHub 上进入 Actions 标签页
2. 选择 "Update Event Documentation" 工作流
3. 点击 "Run workflow" 手动触发
4. 查看执行日志

### 4. 查看生成的文档

工作流成功运行后，会在 `resources/ado-11543/event-structures.plantuml` 生成 PlantUML 文件。

## 📁 文件结构

```
rfcs/
├── .github/workflows/
│   └── update-event-docs.yml          ← 工作流文件（需手动复制到这里）
├── event-sync-automation/
│   ├── scripts/
│   │   └── generate-plantuml.go       ← PlantUML 生成器脚本
│   ├── common-events-example/         ← 示例事件结构（用于测试）
│   ├── README.md                      ← 详细文档
│   ├── INSTALLATION.md                ← 安装指南
│   └── QUICKSTART.md                  ← 本文件
└── resources/ado-11543/
    └── event-structures.plantuml      ← 自动生成的图表（工作流输出）
```

## 🔧 本地测试

在提交到 GitHub 之前，可以本地测试生成器：

```bash
cd event-sync-automation
bash test-generator.sh
```

这会使用示例事件结构生成 PlantUML 文件到 `output/` 目录。

## ❓ 需要帮助？

查看 [INSTALLATION.md](./INSTALLATION.md) 获取详细的安装说明和故障排除指南。

