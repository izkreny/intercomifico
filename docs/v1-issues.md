> 🤖 Written by AI --- read/modified by izkreny! 🤓

# Proposed v1 issues

The breakdown of v1 into issues, plus one issue outside the v1 milestone, proposed through the tracker's create flow and committed for review only: this file is deleted before the pull request merges, once the issues exist. Each section below is one issue: its title, labels, milestone and blockers, then the body as it will be posted, its headings one level deeper here than on GitHub. A `#n` in a body means row n here; the real issue numbers are assigned at creation.

---

## 1. set up the go toolchain

Labels: `infra`. Milestone: `v1`. Blocked by: none.

> 🤖 Written by AI --- read/modified by izkreny! 🤓

### Overview

Create the Go module, the package layout and the tooling of ADR 0010, so every later issue starts with working gates.

### Acceptance criteria

- [ ] `mise install` in a fresh clone installs Go 1.27.1, golangci-lint 2.14.0, task 3.54.0 and govulncheck 1.8.0, as mise.toml pins them
- [ ] The module builds with ADR 0008's layout: the `main` package in `cmd/intercomifico` calling `Run` in the root package `intercomifico`, plus `app`, `internal/intercom` and `internal/ui`
- [ ] `task fmt`, `task lint`, `task test`, `task test:update`, `task audit`, `task ci` and `task run` exist and do what ADR 0010 says
- [ ] `task lint` runs the linters and formatters of bubbletea's .golangci.yml, with test files linted
- [ ] `task ci` runs the formatting check, `task lint`, `task test` and `task audit` in that order, and stops at the first failure
- [ ] `task audit` fails on found vulnerabilities (exit 3); on a network error or its 60-second timeout it fails when `CI` is set and warns and passes otherwise
- [ ] Each gate has been seen failing once: an unformatted file, a lint error, a failing test

### Technical notes

- ADR 0010 for the tools, versions and tasks; ADR 0008 and ADR 0012 for the package layout.
- Copy .golangci.yml from `~/Projects/examples/go/charm/charmbracelet_bubbletea/` and drop its test-file exclusion.
- govulncheck is pinned through mise's `go` backend, since the mise registry does not carry it.

### Dependencies

None.

---

## 2. run task ci in github actions

Labels: `infra`. Milestone: `v1`. Blocked by: #1.

> 🤖 Written by AI --- read/modified by izkreny! 🤓

### Overview

Run the repository's one gate command in CI, and make it required on `main`, so nothing merges red.

### Acceptance criteria

- [ ] A workflow on every push and pull request installs the toolchain from mise.toml and runs `task ci`, adding no check of its own
- [ ] A workflow scheduled weekly, on Monday morning, runs `task audit` on `main`
- [ ] The `task ci` job is a required status check on `main`
- [ ] Dependabot alerts are enabled on the repository, and no Dependabot version-update configuration exists
- [ ] A pull request carrying a lint error shows the check red

### Technical notes

- ADR 0010: CI runs `task ci` and adds nothing; the audit fails on a network problem there because GitHub Actions sets `CI`.
- The spike's plan settled that this issue turns the gates into required status checks.

### Dependencies

Blocked by #1.

---

## 3. load the settings

Labels: `backend`. Milestone: `v1`. Blocked by: #14.

> 🤖 Written by AI --- read/modified by izkreny! 🤓

### Overview

Read the token and the settings with the precedence of ADR 0004: built-in defaults, then the settings file, then environment variables.

### Acceptance criteria

- [ ] `region`, `poll_interval` and `debug_log` take their defaults (`eu`, `2s`, `false`), are overridden by `$XDG_CONFIG_HOME/intercomifico/settings.toml`, and those by `INTERCOMIFICO_REGION`, `INTERCOMIFICO_POLL_INTERVAL` and `INTERCOMIFICO_DEBUG_LOG`
- [ ] With `XDG_CONFIG_HOME` unset, the file is read from `~/.config/intercomifico/`
- [ ] A missing settings file is not an error; a malformed one stops startup with a message naming the file and the problem
- [ ] An invalid value stops startup with a message naming the setting and where it came from
- [ ] A key the settings do not define, a misspelt one or a token, stops startup with a message naming the key and its line
- [ ] A missing `INTERCOMIFICO_TOKEN` stops startup with a message saying how to set it
- [ ] Each setting's value and source are available for the debug log's start-up record

### Technical notes

- ADR 0004 for the settings and the token rule, ADR 0007 for where the file lives.
- The file is parsed with `github.com/pelletier/go-toml/v2`, its decoder set to `DisallowUnknownFields()`, per ADR 0004; `~/Projects/examples/go/go/pelletier_go-toml/` is the reference clone.

