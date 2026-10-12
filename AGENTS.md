# Intercomifico

A terminal client for Intercom customer conversations, written in Go on the Charm v2 stack (bubbletea, bubbles, lipgloss).

## Decisions

Every architecture decision is an ADR in `docs/adr/`, indexed in `docs/adr/README.md`. Read the ADRs a task touches before writing code for it. Work that needs a decision no ADR holds gets that ADR first.

## Commands

The toolchain is pinned in mise.toml and every command is a task in Taskfile.yaml:

- `mise install` installs the toolchain.
- `task ci` runs every gate in order: the formatting check, lint, tests and the dependency audit. Run it before every push.
- `task test:update` regenerates golden files after a deliberate rendering change.

## Rules

- **Load the `go-style` skill before writing Go.** It holds the judgement the linters cannot check, and the ADRs and the linter configuration win wherever it differs from them. For UI work, crush's Bubble Tea rules, named under *Reference code*, apply on top of it.
- **Use only the terms in `docs/VOCABULARY.md`** in code identifiers, docs, issues, commits and UI text, and never call a concept by a name its entry lists under **Never**. A branch that needs a term the file lacks adds it there in the same PR rather than coining one in passing, per ADR 0015.
- **Never edit a golden file by hand.** Regenerate it with `task test:update` and review the changed `.golden` files in the diff.
- **Trust the ADRs and the reference repositories over memory.** bubbletea v2, at `charm.land/bubbletea/v2` with bubbles and lipgloss beside it, is newer than most training data, and many clones below are still on v1: check a clone's go.mod before copying from it.
- **Search the reference repositories, never read them whole.** Some are hundreds of megabytes; grep for the symbol or pattern and read the files it finds.

## Reference code

Local clones of upstream repositories live under `~/Projects/examples/`. Each folder is named `<owner>_<repo>` after its GitHub repository, so `charmbracelet_bubbletea` is `github.com/charmbracelet/bubbletea`.

- **Canonical: `~/Projects/examples/go/charm/charmbracelet_*/`.** Charm's own repositories, and the authority on how the Charm libraries are meant to be used. Check them before writing UI code, and prefer them over community apps and over memory.
  - The libraries' own 'examples/' directories and `charmbracelet_crush` (a large real v2 app) are the first places to look.
  - Before any UI work, read `~/Projects/examples/go/charm/charmbracelet_crush/internal/ui/AGENTS.md`. IMPORTANT NOTE: we do not use its Ultraviolet sections.
- **Community apps: every other folder in `~/Projects/examples/go/charm/`.** Real bubbletea apps, for ideas only; `dlvhdr_gh-dash` is the closest to this app's shape. The folder is not split by Charm version, so v1 and v2 apps sit side by side; `grep -l 'charm.land/bubbletea/v2' ~/Projects/examples/go/charm/*/go.mod` lists the apps on v2.
- **Go: `~/Projects/examples/go/go/`.** Widely used, actively maintained Go projects outside the Charm stack, for how idiomatic Go is written; check them for code that is not UI, and prefer them over memory. Look first, by need:
  - Language and style: `golang_go` (the standard library), `golang_tools`, `golang_example`, `uber-go_guide`.
  - HTTP API clients, models for `internal/intercom`: `google_go-github`, `stripe_stripe-go`, `cli_cli`.
  - Extension through compiled-in modules, the model this app adopts: `caddyserver_caddy`, `coredns_coredns`; other extension models: `hashicorp_go-plugin`, `traefik_traefik`.
  - Terminal apps on other UI libraries, for ideas only: `jesseduffield_lazygit`, `derailed_k9s`, `junegunn_fzf`.
  - Libraries and tools this repository uses or may: `pelletier_go-toml`, `google_go-cmp`, `spf13_cobra`, `dominikh_go-tools`, `golangci_golangci-lint`, `go-task_task`, `goreleaser_goreleaser`.
  - The rest are well-built applications and servers to compare against.
- **Intercom: `~/Projects/examples/SDK/intercom/`.** `intercom_Intercom-OpenAPI` is Intercom's OpenAPI description and the authority on the REST API; the official developer docs at https://developers.intercom.com/ cover what it leaves out, such as authentication, rate limits and webhooks; the Help Center at https://www.intercom.com/help/en/ is the authority on what the product calls things and how it behaves. `docs/VOCABULARY.md` records what each of the three calls every Intercom term. The official SDKs beside it are models for `internal/intercom`, never dependencies.
