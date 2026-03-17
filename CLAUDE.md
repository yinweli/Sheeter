# CLAUDE.md

This file provides guidance for AI CLIs (such as Claude Code or ChatGPT Codex) when working in this repository.

## Language Conventions

- Documentation and comments: Traditional Chinese
- Code naming (variables, functions, classes): English
- Implementation plans: Chinese descriptions with English technical terms

## Workflow

- Use Python scripts instead of long shell command chains for multi-step, repetitive, or cross-file tasks.
- Use Python first for:
  - batch file operations
  - structured data parsing or conversion
  - code generation
  - any task involving loops, conditions, or cross-file text processing

- Generate mock data in batches according to project conventions instead of writing test data manually one by one.
- After running tests or lint checks, summarize only failures and warnings from the raw output.

## Context Management

- Manage context usage proactively during long conversations or when working with large files.
- In long conversations, periodically summarize key decisions, current status, and remaining tasks.
- When working with files longer than 500 lines, handle them in focused sections instead of loading everything at once.
- Before complex multi-step tasks, briefly restate the relevant code style rules and project constraints.
- If context becomes overloaded, warn the user and suggest starting a new conversation or consolidating the current state into a file.

## Code Style

- In Go and C#, `false` checks must be explicit. Negation-style checks are forbidden.
  - Example: `if x == false {}` / `if (x == false) {}`
- In TypeScript, negation-style checks such as `if (!x) {}` are allowed.

- Closing comments are only allowed for control flow blocks: `if`, `for`, `switch`
  - Example: `} // if`
- Closing comments on functions or methods are forbidden.

- Variable names must always be singular, even for slices, arrays, maps, and other collections.
  - Example: `item := []Item{}`, `hero := []Hero{}`
  - Forbidden: `items := []Item{}`, `heroes := []Hero{}`

- Iterator naming:
  - Use `itor` for general iteration
  - Use `k, v` for map iteration

## Project Overview

Sheeter is a Go CLI tool that converts Excel (.xlsx) files into JSON data files and generates reader code for C# / Go. Module: `github.com/yinweli/Sheeter/v3`.

## Development / Build / Common Commands
All commands use [Task](https://taskfile.dev/) (install per official docs).
```bash
task lint                # Format and lint (golangci-lint fmt + run, markdownlint, prettier)
task install             # Install all dev tools (golangci-lint, buf, csharpier, etc.)
```

### Running Tests Directly
```bash
# Build and install
go install ./cmd/sheeter
# All tests (excluding support/ directory)
go test $(go list ./... | grep -v "support")
# Single test
go test ./sheeter/... -run TestEntityName
# With coverage
go test -coverprofile=coverage.txt -covermode=atomic $(go list ./... | grep -v "support")
```

## Architecture
```text
cmd/sheeter/         CLI entry point (Cobra framework)
sheeter/
├── define.go        Constants, tokens, paths
├── builds/          Build pipeline stages: config → initialize → operation → poststep
├── excels/          Excel / sheet reading (xlsxreader)
├── fields/          Field type system — interface + 12 types (bool/int/long/float/double/string, each with array variant)
├── layouts/         JSON layout construction and encapsulation
├── nameds/          Named entity management, field naming, merge / combine logic
├── pipelines/       Pipeline orchestration
├── tmpls/           Code generation templates (C# and Go)
└── utils/           Utilities: file I/O, JSON, strings, terminal colors
```

**Key Patterns:**
* Pipeline architecture: modular build stages chained in sequence
* Interface-based field type registration (`fields/field.go`) — uses a global slice, not a map
* Explicit `Close()` for resource cleanup (Excel / Sheet)
* Config struct with Cobra flag binding for CLI options

**Test data** is located in `testdata/env/`, organized by feature (build, config, excel, etc.). Tests use `testify` for assertions.

## Go Version
Requires Go 1.25.0+.
