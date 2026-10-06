//go:build solution

package channelsynchronization

func Worker(work func(), done chan bool) {
	work()
	done <- true
}

func Start(work func()) chan bool {
	done := make(chan bool, 1)
	go Worker(work, done)
	return done
}

func WaitAll(works ...func()) {
	done := make(chan bool)
	for _, work := range works {
		go Worker(work, done)
	}
	for range works {
		<-done
	}
}

func RunSequence(steps ...func()) {
	// prev is the channel on which the previous step announces that it is
	// done. Nothing comes before the first step, so its signal is preloaded.
	prev := make(chan bool, 1)
	prev <- true
	for _, step := range steps {
		next := make(chan bool, 1)
		go func(wait, done chan bool) {
			<-wait
			step()
			done <- true
		}(prev, next)
		prev = next
	}
	<-prev
}
