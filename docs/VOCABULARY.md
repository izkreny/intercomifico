> 🤖 Written by AI --- read/modified by izkreny! 🤓

# Vocabulary

The project's terms, one name per concept. Code identifiers, docs, issues, commits and UI text use these names, and never call a concept by a name its entry lists under **Never**; such a name may still be another entry's term or a fixed identifier, as `app` is a package and `comment` an Intercom value. A branch that needs a term this file lacks adds it here in the same pull request, per ADR 0015.

A term Intercom's API defines takes Intercom's name, and its source names the schema in Intercom's OpenAPI description 2.16 (`github.com/intercom/Intercom-OpenAPI`, `descriptions/2.16`). General English, Go's own terms and the Charm libraries' names are not listed; their own docs define them.

## acceptance criteria

The checklist on an issue that says when it is done, ticked as each one verifiably lands.

- **Source:** the open issues.
- **Never:** requirements, definition of done.

## acting admin

The admin who owns the token, resolved once at start with `GET /me`; every reply, note and conversation action carries their id.

- **Source:** ADR 0004.
- **Never:** current user, logged-in admin.

## active tab

The tab in front, the one key presses reach.

- **Source:** ADR 0012; #9.
- **Never:** current tab, selected tab.

## admin

An account with access to the workspace, which a conversation can be assigned to.

- **Source:** Intercom, the `admin` schema.
- **Never:** teammate, agent, operator.

## ADR

An architecture decision record in `docs/adr/`, in the format ADR 0001 sets, superseded rather than edited once accepted.

- **Source:** ADR 0001.
- **Never:** design doc, RFC.

## agent

An AI coding agent working the repository under AGENTS.md.

- **Source:** AGENTS.md; ADR 0014.
- **Never:** bot, assistant.

## assign

The conversation action that gives a conversation to an admin or a team, chosen in the assign picker.

- **Source:** Intercom, `message_type` `assignment` on `POST /conversations/{id}/parts`; ADR 0002.
- **Never:** transfer, route.

## close

The conversation action that moves a conversation to the `closed` state.

- **Source:** Intercom, `message_type` `close` on `POST /conversations/{id}/parts`; ADR 0002.
- **Never:** resolve, archive.

## composer

The pane holding the bubbles textarea where a draft is written and sent as a reply or a note.

- **Source:** ADR 0005.
- **Never:** editor, input box, reply box.

## contact

A person the workspace talks to, a lead or a user in Intercom's terms, whose details the customer panel shows.

- **Source:** Intercom, the `contact` schema.
- **Never:** customer, client, end user.

## conversation

A thread between a contact and the workspace, made of conversation parts and always in one state.

- **Source:** Intercom, the `conversation` schema.
- **Never:** ticket, chat, case.

## conversation action

Snooze, assign or close, applied to one conversation.

- **Source:** ADR 0002; ADR 0004.
- **Never:** command, operation.

## conversation part

One item in a conversation's history, such as a reply, a note or an assignment; "part" for short.

- **Source:** Intercom, the `conversation_part` schema.
- **Never:** message, entry.

## custom build

A repository of its own whose `main` calls `intercomifico.Run` with its extra tabs, taking this one's releases with `go get`.

- **Source:** ADR 0012.
- **Never:** fork, company build, outside build.

## customer panel

The pane showing the open contact's name, email and custom attributes; the one place the word "customer" names something.

- **Source:** ADR 0002; ADR 0008.
- **Never:** contact panel, sidebar, profile card.

## debug log

The opt-in, metadata-only `log/slog` log, one file per day in the `logs` folder under `$XDG_STATE_HOME/intercomifico/`.

- **Source:** ADR 0011.
- **Never:** trace, audit log.

## dialog

A box drawn over its owner's view with lipgloss layers, taking every key until it closes; a tab's or the frame's.

- **Source:** ADR 0013.
- **Never:** modal, popup, overlay.

## draft

The composer's text for one conversation, kept in memory until it is sent. A GitHub draft pull request or an issue labelled `draft` is always named in full.

- **Source:** ADR 0005.
- **Never:** unsent reply.

## extra tab

A tab passed to `Run` after the inbox, built outside this repository.

- **Source:** ADR 0012.
- **Never:** outside tab, plugin.

## focus

The one field of the inbox naming the pane, or the open dialog, that receives key presses.

- **Source:** ADR 0008; ADR 0013.
- **Never:** active pane, selected pane.

