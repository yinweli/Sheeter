# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common guildlines
@CLAUDE-common guildlines.md

## Project Overview

Sheeter is a Go CLI tool that converts Excel (.xlsx) files into JSON data files and generates C#/Go reader code. Module: `github.com/yinweli/Sheeter/v3`, Go 1.25.0.

## Common Commands

```bash
# Build & install
go install ./cmd/sheeter

# Run all tests with coverage
go test -coverprofile=coverage.txt -covermode=atomic ./...

# Run a single test
go test -run TestFunctionName ./sheeter/...

# Lint & format (requires: task install)
task lint

# Install dev tools (golangci-lint, csharpier, markdownlint, prettier, goimports)
task install
```

## Architecture

```text
cmd/sheeter/         CLI entry point (Cobra framework)
sheeter/
├── define.go        Constants, tokens, paths
├── builds/          Build pipeline stages: config → initialize → operation → poststep
├── excels/          Excel/sheet reading (xlsxreader)
├── fields/          Field type system — interface + 12 types (bool/int/long/float/double/string, each with array variant)
├── layouts/         JSON layout construction and packing
├── nameds/          Named entity management, field naming, merge/combine logic
├── pipelines/       Pipeline orchestration
├── tmpls/           Code generation templates (C# and Go)
└── utils/           Helpers: file I/O, JSON, string, terminal colors
```

**Key patterns:**

- Pipeline architecture: modular build stages chained sequentially
- Interface-based field type registry (`fields/field.go`) — global slice, not map
- Explicit `Close()` for resource cleanup (Excel/Sheet)
- Config struct with Cobra flag binding for CLI options

**Test data** lives in `testdata/env/` organized by feature (build, config, excel, etc.). Tests use `testify` for assertions.

## Code Style Rules (Strict)

These are **mandatory** — see `CLAUDE-common guildlines.md` for full details.

1. **Boolean checks:** Use `if x == false`, never `if !x`. True checks use `if x`.
2. **Block ending comments:** Add `// if`, `// for`, `// switch` after closing braces of control flow. **Never** on functions/methods.
3. **Variable naming:** Always singular. `item := []Item{}` not `items := []Item{}`.
4. **Iterator naming:** Use `itor` for default iterators, `k, v` for maps.
5. **Language:** Documentation/comments in Traditional Chinese, code naming in English.

## Commit Conventions

Format: `<Type> | <Description in Traditional Chinese>`

Types: `Feature`, `Fix`, `Sheet`, `Message`, `UI`

Branch format: `<account>/<feature-name>`

Feature PRs target `dev` branch.

## CI/CD

GitHub Actions runs on push/PR to main:

- **test.yml** — Go tests + coverage (Codecov)
- **lint.yml** — golangci-lint v2
- **build.yml** — Cross-platform release on version tags (`v*`)

## Workflow Preference

Prefer Python scripts over repeated shell commands for batch operations (file renaming, data transformation, code generation, codebase scanning).
