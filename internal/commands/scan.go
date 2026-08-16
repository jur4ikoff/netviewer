package scanner

import (
	"context"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

const WORKERS_COUNT = 10

type Result struct {
	Port int
	Open bool
}

type ScanRequest struct {
	Host string
	From int
	To   int
}

type Scanner struct {
}

func NewScanner() *Scanner {
	return &Scanner{}
}

func (s Scanner) Scan(ctx context.Context, request *ScanRequest) error {
	jobs := make(chan int)
	results := make(chan Result)

	var wg sync.WaitGroup

	for i := 0; i < WORKERS_COUNT; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.worker(ctx, request, jobs, results)
		}()
	}

	go func() {
		defer close(jobs)

		for port := request.From; port <= request.To; port++ {
			select {
			case jobs <- port:
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	var scanned []Result

	for result := range results {
		scanned = append(scanned, result)
	}

	if err := ctx.Err(); err != nil {
		return err
	}

	for _, res := range scanned {
		if res.Open {
			log.Ctx(ctx).Info().Msgf("OPEN %s:%d", request.Host, res.Port)
		}

	}
	return nil
}

func (s Scanner) scanPort(
	ctx context.Context,
	host string,
	port int,
) bool {
	timeout := 30 * time.Second
	address := net.JoinHostPort(host, strconv.Itoa(port))
	dialer := net.Dialer{Timeout: timeout}

	conn, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func (s *Scanner) worker(
	ctx context.Context,
	request *ScanRequest,
	jobs <-chan int,
	results chan<- Result) {
	for {
		select {
		case <-ctx.Done():
			return
		case port, ok := <-jobs:
			if !ok {
				return
			}

			result := s.scanPort(ctx, request.Host, port)
			select {
			case results <- Result{Port: port, Open: result}:
			case <-ctx.Done():
				return
			}
		}
	}
}
