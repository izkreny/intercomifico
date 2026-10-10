> 🤖 Written by AI --- read/modified by izkreny! 🤓

# Plan: settle the app architecture

Issue: #1. Its questions and its `Done when` list are the scope; this plan only says how the answers get written down.

## Approach

Each question in #1 is settled with the owner one at a time, in an order where earlier answers constrain later ones: scope first, because it decides which API calls matter, then the API, the editor, the stack, and structure and testing last, since the layout follows from everything before it. Questions the answers raise, such as what the app may write to disk and how it is themed, are settled the same way, each in its own ADR.

Every settled decision becomes its own ADR in `docs/adr/` (new), in Michael Nygard's format: Status, Context, Decision, Consequences. An ADR is never edited once accepted; a changed decision is a new ADR that supersedes it, which is why one decision per file matters. Facts about Intercom's API and the Charm libraries are read from their current docs during the spike, never recalled, and each ADR cites what it read.

## Steps

- Add `docs/adr/0001-record-architecture-decisions.md` (new), fixing the ADR format, the numbering, and the rule that an accepted ADR is superseded rather than edited.
- Settle v1 scope with the owner, saved replies included, and record it as an ADR.
- Settle Intercom API access (token supply and storage, pinned API version, regions, how new messages arrive, the saved-replies endpoint) and the client living in this repository, and record them as an ADR.
- Settle the reply composer, its saved-reply picker included, and record it as an ADR.
- Settle the Go version and the Charm library major versions and record them as an ADR.
- Settle testing (coverage, golden files, fakes for the API client) and record it as an ADR.
- Settle the Go tooling (formatter, linter, task runner, toolchain pins) and the exact check commands, and record them as an ADR.
- Settle what the app may write to disk and record it as an ADR.
- Settle the debug log and record it as an ADR.
- Settle the UI architecture and package layout and record them as an ADR.
- Settle theming and record it as an ADR.
- Add `docs/adr/README.md` (new), indexing every ADR with its status, as the entry point later issues link to.
- Add `AGENTS.md` (new), pointing agents at the ADRs and at the reference clones: Charm's own repositories as canonical, community apps for ideas, and Intercom's OpenAPI description and SDKs.
- Add `.agents/gh-solo.md` (new), recording the layer label set, the check commands from the tooling ADR, and the docs-check scope.
- Propose the first implementation issues from the accepted ADRs through the tracker's create flow.

## Verification

- `python3 <skill-dir>/scripts/plan-check.py docs/plans/2026-10-09_GHI-1_settle-app-architecture.md`
- `python3 <skill-dir>/scripts/docs-check.py docs .agents --plans docs/plans`

Neither gate can tell whether a decision is right, only that the files are well formed. Whether each ADR answers its question from #1, and whether the answers fit together, is the owner's judgement on the diff.

## Open questions

None.

## Settled

- CI is decided inside its own `infra` issue; the tooling ADR names the check commands, and that issue turns them into required status checks on `main`.
