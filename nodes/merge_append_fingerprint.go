// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node MergeAppend) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "MergeAppend")

	for _, subNode := range node.Mergeplans {
		subNode.Fingerprint(ctx)
	}
}