### Dependencies

Blocked by #14, which follows #1.

---

## 4. add the intercom read client

Labels: `backend`. Milestone: `v1`. Blocked by: #14.

> 🤖 Written by AI --- read/modified by izkreny! 🤓

### Overview

The read half of the hand-written Intercom client in `internal/intercom`: everything the app fetches.

### Acceptance criteria

- [ ] Every request carries `Authorization: Bearer`, `Intercom-Version: 2.16` and a `context.Context`, and goes to the base URL of the configured region (`eu`, `us`, `au`)
- [ ] `GET /me` returns the acting admin's id and `team_ids`
- [ ] The Mine and Unassigned queues of ADR 0002 are searched through `POST /conversations/search`, every page followed
- [ ] A conversation with its parts, a contact, the admins, the teams and the macros (every page, and changed since a time through `updated_since`) can be fetched
- [ ] A 429 returns a typed error carrying the reset time; any other error returns a typed error carrying Intercom's `request_id` and error codes
- [ ] `httptest` tests cover each call, pagination, the 429 and an error body
- [ ] `internal/intercom` imports nothing from `internal/ui` or `app`

### Technical notes

- ADR 0004 for the client, the endpoints and the errors; ADR 0002 for the queue filters; ADR 0006 for `httptest`.
- `~/Projects/examples/SDK/intercom/intercom_Intercom-OpenAPI` is the authority on the API; the SDKs beside it, and `google_go-github` and `stripe_stripe-go` under `~/Projects/examples/go/go/`, are models.

### Dependencies

Blocked by #14, which follows #1.

---

## 5. add the intercom write calls

Labels: `backend`. Milestone: `v1`. Blocked by: #4.

> 🤖 Written by AI --- read/modified by izkreny! 🤓

### Overview

The write half of the Intercom client: replies, notes and the conversation actions of ADR 0002.

### Acceptance criteria

- [ ] A public reply and an internal note can be posted to a conversation as the acting admin
- [ ] A conversation can be snoozed until a given time, assigned to an admin or a team, and closed
- [ ] Every write carries the acting admin's id
- [ ] A 429 or an error on a write returns the same typed errors as the read client, with no retry
- [ ] `httptest` tests check each request's method, path and body

### Technical notes

- ADR 0004 for `/reply` and `/parts`; ADR 0002 for the actions in scope.

### Dependencies

Blocked by #4.

---

## 6. build the app frame

Labels: `frontend`. Milestone: `v1`. Blocked by: #3, #4.

> 🤖 Written by AI --- read/modified by izkreny! 🤓

### Overview

The frame every screen sits in: `Run`, the public tab contract in `app`, the tab bar, global keys and the Catppuccin Mocha theme.

### Acceptance criteria

- [ ] `intercomifico.Run(extra ...app.Tab)` loads the settings, builds the client and the frame, and shows the inbox tab first and the extra tabs after it
- [ ] `app` holds the `Tab` interface, the open-contact message and the palette type with DaisyUI's twenty roles, and nothing else, per ADR 0012
- [ ] With one tab the tab bar is hidden; with more it shows, and the global keys switch tabs, quit and toggle help
- [ ] No global key is a plain letter; a key press reaches the active tab only, and every other message reaches every tab
- [ ] Every tab gets the palette in a message at start
- [ ] The Mocha palette maps onto the roles as ADR 0009's table says, and the frame paints every cell from the base roles
- [ ] A test checks every text-on-background pair ADR 0009 lists against 4.5:1 in Mocha, and has been seen failing on a deliberately wrong colour
- [ ] The help footer is drawn from the key map with bubbles `help`
- [ ] Golden files cover the frame with one and with two stub tabs at a fixed size
- [ ] The implementing agent proposes every key this issue adds (switching tabs, quitting, help), and the owner approves them before any is coded; each is checked against the textarea's own bindings so none is shadowed

### Technical notes

- ADRs 0008, 0009 and 0012. The inbox tab is a placeholder here; #7 fills it.

### Dependencies

Blocked by #3 and #4.

---

## 7. show the queues

Labels: `frontend`. Milestone: `v1`. Blocked by: #4, #6.

> 🤖 Written by AI --- read/modified by izkreny! 🤓

### Overview

The inbox tab's queue list, kept current by polling: the first build that shows real Intercom data.

### Acceptance criteria

