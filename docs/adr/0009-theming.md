> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 9. Theming

Date: 2026-10-10

## Status

Accepted

## Context

The Charm libraries provide colours and styles, not themes. lipgloss takes any colour, true colour included, and bubbletea detects the terminal's colour profile with `colorprofile.Detect`, downsampling true colour to 256 or 16 colours where the terminal supports no more. Every major terminal emulator supports true colour. bubbletea v2 can also ask the terminal for its background colour, with `tea.RequestBackgroundColor`, and answers with a `tea.BackgroundColorMsg` whose `IsDark()` says whether it is light or dark.

Themes on top of that are app code, and the reference apps solve it the same way: one struct of colours named by role, which the rest of the code reads. crush keeps a `Palette` struct and loads user themes from JSON files, and gh-dash puts a colour theme in its YAML config. huh ships Catppuccin through `github.com/catppuccin/go`, which carries Catppuccin's named colours as Go values; superfile and argonaut ship Catppuccin too.

Two vocabularies are candidates for the role names:

- **DaisyUI's semantic colours**, twenty names in pairs of a colour and the text drawn on it: `primary`, `secondary`, `accent`, `neutral`, `info`, `success`, `warning` and `error`, each with a `-content` partner, plus the surfaces `base-100`, `base-200` and `base-300` and their text, `base-content`.
- **base16**: sixteen slots, `base00` to `base0F`, numbered rather than named, so code reading `base0D` says nothing about what it paints.

WCAG 2.2 Level AA asks for a contrast ratio of at least 4.5:1 between normal text and its background. Catppuccin Mocha meets it, at 7:1 or better, for every pair of text and background the Decision below lists. Its light flavour, Latte, does not: several of its accents fall to about 2:1 against its own background. Of the light themes Ghostty bundles, few pass unmodified; Modus Operandi and GitHub Light Default are among them.

## Decision

- **Colours are named by DaisyUI's twenty semantic names, as they are**, in one palette. Its type is public in `app`, per ADR 0012, so every tab paints with the same roles; the theme's values live in `internal/ui`. No other code names a colour; a screen picks a role, and an internal note, for instance, uses `warning`, the yellow Intercom gives notes.
- **v1 ships one theme, Catppuccin Mocha, in true colour.** Its colours come from `github.com/catppuccin/go`, as in huh, rather than hex values copied into the code.
- **The roles map onto Catppuccin's names:**

  | Role                                                  | Catppuccin                   |
  |-------------------------------------------------------|------------------------------|
  | `primary`                                             | Mauve                        |
  | `secondary`                                           | Pink                         |
  | `accent`                                              | Teal                         |
  | `neutral`                                             | Surface 0                    |
  | `neutral-content`                                     | Text                         |
  | `base-100`, `base-200`, `base-300`                    | Base, Mantle, Crust          |
  | `base-content`                                        | Text                         |
  | `info`, `success`, `warning`, `error`                 | Blue, Green, Yellow, Red     |
  | every other `-content`                                | Base                         |

- **The app paints every cell it draws**, background included, from the base roles, and never shows the terminal's own background through. Mocha's pale text is unreadable on a light background, so on a light terminal the app is a dark window instead.
- **A built-in theme meets WCAG 2.2 Level AA**, 4.5:1, for every pair of text and the background it is drawn on: `base-content` on each of `base-100`, `base-200` and `base-300`; each of `primary`, `secondary`, `accent`, `info`, `success`, `warning` and `error` as text on each of those three; each `-content` on its own role; and `neutral-content` on `neutral`, which is a surface and never text.
- **Downsampling is bubbletea's**: a terminal without true colour gets the nearest colours its profile has, with no code of the app's own.
- **Nerd Font icons come after v1**, through their own ADR. v1 draws with plain Unicode text.

## Consequences

- The app looks the same in every true-colour terminal, whatever theme the terminal wears, and matches one running Catppuccin Mocha exactly.
- A test checks every pair the Decision lists against 4.5:1 in Mocha, so a palette change that breaks AA fails the build.
- `github.com/catppuccin/go` is a dependency, the one huh uses for the same job.
- DaisyUI dims text with opacity, which a terminal lacks; lipgloss's `Faint` attribute does that job, for text such as timestamps.
- v1 asks the terminal nothing about its background, since one theme leaves nothing to choose.
- Every other theme, light ones included, comes as a custom theme, and custom themes arrive after v1 with their own ADR. What that ADR starts from:
  - a theme is a TOML file next to the settings file, per ADR 0007, setting the twenty roles to hex colours;
  - the settings may name a theme for light backgrounds and one for dark backgrounds, and the app picks between them with `tea.RequestBackgroundColor` and `IsDark()`;
  - when no light theme is named, the app uses the default theme on every background;
  - the WCAG rule above binds built-in themes; a custom theme is the user's choice.
