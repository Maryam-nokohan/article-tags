package pkg

import (
	"context"
	"log"
	"sync"

	article "github.com/maryam-nokohan/go-article/proto"
)

type WorkerPool struct {
	workerCount int
	jobs        chan *article.ArticleRequest
	wg          sync.WaitGroup
}

func New(wcount int) WorkerPool {
	return WorkerPool{
		workerCount: wcount,
		jobs:        make(chan *article.ArticleRequest, wcount),
		wg: 		sync.WaitGroup{},
	}
}

func (wp *WorkerPool) Run(ctx context.Context, process func(*article.ArticleRequest) error) {

	for i := 0; i < wp.workerCount; i++ {
		wp.wg.Add(1)
		go func() {
			defer wp.wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case job, ok := <-wp.jobs:
					if !ok {
						return
					}
					err := process(job)
					if err != nil {
						log.Println("Error processing job:", err)
						continue
					}
				}
			}
		}()
	}

}

func (wp *WorkerPool) Submit(job *article.ArticleRequest) {
	wp.jobs <- job
}

func (wp *WorkerPool) Close() {
	close(wp.jobs)
	wp.wg.Wait()
}
