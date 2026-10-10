> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 11. Debug log

Date: 2026-10-10

## Status

Accepted

## Context

A TUI owns the screen, so it cannot print what it is doing; when polling slows down, a request fails or the rate limit bites, the only record is one written to a file. The reference apps do this: bubbletea ships `tea.LogToFile`, gh-dash writes `debug.log` only when started with `--debug`, and crush always logs to `crush.log` in its data directory.

ADR 0007 forbids any file that could carry Intercom or personal data, and a log picks up message text, names and tokens as easily as it picks up timings. What makes a log safe is what it is allowed to contain, not where it is kept.

Intercom answers an error with a body holding a `request_id` and a list of errors, each a `code` and a `message`, and reports the rate limit in the `X-RateLimit-Limit`, `X-RateLimit-Remaining` and `X-RateLimit-Reset` headers.

## Decision

- **The app keeps a debug log, off by default.** The `debug_log` setting (`INTERCOMIFICO_DEBUG_LOG`) turns it on, a boolean defaulting to `false`, with the precedence ADR 0004 sets: the settings file turns it on for every run, the environment for one.
- **It is written with the standard library's `log/slog`** as JSON lines, to `$XDG_STATE_HOME/intercomifico/debug.log` (falling back to `~/.local/state`), beside UI state per ADR 0007.
- **Each start moves the previous log to `debug.log.1`**, so a log left on cannot grow without bound and the run that failed survives one restart.
- **It records metadata only:**
  - at start: the app's version, Go's version, OS and architecture, the colour profile and background bubbletea detected, the window size, and each setting's value and where it came from, with the token reported only as present or absent;
  - per API request: the method, the endpoint as its pattern (`/conversations/{id}`), the status, the duration, the response size, the rate-limit headers, and on an error its `request_id` and error codes;
  - per poll tick and reference-data refresh: what ran, how long it took, how many records came back, and any tick skipped because the previous one was still running;
  - per user action (reply, note, saved reply, snooze, assign, close, queue switch): which action, and whether it succeeded;
  - every error the app shows the user, by type.
- **It never records** request or response bodies, message or note text, error messages, key presses, saved-reply contents, names, emails, any Intercom id, or the token. An id is pseudonymous personal data under GDPR once anyone with access to the workspace can look it up, so the endpoint pattern stands in for it.

## Consequences

- A user reporting a problem can send the log as it is: it holds nothing about customers or teammates.
- Intercom support can trace a failed request from its `request_id`.
- Which conversation a failure belongs to is not in the log; the screen shows it, and the user reports it.
- Key presses stay out of the log because the composer receives text as key presses, so logging them would log the reply.
