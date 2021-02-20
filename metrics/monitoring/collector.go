package monitoring

import (
	"context"
	"fmt"
	"time"

	monitoring "cloud.google.com/go/monitoring/apiv3"
	"github.com/golang/protobuf/ptypes/duration"
	"google.golang.org/api/iterator"
	pb "google.golang.org/genproto/googleapis/monitoring/v3"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Collector implements Collector interface which calls Monitoring API.
type Collector struct {
	client *monitoring.MetricClient
}

// NewCollector returns a new Collector instance.
func NewCollector(ctx context.Context) (*Collector, error) {
	c, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		return nil, err
	}
	return &Collector{client: c}, nil
}

// GetCPUHighPriorityTotal returns CPU utilization (High priority total)
func (c *Collector) GetCPUHighPriorityTotal(ctx context.Context, projectID, instanceID string) (float64, error) {
	filter := fmt.Sprintf(`
resource.type="spanner_instance" AND
resource.labels.instance_id="%s" AND
resource.labels.project_id="%s" AND
metric.type="spanner.googleapis.com/instance/cpu/utilization_by_priority" AND
metric.label.priority="high"
	`, instanceID, projectID)
	return c.getMaxMetricValue(ctx, projectID, instanceID, filter)
}

// GetCPU24HourSmoothedAggregate returns CPU utilization (24-hour smoothed aggregate)
func (c *Collector) GetCPU24HourSmoothedAggregate(ctx context.Context, projectID, instanceID string) (float64, error) {
	filter := fmt.Sprintf(`
resource.type="spanner_instance" AND
resource.labels.instance_id="%s" AND
resource.labels.project_id="%s" AND
metric.type="spanner.googleapis.com/instance/cpu/smoothed_utilization"
	`, instanceID, projectID)
	return c.getMaxMetricValue(ctx, projectID, instanceID, filter)
}

// GetStorageUtilizationPerNode returns usage of Storage
func (c *Collector) GetStorageUtilizationPerNode(ctx context.Context, projectID, instanceID string) (float64, error) {
	filter := fmt.Sprintf(`
resource.type="spanner_instance" AND
resource.labels.instance_id="%s" AND
resource.labels.project_id="%s" AND
metric.type="spanner.googleapis.com/instance/storage/utilization"
	`, instanceID, projectID)
	return c.getMaxMetricValue(ctx, projectID, instanceID, filter)
}

func (c *Collector) getMaxMetricValue(ctx context.Context, projectID, instanceID, filter string) (float64, error) {
	const metricWindowMin = 5
	now := time.Now()
	iter := c.client.ListTimeSeries(ctx, &pb.ListTimeSeriesRequest{
		Name: fmt.Sprintf("projects/%s", projectID),
		Interval: &pb.TimeInterval{
			StartTime: timestamppb.New(now.Add(-time.Minute * metricWindowMin)),
			EndTime:   timestamppb.New(now),
		},
		Aggregation: &pb.Aggregation{
			AlignmentPeriod:    &duration.Duration{Seconds: int64(time.Minute.Seconds())},
			PerSeriesAligner:   pb.Aggregation_ALIGN_MAX,
			CrossSeriesReducer: pb.Aggregation_REDUCE_SUM,
		},
		Filter: filter,
		View:   pb.ListTimeSeriesRequest_FULL,
	})

	max := float64(0)
	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 0, err
		}
		for _, p := range resp.GetPoints() {
			v := p.GetValue().GetDoubleValue() * 100
			if v > max {
				max = v
			}
		}
	}
	return max, nil
}
