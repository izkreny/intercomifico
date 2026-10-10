# gh-solo configuration for intercomifico

## Labels

The gh-solo default taxonomy, unchanged. Layer is the mandatory axis, exactly one per issue except an epic:

- `backend`: the Intercom API client, `internal/intercom`.
- `frontend`: the terminal UI, `internal/ui`.
- `fullstack`: one feature that needs both and cannot be accepted from either side alone.
- `infra`: CI, releases, repository tooling.
- `docs`: prose that is the deliverable, such as an ADR or the README.

The other axes keep their default labels (`bug`, `spike`, `epic`, `urgent`, `someday`, `draft`, `blocked`), each created the first time an issue needs it.

## Issue types

None: the repository belongs to a personal account.

## Check commands

Per ADR 0006:

- `go vet ./...`
- `go test -race ./...`

`go test ./internal/ui/... -update` rewrites golden files after a deliberate rendering change. It is not a gate; the changed `.golden` files are reviewed in the diff. No linter is set up yet; choosing one belongs to the CI issue.

## Docs check

`python3 <skill-dir>/scripts/docs-check.py docs .agents --plans docs/plans` needs no ignore set. `AGENTS.md` sits outside that scope; checking it needs `--ignore '~/*'` for its reference-clone paths.

## Pull request bodies

`## Steps` and `## Verification` are plain bullets, never checkboxes, matching the plan file. `## Open questions` is dropped once nothing is open, and every answered question moves to `## Settled`. Gate results go in the implementation record comment rather than in ticked boxes, and `ready` and `merge` read them there.
