> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 3. Go and Charm library versions

Date: 2026-10-09

## Status

Accepted

## Context

Charm's current stack is v2, under the `charm.land` import paths, while much published example code is still written against v1: `github.com/charmbracelet/bubbletea` imports, `View() string`, `tea.KeyMsg`. The canonical upstream repositories, read at their latest tags, show:

- `charm.land/bubbletea/v2` at v2.1.0, `charm.land/bubbles/v2` at v2.2.1 and `charm.land/lipgloss/v2` at v2.0.6, each requiring Go 1.26 or later in its `go.mod`.
- In v2, `View()` returns a `tea.View` value that also carries alt-screen and mouse mode, and key presses arrive as `tea.KeyPressMsg`. The upgrade guide in `github.com/charmbracelet/bubbletea`, UPGRADE_GUIDE_V2.md, lists every change.
- `github.com/charmbracelet/bubbletea-app-template` is still on v1, so its CI, lint and release layout is reusable but its code is not.

The newest Go release is 1.27.1.

## Decision

- `go.mod` declares Go 1.27, the newest release, and moves to each new Go release once the Charm libraries build on it.
- The UI is built on bubbletea v2, bubbles v2 and lipgloss v2, imported from `charm.land`, at the latest tags above.
- Code written against the v1 API, from any source, is reference for ideas only and is never copied.

## Consequences

- Requiring the newest Go means contributors need it too; with one committer and a mise-managed toolchain, that costs nothing today.
- Examples in community apps still on v1 need translating through the upgrade guide before they help.
