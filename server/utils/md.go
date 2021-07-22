package utils

import "github.com/88250/lute"

func MarkdownToHtml(mdStr string) string {
	luteEngine := lute.New()
	luteEngine.RenderOptions.CodeSyntaxHighlightInlineStyle = true
	return luteEngine.MarkdownStr("", mdStr)
}
