package monitoring

import (
	"context"
	"fmt"

	"github.com/garupanojisan/spanner-autoscaler/metrics"
	"github.com/garupanojisan/spanner-autoscaler/voters"
)

// Voter is implements Poller interface
type Voter struct {
	collector metrics.Collector
}

// NewMonitoringVoter returns new instance of monitoring.Poller
func NewMonitoringVoter(c metrics.Collector) voters.Voter {
	return &Voter{
		collector: c,
	}
}

// Poll returns a recommeneded number of nodes based on metrics below.
// 1. recommended high cpu utilization (https://cloud.google.com/spanner/docs/cpu-utilization#recommended-max)
// 2. recommended limit for storage per node (https://cloud.google.com/spanner/docs/monitoring-cloud#storage)
func (p *Voter) Poll(ctx context.Context, projectID, instanceID string) (int64, error) {
	cpuHighPriority, err := p.collector.GetCPUHighPriorityTotal(ctx, projectID, instanceID)
	if err != nil {
		return 0, err
	}

	cpu24Smoothed, err := p.collector.GetCPU24HourSmoothedAggregate(ctx, projectID, instanceID)
	if err != nil {
		return 0, err
	}

	storageUtilization, err := p.collector.GetStorageUtilizationPerNode(ctx, projectID, instanceID)
	if err != nil {
		return 0, err
	}

	fmt.Printf("CPU(High Pri) = %v %%\n", cpuHighPriority)
	fmt.Printf("CPU(24h) = %v %%\n", cpu24Smoothed)
	fmt.Printf("Storage = %v %%\n", storageUtilization)
	return 0, nil
}
