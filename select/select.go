package select_demo

import (
	"fmt"
	"net/http"
	"slices"
	"time"
)

type RaceResult struct {
	Url      string
	Duration time.Duration
	Success  bool
}

type RaceResults []RaceResult

func Racer(timeout time.Duration, urls ...string) (string, error) {
	results := RaceResults{}
	resultChannel := make(chan RaceResult)

	for _, url := range urls {
		go func(u string) {
			responseTime, success := measureResponseTime(u, timeout)
			resultChannel <- RaceResult{u, responseTime, success}
		}(url)
	}

	peekIntoChannel("Before blocking", resultChannel)

	// ========================================================================
	// Go routines block on receive call e.g. r := <-channel on the channel
	// Once the channel is read from it becomes empty
	// Only way to peek inside channels is the receive statement inside select.
	// ========================================================================
	channelToSlice(resultChannel, &results, len(urls))

	// This does not have any values because the channel was read in channelToSlice function
	peekIntoChannel("After blocking", resultChannel)

	sortResults(&results)

	slowestResult := results[len(results)-1]
	if !slowestResult.Success {
		return "", fmt.Errorf("%q did not respond within %v", slowestResult.Url, timeout)
	}

	return results[0].Url, nil
}

func measureResponseTime(url string, timeout time.Duration) (time.Duration, bool) {
	start := time.Now()
	select {
	case <-ping(url):
		return time.Since(start), true
	case <-time.After(timeout):
		return timeout, false
	}
}

func ping(url string) chan struct{} {
	// always make when creating channels
	// instead of var ch chan struct{}
	// Because, for channels the zero value is nil and if you try and send to it with <-
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func sortResults(results *RaceResults) {
	slices.SortFunc(*results, func(a, b RaceResult) int {
		if a.Duration < b.Duration {
			return -1
		}
		if a.Duration > b.Duration {
			return 1
		}
		return 0
	})
}

func channelToSlice(channel chan RaceResult, slice *RaceResults, size int) {
	for range size {
		r := <-channel
		*slice = append(*slice, RaceResult{r.Url, r.Duration, r.Success})
	}
}
func peekIntoChannel[T any](prefix string, channel <-chan T) {
	select {
	case x, ok := <-channel:
		if ok {
			fmt.Printf("%s: Value %v was read.\n", prefix, x)
		} else {
			fmt.Printf("%s: Channel closed!\n", prefix)
		}
	default:
		fmt.Printf("%s: No value ready, moving on.\n", prefix)
	}
}
