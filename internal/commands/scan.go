package scanner

import "context"

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
	return nil
}
