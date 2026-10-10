# gh-solo configuration for intercomifico

## Layer labels

The default layer labels, mapped onto this repository:

- `backend`: the Intercom API client, `internal/intercom`.
- `frontend`: the terminal UI, `internal/ui`, and the tab contract in `app`.
- `fullstack`: one feature that needs both and cannot be accepted from either side alone.
- `infra`: CI, releases, repository tooling.
- `docs`: prose that is the deliverable, such as an ADR or the README.

## Check commands

Per ADRs 0006 and 0010, run from the repository root with the toolchain mise.toml pins:

- `task lint`
- `task test`

`task test:update` rewrites golden files after a deliberate rendering change. It is not a gate; the changed `.golden` files are reviewed in the diff.

## Docs check

`python3 <skill-dir>/scripts/docs-check.py docs .agents --plans docs/plans` needs no ignore set. `AGENTS.md` sits outside that scope; checking it needs `--ignore '~/*'` for its reference-clone paths.
