package job

import (
	"context"
	"sync"
	"time"
)

type HandlerTask func(context *JobContext)

type JobEngine struct {
	Context   *JobContext
	Mutex     *sync.Mutex
	Timer     *time.Timer
	AwaitTime time.Duration
	ErrorTime time.Duration
	Task      []HandlerTask
}

type JobContext struct {
	ctx context.Context

	ObjectMap map[string]any
}

func NewJobEngine(ctx context.Context, awaitTime, errorTime int) *JobEngine {
	timer := time.NewTimer(0)
	awaitTimeout := time.Duration(awaitTime) * time.Second
	errotTimeout := time.Duration(errorTime) * time.Second
	return &JobEngine{
		Context:   &JobContext{ctx: ctx},
		Timer:     timer,
		AwaitTime: awaitTimeout,
		ErrorTime: errotTimeout,
	}
}

func (j *JobEngine) JobTimingHandle() {

	for {
		select {
		case <-j.Context.ctx.Done():
			return
		case <-j.Timer.C: // wait for timer triggered
		}
		for _, fc := range j.Task {
			fc(j.Context)
		}

	}
}

func (j *JobEngine) AddTask(tasks ...HandlerTask) {
	if tasks == nil {
		j.Task = make([]HandlerTask, 0)
	}
	j.Task = append(j.Task, tasks...)
}

func (j *JobEngine) ErrorTimeout() {
	j.Timer.Reset(j.ErrorTime)
}

func (j *JobEngine) AwaitTimeout() {
	j.Timer.Reset(j.AwaitTime)
}

func (c *JobContext) SetContextMap(key string, value any) {
	if c.ObjectMap == nil {
		c.ObjectMap = make(map[string]any)
	}
	c.ObjectMap[key] = value
}

func (c *JobContext) GetContextObject(key string) (value any, exists bool) {

	value, exists = c.ObjectMap[key]
	return
}
