package workerpool

import (
	"context"
	"sync"

	article "github.com/maryam-nokohan/go-article/proto"
)

type WorkerPool struct {
	workerCount int
	jobs        chan *article.ArticleRequest
	results     chan *article.ArticleResponse
	Done        chan struct{}
}

func New(wcount int) WorkerPool {
	return WorkerPool{
		workerCount: wcount,
		jobs:        make(chan *article.ArticleRequest, wcount),
		results:     make(chan *article.ArticleResponse, wcount),
		Done:        make(chan struct{}),
	}
}

func (wp WorkerPool) Run(ctx context.Context , process func(*article.ArticleRequest) error) {
	var wg sync.WaitGroup

	for i := 0; i < wp.workerCount; i++ {
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			for {
				select {
				case <- ctx.Done():
						return
				case job , ok :=  <- wp.jobs:
					if !ok {
						return
					}
					_ = process(job)
				}
			}
		}()
	}

	wg.Wait()
	close(wp.Done)
}

func (wp *WorkerPool) Submit(job *article.ArticleRequest) {
	wp.jobs <- job
}

func (wp *WorkerPool) Close() {
	close(wp.jobs)
	<-wp.Done
}