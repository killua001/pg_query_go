// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node A_ArrayExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "A_ArrayExpr")

	for _, subNode := range node.Elements {
		subNode.Fingerprint(ctx)
	}
}
