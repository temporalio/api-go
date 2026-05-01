package main

import (
	"path/filepath"
	"slices"
	"strings"

	"go.temporal.io/api/internal/strcase"
)

const experimentalNumberStart = 1000

type descriptorChanges struct {
	Methods  []methodInfo
	Overlays []messageOverlay
	Enums    []enumInfo
}

type messageOverlay struct {
	StableMessage  string
	OverlayMessage string
	SourceMessage  string
	ProtoPackage   string
	ProtoFile      string
	StableImport   string
	GoPackage      string
	RelDir         string
	FieldName      string
	SourceName     string
	SourceNumber   int
	FieldNumber    int
}

type messageOverlayGroup struct {
	StableMessage  string
	OverlayMessage string
	VariableName   string
	ProtoPackage   string
	ProtoFile      string
	StableImport   string
	GoPackage      string
	RelDir         string
	Fields         []messageOverlay
}

type enumInfo struct {
	StableEnum   string
	SourceEnum   string
	ValueName    string
	SourceNumber int
	ValueNumber  int
}

func collectChanges(base descriptorSnapshot, source descriptorSnapshot) descriptorChanges {
	changes := descriptorChanges{
		Methods:  make([]methodInfo, 0),
		Overlays: make([]messageOverlay, 0),
		Enums:    make([]enumInfo, 0),
	}

	for name, method := range source.WorkflowServiceMethods {
		if _, exists := base.WorkflowServiceMethods[name]; exists {
			continue
		}
		changes.Methods = append(changes.Methods, method)
	}

	for messageName, sourceMessage := range source.WorkflowMessages {
		baseMessage, ok := base.WorkflowMessages[messageName]
		if !ok {
			continue
		}
		usedNumbers := make(map[int]struct{}, len(baseMessage.Fields))
		baseFields := make(map[string]protoFieldInfo, len(baseMessage.Fields))
		for _, field := range baseMessage.Fields {
			usedNumbers[field.Number] = struct{}{}
			baseFields[field.Name] = field
		}
		addedFields := make([]protoFieldInfo, 0, len(sourceMessage.Fields))
		for _, field := range sourceMessage.Fields {
			if _, exists := baseFields[field.Name]; exists {
				continue
			}
			addedFields = append(addedFields, field)
		}
		slices.SortFunc(addedFields, func(a protoFieldInfo, b protoFieldInfo) int {
			return a.Number - b.Number
		})
		nextNumber := nextAvailableNumber(usedNumbers)
		for _, field := range addedFields {
			sourceFile := source.WorkflowMessageFiles[messageName]
			importPath, packageName := splitGoPackage(sourceFile.GetOptions().GetGoPackage())
			changes.Overlays = append(changes.Overlays, messageOverlay{
				StableMessage:  messageName,
				OverlayMessage: messageName + "Overlay",
				SourceMessage:  messageName,
				ProtoPackage:   sourceFile.GetPackage(),
				ProtoFile:      sourceFile.GetName(),
				StableImport:   importPath,
				GoPackage:      packageName,
				RelDir:         relDirFromProtoFile(sourceFile.GetName()),
				FieldName:      strcase.ToCamel(field.Name),
				SourceName:     field.Name,
				SourceNumber:   field.Number,
				FieldNumber:    nextNumber,
			})
			usedNumbers[nextNumber] = struct{}{}
			nextNumber = nextAvailableNumber(usedNumbers)
		}
	}

	for enumName, sourceValues := range source.WorkflowEnums {
		baseValues, ok := base.WorkflowEnums[enumName]
		if !ok {
			continue
		}
		usedNumbers := make(map[int]struct{}, len(baseValues))
		for _, number := range baseValues {
			usedNumbers[int(number)] = struct{}{}
		}
		names := make([]string, 0, len(sourceValues))
		for name := range sourceValues {
			if _, exists := baseValues[name]; exists {
				continue
			}
			names = append(names, name)
		}
		slices.Sort(names)
		nextNumber := nextAvailableNumber(usedNumbers)
		for _, name := range names {
			changes.Enums = append(changes.Enums, enumInfo{
				StableEnum:   enumName,
				SourceEnum:   enumName,
				ValueName:    name,
				SourceNumber: int(sourceValues[name]),
				ValueNumber:  nextNumber,
			})
			usedNumbers[nextNumber] = struct{}{}
			nextNumber = nextAvailableNumber(usedNumbers)
		}
	}

	slices.SortFunc(changes.Methods, func(a methodInfo, b methodInfo) int {
		return strings.Compare(a.Method, b.Method)
	})
	return changes
}

func nextAvailableNumber(usedNumbers map[int]struct{}) int {
	for next := experimentalNumberStart; ; next++ {
		if _, exists := usedNumbers[next]; !exists {
			return next
		}
	}
}

func groupMessageOverlays(overlays []messageOverlay) []messageOverlayGroup {
	groups := make([]messageOverlayGroup, 0)
	groupByMessage := make(map[string]int)
	for _, overlay := range overlays {
		idx, ok := groupByMessage[overlay.StableMessage]
		if !ok {
			idx = len(groups)
			groupByMessage[overlay.StableMessage] = idx
			groups = append(groups, messageOverlayGroup{
				StableMessage:  overlay.StableMessage,
				OverlayMessage: overlay.OverlayMessage,
				VariableName:   lowerFirst(overlay.OverlayMessage),
				ProtoPackage:   overlay.ProtoPackage,
				ProtoFile:      overlay.ProtoFile,
				StableImport:   overlay.StableImport,
				GoPackage:      overlay.GoPackage,
				RelDir:         overlay.RelDir,
			})
		}
		groups[idx].Fields = append(groups[idx].Fields, overlay)
	}
	return groups
}

func lowerFirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func splitGoPackage(goPackage string) (importPath string, packageName string) {
	importPath, packageName, _ = strings.Cut(goPackage, ";")
	if packageName == "" {
		packageName = filepath.Base(importPath)
	}
	return importPath, packageName
}

func relDirFromProtoFile(name string) string {
	name = strings.TrimPrefix(filepath.ToSlash(name), "temporal/api/")
	return filepath.Dir(name)
}
