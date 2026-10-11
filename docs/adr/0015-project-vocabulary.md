> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 15. Project vocabulary

Date: 2026-10-11

## Status

Accepted

## Context

Agents write most of this repository, and an agent that meets a concept it has no word for coins one. Each coinage is harmless alone; together they leave the code, the docs and the tracker calling one thing by several names. An inventory of the ADRs, AGENTS.md and the open issues already finds it: Intercom's macro is a saved reply across the ADRs and its admin a teammate, and ADR 0012 alone calls a build carrying extra tabs a custom build, a company's own build and an outside build.

Intercom's OpenAPI description (`github.com/intercom/Intercom-OpenAPI`, `descriptions/2.16`) names every object the app reads and writes, and `internal/intercom` mirrors those names, so a different word in the UI or the docs is a translation every reader has to make.

The reviewer of every pull request reads AGENTS.md as the repository's standards, so a rule there is enforced in review with no further tooling.

## Decision

- **`docs/VOCABULARY.md` is the closed list of the project's terms.** Each concept has one name, one entry, a definition, its source, and the names it is never called.
- **A term Intercom's API defines takes Intercom's name**, citing the schema in the OpenAPI description.
- **A branch that introduces a term adds it to `docs/VOCABULARY.md` in the same pull request**, and the owner's review of that pull request is the term's approval. A term is never coined in passing.
- **Code identifiers, docs, issues, commits and UI text use only the vocabulary's terms**, and AGENTS.md says so.
- **Renaming a term is an edit to `docs/VOCABULARY.md`**, reviewed like any other. Accepted ADRs keep the words they were written with, per ADR 0001; where one differs from the vocabulary, the vocabulary wins.

## Consequences

- A word missing from the vocabulary, or one listed under **Never**, becomes a review finding.
- Every term is approved before it spreads, at the cost of one more file in the diff of the branch that needs it.
- Older ADRs read with words the vocabulary has since ruled out, such as saved reply; their meaning is unchanged and the vocabulary names the term that replaced each.
- General English, Go's terms and the Charm libraries' names stay out, so the list stays short enough to read whole.