## frame

The app's only top-level Bubble Tea model: it holds the tabs, draws the tab bar, owns the global keys, the window size and the theme, and can own a dialog.

- **Source:** ADR 0012.
- **Never:** shell, window, root view.

## gate

A check with an exit code a branch must pass before it is ready: the formatting check, `task lint`, `task test` and `task audit`, all run by `task ci`.

- **Source:** ADR 0010.
- **Never:** hook, guard.

## global key

A key binding the frame owns, for switching tabs, quitting and help, and never a plain letter.

- **Source:** ADR 0012.
- **Never:** hotkey, shortcut.

## golden file

A `testdata/<TestName>.golden` file holding a view's rendered output, rewritten only by `task test:update`.

- **Source:** ADR 0006.
- **Never:** snapshot, fixture.

## help footer

The line of key bindings drawn from the key map with bubbles `help`.

- **Source:** ADR 0008.
- **Never:** status bar, hint bar.

## inbox

The app's v1 tab, one model holding the queue list, the conversation, the composer and the customer panel. A team's inbox in Intercom is always named "a team's inbox".

- **Source:** ADR 0008; ADR 0012.
- **Never:** main view, home.

## issue

A GitHub issue in this repository's tracker.

- **Source:** AGENTS.md; ADR 0001.
- **Never:** ticket, task, story.

## key binding

One `key.Binding`: the keys that trigger an action and its help text.

- **Source:** ADR 0008.
- **Never:** shortcut, hotkey.

## key map

The set of key bindings a model routes key presses by.

- **Source:** ADR 0008.
- **Never:** keymap, bindings table.

## keyring

The system's secret store, the freedesktop Secret Service on Linux, holding the token in the entry whose `service` is `intercomifico` and whose `username` is `token`.

- **Source:** ADR 0004.
- **Never:** keychain, wallet, password store.

## loading state

What the inbox shows from start until the first calls return.

- **Source:** ADR 0007.
- **Never:** splash screen.

## macro

A pre-defined response the composer inserts at its cursor, listed by the macro picker.

- **Source:** Intercom, the `macro` schema.
- **Never:** saved reply, canned response, template, snippet.

## Mine

The queue of conversations assigned to the acting admin, open ones first and snoozed ones after them.

- **Source:** ADR 0002.
- **Never:** My conversations, assigned to me.

## note

An internal conversation part that only admins see, drawn in the `warning` role.

- **Source:** Intercom, `message_type` `note` on `POST /conversations/{id}/reply`; ADR 0005.
- **Never:** comment, internal comment, private reply.

## open contact

The contact of the open conversation, sent to every tab in the open-contact message whenever it changes.

- **Source:** ADR 0012.
- **Never:** current customer, selected contact.

## open conversation

The conversation the inbox shows, refetched on every poll tick.

- **Source:** ADR 0004; ADR 0012.
- **Never:** current conversation, selected conversation, active conversation.

## owner

The person who owns the repository and is its only committer, whose review approves every pull request and every new term.

- **Source:** ADR 0015; the open issues.
- **Never:** maintainer, lead.

## palette

The twenty roles with their colours, a public type in `app` sent to every tab at start.

- **Source:** ADR 0009; ADR 0012.
- **Never:** colour scheme.

## pane

One region of the inbox: the queue list, the conversation, the composer or the customer panel, each a plain struct with no `Update` of its own.

- **Source:** ADR 0008.
- **Never:** panel, widget, component, sub-model.

## picker

A dialog that is a filterable bubbles `list`: macro, assign, snooze or log cleanup.

- **Source:** ADR 0013.
- **Never:** selector, chooser, menu.

## placeholder

A field in a macro's text that the app fills from the contact or the acting admin, or leaves visible for the user to edit.

- **Source:** ADR 0005.
- **Never:** variable, template tag.

## plan

A branch's plan file in `docs/plans/`, named by its date and issue: how the issue will be done, while the issue says what.

- **Source:** the plan files in `docs/plans/`, such as `docs/plans/2026-10-11_GHI-4_add-go-style-skill.md`.
- **Never:** spec, design doc.

## poll tick

One run of polling, every `poll_interval`: the visible queue's search again and the open conversation refetched.

- **Source:** ADR 0004.
- **Never:** refresh, sync.

## queue

A conversation search shown as a list: Mine or Unassigned.

