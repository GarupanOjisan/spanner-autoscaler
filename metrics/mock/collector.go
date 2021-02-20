package mock

import (
	"context"
)

// Collector is mock of metric collector
type Collector struct {
	ResultOfGetCPUHighPriorityTotal       float64
	ResultOfGetCPU24HourSmoothedAggregate float64
	ResultOfGetStorageUtilizationPerNode  float64
}

// GetCPUHighPriorityTotal returns values which predefined by caller
func (c *Collector) GetCPUHighPriorityTotal(ctx context.Context, projectID, instanceID string) (float64, error) {
	return c.ResultOfGetCPUHighPriorityTotal, nil
}

// GetCPU24HourSmoothedAggregate returns values which predefined by caller
func (c *Collector) GetCPU24HourSmoothedAggregate(ctx context.Context, projectID, instanceID string) (float64, error) {
	return c.ResultOfGetCPU24HourSmoothedAggregate, nil
}

// GetStorageUtilizationPerNode returns values which predefined by caller
func (c *Collector) GetStorageUtilizationPerNode(ctx context.Context, projectID, instanceID string) (float64, error) {
	return c.ResultOfGetStorageUtilizationPerNode, nil
}
