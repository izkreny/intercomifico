> 🤖 Written by AI --- read/modified by izkreny! 🤓

# 9. Theming

Date: 2026-10-10

## Status

Accepted

## Context

The Charm libraries provide colours and styles, not themes. lipgloss offers the terminal's 16 ANSI colours as constants (`lipgloss.Red`, `lipgloss.BrightBlue` and the rest), true-colour values, and `lipgloss.LightDark` to choose between two colours; bubbletea v2 can ask the terminal for its background colour. Themes on top of that are app code: crush keeps every colour in one `Palette` struct and loads user themes from JSON files that inherit from a base, and gh-dash puts a colour theme in its YAML config.

Two cross-application standards exist:

- **The terminal's 16 ANSI colours** plus its default foreground and background. Every terminal theme defines them. Omarchy, for one, generates its terminal configuration from a theme's colors.toml file (`color0` to `color15`, accent, foreground, background), so an app painting only with ANSI colours follows whatever theme the terminal wears.
- **base16** (the styling spec in `github.com/tinted-theming/home`): sixteen named slots, `base00` to `base0F`, each with a role, mapped onto the ANSI colours.

## Decision

- **The built-in theme uses only the terminal's 16 ANSI colours and its default foreground and background.** No true-colour or 256-colour value appears in it.
- **Colours are named by role, in one palette** in the styles file ADR 0008 places in `internal/ui`: roles such as accent, muted text, error, public reply and internal note, each mapped onto an ANSI colour. No other code names a colour.
- **Light and dark come from the terminal's own theme**, so v1 does no background detection and ships one theme.

## Consequences

- The app matches Ghostty's theme, Omarchy's, or any other the user switches to, with no configuration.
- How good it looks depends on the terminal theme's ANSI colours; a theme whose yellow is unreadable on its background makes internal notes hard to read.
- A user theme is a later addition: a TOML file next to the settings file, per ADR 0007, overriding roles with hex colours, possibly through base16 slot names. It gets its own ADR.
