> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 1. Record architecture decisions

Date: 2026-10-09

## Status

Accepted

## Context

Intercomifico starts from a blueprint written in an earlier session, which mixes decisions, example code and defects in one file. Every later issue needs a place to cite for why the app is built the way it is, and a decision that changes later must not silently rewrite the reasoning behind work already done.

## Decision

- Every architecture decision is an ADR in `docs/adr/`, one decision per file, named by a four-digit number in order of acceptance and a kebab-case title, like `0002-scope-of-v1.md`.
- Each ADR follows Michael Nygard's format: Status, Context, Decision, Consequences. Its Context cites what it was decided from: the docs read and the example code looked at.
- An ADR lands through a reviewed pull request, and the merge is its acceptance, so it is committed with the status Accepted.
- An accepted ADR is never edited except for its status line. A changed decision is a new ADR that says `Supersedes NNNN`, and the old one's status becomes `Superseded by NNNN`.
- `docs/adr/README.md` indexes every ADR with its status and is the entry point issues link to.

## Consequences

- Every implementation issue cites the ADR it implements. Work that needs a decision no ADR holds gets that ADR first.
- The blueprint stays unversioned. Anything in it that no ADR adopts is dropped.
- Changing one's mind costs a new file rather than an edit, which is the point: the old reasoning stays readable next to the new.
