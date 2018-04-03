package hat

import "strings"

// urlJoin is used instead of a more formal url.URL concat so
// that multiple paths
func urlJoin(elems ...string) string {
	var u strings.Builder
	for i, e := range elems {
		e = strings.Trim(e, "/")

		u.WriteString(e)
		if i != len(elems)-1 {
			u.WriteString("/")
		}
	}
	return u.String()
}
