// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node MergeJoin) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "MergeJoin")

	for _, subNode := range node.Mergeclauses {
		subNode.Fingerprint(ctx)
	}
}
