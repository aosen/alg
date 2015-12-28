package link

import (
	"testing"
)

func TestLink(t *testing.T) {
	link := NewLink()
	link.HeadInsert(100)
	link.TailInsert(200)
}
