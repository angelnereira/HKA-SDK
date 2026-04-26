// Package internal provides unexported helpers used across the SDK.
package internal

import "strings"

// Tag writes a non-empty XML element. If value is empty the tag is omitted.
func Tag(sb *strings.Builder, ns, name, value string) {
	if value == "" {
		return
	}
	sb.WriteString("<")
	sb.WriteString(ns)
	sb.WriteString(":")
	sb.WriteString(name)
	sb.WriteString(">")
	sb.WriteString(EscapeXML(value))
	sb.WriteString("</")
	sb.WriteString(ns)
	sb.WriteString(":")
	sb.WriteString(name)
	sb.WriteString(">")
}

// TagForce writes an XML element even when value is empty.
// Use only for fields that must be present as empty tags per HKA spec.
func TagForce(sb *strings.Builder, ns, name, value string) {
	sb.WriteString("<")
	sb.WriteString(ns)
	sb.WriteString(":")
	sb.WriteString(name)
	sb.WriteString(">")
	sb.WriteString(EscapeXML(value))
	sb.WriteString("</")
	sb.WriteString(ns)
	sb.WriteString(":")
	sb.WriteString(name)
	sb.WriteString(">")
}

// Open writes an opening XML element tag.
func Open(sb *strings.Builder, ns, name string) {
	sb.WriteString("<")
	sb.WriteString(ns)
	sb.WriteString(":")
	sb.WriteString(name)
	sb.WriteString(">")
}

// Close writes a closing XML element tag.
func Close(sb *strings.Builder, ns, name string) {
	sb.WriteString("</")
	sb.WriteString(ns)
	sb.WriteString(":")
	sb.WriteString(name)
	sb.WriteString(">")
}

// EscapeXML replaces XML special characters with their entity references.
func EscapeXML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "\"", "&quot;")
	s = strings.ReplaceAll(s, "'", "&apos;")
	return s
}
