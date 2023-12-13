package job

import (
	"context"
	"sync"
	"time"
)

type HandlerTask func(context *JobContext)

type JobEngine struct {
	Context *JobContext
	mutex   *sync.Mutex
	Wg      sync.WaitGroup

	Timer      *time.Timer
	AwaitTime  time.Duration
	PrefixTask []HandlerTask
	Task       []HandlerTask
}

type JobContext struct {
	ctx       context.Context
	ObjectMap map[string]any
}

func NewJobEngine(ctx context.Context, awaitTime int) *JobEngine {
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
		for _, fc := range j.PrefixTask {
			fc(j.Context)
		}
		for _, fc := range j.Task {
			j.Wg.Add(1)
			go func(fc HandlerTask) {
				defer j.Wg.Done()
				fc(j.Context)
			}(fc)
		}
		j.Wg.Wait()
		j.Timer.Reset(j.AwaitTime)

	}
}

func (j *JobEngine) AddTask(tasks ...HandlerTask) {
	j.mutex.Lock()
	defer j.mutex.Unlock()

	if tasks == nil {
		j.Task = make([]HandlerTask, 0)
	}
	j.Task = append(j.Task, tasks...)
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
