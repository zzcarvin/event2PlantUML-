package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	inputDir := flag.String("input", "./events", "Input directory containing Go files")
	outputFile := flag.String("output", "./event-structures.plantuml", "Output PlantUML file path")
	flag.Parse()

	fset := token.NewFileSet()
	var allStructs []StructInfo

	// Walk through all Go files in the input directory
	err := filepath.Walk(*inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		src, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		file, err := parser.ParseFile(fset, path, src, parser.ParseComments)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to parse %s: %v\n", path, err)
			return nil
		}

		// Extract structs from the AST
		ast.Inspect(file, func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.GenDecl:
				if x.Tok == token.TYPE {
					for _, spec := range x.Specs {
						if ts, ok := spec.(*ast.TypeSpec); ok {
							if st, ok := ts.Type.(*ast.StructType); ok {
								structInfo := parseStruct(ts.Name.Name, st, fset, src)
								// Only include exported structs or structs ending with Event
								if ast.IsExported(ts.Name.Name) || strings.HasSuffix(ts.Name.Name, "Event") {
									allStructs = append(allStructs, structInfo)
								}
							}
						}
					}
				}
			}
			return true
		})

		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if len(allStructs) == 0 {
		fmt.Fprintf(os.Stderr, "Warning: No structs found in %s\n", *inputDir)
		return
	}

	// Generate PlantUML
	plantuml := generatePlantUML(allStructs)

	// Write to output file
	outputDir := filepath.Dir(*outputFile)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(*outputFile, []byte(plantuml), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated PlantUML file: %s\n", *outputFile)
	fmt.Printf("Found %d event structures\n", len(allStructs))
}

type StructInfo struct {
	Name    string
	Fields  []FieldInfo
	Comment string
}

type FieldInfo struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

func parseStruct(name string, st *ast.StructType, fset *token.FileSet, src []byte) StructInfo {
	info := StructInfo{
		Name:   name,
		Fields: []FieldInfo{},
	}

	if st.Fields != nil {
		for _, field := range st.Fields.List {
			// 使用 FileSet 的偏移量来安全地从源代码中截取类型字符串，避免 slice 越界
			start := fset.Position(field.Type.Pos()).Offset
			end := fset.Position(field.Type.End()).Offset

			if start < 0 {
				start = 0
			}
			if end > len(src) {
				end = len(src)
			}
			if start > end {
				start, end = end, end
			}

			fieldType := ""
			if start < end {
				fieldType = string(src[start:end])
			}

			for _, fieldName := range field.Names {
				fieldInfo := FieldInfo{
					Name: fieldName.Name,
					Type: fieldType,
				}

				if field.Tag != nil {
					fieldInfo.Tag = field.Tag.Value
				}

				if field.Comment != nil {
					fieldInfo.Comment = strings.TrimSpace(field.Comment.Text())
				}

				info.Fields = append(info.Fields, fieldInfo)
			}

			// Handle embedded/anonymous fields
			if len(field.Names) == 0 {
				fieldInfo := FieldInfo{
					Name: "",
					Type: fieldType,
				}
				if field.Comment != nil {
					fieldInfo.Comment = strings.TrimSpace(field.Comment.Text())
				}
				info.Fields = append(info.Fields, fieldInfo)
			}
		}
	}

	return info
}

func generatePlantUML(structs []StructInfo) string {
	var sb strings.Builder

	sb.WriteString("@startuml\n")
	sb.WriteString("!theme plain\n")
	sb.WriteString("skinparam backgroundColor #FFFFFF\n")
	sb.WriteString("skinparam classAttributeIconSize 0\n\n")
	sb.WriteString("title Event Structures from Common Events Package\n\n")

	// Generate class diagrams for each struct
	for _, s := range structs {
		sb.WriteString(fmt.Sprintf("class %s {\n", s.Name))

		for _, field := range s.Fields {
			fieldType := cleanType(field.Type)
			fieldName := field.Name
			visibility := "+"

			if fieldName == "" {
				// Embedded field
				fieldName = fieldType
				visibility = "^"
			}

			// Extract json tag if available
			jsonTag := extractJSONTag(field.Tag)
			fieldDisplay := fieldName
			if jsonTag != "" && jsonTag != fieldName {
				fieldDisplay = fmt.Sprintf("%s (%s)", fieldName, jsonTag)
			}

			sb.WriteString(fmt.Sprintf("  %s %s : %s\n", visibility, fieldDisplay, fieldType))
		}

		sb.WriteString("}\n\n")
	}

	// Add relationships based on embedded fields
	relationships := make(map[string]map[string]bool)
	for _, s := range structs {
		for _, field := range s.Fields {
			if field.Name == "" {
				// Check if this is an embedded field that matches another struct
				embeddedType := cleanType(field.Type)
				for _, otherStruct := range structs {
					if otherStruct.Name == embeddedType {
						if relationships[s.Name] == nil {
							relationships[s.Name] = make(map[string]bool)
						}
						relationships[s.Name][otherStruct.Name] = true
						sb.WriteString(fmt.Sprintf("%s <|-- %s\n", otherStruct.Name, s.Name))
					}
				}
			}
		}
	}

	sb.WriteString("\n@enduml\n")
	return sb.String()
}

func cleanType(t string) string {
	// Remove common prefixes and clean up the type string
	t = strings.TrimSpace(t)
	t = strings.TrimPrefix(t, "*")
	t = strings.TrimPrefix(t, "[]")
	return strings.TrimSpace(t)
}

func extractJSONTag(tag string) string {
	if tag == "" {
		return ""
	}
	// Simple extraction of json tag
	tag = strings.Trim(tag, "`")
	parts := strings.Split(tag, " ")
	for _, part := range parts {
		if strings.HasPrefix(part, "json:") {
			jsonVal := strings.TrimPrefix(part, "json:")
			jsonVal = strings.Trim(jsonVal, "\"")
			if idx := strings.Index(jsonVal, ","); idx != -1 {
				jsonVal = jsonVal[:idx]
			}
			if jsonVal == "-" {
				return ""
			}
			return jsonVal
		}
	}
	return ""
}
