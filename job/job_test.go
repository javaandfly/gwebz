package job

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestJobEngine(t *testing.T) {
	start := time.Now()
	job := NewJobEngine(context.Background(), 10000000, false)
	for i := 0; i < 10000; i++ {
		job.AddTask(add)
	}
	job.JobTimingHandle()

	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)

}

func add(context *JobContext) {
	for i := 0; i < 10; i++ {

	}
}
