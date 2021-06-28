package TaskPool

import (
	"fmt"
	"os"
	"runtime"
	"sync/atomic"
)

type TaskItem struct {
	Item func()
}

type Option struct {
	WorkSize int
}

type TaskPool struct {
	opt          Option
	taskChanPool []chan TaskItem

	taskNum uint64

	exit chan struct{}
	stop bool
}

func NewTaskPool(o Option) *TaskPool {
	taskPool := &TaskPool{
		opt:     o,
		taskNum: 0,
		exit:    make(chan struct{}),
	}
	for i := 0; i < taskPool.opt.WorkSize; i++ {
		taskChan := make(chan TaskItem, 100)
		taskPool.taskChanPool[i] = taskChan
		go taskPool.work(taskChan, taskPool.exit)
	}
	return taskPool
}

func (pThis *TaskPool) work(taskChan chan TaskItem, exit chan struct{}) {
	for {
		select {
		case taskItem, ok := <-taskChan:
			if ok {
				pThis.handle(taskItem)
			} else {
				return
			}
		case <-exit:
			return
		}
	}
}

func (pThis *TaskPool) handle(task TaskItem) {
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			pl := fmt.Sprintf("http call panic: %v\n%s\n", err, buf)
			fmt.Fprintf(os.Stderr, pl)
		}
	}()
	task.Item()
}

func (pThis *TaskPool) Push(task TaskItem) {
	pThis.taskChanPool[atomic.AddUint64(&pThis.taskNum, 1)%uint64(pThis.opt.WorkSize)] <- task
}

func (pThis *TaskPool) Close() {
	if pThis.stop {
		return
	}
	pThis.stop = true
	close(pThis.exit)
}
