# Installation Guide

## Quick Start

### Step 1: Copy GitHub Action Workflow

The GitHub Action workflow file needs to be placed in the root of your RFC repository (not in the `event-sync-automation` folder).

**Copy this file:**
- From: `event-sync-automation/.github/workflows/update-event-docs.yml`
- To: `.github/workflows/update-event-docs.yml` (in the root of your RFC repository)

If the `.github/workflows/` directory doesn't exist, create it first.

### Step 2: Configure GitHub Secrets

In your GitHub repository settings (Settings → Secrets and variables → Actions), add:

1. **COMMON_EVENTS_TOKEN** (Repository Secret)
   - Go to: Settings → Secrets and variables → Actions → New repository secret
   - Name: `COMMON_EVENTS_TOKEN`
   - Value: A GitHub Personal Access Token with read access to the common-events repository
   - Required scopes: `repo` (read access)

2. **COMMON_EVENTS_REPO** (Repository Variable, optional)
   - Go to: Settings → Secrets and variables → Actions → Variables → New repository variable
   - Name: `COMMON_EVENTS_REPO`
   - Value: The full repository name (e.g., `your-org/common-events`)
   - If not set, defaults to `your-org/common-events`

### Step 3: Test the Workflow

1. Go to your repository on GitHub
2. Navigate to Actions tab
3. Select "Update Event Documentation" workflow
4. Click "Run workflow" → "Run workflow"
5. Monitor the workflow execution

### Step 4: Set Up Webhook in Common Events Repository (Optional)

To automatically trigger documentation updates when event structures change:

1. In your common-events repository, create `.github/workflows/notify-docs-update.yml`
2. Copy the content from `event-sync-automation/.github/workflows/notify-docs-update.example.yml`
3. Update the `repository` field to point to your RFC repository
4. Add `RFC_REPO_TOKEN` secret in the common-events repository settings

## File Structure After Installation

```
rfcs/
├── .github/
│   └── workflows/
│       └── update-event-docs.yml          ← Copy workflow here
├── event-sync-automation/                 ← Keep as reference
│   ├── scripts/
│   │   └── generate-plantuml.go
│   ├── common-events-example/
│   ├── README.md
│   └── ...
└── resources/
    └── ado-11543/
        └── event-structures.plantuml      ← Generated here
```

## Troubleshooting

### Workflow doesn't appear in Actions tab
- Make sure the workflow file is in `.github/workflows/` at the root
- Check that the file has `.yml` or `.yaml` extension
- Verify the YAML syntax is correct

### "Repository not found" error
- Verify `COMMON_EVENTS_TOKEN` has correct permissions
- Check that `COMMON_EVENTS_REPO` variable is set correctly (if using)
- Ensure the token has access to the common-events repository

### "No structs found" warning
- Check that the common-events repository has Go files in the `events/` directory
- Verify the structs are exported (capitalized) or end with "Event"
- Check workflow logs for detailed error messages

## Testing Locally

Before pushing to GitHub, you can test the generator locally:

```bash
cd event-sync-automation
chmod +x test-generator.sh  # On Unix/Mac
./test-generator.sh         # On Unix/Mac
# or on Windows:
bash test-generator.sh
```

This will use the example event structures to generate a PlantUML file.

