package cache

import (
	"time"

	"github.com/jonboulle/clockwork"
)

/*
	This package represents a way of testing time-sensitive code using clockwork package.
	go 1.25 introduced synctest into the standard which allows us to easily test such code with standard library.
	(see synctest example).
*/

type ID string

type Entry struct {
	CreatedAt time.Time
	Value     string
}

type Cache struct {
	entries map[ID]Entry
	TTL     time.Duration
	Clock   clockwork.Clock // Adding controllable clock
}

func (c *Cache) Get(id ID) (string, bool) {
	entry, ok := c.entries[id]
	if entry.CreatedAt.Add(c.TTL).Before(c.Clock.Now()) {
		delete(c.entries, id)
		return "", false
	}

	return entry.Value, ok
}

func (c *Cache) Set(id ID, value string) {
	if c.entries == nil {
		c.entries = make(map[ID]Entry)
	}

	c.entries[id] = Entry{
		CreatedAt: c.Clock.Now(),
		Value:     value,
	}
}
