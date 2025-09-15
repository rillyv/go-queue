package goqueue

import (
	"sync"
	"time"
	"github.com/google/uuid"
)

type Job struct {
	ID        string
	Payload   []byte
	Tries     int
	VisibleAt time.Time
}

type Queue struct {
	mu   sync.Mutex
	jobs []*Job
}

func NewQueue() *Queue {
	return &Queue{jobs: make([]*Job, 0)}
}

func (q *Queue) Enqueue(payload []byte) string {
	q.mu.Lock()
	defer q.mu.Unlock()
	job := &Job{
		ID: uuid.New().String(),
		Payload: payload,
		VisibleAt: time.Now(),
	}

	q.jobs = append(q.jobs, job)
	return job.ID
}

func (q *Queue) Dequeue() *Job {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.jobs) == 0 {
		return nil
	}
	job := q.jobs[0]
	q.jobs = q.jobs[1:]
	return job
}
