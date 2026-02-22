package workerpool

import (
	"context"
	"sync"

	article "github.com/maryam-nokohan/go-article/proto"
)

type WorkerPool struct {
	workerCount int
	jobs        chan article.ArticleRequest
	results     chan article.ArticleResponse
	Done        chan struct{}
}

func New(wcount int) WorkerPool {
	return WorkerPool{
		workerCount: wcount,
		jobs:        make(chan article.ArticleRequest, wcount),
		results:     make(chan article.ArticleResponse, wcount),
		Done:        make(chan struct{}),
	}
}

func (wp WorkerPool) Run(ctx context.Context) {
	var wg sync.WaitGroup

	for i := 0; i < wp.workerCount; i++ {
		wg.Add(1)
		// fan out worker goroutines
		//reading from jobs channel and
		//pushing calcs into results channel
		// go worker(ctx, &wg, wp.jobs, wp.results)
	}

	wg.Wait()
	close(wp.Done)
	close(wp.results)
}
