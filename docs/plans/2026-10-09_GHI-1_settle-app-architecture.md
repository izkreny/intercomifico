> 🤖 Written by AI --- read/modified by izkreny! 🤓

# Plan: settle the app architecture

Issue: #1. Its questions and its `Done when` list are the scope; this plan only says how the answers get written down.

## Approach

The first blueprint is input, not a spec. Each question in #1 is settled with the owner one at a time, in an order where earlier answers constrain later ones: scope first, because it decides which API calls matter, then the API, the editor, the stack, and structure and testing last, since the layout follows from everything before it.

Every settled decision becomes its own ADR in `docs/adr/` (new), in Michael Nygard's format: Status, Context, Decision, Consequences. An ADR is never edited once accepted; a changed decision is a new ADR that supersedes it, which is why one decision per file matters. Facts about Intercom's API and the Charm libraries are read from their current docs during the spike, never recalled, and each ADR cites what it read.

The blueprint's defects listed in #1 are resolved inside the ADR whose decision they belong to, so none needs an ADR of its own.

## Steps

- Add `docs/adr/0001-record-architecture-decisions.md` (new), fixing the ADR format, the numbering, and the rule that an accepted ADR is superseded rather than edited.
- Settle v1 scope with the owner and record it as an ADR.
- Settle Intercom API access (token supply and storage, pinned API version, how new messages arrive) and record it as an ADR.
- Settle the reply editor (in-TUI textarea, external editor, or both in sequence, and how the editor is chosen) and record it as an ADR.
- Settle the Go version and the Charm library major versions and record them as an ADR.
- Settle package layout, the API client interface used for fakes, and the coverage target, and record them as an ADR that names the exact check commands.
- Add `docs/adr/README.md` (new), indexing every ADR with its status, as the entry point later issues link to.
- Add `.agents/gh-solo.md` (new), recording the layer label set, the check commands from the testing ADR, and the docs-check ignore set.
- Propose the first implementation issues from the accepted ADRs through the tracker's create flow.

## Verification

- `python3 <skill-dir>/scripts/plan-check.py docs/plans/2026-10-09_GHI-1_settle-app-architecture.md`
- `python3 <skill-dir>/scripts/docs-check.py docs .agents --plans docs/plans`

Neither gate can tell whether a decision is right, only that the files are well formed. Whether each ADR answers its question from #1, and whether the answers fit together, is the owner's judgement on the diff.

## Open questions

None.

## Settled

- CI is decided inside its own `infra` issue; the testing ADR names the check commands, and that issue turns them into required status checks on `main`.
