package remotepackage_test

import (
	"testing"

	cm "github.com/easierway/concurrent_map"
)

func TestCocurrentMao(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}