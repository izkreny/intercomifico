> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 7. What the app writes to disk

Date: 2026-10-09

## Status

Accepted

## Context

The app fetches conversations, contacts, admins, teams and macros from Intercom, and could keep any of them on disk between runs. Against that:

- Intercom is the source of truth, and ADR 0004 already polls to keep the app current, so a stored copy is stale as soon as an admin acts elsewhere.
- Conversations and contacts are contacts' personal data, and admin and team lists carry admins' names and emails. The workspace is hosted in the EU, so any of it on disk would fall under GDPR and need encryption, expiry and deletion on request.
- Unsent drafts are the user's own text, but written to a contact, so they routinely repeat the contact's details. Log files pick up message text and tokens just as easily.
- Starting from nothing costs a handful of calls (`GET /me`, one queue search, admins, teams and every page of macros), far inside the rate limit.
- The reference repositories keep remote data in memory too: gh-dash stores only the user's own bookmarks, dhth_prs only a 30-second HTTP cache, and crush's SQLite database holds crush's own sessions rather than a copy of a remote service.

## Decision

- **The app never writes Intercom data or personal data to disk**: no conversations, contacts, admins, teams, macros, drafts, caches or logs that could carry any of them. Fetched data lives in memory for the length of a run and is fetched again at the next start.
- **The access token is never written anywhere**, per ADR 0004.
- **The only files the app may ever use are its own:**
  - settings in `$XDG_CONFIG_HOME/intercomifico/settings.toml` (falling back to `~/.config` when the variable is unset), written by the user and only ever read by the app;
  - UI state, such as the last visible queue, in `$XDG_STATE_HOME/intercomifico/` (falling back to `~/.local/state`), written by the app;
  - the debug logs in a `logs` folder of the same directory, holding metadata only, per ADR 0011.
- **v1 reads the settings file**, with the precedence ADR 0004 sets, writes the debug log when it is turned on, and keeps no UI state.

## Consequences

- No contact or admin data is left on the machine after the app exits, so the app has no data to secure, expire or delete.
- Every start shows a short loading state while the first calls return.
- UI state arrives with the first feature that needs it, through that feature's ADR and inside this decision.
- Saving drafts across a crash and caching macros are ruled out rather than deferred. Revisiting any of them means superseding this ADR.