- **Source:** ADR 0002.
- **Never:** inbox, view, folder, filter.

## queue list

The pane listing the visible queue's conversations.

- **Source:** ADR 0008.
- **Never:** conversation list, list view.

## reference data

The admins, teams and macros, fetched at start and refreshed on a slow timer.

- **Source:** ADR 0004.
- **Never:** lookup data, cache.

## reference repository

A local clone of an upstream repository, listed under *Reference code* in AGENTS.md, read for how things are built and never depended on.

- **Source:** ADR 0006; ADR 0010; ADR 0014.
- **Never:** reference clone, reference app, example repo.

## refresh key

The key that runs a poll tick and the reference-data refresh at once and restarts both timers.

- **Source:** ADR 0004.
- **Never:** reload key.

## region

The `region` setting, `eu`, `us` or `au`, choosing Intercom's host.

- **Source:** ADR 0004.
- **Never:** data centre, zone.

## reply

A public conversation part addressed to the contact, sent from the composer.

- **Source:** Intercom, `message_type` `comment` on `POST /conversations/{id}/reply`; ADR 0005.
- **Never:** comment, response, answer.

## role

One of the twenty semantic names DaisyUI gives its colours, such as `primary` or `base-100`, which the app paints with. A contact's `role` field in Intercom is always written as the field.

- **Source:** ADR 0009.
- **Never:** colour name, slot, token.

## setting

A value with a built-in default, overridden by the settings file and that by an environment variable: `region`, `poll_interval`, `debug_log`.

- **Source:** ADR 0004.
- **Never:** option, config, preference.

## settings file

`$XDG_CONFIG_HOME/intercomifico/settings.toml`, written by the user and only ever read by the app.

- **Source:** ADR 0004; ADR 0007.
- **Never:** config file.

## snooze

The conversation action that moves a conversation to the `snoozed` state until a time chosen in the snooze picker.

- **Source:** Intercom, `message_type` `snoozed` on `POST /conversations/{id}/parts`; ADR 0002.
- **Never:** defer, postpone.

## spike

An issue that answers a question rather than shipping a feature.

- **Source:** the open issues.
- **Never:** research task, investigation.

## state

A conversation's `state`: `open`, `snoozed` or `closed`.

- **Source:** Intercom, the `conversation` schema.
- **Never:** status.

## tab

What the frame holds and switches between, implementing `app.Tab`: the inbox, or an extra tab.

- **Source:** ADR 0012.
- **Never:** page, screen, view, plugin.

## tab bar

The row of tab titles the frame draws, hidden while there is one tab.

- **Source:** ADR 0012.
- **Never:** tab strip, header.

## task

A command in Taskfile.yaml, run as `task <name>`.

- **Source:** ADR 0010.
- **Never:** script, target; and never an issue.

## team

A set of admins, which a conversation can be assigned to.

- **Source:** Intercom, the `team` schema.
- **Never:** group, squad.

## theme

The colours given to the twenty roles; v1 ships Catppuccin Mocha alone.

- **Source:** ADR 0009.
- **Never:** skin, colour scheme.

## ticket

Intercom's ticket object, out of v1; never a name for a conversation or an issue.

- **Source:** Intercom, the `ticket` schema; ADR 0002.

## token

The access token of the private Intercom app, read from `INTERCOMIFICO_TOKEN` or the keyring and never written anywhere.

- **Source:** ADR 0004.
- **Never:** API key, password, secret.

## UI state

What the app may write about itself, such as the last visible queue, under `$XDG_STATE_HOME/intercomifico/`; v1 keeps none.

- **Source:** ADR 0007.
- **Never:** session, cache.

## Unassigned

The queue of open conversations assigned to no admin, and either to no team or to one of the acting admin's teams.

- **Source:** ADR 0002.
- **Never:** unclaimed, triage.

## user

The person running the app, who acts in Intercom as the acting admin; never a contact, though Intercom gives some contacts the `role` `user`.

- **Source:** ADR 0002; ADR 0004; ADR 0005.
- **Never:** operator, agent.

## visible queue

The queue the inbox shows, whose search each poll tick runs again.

- **Source:** ADR 0004; #10.
- **Never:** open queue, current queue, selected queue.

## workspace

The Intercom account whose conversations the app shows, one per run; the API's `app` object, which this repository never calls an app.

- **Source:** Intercom, the `app` schema; ADR 0002.
- **Never:** app, account, instance.
