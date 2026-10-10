> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 2. Scope of v1

Date: 2026-10-09

## Status

Accepted

## Context

Intercom's REST API (OpenAPI description 2.16, `github.com/intercom/Intercom-OpenAPI`) sets what a terminal inbox can be built on:

- Conversations carry no inbox field. A queue is a filter on `POST /conversations/search`: `state`, `admin_assignee_id`, `team_assignee_id` and similar fields.
- `sla_applied` carries only a name and a status (`hit`, `missed`, `active`, `cancelled`), never a deadline, so no time remaining can be shown.
- Webhooks need a public HTTPS endpoint, which a terminal app does not have.

v1 has to be the smallest loop a teammate can work their queue with.

## Decision

v1 lets one teammate, in one workspace:

- Browse two queues, each a conversation search:
  - **Mine**: conversations assigned to me, open ones first and snoozed ones after them. Closed conversations never appear.
  - **Unassigned**: open conversations assigned to no teammate.
- Read a conversation with its parts, next to a customer panel showing the contact's name, email and custom attributes.
- Send a public reply or add an internal note.
- Insert a saved reply (an Intercom macro) into the composer.
- Snooze, assign (to me, another teammate or a team) and close a conversation.
- See new activity by polling, at an interval the API decision sets.

Out of v1: further queues, such as my teams' conversations assigned to no teammate or views by priority; tickets, tags, attachments, the user's own local saved replies, SLA timers, free-text conversation search, more than one workspace, and real-time push through a webhook relay.

## Consequences

- No SLA countdown is shown, since the API exposes no deadline to count down to.
- Each action above becomes at least one implementation issue; anything on the out list needs a new ADR before work starts on it.
- A webhook relay stays a `someday` idea: it needs its own hosted service, so it gets its own ADR if polling ever proves too slow.
