package main

import (
	"context"
	"log"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type RateLimiter interface {
	Wait(ctx context.Context) error
}

type MultiLimiter struct {
	limiters []RateLimiter
}

func newMultiLimiter(limiters ...RateLimiter) RateLimiter {
	return &MultiLimiter{
		limiters: limiters,
	}
}

func (l *MultiLimiter) Wait(ctx context.Context) error {
	for _, limiter := range l.limiters {
		if err := limiter.Wait(ctx); err != nil {
			return err
		}
	}
	return nil
}

type Conn struct {
	limiter RateLimiter
}

func Open() *Conn {
	r1 := rate.NewLimiter(rate.Every(time.Second/1), 4)
	r2 := rate.NewLimiter(rate.Every(time.Minute/10), 8)
	return &Conn{
		limiter: newMultiLimiter(r1, r2),
	}
}

func (c *Conn) Login(ctx context.Context, id int) error {
	if err := c.limiter.Wait(ctx); err != nil {
		return err
	}
	log.Printf("Login id=%d\n", id)
	return nil
}

func (c *Conn) Register(ctx context.Context, id int) error {
	if err := c.limiter.Wait(ctx); err != nil {
		return err
	}
	log.Printf("Register id=%d\n", id)
	return nil
}

func main() {
	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(20)
	conn := Open()
	for i := 0; i < 10; i++ {
		id := i
		go func() {
			defer wg.Done()
			if err := conn.Login(ctx, id); err != nil {
				log.Println("Login err=", err)
				return
			}
		}()
		go func(id int) {
			defer wg.Done()
			if err := conn.Register(ctx, id); err != nil {
				log.Println("Register err=", err)
				return
			}
		}(i)
	}
	wg.Wait()
}
