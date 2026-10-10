> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 10. Go tooling

Date: 2026-10-10

## Status

Accepted

## Context

Besides the compiler, Charm's own repositories rely on three tools: a formatter, a linter and a task runner. A tally of the 50 golangci-lint configurations among the reference repositories shows:

- 15 of Charm's 22 repositories share one configuration, the one in `github.com/charmbracelet/bubbletea`: golangci-lint v2, the standard linters plus 22 more (bodyclose, exhaustive, goconst, godot, gomoddirectives, goprintffuncname, gosec, misspell, nakedret, nestif, nilerr, noctx, nolintlint, prealloc, revive, rowserrcheck, sqlclosecheck, tparallel, unconvert, unparam, whitespace, wrapcheck), with gofumpt and goimports as formatters.
- That shared configuration turns off linting of test files. crush, soft-serve and ultraviolet keep golangci-lint's default, which lints them, as most community apps do.
- Community apps vary from no extra linters to 68, and their most frequent picks overlap Charm's set.
- bubbletea, crush and glow run their commands through a Taskfile.yaml (go-task). mise tasks would do the same job from the mise.toml that pins the toolchain, at the cost of departing from Charm's convention.

## Decision

- **golangci-lint v2 formats and lints**: `golangci-lint fmt` with gofumpt and goimports, and `golangci-lint run` with the linters and settings of bubbletea's .golangci.yml.
- **Test files are linted too**, unlike bubbletea's configuration: ADR 0006 makes tests a large share of the code, so they meet the same bar.
- **A Taskfile.yaml names every command**, following Charm's convention:
  - `task fmt` formats the code.
  - `task lint` runs the linters.
  - `task test` runs `go vet ./...` and `go test -race ./...`.
  - `task test:update` rewrites golden files with `go test ./internal/ui/... -update`, per ADR 0006.
  - `task run` starts the app.
- **A project-local mise.toml pins the toolchain** at exact versions: Go 1.27.1, golangci-lint 2.14.0 and task 3.54.0, the newest releases today.
- **`task lint` and `task test` are the repository's gates**, locally and in CI; the CI issue runs them and adds nothing of its own.

## Consequences

- Anyone working on the repository needs mise, which installs the rest from mise.toml.
- A new codebase meets the strict rules from its first line, so there is no backlog of old warnings to suppress.
- `exhaustive` flags a `switch` over an enum that misses a value, which guards the single `Update` and its focus field from ADR 0008.
- Moving the tasks into mise.toml and dropping go-task can be reconsidered later, in a new ADR.
