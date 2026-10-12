> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 13. Dialogs

Date: 2026-10-10

## Status

Accepted

## Context

Several v1 actions ask the user to pick one item from a list: inserting a macro (ADR 0005), assigning a conversation to me, another admin or a team, snoozing it until a chosen time (ADR 0002), and deciding whether to clean up oversized logs (ADR 0011). Each needs room for its list while what it acts on stays in view.

Bubble Tea has no modal concept; a dialog is the app's own state. The Charm libraries supply the drawing:

- lipgloss v2 places blocks of text as layers, each at a position and depth (`lipgloss.NewLayer` with `X`, `Y` and `Z`), and `lipgloss.NewCompositor` flattens them into one string, so a box can sit over a view instead of replacing it. `lipgloss.Place` only centres a block in empty space.
- crush keeps a `Dialog` interface (`HandleMsg`, `Draw`) and a stack of open dialogs, drawn on an Ultraviolet screen, which ADR 0008 leaves out of v1.

Three other shapes were weighed: a picker that takes over a pane, which hides the context it acts on and shrinks with that pane; a picker inline in the composer, which serves macros only; and crush's Ultraviolet overlay, which overturns ADR 0008 for mouse support and stacking that v1 does not need.

## Decision

- **A dialog is a box drawn over the view of whoever opened it** with lipgloss layers, centred and sized to that area, with a maximum width. A tab's dialog covers that tab only, never the tab bar, since the tab of ADR 0012 that opened it owns it; a dialog about the whole app, such as log cleanup, belongs to the frame and covers the whole screen.
- **An open dialog takes every key until it closes.** A tab's dialog is a value of that tab's focus field from ADR 0008, so no key reaches the panes behind it; the frame's dialog sits in a slot the frame holds and checks before passing any key on, so no key reaches a tab. `esc` closes either with no effect; confirming closes it and hands the choice, as a message, to whichever of the tab or the frame opened it.
- **One dialog is open at a time.** No v1 action opens a dialog from a dialog.
- **A dialog about a conversation acts on the conversation it was opened for**, by id, as a draft does under ADR 0005, even when polling changes what is on screen behind it.
- **The v1 dialogs are pickers**, each a filterable bubbles `list`:
  - macros, listed and inserted per ADR 0005;
  - assign, listing me first, then admins and teams from the reference data of ADR 0004;
  - snooze, listing in 1 hour, in 4 hours, tomorrow at 9:00 and next Monday at 9:00, in the machine's local time;
  - log cleanup, offering to delete every log file but today's and yesterday's, or to keep them all, per ADR 0011.

## Consequences

- The conversation stays visible behind every picker.
- Golden files capture an open dialog like any other view, per ADR 0006.
- Polling keeps running while a dialog is open, so the view behind it can change underneath.
- Snoozing until a date and time the user types needs an input field no picker has, so it comes after v1.
- Inserting a macro from a trigger typed in the composer, with the same filterable list, is a later addition with its own ADR.
- A dialog stack, mouse support or Ultraviolet each come with a new ADR, superseding the parts of this one and ADR 0008 they change.
