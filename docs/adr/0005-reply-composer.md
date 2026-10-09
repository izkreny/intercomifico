> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 5. Reply composer

Date: 2026-10-09

## Status

Accepted

## Context

The first blueprint planned a bubbles textarea first and a full-screen Neovim handoff second, with the conversation history in a split window. Its handoff code passed customer text through `sh -c`, which lets a customer message run commands on the replying teammate's machine.

How comparable apps compose text:

- gh-dash comments through its `inputbox` component, a wrapper around the bubbles textarea: `ctrl+d` submits, `esc` cancels, and no external editor is involved.
- crush composes in a bubbles textarea where `enter` sends and `shift+enter` or `ctrl+j` adds a newline, with `ctrl+o` opening an external editor as an extra.
- The bubbles textarea (`charm.land/bubbles/v2/textarea`) is a multi-line input with word wrap, cursor movement, paste and a configurable key map. It binds many control keys itself, `ctrl+n`, `ctrl+d` and `ctrl+e` among them.

## Decision

- **v1 composes replies and notes in a bubbles textarea inside the app, and nowhere else.** No external editor and no subprocess.
- `enter` inserts a newline, since support replies are usually several lines; a dedicated key sends, and sending is never implicit.
- Leaving the composer keeps the draft. Nothing the user typed is discarded without an explicit action.
- The composer shows whether it holds a public reply or an internal note, and switching between them keeps the text.
- The exact key bindings are set in the composer's implementation issue and checked against the textarea's own bindings, so none of them is shadowed.

## Consequences

- Composing works the same in every terminal and needs no editor configuration.
- Long replies are less comfortable than in a real editor. Opening `$VISUAL` or `$EDITOR` on the draft, with the conversation history below a cut line as `git commit` does, is the planned extension, and it gets its own ADR when it is picked up.
- The blueprint's Neovim handoff, and with it the shell-injection risk, is dropped.
