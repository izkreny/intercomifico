> 🤖 Written by AI --- read/modified by izkreny! 🤓

# Plan: add the project vocabulary

Issue: #18. Its acceptance criteria are the scope; ADR 0001 sets the format ADR 0015 follows.

## Approach

The vocabulary is built from an inventory, not from memory. Every ADR, AGENTS.md and the body of every open issue is read for words the project gives one specific meaning, in the three areas the issue names: the app's concepts, Intercom's API, and how the repository is worked. General English, Go's own terms and the Charm libraries' names stay out. Each candidate is recorded with where it appears and any other name the same concept goes by, which is what turns up the conflicts.

A term Intercom's API defines takes the name `intercom_Intercom-OpenAPI` gives it, and its entry cites the schema it comes from. Every other term cites the ADR that introduces it, or AGENTS.md.

`docs/VOCABULARY.md` (new) holds one entry per term, sorted alphabetically, each a heading with the definition, the source and a "never" line naming the synonyms it must not go by. A heading per entry rather than a table row, so a term is linkable and a definition can run to a full sentence.

Where the existing docs use two names for one concept, the vocabulary picks one, preferring Intercom's name and otherwise the one the ADRs use more, and lists the other under "never". The choice is not silent: a PR comment lists each conflict as a table row, with both names, where each appears and which one the vocabulary took, so the owner overrules any of them in review. Accepted ADRs are never edited, so a conflicting word in an ADR stays where it is and the vocabulary is what settles it; rewording open issues is left to the issues themselves.

ADR 0015 (`docs/adr/0015-project-vocabulary.md` (new)) records the rule, and AGENTS.md gets a rule under *Rules* pointing agents at the file.

## Steps

- Inventory the project terms used in `docs/adr/`, AGENTS.md and every open issue body, with where each appears and every other name it goes by.
- Check each Intercom term against `intercom_Intercom-OpenAPI` and record the schema name it cites.
- Write `docs/VOCABULARY.md` (new): one alphabetical entry per term with its definition, source and "never" synonyms.
- Write ADR 0015 at `docs/adr/0015-project-vocabulary.md` (new): a branch that introduces a term adds it to `docs/VOCABULARY.md` (new) in the same PR, and the owner's review is the approval; add its row to `docs/adr/README.md`.
- Add a rule to AGENTS.md: use only the vocabulary's terms in code identifiers, docs, issues, commits and UI text, and add a needed term in the same PR rather than coin one in passing.
- Post the conflict list on the PR as a comment, one table row per conflict.

## Verification

- `python3 <skill-dir>/scripts/plan-check.py docs/plans/2026-10-11_GHI-18_add-project-vocabulary.md`
- `python3 <skill-dir>/scripts/docs-check.py docs .agents --plans docs/plans`
- `python3 <skill-dir>/scripts/docs-check.py AGENTS.md --ignore '~/*'`
- `task ci`

No gate can judge whether the inventory missed a term or whether a definition is right; that is the review round's read and the owner's. `task ci` touches no Go this branch changes, and runs because every push runs it.

## Open questions

None.

## Settled

None yet.
