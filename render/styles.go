package render

import (
	"strings"

	"charm.land/lipgloss/v2"
)

// Styles holds the lipgloss styles for every chrome element a2tea renders.
// The zero value is not valid — always start from DefaultStyles and customize
// from there:
//
//	st := render.DefaultStyles()
//	st.Heading = st.Heading.Foreground(lipgloss.Color("99"))
//	surf := render.NewSurface(id, comps, render.WithStyles(st))
//
// DefaultStyles is deliberately monochrome: a2tea surfaces render inside a host
// TUI (crush) that owns the color theme, so the default chrome sticks to
// attributes that read correctly on any background — borders, bold, faint, and
// reverse-video for focus. A host that wants to inject its palette builds on
// DefaultStyles and overrides specific fields.
type Styles struct {
	// CardBorder draws the rounded box around a Card's child.
	CardBorder lipgloss.Style

	// Heading renders Text variants h1–h3.
	Heading lipgloss.Style
	// Subheading renders Text variants h4–h5.
	Subheading lipgloss.Style
	// Caption renders caption-variant Text and other secondary chrome
	// (field labels, media placeholders).
	Caption lipgloss.Style

	// Button renders an idle button's label.
	Button lipgloss.Style
	// ButtonFocused renders the focused button's label.
	ButtonFocused lipgloss.Style

	// MarkdownRenderer, when non-nil, renders the text of a body-variant
	// (or unstyled) Text component as Markdown before wrapping. The host
	// injects its configured renderer (e.g. glamour) so a2tea stays free
	// of markdown dependencies. Heading and caption variants are NOT
	// passed through this renderer — they are short labels that the
	// variant styles already handle. When nil, text renders as-is
	// (backward compatible).
	MarkdownRenderer func(text string, width int) string
}

// DefaultStyles returns the monochrome style set used when no host palette is
// provided. Do not add hardcoded colors here — the host overrides specific
// fields via WithStyles.
func DefaultStyles() Styles {
	return Styles{
		CardBorder: lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0, 1),

		Heading:    lipgloss.NewStyle().Bold(true),
		Subheading: lipgloss.NewStyle().Bold(true).Faint(true),
		Caption:    lipgloss.NewStyle().Faint(true),

		Button:        lipgloss.NewStyle().Bold(true),
		ButtonFocused: lipgloss.NewStyle().Bold(true).Reverse(true),
	}
}

// wrapTo wraps s to width w when w > 0; otherwise returns s unchanged.
func wrapTo(s string, w int) string {
	if w <= 0 {
		return s
	}
	return lipgloss.NewStyle().Width(w).Render(s)
}

// hr draws a horizontal rule of width w (minimum 1 cell).
func hr(w int) string {
	if w < 1 {
		w = 1
	}
	return strings.Repeat("─", w)
}
