> 🤖 Written by AI --- read/modified by izkreny! 🤓

# Plan: add the go style skill

Issue: #4. Its acceptance criteria are the scope; ADR 0014 holds the decision and ADR 0010 the linters the skill must not repeat.

## Approach

The branch has two deliverables in two places. The `go-style` skill lives outside this repository, with the owner's other agent skills, as ADR 0014 decides; this branch carries only the plan and the AGENTS.md line that names the skill. The skill is written and checked here, and committing it in its own tree is the owner's step, so the review round on this PR sees the AGENTS.md line alone.

"Does not restate a rule golangci-lint already enforces" is only checkable against a written list, so the first step builds one: the linters golangci-lint v2 runs by default, the ones `.golangci.yml` enables, revive's default rule set (it runs unconfigured here) and staticcheck's checks, read from the golangci-lint and revive sources rather than recalled. Each guide's rules are then sorted against that list, and a rule a linter catches is dropped, or, where the linter catches only part of it, kept for the part it misses.

The skill is repository-neutral, since one copy serves every repository that uses it: it says that the repository's ADRs and linter configuration win where they differ from it, without naming this one's. Bubble Tea's rules stay out of it; the AGENTS.md line sends UI work to crush's UI rules on top of it.

## Steps

- Write down what the configured linters enforce: golangci-lint v2's default linters, the ones `.golangci.yml` enables with its settings, revive's default rules and staticcheck's checks, each from its source.
- Distil Effective Go, Go's Code Review Comments, Google's Go Style Guide and Uber's Go guide into rules on package and interface design, error design, concurrency and `context`, and names and comments, dropping each rule the list above already covers.
- Write `go-style` as a new skill with the owner's other agent skills, structured per the skill-authoring conventions: a short SKILL.md that routes to one reference per area, and a section saying the repository's ADRs and linter configuration win.
- Run the skill check on `go-style` and fix what it reports.
- Add a rule to AGENTS.md telling agents to load `go-style` before writing Go, with crush's Bubble Tea rules on top for UI work.

## Verification

- `python3 <skill-dir>/scripts/plan-check.py docs/plans/2026-10-11_GHI-4_add-go-style-skill.md`
- `python3 <skill-dir>/scripts/docs-check.py docs .agents --plans docs/plans`
- `python3 <skill-dir>/scripts/docs-check.py AGENTS.md --ignore '~/*'`
- `node <skills-guru-dir>/scripts/check.js <go-style-dir>`
- `task ci`

No gate here can judge the skill's content: the review round sees only AGENTS.md, so whether each rule is right, and whether one slipped through that a linter already enforces, is the owner's read of the skill, or a `skills-guru review` of it. `task ci` touches no Go this branch changes, and runs because every push runs it. The skill's commit in its own tree is not a gate on this branch.

## Open questions

None.

## Settled

None yet.
