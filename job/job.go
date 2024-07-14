package job

import (
	"context"
	"sync"
	"time"
)

type HandlerTask func(context *JobContext)

type JobContextObj interface {
	GetContextObject(context *JobContext)
}

type JobEngine struct {
	Context               *JobContext
	mutex                 *sync.Mutex
	Wg                    sync.WaitGroup
	isEnableManyGoroutine bool

	Timer     *time.Timer
	AwaitTime time.Duration
	task      []HandlerTask
	multTask  []HandlerTask
}

type JobContext struct {
	ctx       context.Context
	ObjectMap map[string]any
}

func NewJobEngine(ctx context.Context, awaitTime int, isEnableManyGoroutine bool) *JobEngine {
	timer := time.NewTimer(0)
	awaitTimeout := time.Duration(awaitTime) * time.Second
	// errotTimeout := time.Duration(errorTime) * time.Second
	return &JobEngine{
		Context:   &JobContext{ctx: ctx},
		mutex:     &sync.Mutex{},
		Wg:        sync.WaitGroup{},
		Timer:     timer,
		AwaitTime: awaitTimeout,
	}
}

func (j *JobEngine) JobTimingHandle() {

	for {
		select {
		case <-j.Context.ctx.Done():
			return
		case <-j.Timer.C: // wait for timer triggered
		}
		if j.isEnableManyGoroutine {
			for _, fc := range j.multTask {
				j.Wg.Add(1)
				go func(fc HandlerTask) {
					defer j.Wg.Done()
					fc(j.Context)
				}(fc)
			}
			j.Wg.Wait()
		} else {
			for _, fc := range j.task {
				fc(j.Context)
			}
		}

		j.Timer.Reset(j.AwaitTime)

	}
}

func (j *JobEngine) AddTask(tasks ...HandlerTask) {
	j.mutex.Lock()
	defer j.mutex.Unlock()

	if j.task == nil {
		j.task = make([]HandlerTask, 0)
	}
	j.task = append(j.task, tasks...)
}

func (j *JobEngine) AddMultTask(tasks ...HandlerTask) {
	j.mutex.Lock()
	defer j.mutex.Unlock()

	if j.multTask == nil {
		j.task = make([]HandlerTask, 0)
	}
	j.multTask = append(j.multTask, tasks...)
}

func (j *JobEngine) Stop() {
	j.mutex.Lock()
	defer j.mutex.Unlock()
	j.Context.ctx.Done()
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
