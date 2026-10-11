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
- Desktop keyrings on Linux share one interface, the freedesktop Secret Service over D-Bus, served by KWallet, GNOME Keyring and KeePassXC alike; `secret-tool` is its command-line client. `github.com/zalando/go-keyring`, which GitHub's `gh` uses, reads it, and the macOS Keychain and Windows Credential Manager too. It finds an entry by its `service` and `username` attributes, and asks the keyring to unlock a locked collection, which shows the keyring's own password dialog.
- No Go SDK is current: `github.com/intercom/intercom-go` targets API 1.3 and last released in 2015. The SDKs generated from the OpenAPI description (TypeScript, Python, Java, PHP) and the hand-written Ruby one are current.

## Decision

- **A hand-written client in `internal/intercom`**, covering only the endpoints v1 uses, on `net/http` with a `context.Context` on every call. It imports nothing from the UI, so it can move into its own repository unchanged if it ever gains a second user. The official SDKs are models for its naming, pagination and errors, never dependencies.
- **Every request pins `Intercom-Version: 2.16`.** Moving to a newer version is a new ADR.
- **The token comes from the `INTERCOMIFICO_TOKEN` environment variable when it is set, and otherwise from the system keyring, and nowhere else.** The app reads the keyring with `github.com/zalando/go-keyring`, looking up the entry whose `service` is `intercomifico` and whose `username` is `token`; both values are fixed in the code, not settings. The user stores that entry once, on Linux with `secret-tool store --label='Intercomifico token' service intercomifico username token`, and the app only ever reads it. The app never writes the token anywhere, never logs it, and refuses to start when neither source has it, with a message giving both ways to set it. The environment variable serves a machine with no keyring, such as one reached over SSH; handing it to the app's process alone, rather than exporting it into the shell, keeps it from every other program that shell starts, AI agents included.
- **Every other setting has a built-in default, overridden by the settings file, overridden in turn by an environment variable.** The file is `$XDG_CONFIG_HOME/intercomifico/settings.toml` per ADR 0007, and a missing file is not an error. The file holds what the user always wants; the environment holds what one run wants. The file is parsed with `github.com/pelletier/go-toml/v2`, its decoder set to `DisallowUnknownFields()` so that it rejects any key the settings do not define, so the app refuses to start on a misspelt key or on a token in the file, and the token can never end up in it.
- **The region is the `region` setting** (`INTERCOMIFICO_REGION`), one of `eu`, `us` or `au`, defaulting to `eu`, where the owner's workspace is hosted.
- **The acting admin is resolved once at startup with `GET /me`**, and its id goes into every reply and conversation action.
- **Queues are refreshed by polling** at the `poll_interval` setting (`INTERCOMIFICO_POLL_INTERVAL`; a Go duration, default `2s`, `0` to disable): each tick re-runs the visible queue's search and refetches the open conversation with `GET /conversations/{id}`, whether or not the search still returns it, so a conversation closed, snoozed or reassigned elsewhere shows its new state.
- **Reference data is refreshed on a slow timer**, so a session that runs for days stays current: every 15 minutes the app fetches the macros changed since its last fetch (`updated_since`) along with the admins and teams, and every hour it re-fetches all macros, so deleted ones disappear.
- **A refresh key runs a poll tick and the reference-data refresh at once**, and restarts both timers. The owner chooses the key in the inbox's implementation issue, checked against the textarea's own bindings, as ADR 0005 does for the composer's keys.
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

- Every setting but the token can live in the settings file or the environment; the token lives in the environment or the keyring only. Later settings, such as user-defined key bindings, join the same file and the same precedence.
- Whether starting the app asks for a password is the keyring's own setting, not the app's: a keyring unlocked at login asks nothing, and one set to lock again after use, such as KWallet closing the wallet when the last application stops using it, shows its unlock dialog at every start.
- On a desktop with no Secret Service running, such as a bare window manager without GNOME Keyring or KeePassXC, the keyring lookup fails and the environment variable is the only source.
- Two-second polling, one search and one conversation fetch a tick, is sixty calls a minute: under 1% of the 10,000 a minute each app may make, and of the 25,000 a minute the workspace shares across all its apps.
- Intercom documents no freshness guarantee for conversation search, so a change made elsewhere shows up within one poll interval at best.
- Which scopes the private app needs is only partly documented: "Read conversations" and "Write conversations", plus read access to admins and contacts. The first working build confirms the set, and the README records it.
