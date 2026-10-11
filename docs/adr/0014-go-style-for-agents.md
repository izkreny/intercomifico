> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 14. Go style for agents

Date: 2026-10-11

## Status

Accepted

## Context

Agents write most of this repository's Go. ADR 0010's linters enforce what a machine can check: formatting, imports, many of the rules in Go's Code Review Comments through `revive`, and classic bug patterns such as unchecked errors, unclosed response bodies and requests without a context. What they cannot check is judgement: how packages and interfaces are shaped, how errors are designed and wrapped, how goroutines and `context` cancellation are owned, and whether names and comments say anything.

Go's published guides cover that judgement: Effective Go, Go's Code Review Comments, Google's Go Style Guide and Uber's Go guide. Several of the reference repositories hand their agents rules of this kind in an agent file; crush keeps its Bubble Tea rules in `~/Projects/examples/go/charm/charmbracelet_crush/internal/ui/AGENTS.md`, which AGENTS.md already points UI work at.

## Decision

- **Agents load the `go-style` skill before writing Go for this repository.** It is distilled from the four guides above and covers only what the linters of ADR 0010 do not enforce.
- **The skill lives outside this repository**, with the owner's other agent skills, and AGENTS.md names it.
- **The ADRs and the linter configuration win** wherever the skill differs from them, and the skill says so.
- **Bubble Tea's rules stay out of the skill.** UI work reads crush's UI rules on top of it, as AGENTS.md says.
- **The skill exists before any Go is written**: its issue follows the toolchain issue and comes before every other Go issue.

## Consequences

- An agent writing Go meets the same judgement rules in every repository that uses the skill, not one copy per repository.
- The skill changes without a review in this repository; anything this repository requires of its Go beyond the skill is an ADR.
- A contributor without the skill still has the linters and the ADRs, which bind everyone.
