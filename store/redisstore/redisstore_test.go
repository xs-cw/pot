package redisstore

import (
	"testing"
	"time"

	"github.com/go-pot/pot/store"
	ffmt "gopkg.in/ffmt.v1"
)

func TestA(t *testing.T) {
	s, err := store.NewStore("redis://db.gs.wzsm.studio:6379?db=3&password=&name=s&keyPairs=sjapp&keyPrefix=app:session_")
	if err != nil {
		ffmt.Mark(err)
		return
	}
	err = s.Load("3", time.Now())
	if err != nil {
		ffmt.Mark(err)
		return
	}

	s.Save("1", time.Now())

	s.Delete("2")
}
