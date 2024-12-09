package pot // import "github.com/go-pot/pot"

import (
	"github.com/go-pot/pot/negroni"
)

type Pot struct {
	negroni.Negroni // 基础
}

func NewPot(h ...negroni.Handler) *Pot {
	neg := negroni.New(h...)

	p := &Pot{
		Negroni: *neg,
	}
	return p
}
