> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 13. Dialogs

Date: 2026-10-10

## Status

Accepted

## Context

Several v1 actions ask the user to pick one item from a list: inserting a saved reply (ADR 0005), assigning a conversation to me, a teammate or a team, snoozing it until a chosen time (ADR 0002), and deciding whether to clean up oversized logs (ADR 0011). Each needs room for its list while what it acts on stays in view.

Bubble Tea has no modal concept; a dialog is the app's own state. The Charm libraries supply the drawing:

- lipgloss v2 places blocks of text as layers, each at a position and depth (`lipgloss.NewLayer` with `X`, `Y` and `Z`), and `lipgloss.NewCompositor` flattens them into one string, so a box can sit over a view instead of replacing it. `lipgloss.Place` only centres a block in empty space.
- crush keeps a `Dialog` interface (`HandleMsg`, `Draw`) and a stack of open dialogs, drawn on an Ultraviolet screen, which ADR 0008 leaves out of v1.

Three other shapes were weighed: a picker that takes over a pane, which hides the context it acts on and shrinks with that pane; a picker inline in the composer, which serves saved replies only; and crush's Ultraviolet overlay, which overturns ADR 0008 for mouse support and stacking that v1 does not need.

## Decision

- **A dialog is a box drawn over the view of whoever opened it** with lipgloss layers, centred and sized to that area, with a maximum width. A tab's dialog covers that tab only, never the tab bar, since the tab of ADR 0012 that opened it owns it; a dialog about the whole app, such as log cleanup, belongs to the frame and covers the whole screen.
- **An open dialog is a value of the focus field** from ADR 0008, so every key goes to the dialog and none to the panes behind it. `esc` closes it with no effect; confirming closes it and hands the choice to the tab as a message.
- **One dialog is open at a time.** No v1 action opens a dialog from a dialog.
- **A dialog about a conversation acts on the conversation it was opened for**, by id, as a draft does under ADR 0005, even when polling changes what is on screen behind it.
- **The v1 dialogs are pickers**, each a filterable bubbles `list`:
  - saved replies, listed and inserted per ADR 0005;
  - assign, listing me first, then teammates and teams from the reference data of ADR 0004;
  - snooze, listing preset times that its implementation issue sets;
  - log cleanup, offering to delete every log file but today's and yesterday's, or to keep them all, per ADR 0011.

## Consequences

- The conversation stays visible behind every picker.
- Golden files capture an open dialog like any other view, per ADR 0006.
- Polling keeps running while a dialog is open, so the view behind it can change underneath.
- Inserting a saved reply from a trigger typed in the composer, with the same filterable list, is a later addition with its own ADR.
- A dialog stack, mouse support or Ultraviolet each come with a new ADR, superseding the parts of this one and ADR 0008 they change.
