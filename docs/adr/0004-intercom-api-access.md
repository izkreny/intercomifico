> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 4. Intercom API access

Date: 2026-10-09

## Status

Accepted

## Context

v1 needs about ten REST endpoints, per ADR 0002. What Intercom offers, read from its OpenAPI description (`github.com/intercom/Intercom-OpenAPI`, `descriptions/2.16`) and its developer docs at https://developers.intercom.com/:

- The newest stable API version is 2.16, chosen per request with the `Intercom-Version` header; without it, the version set on the app in the Developer Hub applies. Preview, formerly Unstable, can change without a new version number.
- A tool for one's own workspace authenticates with an access token from a private app in the Developer Hub, sent as `Authorization: Bearer <token>`. OAuth is meant for public apps.
- Each region has its own host: `api.intercom.io` (US), `api.eu.intercom.io` (EU) and `api.au.intercom.io` (AU).
- Admin actions need the acting admin's `admin_id`, which `GET /me` returns for the token's owner.
- Lists paginate with a `starting_after` cursor. Rate limits are 10,000 calls a minute per app, enforced in 10-second windows, with `X-RateLimit-Remaining` and `X-RateLimit-Reset` headers and a 429 once exceeded.
- No Go SDK is current: `github.com/intercom/intercom-go` targets API 1.3 and last released in 2015. The SDKs generated from the OpenAPI description (TypeScript, Python, Java, PHP) and the hand-written Ruby one are current.

## Decision

- **A hand-written client in `internal/intercom`**, covering only the endpoints v1 uses, on `net/http` with a `context.Context` on every call. It imports nothing from the UI, so it can move into its own repository unchanged if it ever gains a second user. The official SDKs are models for its naming, pagination and errors, never dependencies.
- **Every request pins `Intercom-Version: 2.16`.** Moving to a newer version is a new ADR.
- **The token comes from the `INTERCOMIFICO_TOKEN` environment variable and nowhere else.** The app never writes it to disk, never logs it, and refuses to start without it. Where it is stored is the user's choice, such as, in Fish, `set -x INTERCOMIFICO_TOKEN (secret-tool lookup intercom token)`.
- **Every other setting has a built-in default, overridden by the settings file, overridden in turn by an environment variable.** The file is `$XDG_CONFIG_HOME/intercomifico/settings.toml` per ADR 0007, and a missing file is not an error. The file holds what the user always wants; the environment holds what one run wants. The app refuses to start when the file contains a token, so the token can never end up in it.
- **The region is the `region` setting** (`INTERCOMIFICO_REGION`), one of `eu`, `us` or `au`, defaulting to `eu`, where the owner's workspace is hosted.
- **The acting admin is resolved once at startup with `GET /me`**, and its id goes into every reply and conversation action.
- **Queues are refreshed by polling** at the `poll_interval` setting (`INTERCOMIFICO_POLL_INTERVAL`; a Go duration, default `2s`, `0` to disable): each tick re-runs the visible queue's search and refetches the open conversation with `GET /conversations/{id}`, whether or not the search still returns it, so a conversation closed, snoozed or reassigned elsewhere shows its new state.
- **Reference data is refreshed on a slow timer**, so a session that runs for days stays current: every 15 minutes the app fetches the macros changed since its last fetch (`updated_since`) along with the admins and teams, and every hour it re-fetches all macros, so deleted ones disappear.
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
| Saved replies                        | `GET /macros`, every page at startup, then refreshed on the reference-data timer               |

## Consequences

- Every setting but the token can live in the settings file or the environment; the token lives in the environment only. Later settings, such as user-defined key bindings, join the same file and the same precedence.
- Two-second polling, one search and one conversation fetch a tick, is sixty calls a minute: under 1% of the 10,000 a minute each app may make, and of the 25,000 a minute the workspace shares across all its apps.
- Intercom documents no freshness guarantee for conversation search, so a change made elsewhere shows up within one poll interval at best.
- Which scopes the private app needs is only partly documented: "Read conversations" and "Write conversations", plus read access to admins and contacts. The first working build confirms the set, and the README records it.
