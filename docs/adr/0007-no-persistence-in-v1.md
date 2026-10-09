> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 7. No persistence in v1

Date: 2026-10-09

## Status

Accepted

## Context

The app fetches conversations, contacts, admins and teams from Intercom, and could keep them on disk between runs. Against that:

- Intercom is the source of truth, and ADR 0004 already polls to keep the app current, so a stored copy is stale as soon as a teammate acts elsewhere.
- Conversations and contacts are customer personal data. The workspace is hosted in the EU, so a copy on disk would fall under GDPR and need encryption, expiry and deletion on request.
- Starting from nothing costs a handful of calls (`GET /me`, one queue search, admins and teams), far inside the rate limit.
- The reference apps keep remote data in memory too: gh-dash stores only the user's own bookmarks, dhth_prs only a 30-second HTTP cache, and crush's SQLite database holds crush's own sessions rather than a copy of a remote service.

## Decision

- **v1 writes nothing it fetched from Intercom to disk.** All fetched data lives in memory for the length of a run and is fetched again at the next start.
- **No database, cache directory or state file in v1.**

## Consequences

- No customer data is left on the machine after the app exits.
- Every start shows a short loading state while the first calls return.
- Two later needs will reopen this, each with its own ADR: saved replies (macros, out of v1 per ADR 0002), which will want a local copy, and the user's own unsent drafts, which are worth surviving a crash. Neither decision is made here.
