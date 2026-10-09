> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 2. Scope of v1

Date: 2026-10-09

## Status

Accepted

## Context

The first blueprint mocked queues named "Open Tickets", "Snoozed Queue" and "Unassigned", an "SLA remaining" timer and a status bar. Reading Intercom's REST API (OpenAPI description 2.16, `github.com/intercom/Intercom-OpenAPI`) showed that two of those do not exist as data:

- Conversations carry no inbox field. A queue is a filter on `POST /conversations/search`: `state`, `admin_assignee_id`, `team_assignee_id` and similar fields.
- `sla_applied` carries only a name and a status (`hit`, `missed`, `active`, `cancelled`), never a deadline, so no time remaining can be shown.
- Webhooks need a public HTTPS endpoint, which a terminal app does not have.

v1 has to be the smallest loop a teammate can work their queue with.

## Decision

v1 lets one teammate, in one workspace:

- Browse three queues, each a conversation search: open conversations assigned to me, open unassigned conversations, and snoozed conversations assigned to me.
- Read a conversation with its parts, next to a customer panel showing the contact's name, email and custom attributes.
- Send a public reply or add an internal note.
- Snooze, assign (to me, another teammate or a team) and close a conversation.
- See new activity by polling, at an interval the API decision sets.

Out of v1: tickets, tags, attachments, saved replies (Intercom macros, which only the Preview API offers), SLA timers, free-text conversation search, more than one workspace, and real-time push through a webhook relay.

## Consequences

- The mock's "SLA remaining" timer is dropped rather than approximated.
- Each action above becomes at least one implementation issue; anything on the out list needs a new ADR before work starts on it.
- A webhook relay stays a `someday` idea: it needs its own hosted service, so it gets its own ADR if polling ever proves too slow.
