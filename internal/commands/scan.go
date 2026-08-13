package commands

import "context"

type ScanRequest struct {
	host string
	from int
	to   int
}

type Scanner struct {
}

func NewScanner() *Scanner {
	return &Scanner{}
}

func (s Scanner) Scan(ctx context.Context, request ScanRequest) error {
	return nil
}
