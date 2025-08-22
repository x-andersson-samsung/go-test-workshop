package synctest

import (
	"log/slog"
	"sync"
	"testing"
	"testing/synctest"
	"time"
)

// Despite the fact that we are waiting for 4 seconds in the longest routine, the test will still barely take any time.
func TestSynctestTime(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		wg := sync.WaitGroup{}
		wg.Add(3)
		go func() {
			time.Sleep(1 * time.Second)
			slog.Info("1s go routine")
			wg.Done()
		}()
		go func() {
			time.Sleep(2 * time.Second)
			slog.Info("2s go routine")
			wg.Done()
		}()
		go func() {
			time.Sleep(4 * time.Second)
			slog.Info("4s go routine")
			wg.Done()
		}()

		time.Sleep(3 * time.Second)
		slog.Info("3s main routing")

		wg.Wait()
	})
}
