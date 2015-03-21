package canonical

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/zenazn/goji/web"
)

func Canonical(hostname string) func(c *web.C, h http.Handler) http.Handler {
	return func(c *web.C, h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if !strings.HasPrefix(r.Host, "localhost") {
				if r.Host != hostname {
					http.Redirect(w, r, fmt.Sprintf("http://%s%s", hostname, r.URL.Path), http.StatusMovedPermanently)
					return
				}
			}

			h.ServeHTTP(w, r)
			return
		}

		return http.HandlerFunc(fn)
	}
}
