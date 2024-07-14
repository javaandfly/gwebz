package job

import (
	"context"
	"testing"
	"time"
)

func TestJobEngine(t *testing.T) {
	start := time.Now()
	job := NewJobEngine(context.Background(), 10000, false)
	for i := 0; i < 10000; i++ {
		job.AddTask(add)
	}
	go job.JobTimingHandle()

	time.Sleep(10 * time.Second)

	job.Stop()

	elapsed := time.Since(start)
	t.Logf("该函数执行完成耗时：%v", elapsed)

}

func add(context *JobContext) {
	for i := 0; i < 10; i++ {
		i = i + 1
	}
}
