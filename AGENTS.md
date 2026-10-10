# Intercomifico

A terminal client for Intercom customer conversations, written in Go on the Charm stack (bubbletea, bubbles, lipgloss).

## Decisions

Every architecture decision is an ADR in `docs/adr/`, indexed in `docs/adr/README.md`. Read the ADRs a task touches before writing code for it. Work that needs a decision no ADR holds gets that ADR first, per ADR 0001.

## Reference code

Local clones of upstream repositories live under `~/Projects/examples/`. Each folder is named `<owner>_<repo>` after its GitHub repository, so `charmbracelet_bubbletea` is `github.com/charmbracelet/bubbletea`.

- **Canonical: `~/Projects/examples/go/charm/charmbracelet_*/`.** Charm's own repositories, and the authority on how the Charm libraries are meant to be used. Check them before writing UI code, and prefer them over community apps and over memory. The libraries' own `examples/` directories and `charmbracelet_crush` (a large real v2 app) are the first places to look.
- **Community apps: every other folder in `~/Projects/examples/go/charm/`.** Real bubbletea apps, for ideas only; `dlvhdr_gh-dash` is the closest to this app's shape. Many are still on the v1 API, which ADR 0003 rules out copying.
- **Go: `~/Projects/examples/go/go/`.** Widely used, actively maintained Go projects outside the Charm stack, for how idiomatic Go is written; check them for code that is not UI, and prefer them over memory. Look first, by need:
  - Language and style: `golang_go` (the standard library), `golang_tools`, `golang_example`, `uber-go_guide`.
  - HTTP API clients, models for `internal/intercom`: `google_go-github`, `stripe_stripe-go`, `cli_cli`.
  - Extension through compiled-in modules, as ADR 0012 adopts: `caddyserver_caddy`, `coredns_coredns`; other extension models: `hashicorp_go-plugin`, `traefik_traefik`.
  - Terminal apps on other UI libraries, for ideas only: `jesseduffield_lazygit`, `derailed_k9s`, `junegunn_fzf`.
  - Libraries and tools this repository uses or may: `BurntSushi_toml`, `google_go-cmp`, `spf13_cobra`, `dominikh_go-tools`, `golangci_golangci-lint`, `go-task_task`, `goreleaser_goreleaser`.
  - The rest are well-built applications and servers to compare against.
- **Intercom: `~/Projects/examples/SDK/intercom/`.** `intercom_Intercom-OpenAPI` is Intercom's OpenAPI description and the authority on the REST API; the official developer docs at https://developers.intercom.com/ cover what it leaves out, such as authentication, rate limits and webhooks. The official SDKs beside it are models for `internal/intercom`, never dependencies.
