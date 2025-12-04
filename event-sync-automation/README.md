# Event Structure Synchronization Automation

This directory contains automation tools and workflows to keep event structure documentation synchronized with the common-events package.

## Overview

This automation solution ensures that PlantUML diagrams in the RFC documentation automatically reflect changes to event structures defined in the common-events package.

## Components

### 1. GitHub Action Workflow
- **Location**: `.github/workflows/update-event-docs.yml`
- **Purpose**: Automatically generates PlantUML diagrams from Go struct definitions in the common-events package
- **Triggers**:
  - Manual workflow dispatch
  - Repository dispatch events (triggered by common-events repo)
  - Scheduled daily checks (2 AM UTC)

### 2. PlantUML Generator Script
- **Location**: `scripts/generate-plantuml.go`
- **Purpose**: Parses Go source files and generates PlantUML class diagrams
- **Usage**: 
  ```bash
  go run scripts/generate-plantuml.go \
    --input ./events \
    --output ../resources/ado-11543/event-structures.plantuml
  ```

## Setup Instructions

**⚠️ Important**: Before following these steps, see [INSTALLATION.md](./INSTALLATION.md) for detailed installation instructions, including where to place the workflow file.

### Step 1: Copy Workflow File to Correct Location

The GitHub Action workflow file must be placed in the root of your RFC repository:
- **Source**: `event-sync-automation/.github/workflows/update-event-docs.yml`
- **Destination**: `.github/workflows/update-event-docs.yml` (at repository root)

### Step 2: Configure Secrets and Variables

In your GitHub repository settings, add the following secrets:

1. **COMMON_EVENTS_TOKEN** (Secret)
   - A GitHub Personal Access Token with read access to the common-events repository
   - Required scopes: `repo` (read access)

2. **COMMON_EVENTS_REPO** (Variable, optional)
   - The full repository name (e.g., `your-org/common-events`)
   - Default: `your-org/common-events`

### Step 2: Create Common Events Package Structure

In your common-events repository, create the following structure:

```
common-events/
├── events/
│   ├── device.go          # Device-related event structures
│   ├── client.go          # Client-related event structures
│   └── common.go          # Common event structures
└── go.mod
```

### Step 3: Add Webhook Notification (Optional)

To automatically trigger documentation updates when event structures change, add a workflow in the common-events repository:

**File**: `common-events/.github/workflows/notify-docs-update.yml`

```yaml
name: Notify Documentation Update

on:
  push:
    branches:
      - main
      - master
    paths:
      - 'events/**/*.go'

jobs:
  notify:
    runs-on: ubuntu-latest
    steps:
      - name: Trigger RFC documentation update
        uses: peter-evans/repository-dispatch@v2
        with:
          token: ${{ secrets.RFC_REPO_TOKEN }}
          repository: your-org/rfcs
          event-type: event-structure-updated
          client-payload: |
            {
              "sha": "${{ github.sha }}",
              "ref": "${{ github.ref }}",
              "author": "${{ github.actor }}",
              "message": "${{ github.event.head_commit.message }}"
            }
```

### Step 4: Test the Workflow

1. Make a test change to an event structure in the common-events repository
2. Push the change to trigger the webhook (or manually dispatch the workflow)
3. Check the Actions tab in the RFC repository to see the workflow running
4. A Pull Request should be automatically created with updated PlantUML diagrams

## Manual Usage

You can also run the generator manually:

```bash
# From the event-sync-automation directory
cd scripts
go run generate-plantuml.go \
  --input /path/to/common-events/events \
  --output /path/to/resources/ado-11543/event-structures.plantuml
```

## Output

The generated PlantUML file will be saved to:
- `resources/ado-11543/event-structures.plantuml`

This file contains class diagrams representing all event structures found in the common-events package.

## Features

- **Automatic Parsing**: Extracts struct definitions from Go source files
- **JSON Tag Support**: Includes JSON field tags in the diagram
- **Embedded Field Support**: Handles embedded struct fields and relationships
- **Export Filtering**: Only includes exported structs or structs ending with "Event"
- **Automatic PR Creation**: Creates pull requests for documentation updates

## Troubleshooting

### Workflow fails with "repository not found"
- Check that `COMMON_EVENTS_TOKEN` has the correct permissions
- Verify the repository name in `COMMON_EVENTS_REPO` variable

### No structs found
- Ensure the input directory path is correct
- Check that Go files are in the specified directory
- Verify structs are exported (capitalized) or end with "Event"

### PlantUML file not generated
- Check workflow logs for parsing errors
- Ensure output directory exists or can be created
- Verify file permissions

## Future Enhancements

- Support for interface definitions
- Generate sequence diagrams for event flows
- Support for custom annotations to control diagram generation
- Integration with other documentation formats

