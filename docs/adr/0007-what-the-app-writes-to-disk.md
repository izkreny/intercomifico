> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 7. What the app writes to disk

Date: 2026-10-09

## Status

Accepted

## Context

The app fetches conversations, contacts, admins, teams and saved replies from Intercom, and could keep any of them on disk between runs. Against that:

- Intercom is the source of truth, and ADR 0004 already polls to keep the app current, so a stored copy is stale as soon as a teammate acts elsewhere.
- Conversations and contacts are customers' personal data, and admin and team lists carry teammates' names and emails. The workspace is hosted in the EU, so any of it on disk would fall under GDPR and need encryption, expiry and deletion on request.
- Unsent drafts are the user's own text, but written to a customer, so they routinely repeat the customer's details. Log files pick up message text and tokens just as easily.
- Starting from nothing costs a handful of calls (`GET /me`, one queue search, admins, teams and every page of macros), far inside the rate limit.
- The reference apps keep remote data in memory too: gh-dash stores only the user's own bookmarks, dhth_prs only a 30-second HTTP cache, and crush's SQLite database holds crush's own sessions rather than a copy of a remote service.

## Decision

- **The app never writes Intercom data or personal data to disk**: no conversations, contacts, admins, teams, saved replies, drafts, caches or logs that could carry any of them. Fetched data lives in memory for the length of a run and is fetched again at the next start.
- **The access token is never written anywhere**, per ADR 0004.
- **The only files the app may ever use are its own:**
  - settings in `$XDG_CONFIG_HOME/intercomifico/settings.toml` (falling back to `~/.config` when the variable is unset), written by the user and only ever read by the app;
  - UI state, such as the last open queue, in `$XDG_STATE_HOME/intercomifico/` (falling back to `~/.local/state`), written by the app.
- **v1 uses neither.** Its settings are ADR 0004's environment variables, and it keeps no UI state.

## Consequences

- No customer or teammate data is left on the machine after the app exits, so the app has no data to secure, expire or delete.
- Every start shows a short loading state while the first calls return.
- Settings move into the settings file with the first setting that needs a file, and UI state with the first feature that needs it; each arrives through the ADR for that feature, inside this decision.
- Saving drafts across a crash, caching saved replies and a debug log are ruled out rather than deferred. Revisiting any of them means superseding this ADR.