- [ ] The inbox shows the Mine queue (open conversations first, then snoozed) and the Unassigned queue, and the user can switch between them
- [ ] A loading state shows until the first calls return
- [ ] Every `poll_interval` the visible queue's search runs again; `0` turns polling off
- [ ] Admins, teams and macros are fetched at start and refreshed every 15 minutes (macros through `updated_since`), with every macro fetched again each hour
- [ ] The refresh key runs a poll and the reference-data refresh at once and restarts both timers
- [ ] A 429 pauses polling until its reset time and says so on screen
- [ ] The layout follows the window size, and golden files cover both queues, the loading state and the 429 notice
- [ ] The README lists the scopes the private Intercom app needs, confirmed against a real workspace
- [ ] The implementing agent proposes every key this issue adds (switching queues, the refresh key), and the owner approves them before any is coded; each is checked against the textarea's own bindings so none is shadowed

### Technical notes

- ADR 0002 for the queues, ADR 0004 for polling, reference data and the refresh key, ADR 0008 for the model and focus, ADR 0006 for the tests.

### Dependencies

Blocked by #4 and #6.

---

## 8. show a conversation

Labels: `frontend`. Milestone: `v1`. Blocked by: #7.

> 🤖 Written by AI --- read/modified by izkreny! 🤓

### Overview

Opening a conversation from a queue shows its parts and the customer panel.

### Acceptance criteria

- [ ] Selecting a conversation shows its parts in order, each with its author and time, with internal notes set apart in the `warning` role
- [ ] A long conversation scrolls
- [ ] The customer panel shows the contact's details
- [ ] Each poll tick fetches the open conversation again, so a conversation closed, snoozed or reassigned elsewhere shows its new state
- [ ] Every tab receives the open-contact message when the open conversation changes, and a message with no contact when none is open
- [ ] The focus field routes keys between the queue list, the conversation and the customer panel
- [ ] Golden files cover a conversation with a note and the customer panel
- [ ] The implementing agent proposes every key this issue adds (moving focus between panes, scrolling), and the owner approves them before any is coded; each is checked against the textarea's own bindings so none is shadowed

### Technical notes

- ADR 0002, ADR 0004 for the per-tick fetch, ADR 0008, ADR 0012 for the open-contact message.

### Dependencies

Blocked by #7.

---

## 9. add the reply composer

Labels: `frontend`. Milestone: `v1`. Blocked by: #5, #8.

> 🤖 Written by AI --- read/modified by izkreny! 🤓

### Overview

Write and send a public reply or an internal note from a bubbles textarea, with a draft kept per conversation.

### Acceptance criteria

- [ ] `enter` inserts a newline; `ctrl+enter` sends, and `ctrl+s` sends too
- [ ] The composer shows whether it holds a reply or a note, and switching keeps the text
- [ ] Each conversation keeps its own draft in memory; opening another conversation shows that one's draft, and leaving the composer keeps it
- [ ] A send goes to the conversation the draft belongs to; a failed send keeps the draft and shows the error
- [ ] Tests cover sending, the reply/note switch and drafts across conversations through `Update`, plus golden files
- [ ] The implementing agent proposes every key this issue adds (the reply/note switch and every binding beyond `enter`, `ctrl+enter` and `ctrl+s`), and the owner approves them before any is coded; each is checked against the textarea's own bindings so none is shadowed

### Technical notes

- ADR 0005; ADR 0006 for the tests.

### Dependencies

Blocked by #5 and #8.

---

## 10. insert saved replies from a dialog

Labels: `frontend`. Milestone: `v1`. Blocked by: #9.

> 🤖 Written by AI --- read/modified by izkreny! 🤓

### Overview

The dialog mechanism of ADR 0013, and its first dialog: a picker that inserts a saved reply into the composer.

### Acceptance criteria

- [ ] A dialog draws over its owner's view with lipgloss layers: a tab's over that tab only, the frame's over the whole screen
- [ ] While a dialog is open every key goes to it, the frame's included, which receives each key before any tab does; `esc` closes it with no effect, and confirming closes it and hands the choice as a message to whichever opened it
- [ ] Only one dialog is open at a time
- [ ] The saved-reply picker lists, filterably, the macros visible to the acting admin
- [ ] Choosing one inserts it at the composer's cursor, with placeholders filled from the contact and the acting admin, and any other placeholder left visible
- [ ] A dialog about a conversation acts on the conversation it was opened for, even when polling changes the screen behind it
- [ ] Golden files cover the open picker
- [ ] The implementing agent proposes every key this issue adds (opening the picker), and the owner approves them before any is coded; each is checked against the textarea's own bindings so none is shadowed

### Technical notes

- ADR 0013 for dialogs, ADR 0005 for the picker and placeholders.

### Dependencies

Blocked by #9.

---

## 11. add conversation actions

Labels: `frontend`. Milestone: `v1`. Blocked by: #5, #10.

> 🤖 Written by AI --- read/modified by izkreny! 🤓

### Overview

Snooze, assign and close the open conversation, through pickers built on #10's dialogs.

