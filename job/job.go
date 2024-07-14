package job

import (
	"context"
	"sync"
	"time"
)

type HandlerTask func(context *JobContext)

type JobEngine struct {
	Context               *JobContext
	mutex                 *sync.Mutex
	Wg                    sync.WaitGroup
	isEnableManyGoroutine bool

	timer     *time.Timer
	awaitTime time.Duration
	batchTask []HandlerTask
	task      []HandlerTask
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
		Context:               &JobContext{ctx: ctx},
		mutex:                 &sync.Mutex{},
		Wg:                    sync.WaitGroup{},
		timer:                 timer,
		awaitTime:             awaitTimeout,
		isEnableManyGoroutine: isEnableManyGoroutine,
	}
}

func (j *JobEngine) JobTimingHandle() {

	for {
		select {
		case <-j.Context.ctx.Done():
			return
		case <-j.timer.C: // wait for timer triggered
		}

		doTask := func() {
			j.Wg.Add(1)
			go func() {
				defer j.Wg.Done()
				for _, fc := range j.task {
					fc(j.Context)
				}

			}()
		}

		doMultTask := func() {
			if j.isEnableManyGoroutine {
				for _, fc := range j.batchTask {
					j.Wg.Add(1)
					go func(fc HandlerTask) {
						defer j.Wg.Done()
						fc(j.Context)
					}(fc)
				}
			}
		}

		doTask()
		doMultTask()

		j.Wg.Wait()

		j.timer.Reset(j.awaitTime)

	}
}

func (j *JobEngine) AddMultTask(tasks ...HandlerTask) {

	j.mutex.Lock()
	defer j.mutex.Unlock()

	if j.batchTask == nil {
		j.batchTask = make([]HandlerTask, 0)
	}
	j.batchTask = append(j.batchTask, tasks...)
}

func (j *JobEngine) AddTask(tasks ...HandlerTask) {
	j.mutex.Lock()
	defer j.mutex.Unlock()

	if j.task == nil {
		j.task = make([]HandlerTask, 0)
	}
	j.task = append(j.task, tasks...)
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
