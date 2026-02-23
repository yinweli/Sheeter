# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What This Project Does

Sheeter is a CLI tool written in Go that reads Excel (`.xlsx`) files and generates:
- **JSON data files** from the sheet's data rows
- **C# reader code** (`codeCs/`) for loading those JSON files at runtime
- **Go reader code** (`codeGo/`) for the same purpose

## Commands

### Development Tools Setup
```bash
task install   # installs golangci-lint, goimports, csharpier, prettier
```

### Build & Run
```bash
go build -o sheeter cmd/sheeter/*.go
sheeter build --config <config.yaml>
```

### Testing
```bash
go test ./... -cover          # all packages
go test ./... -bench=. -benchmem  # benchmarks

# Run a single package's tests:
go test ./sheeter/builds/... -cover
go test ./sheeter/fields/... -run TestFieldName
```

### Linting & Formatting
```bash
task lint   # csharpier, gofmt, goimports, golangci-lint, markdownlint, prettier (all-in-one)
```

## Architecture

### Build Pipeline (3 phases)

The `sheeter build` command runs three sequential phases, each implemented in `sheeter/builds/`:

1. **`Initialize`** — discovers Excel files/sheets from `source` paths, validates names, produces `[]InitializeData`
2. **`Operation`** — for each sheet: extracts fields, converts rows to JSON, generates C#/Go code
3. **`Poststep`** — post-processing (e.g., copying shared helper files to output directories)

### Concurrent Execution via `pipelines`

`sheeter/pipelines/pipelines.go` provides a generic `Pipeline[T]` function that fans out work across goroutines with a progress bar. Each phase passes its slice of items and a list of `Execute[T]` step functions. If any step in a pipeline item fails, subsequent steps for that item are skipped.

### Core Processing Packages

| Package | Role |
|---|---|
| `sheeter/builds/` | Orchestrates the three build phases; owns `Config` |
| `sheeter/excels/` | Opens/caches XLSX files via `xlsxreader`; `CloseAll()` called at end |
| `sheeter/fields/` | 12 field types (bool, int, long, float, double, string + array variants); each implements parsing and C#/Go type name methods |
| `sheeter/layouts/` | Serializes processed data rows to JSON (`jsonPack.go`) |
| `sheeter/tmpls/` | `text/template`-based C# (`tmplCs.go`) and Go (`tmplGo.go`) code generation |
| `sheeter/nameds/` | Name normalization (CamelCase), merge-term parsing |
| `sheeter/utils/` | Shared utilities: unique slice, SheetTerm/MergeTerm parsing, colored stdout |
| `sheeter/pipelines/` | Generic concurrent pipeline with progress bar |

### Configuration (`Config` struct)

Config is loaded from a YAML file (`--config`) and/or CLI flags. Key fields:

```yaml
source: [path/to/excel.xlsx, path/to/folder/]  # files or directories
output: path/to/output/                         # output root
tag: server                                     # column tag filter ("ignore" = skip column)
lineOfTag: 1                                    # row number (1-based) for tags
lineOfName: 2                                   # row number for field names
lineOfNote: 3                                   # row number for comments
lineOfField: 4                                  # row number for field types
lineOfData: 5                                   # first data row
merge: []                                       # merge sheets: "name$excel#sheet&..."
exclude: []                                     # exclude sheets: "excel#sheet"
```

### Output Directory Layout

Under the configured `output` path:
- `json/` — one `.json` per sheet
- `codeCs/` — C# reader classes
- `codeGo/` — Go reader packages

### Key Constants (`sheeter/define.go`)

- `TokenIgnore = "ignore"` — tag/output value that skips a column
- `TokenArray = ","` — array element separator within a cell
- `TokenName = "$"`, `TokenTerm = "&"`, `TokenExcel = "#"` — merge/exclude term syntax
- `IndexOutput = 0`, `IndexPrimary = 1` — fixed column positions for output flag and primary key

## Module Path

`github.com/yinweli/Sheeter/v3` — use this prefix for all internal imports.

## Code Style Requirements

### Boolean Checks - **Strict Rule**

**MUST** use explicit `false` checks, while `true` checks use standard form.

```go
if x == false { ... }  // Required: explicit false check
if x { ... }           // Standard: check for true
// Forbidden: if !x
```

### Block Ending Comments - **Strict Rule**

Reserve ending comments for control flow only (`if`, `for`, `switch`). **NEVER** add ending comments for functions/methods.

```go
if condition {
    // logic
} // if

for i := range items {
    // logic
} // for

switch x {
case 1:
    // logic
} // switch

// NEVER add ending comments for functions/methods
```

### Iterator Naming

```go
for itor := range items {
    // logic
} // for

for k, v := range someMap {
    // logic
} // for
```

- `itor`: Default iterator variable name
- `k, v`: Use when iterating over Maps

### Language Conventions

- **Documentation & Comments**: Traditional Chinese
- **Code Naming (variables, functions)**: English
- **Implementation Plans**: Chinese description + English technical terms

## Interaction Guidelines

Communicate with users in Traditional Chinese
