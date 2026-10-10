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
- Go's own vulnerability scanner, `govulncheck` from `golang.org/x/vuln`, checks the modules a build uses against the Go vulnerability database at vuln.go.dev, and reports a vulnerability only where the code calls the affected function. It exits 3 when it finds vulnerabilities and 1 on any other error, such as an unreachable database. golangci-lint's `gosec` reads this repository's own code, not its dependencies.
- bubbletea, crush and glow run their commands through a Taskfile.yaml (go-task). mise tasks would do the same job from the mise.toml that pins the toolchain, at the cost of departing from Charm's convention.

## Decision

- **golangci-lint v2 formats and lints**: `golangci-lint fmt` with gofumpt and goimports, and `golangci-lint run` with the linters and settings of bubbletea's .golangci.yml.
- **Test files are linted too**, unlike bubbletea's configuration: ADR 0006 makes tests a large share of the code, so they meet the same bar.
- **A Taskfile.yaml names every command**, following Charm's convention:
  - `task fmt` formats the code.
  - `task lint` runs the linters.
  - `task test` runs `go vet ./...` and `go test -race ./...`.
  - `task audit` runs `govulncheck ./...` under a 60-second `timeout`. Vulnerabilities found (exit 3) always fail it. A network problem (exit 1, or the timeout) fails it in CI, recognised by the `CI` environment variable GitHub Actions sets, and elsewhere prints a warning and passes, so working offline never blocks.
  - `task ci` runs every gate in order and stops at the first that fails: `golangci-lint fmt --diff`, which prints what formatting would change and exits non-zero if anything would, then `task lint`, `task test` and `task audit`, last so that a network problem never hides the results about the change itself.
  - `task test:update` rewrites golden files with `go test ./internal/ui/... -update`, per ADR 0006.
  - `task run` starts the app.
- **A project-local mise.toml pins the toolchain** at exact versions: Go 1.27.1, golangci-lint 2.14.0, task 3.54.0 and govulncheck 1.8.0 (through mise's `go` backend, as the mise registry does not carry it), the newest releases today.
- **`task lint`, `task audit` and `task test` are the repository's gates**, with the formatting check run by `task ci`. `task ci` is the one command for all of them, locally and in CI; the CI issue runs it and adds nothing of its own.

## Consequences

- Anyone working on the repository needs mise, which installs the rest from mise.toml.
- `task audit` needs the network to reach the vulnerability database, so it is a task of its own rather than part of `task lint`, which runs offline.
- A new advisory for a dependency or for Go itself can turn `task audit` red on code that has not changed, until the dependency or Go is upgraded.
- An audit skipped offline is caught by CI, and the CI issue can also run `task audit` on a schedule, so a new advisory shows up on `main` without waiting for a pull request.
- A new codebase meets the strict rules from its first line, so there is no backlog of old warnings to suppress.
- `exhaustive` flags a `switch` over an enum that misses a value, which guards the single `Update` and its focus field from ADR 0008.
- Moving the tasks into mise.toml and dropping go-task can be reconsidered later, in a new ADR.
