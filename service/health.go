package service

import (
	"context"
	pbhealth "github.com/zhs/esb-protobufs/go/health"
	"sync"
)

func NewHealthService() *HealthService {
	return &HealthService{}
}

type HealthService struct {
	mu sync.Mutex
}

func (s *HealthService) Check(ctx context.Context, in *pbhealth.HealthCheckRequest, res *pbhealth.HealthCheckResponse) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	res.Status = pbhealth.HealthCheckResponse_SERVING
	return nil
}
