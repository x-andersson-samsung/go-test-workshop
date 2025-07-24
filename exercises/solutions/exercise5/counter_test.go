package exercise5

import (
	"sync"
	"testing"
	"testing/synctest"

	"github.com/stretchr/testify/require"
)

func TestCounter_Inc(t *testing.T) {
	// Try commenting out c.Lock() and defer c.Unlock() from Inc method.
	// Then try running tests both with and without -race flag.
	t.Run("single", func(t *testing.T) {
		count := 4
		c := NewCounter()
		for range count {
			c.Inc()
		}
		require.Equal(t, count, c.Value())
	})
	t.Run("parallel_waitgroup", func(t *testing.T) {
		count := 4
		wg := &sync.WaitGroup{}
		wg.Add(count)

		c := NewCounter()
		for range count {
			go func() {
				c.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		require.Equal(t, count, c.Value())
	})
	t.Run("parallel_channel", func(t *testing.T) {
		count := 4
		ch := make(chan struct{})

		c := NewCounter()
		for range count {
			go func() {
				c.Inc()
				ch <- struct{}{}
			}()
		}
		for range count {
			// Read tokens from ch
			<-ch
		}

		require.Equal(t, count, c.Value())
	})
	t.Run("parallel_synctest", func(t *testing.T) {
		// Requires go1.25+ or running with GOEXPERIMENT=synctest for go1.24
		synctest.Run(func() {
			count := 4

			c := NewCounter()
			for range count {
				go c.Inc()
			}
			synctest.Wait()

			require.Equal(t, count, c.Value())
		})
	})
}

func TestCounter_Dec(t *testing.T) {
	t.Run("single", func(t *testing.T) {
		count := 4
		c := NewCounter()
		for range count {
			c.Dec()
		}
		require.Equal(t, -count, c.Value())
	})
	t.Run("parallel_waitgroup", func(t *testing.T) {
		count := 4
		wg := &sync.WaitGroup{}
		wg.Add(count)

		c := NewCounter()
		for range count {
			go func() {
				c.Dec()
				wg.Done()
			}()
		}
		wg.Wait()

		require.Equal(t, -count, c.Value())
	})
	t.Run("parallel_channel", func(t *testing.T) {
		count := 4
		ch := make(chan struct{})

		c := NewCounter()
		for range count {
			go func() {
				c.Dec()
				ch <- struct{}{}
			}()
		}
		for range count {
			// Read tokens from ch
			<-ch
		}

		require.Equal(t, -count, c.Value())
	})
	t.Run("parallel_synctest", func(t *testing.T) {
		// Requires go1.25+ or running with GOEXPERIMENT=synctest for go1.24
		synctest.Run(func() {
			count := 4

			c := NewCounter()
			for range count {
				go c.Dec()
			}
			synctest.Wait()

			require.Equal(t, -count, c.Value())
		})
	})
}

func TestCounter_Reset(t *testing.T) {
	t.Run("single", func(t *testing.T) {
		c := NewCounter()
		c.value = 4
		c.Reset()
		require.Equal(t, 0, c.Value())
	})
	t.Run("parallel_waitgroup", func(t *testing.T) {
		count := 4
		wg := &sync.WaitGroup{}
		wg.Add(count)

		c := NewCounter()
		c.value = 4

		for range count {
			go func() {
				c.Reset()
				wg.Done()
			}()
		}
		wg.Wait()

		require.Equal(t, 0, c.Value())
	})
	t.Run("parallel_channel", func(t *testing.T) {
		count := 4
		ch := make(chan struct{})

		c := NewCounter()
		c.value = 4

		for range count {
			go func() {
				c.Reset()
				ch <- struct{}{}
			}()
		}
		for range count {
			// Read tokens from ch
			<-ch
		}

		require.Equal(t, 0, c.Value())
	})
	t.Run("parallel_synctest", func(t *testing.T) {
		// Requires go1.25+ or running with GOEXPERIMENT=synctest for go1.24
		synctest.Run(func() {
			count := 4

			c := NewCounter()
			c.value = 4

			for range count {
				go c.Reset()
			}
			synctest.Wait()

			require.Equal(t, 0, c.Value())
		})
	})
}

func TestCounter_Value(t *testing.T) {
	t.Run("single", func(t *testing.T) {
		c := NewCounter()
		c.value = 4

		require.Equal(t, 4, c.Value())
	})
	t.Run("parallel_waitgroup", func(t *testing.T) {
		count := 4
		wg := &sync.WaitGroup{}
		wg.Add(count)

		c := NewCounter()
		c.value = 4

		for range count {
			go func() {
				c.Value()
				wg.Done()
			}()
		}
		wg.Wait()

		require.Equal(t, 4, c.Value())
	})
	t.Run("parallel_channel", func(t *testing.T) {
		count := 4
		ch := make(chan struct{})

		c := NewCounter()
		c.value = 4

		for range count {
			go func() {
				c.Value()
				ch <- struct{}{}
			}()
		}
		for range count {
			// Read tokens from ch
			<-ch
		}

		require.Equal(t, 4, c.Value())
	})
	t.Run("parallel_synctest", func(t *testing.T) {
		// Requires go1.25+ or running with GOEXPERIMENT=synctest for go1.24
		synctest.Run(func() {
			count := 4

			c := NewCounter()
			c.value = 4

			for range count {
				go c.Value()
			}
			synctest.Wait()

			require.Equal(t, 4, c.Value())
		})
	})
}
