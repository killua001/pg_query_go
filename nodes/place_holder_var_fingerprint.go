// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node PlaceHolderVar) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "PlaceHolderVar")
	if node.Phexpr != nil {
		node.Phexpr.Fingerprint(ctx)
	}
}
