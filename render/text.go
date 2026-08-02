package render

import (
	"strings"

	a2ui "github.com/tmc/a2ui"
)

// renderText renders a Text component: the resolved text styled per variant
// (h1–h3 heading, h4–h5 subheading, caption faint; body and unknown variants
// plain), wrapped to the surface width.
//
// Body and unstyled Text whose content looks like Markdown is passed through
// the host's MarkdownRenderer when one is configured. Headings and captions
// are short labels that the variant styles already handle, so they skip the
// renderer.
func (s *Surface) renderText(c a2ui.Component) string {
	text := s.dynString(c.Text.Text)
	switch c.Text.Variant {
	case a2ui.TextVariantH1, a2ui.TextVariantH2, a2ui.TextVariantH3:
		text = s.styles.Heading.Render(text)
	case a2ui.TextVariantH4, a2ui.TextVariantH5:
		text = s.styles.Subheading.Render(text)
	case a2ui.TextVariantCaption:
		text = s.styles.Caption.Render(text)
	default:
		if s.styles.MarkdownRenderer != nil && looksLikeMarkdown(text) {
			return s.styles.MarkdownRenderer(text, s.width)
		}
	}
	return wrapTo(text, s.width)
}

// looksLikeMarkdown reports whether the text contains common Markdown markers
// that would benefit from a Markdown renderer.
func looksLikeMarkdown(s string) bool {
	return strings.Contains(s, "# ") ||
		strings.Contains(s, "**") ||
		strings.Contains(s, "| ") ||
		strings.Contains(s, "```") ||
		strings.Contains(s, "\n- ") ||
		strings.Contains(s, "\n1. ")
}
