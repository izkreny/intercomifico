> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 4. Intercom API access

Date: 2026-10-09

## Status

Accepted

## Context

v1 needs about ten REST endpoints, per ADR 0002. What Intercom offers, read from its OpenAPI description (`github.com/intercom/Intercom-OpenAPI`, `descriptions/2.16`) and its developer docs:

- The newest stable API version is 2.16, chosen per request with the `Intercom-Version` header; without it, the version set on the app in the Developer Hub applies. Preview, formerly Unstable, can change without a new version number.
- A tool for one's own workspace authenticates with an access token from a private app in the Developer Hub, sent as `Authorization: Bearer <token>`. OAuth is meant for public apps.
- Each region has its own host: `api.intercom.io` (US), `api.eu.intercom.io` (EU) and `api.au.intercom.io` (AU).
- Admin actions need the acting admin's `admin_id`, which `GET /me` returns for the token's owner.
- Lists paginate with a `starting_after` cursor. Rate limits are 10,000 calls a minute per app, enforced in 10-second windows, with `X-RateLimit-Remaining` and `X-RateLimit-Reset` headers and a 429 once exceeded.
- No Go SDK is current: `github.com/intercom/intercom-go` targets API 1.3 and last released in 2015. The SDKs generated from the OpenAPI description (TypeScript, Python, Java, PHP) and the hand-written Ruby one are current.

## Decision

- **A hand-written client in `internal/intercom`**, covering only the endpoints v1 uses, on `net/http` with a `context.Context` on every call. It imports nothing from the UI, so it can move into its own repository unchanged if it ever gains a second user. The official SDKs are models for its naming, pagination and errors, never dependencies.
- **Every request pins `Intercom-Version: 2.16`.** Moving to a newer version is a new ADR.
- **The token comes from the `INTERCOMIFICO_TOKEN` environment variable and nowhere else.** The app never writes it to disk, never logs it, and refuses to start without it. Where it is stored is the user's choice, such as `set -x INTERCOMIFICO_TOKEN (secret-tool lookup intercom token)`.
- **The region comes from `INTERCOMIFICO_REGION`**, one of `eu`, `us` or `au`, defaulting to `eu`, where the owner's workspace is hosted.
- **The acting admin is resolved once at startup with `GET /me`**, and its id goes into every reply and conversation action.
- **Queues are refreshed by polling** every `INTERCOMIFICO_POLL_INTERVAL` (a Go duration, default `10s`, `0` to disable): each tick re-runs the visible queue's search and refetches the open conversation when its `updated_at` moved.
- **A 429 becomes a typed error carrying the reset time.** Polling pauses until then, and a user action that hits it shows the error instead of retrying behind the user's back.

The v1 endpoints:

| Need                                 | Call                                                                                           |
|--------------------------------------|------------------------------------------------------------------------------------------------|
| Who am I                             | `GET /me`                                                                                      |
| Queue contents                       | `POST /conversations/search`, filtered on `state`, `admin_assignee_id`, `team_assignee_id`     |
| One conversation and its parts       | `GET /conversations/{id}`                                                                      |
| Public reply, internal note          | `POST /conversations/{id}/reply`, `message_type` `comment` or `note`                           |
| Snooze, assign, close                | `POST /conversations/{id}/parts`, `message_type` `snoozed`, `assignment` or `close`            |
| Customer panel                       | `GET /contacts/{id}`                                                                           |
| Assignment targets                   | `GET /admins`, `GET /teams`                                                                    |

## Consequences

- Configuration in v1 is three environment variables. A config file arrives with its own ADR when something needs one, such as user-defined key bindings.
- Ten seconds of polling over one queue and one conversation is a few dozen calls a minute, far inside the rate limit.
- Search results can lag a few minutes behind changes made elsewhere, per Intercom's own description of contact search, so a refresh is near real-time at best.
- Which scopes the private app needs is only partly documented: "Read conversations" and "Write conversations", plus read access to admins and contacts. The first working build confirms the set, and the README records it.
