package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			// Send statement
			resultChannel <- result{url, wc(url)}
		}()
	}

	for range urls {
		// Receive statement
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
