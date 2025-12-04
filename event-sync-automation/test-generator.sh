#!/bin/bash
# Test script for the PlantUML generator
# This script can be run locally to test the generator before using it in GitHub Actions

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
EXAMPLE_DIR="$SCRIPT_DIR/common-events-example"
OUTPUT_DIR="$SCRIPT_DIR/output"
OUTPUT_FILE="$OUTPUT_DIR/event-structures.plantuml"

echo "Testing PlantUML Generator"
echo "=========================="
echo ""

# Create output directory
mkdir -p "$OUTPUT_DIR"

# Run the generator
echo "Running generator..."
echo "Input: $EXAMPLE_DIR/events"
echo "Output: $OUTPUT_FILE"
echo ""

cd "$SCRIPT_DIR"
go run scripts/generate-plantuml.go \
  --input "$EXAMPLE_DIR/events" \
  --output "$OUTPUT_FILE"

if [ $? -eq 0 ]; then
    echo ""
    echo "✓ Successfully generated PlantUML file"
    echo ""
    echo "Preview of generated file:"
    echo "---------------------------"
    head -20 "$OUTPUT_FILE"
    echo ""
    echo "Full file saved to: $OUTPUT_FILE"
    echo ""
    echo "To view the diagram, you can:"
    echo "1. Copy the content to http://www.plantuml.com/plantuml/uml/"
    echo "2. Or use a PlantUML viewer in your IDE"
else
    echo "✗ Generation failed"
    exit 1
fi

