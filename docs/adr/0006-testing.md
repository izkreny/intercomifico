> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 6. Testing

Date: 2026-10-09

## Status

Accepted

## Context

The first blueprint set 100% coverage of `Update` as its target. A scan of the test files and CI configuration of the 72 reference repositories found:

- Coverage gates are rare: three repositories fail CI on a coverage percentage, `charmbracelet/soft-serve` the only one of Charm's, and none aims at 100%.
- Golden files are Charm's own practice for rendered output, in `charmbracelet/x`, bubbles, lipgloss, glamour, crush, fang and bubbletea, through `github.com/charmbracelet/x/exp/golden`: `golden.RequireEqual` compares output with `testdata/<TestName>.golden`, and `go test ./... -update` rewrites the files.
- `teatest`, which drives a running program over time, is rare and partly abandoned: bubbletea's own example and gh-dash's tests have it commented out.
- Putting the API client behind an interface, so the UI is tested against a fake, appears only in kl; gh-dash's package-global client is the counter-example.

In a terminal app the UI tests are cheap: `Update` is a function call and `View()` returns a `tea.View` whose `Content` is plain text, both inside the `go test` process.

## Decision

- **No coverage gate.** Coverage is reported, never enforced. Every behaviour change carries a test, and the aim is as many unit tests as are practical.
- **`Update` is tested directly**: a test builds the model, sends it a message (a key press, an API result, an error) and checks the new state and the returned command.
- **`View()` is tested with golden files**: rendered at a fixed terminal size, the `Content` of the returned `tea.View` is compared with `golden.RequireEqual`. A deliberate change to the output is accepted by running `go test ./internal/ui/... -update`, scoped to the package that imports `golden`, since only its test binaries define the `-update` flag, and reviewing the changed `.golden` files in the diff.
- **The UI depends on an interface for the Intercom calls it makes**, declared where it is used, and its tests pass a hand-written fake returning prepared data or errors. UI tests never touch HTTP.
- **`internal/intercom` is tested against `net/http/httptest`**: a local server inside the test answers with JSON taken from the OpenAPI examples, and the test checks the request the client sent, including the `Intercom-Version` and `Authorization` headers. No test reaches the network.
- **No `teatest`.**
- **The checks are `go vet ./...` and `go test -race ./...`**, and both must pass before a branch is ready.

## Consequences

- The test suite runs in seconds and needs no Intercom account, token or network.
- Golden files make every rendering change visible in review, at the cost of updating them on purpose whenever the layout changes.
- Nothing exercises the real Intercom API automatically. Recording real responses for replay, as crush does, is possible later and needs its own ADR, since it involves a real token and real customer data.
- A linter beyond `go vet` is a separate decision, made together with CI in its own `infra` issue.
