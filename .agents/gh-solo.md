# gh-solo configuration for intercomifico

## Layer labels

The default layer labels, mapped onto this repository:

- `backend`: the Intercom API client, `internal/intercom`.
- `frontend`: the terminal UI, `internal/ui`, and the tab contract in `app`.
- `fullstack`: one feature that needs both and cannot be accepted from either side alone.
- `infra`: CI, releases, repository tooling.
- `docs`: prose that is the deliverable, such as an ADR or the README.

## Check commands

- `task lint` runs the linters.
- `task test` runs `go vet` and the tests with the race detector.
- `task audit` checks the dependencies for known vulnerabilities.

`task ci` runs them all in order, after the formatting check, and stops at the first failure. Run it before every push.

## Docs check

`python3 <skill-dir>/scripts/docs-check.py docs .agents --plans docs/plans` needs no ignore set. `AGENTS.md` sits outside that scope; checking it needs `--ignore '~/*'` for its reference-clone paths.