### Acceptance criteria

- [ ] An assign picker lists me first, then teammates and teams, and assigns the conversation to the choice
- [ ] A snooze picker lists in 1 hour, in 4 hours, tomorrow at 9:00 and next Monday at 9:00, in the machine's local time, and snoozes the conversation until the choice
- [ ] A close action closes the conversation
- [ ] The queues show each change no later than the next poll
- [ ] A failed action, a 429 included, shows its error and changes nothing on screen
- [ ] Golden files cover both pickers
- [ ] The implementing agent proposes every key this issue adds (assign, snooze, close), and the owner approves them before any is coded; each is checked against the textarea's own bindings so none is shadowed

### Technical notes

- ADR 0002 for the actions, ADR 0013 for the pickers, ADR 0004 for errors.

### Dependencies

Blocked by #5 and #10.

---

## 12. write the debug log

Labels: `fullstack`. Milestone: `v1`. Blocked by: #3, #10.

> 🤖 Written by AI --- read/modified by izkreny! 🤓

### Overview

The opt-in, metadata-only debug log of ADR 0011, and the clean-up offer when the logs grow past 1 GB.

### Acceptance criteria

- [ ] With `debug_log` off no log file is created; with it on, `log/slog` JSON lines go to one file per day in the `logs` folder under `$XDG_STATE_HOME/intercomifico/`, appended by every run that day
- [ ] It records what ADR 0011 lists: the start-up record, each API request, each poll tick and reference-data refresh, each user action, and each error shown
- [ ] A test runs requests carrying known message text, names, emails, conversation, contact, admin, team and conversation-part ids, and a token through the log, and finds none of them in its output, while an error's `request_id` is there
- [ ] At start and whenever the date changes, a `logs` folder over 1 GB opens a frame dialog offering to delete every log file but today's and yesterday's, or to keep them all

### Technical notes

- ADR 0011; ADR 0007 for what may be written; ADR 0013 for the frame's dialog.

### Dependencies

Blocked by #3 and #10.

---

## 13. add a contributing guide

Labels: `docs`. Milestone: none. Blocked by: none.

> 🤖 Written by AI --- read/modified by izkreny! 🤓

### Overview

A CONTRIBUTING.md for outside contributors, which GitHub links from the new-issue and new-pull-request pages, stating how AI-written contributions are handled.

### Acceptance criteria

- [ ] CONTRIBUTING.md exists at the repository root
- [ ] It states that an agent publishes nothing on its own initiative: no pushed commit, issue, pull request or comment unless the person it works for asked for that act, and that person answers for it
- [ ] It states that everything an agent writes for the repository discloses it: a commit body, issue, pull request or comment opens with a blockquote line starting `> 🤖 Written by AI`, completed with the contributor's own name
- [ ] GitHub's new-issue page links to it

### Technical notes

- Modelled on the AI contribution rules in `~/Projects/examples/go/go/caddyserver_caddy/AGENTS.md` and `~/Projects/examples/go/go/traefik_traefik/AGENTS.md`.
- The `> 🤖` prefix is the one gh-solo's gates test, so a contributor's line and the owner's agents' line read the same way.

### Dependencies

None.

---

## 14. add the go style skill

Labels: `infra`. Milestone: `v1`. Blocked by: #1.

> 🤖 Written by AI --- read/modified by izkreny! 🤓

### Overview

A `go-style` agent skill that teaches the Go practice golangci-lint cannot check, so that every Go issue after the toolchain is written to it. The skill lives outside this repository; this issue's change here is the line in AGENTS.md that tells agents to load it.

### Acceptance criteria

- [ ] A `go-style` skill exists, distilled from Effective Go, Go's Code Review Comments, Google's Go Style Guide and Uber's Go guide
- [ ] It covers what the linters of ADR 0010 do not enforce: package and interface design, error design, concurrency and `context`, and the quality of names and comments
- [ ] It does not restate a rule golangci-lint already enforces
- [ ] Where it differs from an ADR or from the linter configuration, the ADR and the configuration win, and the skill says so
- [ ] AGENTS.md tells agents to load `go-style` before writing Go, with crush's Bubble Tea rules on top for UI work

### Technical notes

- ADR 0014 records this decision.
- `~/Projects/examples/go/go/uber-go_guide` is the local copy of Uber's guide; the other sources are on go.dev and google.github.io.
- Bubble Tea's own rules stay out of this skill: they are `~/Projects/examples/go/charm/charmbracelet_crush/internal/ui/AGENTS.md`, which AGENTS.md already points at.

### Dependencies

Blocked by #1, so the skill is checked against the real linter configuration. #3 and #4 wait on it, and every later Go issue waits through them.
