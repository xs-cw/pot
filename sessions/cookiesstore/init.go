package cookiesstore

import (
	"net/url"

	"github.com/go-pot/pot/sessions"
)

func init() {
	sessions.RegisterSession("cookie", func(ur *url.URL) (sessions.Store, error) {
		query := ur.Query()
		pairs := []byte(query.Get("keyPairs"))

		return New(pairs), nil
	})
}
