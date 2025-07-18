package cache

import (
	"testing"
	"time"

	"github.com/jonboulle/clockwork"
)

func TestCache(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ch := Cache{
			TTL:   time.Hour,
			Clock: clockwork.NewFakeClock(),
		}
		expected := "value"

		ch.Set("key", expected)
		got, ok := ch.Get("key")
		if !ok {
			t.Error("key not found in cache")
		}
		if got != expected {
			t.Errorf("got %s, want %s", got, expected)
		}
	})
	t.Run("expired", func(t *testing.T) {
		clock := clockwork.NewFakeClock()
		ch := Cache{
			TTL:   time.Hour,
			Clock: clock,
		}
		expected := "value"

		ch.Set("key", expected)
		clock.Advance(2 * time.Hour)

		_, ok := ch.Get("key")
		if ok {
			t.Error("expected key to expire")
		}
	})
}
