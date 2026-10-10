> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 8. UI architecture and package layout

Date: 2026-10-09

## Status

Accepted

## Context

Two shapes for a multi-pane Bubble Tea app appear in the reference code:

- **crush** (`charmbracelet/crush`, in the UI conventions file of its `internal/ui` package): one top-level Bubble Tea model is the only model. It owns the state, routes every message in one `Update`, tracks focus in an explicit field, and computes the layout. Sub-components are plain structs with methods the main model calls, returning a `tea.Cmd` when they need a side effect; none takes part in the message loop. Logic is split across files, never into nested models. It never does IO in `Update` and never changes state inside a command.
- **gh-dash** (`dlvhdr/gh-dash`): each pane is a sub-model with its own `Update`, composed by a root model. Its root file grew to 1946 lines anyway, and focus is spread across boolean checks on several components.

ADR 0006 tests behaviour by calling `Update` directly, which reaches the whole app when one model handles every message.

## Decision

- **The inbox tab is one model that handles every message reaching it**, following crush's conventions above. It sits inside the frame of ADR 0012, which is the app's only top-level model. The panes (queue list, conversation, composer, customer panel) are plain structs with methods; none has its own `Update`.
- **Focus is one explicit field** with a value per pane, and every key press is routed by it, so a letter typed into the composer never triggers an action key.
- **Layout follows `tea.WindowSizeMsg`**: pane sizes are computed from the terminal size, never fixed.
- **Rendering composes strings with lipgloss.** crush's Ultraviolet screen buffer is not adopted in v1; it earns its place with overlays and mouse hit-testing that v1 does not have.
- **Key bindings are a `key.Binding` key map** from `charm.land/bubbles/v2/key`, shown in a footer through `charm.land/bubbles/v2/help`.
- **ANSI-aware string handling uses `github.com/charmbracelet/x/ansi`**, never byte slicing.
- **The package layout:**
  - The `main` package in `cmd/intercomifico` calls `Run`, per ADR 0012.
  - The root package `intercomifico` holds `Run`, which reads the settings and token from ADR 0004, builds the client, the frame and the inbox, and starts the program.
  - `app` is the public tab contract from ADR 0012.
  - `internal/intercom` is the API client from ADR 0004.
  - `internal/ui` holds the frame and the inbox tab: one file per pane plus the focus, layout, key map and styles, and the interface for the Intercom calls it makes, per ADR 0006.

## Consequences

- The inbox's `Update` is the single entry point for its behaviour, so tests reach any of it by sending a message.
- The inbox model will grow. Splitting it by file, as crush does, is what keeps it readable; a nested model is not the remedy.
- Adopting Ultraviolet later, for dialogs or mouse support, is a new ADR.
