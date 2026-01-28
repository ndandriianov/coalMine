package pauseController

import (
	"context"
	"sync"
)

type PauseController struct {
	isMineOnPause bool
	mtx           sync.Mutex
	cond          *sync.Cond
}

// NewPauseController returns the pointer to a new PauseController with initial value false.
func NewPauseController() *PauseController {
	pc := &PauseController{
		isMineOnPause: false,
		mtx:           sync.Mutex{},
	}
	pc.cond = sync.NewCond(&pc.mtx)

	return pc
}

func (c *PauseController) Pause() {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.isMineOnPause = true
}

func (c *PauseController) Resume() {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.isMineOnPause = false
	c.cond.Broadcast()
}

func (c *PauseController) WaitIfPaused(ctx context.Context) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	for c.isMineOnPause {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		c.cond.Wait()
	}

	return nil
}

func (c *PauseController) IsMineOnPause() bool {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	return c.isMineOnPause
}
