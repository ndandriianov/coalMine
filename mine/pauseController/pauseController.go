package pauseController

import (
	"context"
	"sync"
)

type PauseController struct {
	mtx    sync.RWMutex
	pause  chan struct{}
	resume chan struct{}
}

// NewPauseController returns the pointer to a new PauseController initially not on pause.
func NewPauseController() *PauseController {
	pause := make(chan struct{})
	resume := make(chan struct{})
	close(resume)

	return &PauseController{
		pause:  pause,
		resume: resume,
	}
}

func (c *PauseController) Pause() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	select {
	case <-c.pause:
	default:
		close(c.pause)
		c.resume = make(chan struct{})
	}
}

func (c *PauseController) Resume() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	select {
	case <-c.resume:
	default:
		close(c.resume)
		c.pause = make(chan struct{})
	}
}

func (c *PauseController) WaitIfPaused(ctx context.Context) error {
	c.mtx.RLock()
	resume := c.resume
	c.mtx.RUnlock()

	select {
	case <-resume:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (c *PauseController) PauseChan() <-chan struct{} {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.pause
}
