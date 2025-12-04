# Files Overview

This document lists all files in the `event-sync-automation` directory and their purposes.

## 📁 Directory Structure

```
event-sync-automation/
├── .github/
│   └── workflows/
│       ├── update-event-docs.yml              # Main GitHub Action workflow (copy to repo root)
│       └── notify-docs-update.example.yml     # Example webhook workflow for common-events repo
├── scripts/
│   └── generate-plantuml.go                   # Go script to generate PlantUML from Go structs
├── common-events-example/
│   ├── events/
│   │   ├── device.go                          # Example device event structures
│   │   └── client.go                          # Example client event structures
│   └── go.mod                                 # Example go.mod for common-events package
├── .gitignore                                 # Git ignore rules for this directory
├── README.md                                  # Main documentation
├── INSTALLATION.md                            # Detailed installation guide
├── QUICKSTART.md                              # Quick start guide
├── FILES.md                                   # This file
└── test-generator.sh                          # Test script for local testing

```

## 📄 File Descriptions

### Workflow Files

#### `.github/workflows/update-event-docs.yml`
- **Purpose**: Main GitHub Action workflow that generates PlantUML diagrams
- **Location after setup**: Copy to `.github/workflows/update-event-docs.yml` at repository root
- **Triggers**: 
  - Manual workflow dispatch
  - Repository dispatch events
  - Daily scheduled runs (2 AM UTC)

#### `.github/workflows/notify-docs-update.example.yml`
- **Purpose**: Example workflow for the common-events repository to trigger RFC updates
- **Location**: Copy to `common-events/.github/workflows/notify-docs-update.yml`
- **Triggers**: On push to events/**/*.go files

### Scripts

#### `scripts/generate-plantuml.go`
- **Purpose**: Parses Go source files and generates PlantUML class diagrams
- **Language**: Go
- **Usage**: 
  ```bash
  go run scripts/generate-plantuml.go --input ./events --output ../output.plantuml
  ```

#### `test-generator.sh`
- **Purpose**: Test script for local testing of the PlantUML generator
- **Usage**: `bash test-generator.sh`
- **Output**: Generates test PlantUML file in `output/` directory

### Example Files

#### `common-events-example/`
- **Purpose**: Example structure and event definitions for testing
- **Contains**: Sample Go files with event structures that can be used for local testing
- **Files**:
  - `events/device.go`: Example device-related events
  - `events/client.go`: Example client-related events
  - `go.mod`: Example Go module file

### Documentation

#### `README.md`
- **Purpose**: Comprehensive documentation including overview, components, setup instructions, and troubleshooting

#### `INSTALLATION.md`
- **Purpose**: Detailed step-by-step installation guide with file locations and configuration

#### `QUICKSTART.md`
- **Purpose**: Quick start guide for 5-minute setup

#### `FILES.md`
- **Purpose**: This file - overview of all files in the directory

### Configuration

#### `.gitignore`
- **Purpose**: Git ignore rules for build artifacts, output files, and IDE files

## 🚀 Next Steps

1. **Read**: Start with [QUICKSTART.md](./QUICKSTART.md) for quick setup
2. **Install**: Follow [INSTALLATION.md](./INSTALLATION.md) for detailed steps
3. **Test**: Use `test-generator.sh` to test locally
4. **Deploy**: Copy workflow files to their target locations and configure secrets

## 📝 Notes

- The `.github` directory in this folder is for reference only
- The actual workflow file must be copied to the repository root `.github/workflows/`
- The example files can be used for local testing before connecting to the actual common-events repository

