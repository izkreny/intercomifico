> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 9. Theming

Date: 2026-10-10

## Status

Accepted

## Context

The Charm libraries provide colours and styles, not themes. lipgloss takes any colour, true colour included, and `lipgloss.LightDark` picks one of two by background; bubbletea v2 asks the terminal for its background colour with `tea.RequestBackgroundColor` and answers with a `tea.BackgroundColorMsg` whose `IsDark()` says which kind it is; and bubbletea detects the terminal's colour profile with `colorprofile.Detect`, downsampling true colour to 256 or 16 colours where the terminal supports no more. Every major terminal emulator supports true colour.

Themes on top of that are app code, and the reference apps solve it the same way: one struct of colours named by role, which the rest of the code reads. crush keeps a `Palette` struct and loads user themes from JSON files, and gh-dash puts a colour theme in its YAML config. huh ships Catppuccin through `github.com/catppuccin/go`, choosing the light Latte flavour or the dark Mocha by `isDark`; superfile and argonaut ship Catppuccin too.

Two vocabularies are candidates for the role names:

- **DaisyUI's semantic colours**, twenty names in pairs of a colour and the text drawn on it: `primary`, `secondary`, `accent`, `neutral`, `info`, `success`, `warning` and `error`, each with a `-content` partner, plus the surfaces `base-100`, `base-200` and `base-300` and their text, `base-content`.
- **base16**: sixteen slots, `base00` to `base0F`, numbered rather than named, so code reading `base0D` says nothing about what it paints.

## Decision

- **Colours are named by DaisyUI's twenty semantic names, as they are**, in one palette. Its type is public in `app`, per ADR 0012, so every tab paints with the same roles; the Catppuccin values and the background detection live in `internal/ui`. No other code names a colour; a screen picks a role, and an internal note, for instance, uses `warning`, the yellow Intercom gives notes.
- **The built-in theme is Catppuccin, in true colour**: Latte on a light background and Mocha on a dark one. Its colours come from `github.com/catppuccin/go`, as in huh, rather than hex values copied into the code.
- **The roles map onto Catppuccin's names**, which are the same in both flavours:

  | Role                                                  | Catppuccin                   |
  |-------------------------------------------------------|------------------------------|
  | `primary`                                             | Mauve                        |
  | `secondary`                                           | Pink                         |
  | `accent`                                              | Teal                         |
  | `neutral`                                             | Surface 1                    |
  | `neutral-content`                                     | Text                         |
  | `base-100`, `base-200`, `base-300`                    | Base, Mantle, Crust          |
  | `base-content`                                        | Text                         |
  | `info`, `success`, `warning`, `error`                 | Blue, Green, Yellow, Red     |
  | every other `-content`                                | Base                         |

- **The flavour follows the terminal's background**: the app sends `tea.RequestBackgroundColor` at start and picks Latte or Mocha from `IsDark()`. Until the answer arrives, or when the terminal never answers, it uses Mocha.
- **Downsampling is bubbletea's**: a terminal without true colour gets the nearest colours its profile has, with no code of the app's own.
- **Nerd Font icons come after v1**, through their own ADR. v1 draws with plain Unicode text.

## Consequences

- The app looks the same in every true-colour terminal, whatever theme the terminal wears, and matches it exactly in a terminal running Catppuccin Latte or Mocha.
- `github.com/catppuccin/go` is a dependency, the one huh uses for the same job.
- DaisyUI dims text with opacity, which a terminal lacks; lipgloss's `Faint` attribute does that job, for text such as timestamps.
- A user theme is a later addition: a TOML file next to the settings file, per ADR 0007, setting the twenty roles to hex colours. It gets its own ADR.
