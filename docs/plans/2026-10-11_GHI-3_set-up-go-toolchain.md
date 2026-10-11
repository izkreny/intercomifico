> 🤖 Written by AI --- read/modified by izkreny! 🤓

# Plan: set up the Go toolchain

Issue: #3. Its acceptance criteria are the scope; ADR 0010 holds the tools, versions and tasks, ADR 0008 and ADR 0012 the package layout, and ADR 0003 the Go version in `go.mod` (new).

## Approach

The toolchain is pinned in `mise.toml` (new), the commands live in `Taskfile.yaml` (new), and the lint configuration in `.golangci.yml` (new) is bubbletea's with its `run.tests: false` line removed, so test files are linted. The packages are stubs that build and nothing more: each carries only its package comment, and `Run` in the root package takes no arguments and returns nil. #9 gives `Run` its `extra ...app.Tab` signature and puts the `Tab` contract into `app`; defining either here would be code with no caller.

`task audit` is the one task with logic of its own. It runs `govulncheck ./...` under `timeout 60` and branches on the exit status: 0 passes, 3 fails, and anything else (1, or 124 from the timeout) fails when `CI` is set and otherwise prints a warning naming an unknown error and passes.

Dependabot alerts, also in ADR 0010, are a repository setting rather than a file and are not among #3's criteria, so this branch leaves them alone.

## Steps

- Add `mise.toml` (new) pinning `go` 1.27.1, `golangci-lint` 2.14.0, `task` 3.54.0 and `go:golang.org/x/vuln/cmd/govulncheck` 1.8.0, then run `mise install` against it.
- Add `go.mod` (new) for the module `github.com/izkreny/intercomifico` declaring `go 1.27`, per ADR 0003.
- Add the root package in `intercomifico.go` (new) with `Run() error`, and the `main` package in `cmd/intercomifico/main.go` (new) calling it and exiting non-zero on an error.
- Add `app/doc.go` (new), `internal/intercom/doc.go` (new) and `internal/ui/doc.go` (new), each holding the package comment that says what ADR 0008 and ADR 0012 put there.
- Add `.golangci.yml` (new), copied from bubbletea's with the test-file exclusion dropped.
- Add `Taskfile.yaml` (new) with `fmt`, `lint`, `test`, `test:update`, `audit`, `ci` and `run` as ADR 0010 states them, `ci` calling the formatting check and the other gates in order.
- See each gate fail once and recover: an unformatted file under `task ci`, a lint error under `task lint`, and a failing test under `task test`, each reverted afterwards and recorded in the handoff.
- See `task audit` take each branch by putting a stand-in `govulncheck` first on `PATH` that exits 3, exits 1 and sleeps past the timeout, with and without `CI` set.
- Record `task lint`, `task audit` and `task test` as the check commands in `.agents/gh-solo.md`, with `task ci` running them all.

## Verification

- `python3 <skill-dir>/scripts/plan-check.py docs/plans/2026-10-11_GHI-3_set-up-go-toolchain.md`
- `python3 <skill-dir>/scripts/docs-check.py docs .agents --plans docs/plans`
- `mise install`
- `task ci`
- `task test:update`
- `task run`

`task ci` going green proves the gates pass on this code, not that they would catch anything; the failures seen in Steps are that evidence. No gate checks that the lint set matches bubbletea's, which is a reading of the diff. `task test:update` has no golden files to rewrite yet, so its green only shows the command is well formed.

## Open questions

None.

## Settled

None yet.
