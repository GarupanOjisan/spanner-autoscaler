package monitoring

import (
	"context"
	"os"
	"testing"

	monitoring "cloud.google.com/go/monitoring/apiv3"
)

func TestCollector_GetCPUHighPriorityTotal(t *testing.T) {
	ctx := context.Background()
	client, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		t.Fatal(err)
	}
	type fields struct {
		client *monitoring.MetricClient
	}
	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				client: client,
			},
			args: args{
				ctx:        ctx,
				projectID:  os.Getenv("GCP_PROJECT_ID"),
				instanceID: os.Getenv("SPANNER_INSTANCE_ID"),
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Collector{
				client: tt.fields.client,
			}
			got, err := c.GetCPUHighPriorityTotal(tt.args.ctx, tt.args.projectID, tt.args.instanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collector.GetCPUHighPriorityTotal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got <= tt.want {
				t.Errorf("Collector.GetCPUHighPriorityTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollector_GetCPU24HourSmoothedAggregate(t *testing.T) {
	ctx := context.Background()
	client, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		t.Fatal(err)
	}
	type fields struct {
		client *monitoring.MetricClient
	}
	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				client: client,
			},
			args: args{
				ctx:        ctx,
				projectID:  os.Getenv("GCP_PROJECT_ID"),
				instanceID: os.Getenv("SPANNER_INSTANCE_ID"),
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Collector{
				client: tt.fields.client,
			}
			got, err := c.GetCPU24HourSmoothedAggregate(tt.args.ctx, tt.args.projectID, tt.args.instanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collector.GetCPU24HourSmoothedAggregate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got <= tt.want {
				t.Errorf("Collector.GetCPU24HourSmoothedAggregate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollector_GetStorageUtilizationPerNode(t *testing.T) {
	ctx := context.Background()
	client, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		t.Fatal(err)
	}
	type fields struct {
		client *monitoring.MetricClient
	}
	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				client: client,
			},
			args: args{
				ctx:        ctx,
				projectID:  os.Getenv("GCP_PROJECT_ID"),
				instanceID: os.Getenv("SPANNER_INSTANCE_ID"),
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Collector{
				client: tt.fields.client,
			}
			got, err := c.GetStorageUtilizationPerNode(tt.args.ctx, tt.args.projectID, tt.args.instanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collector.GetStorageUtilizationPerNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got <= tt.want {
				t.Errorf("Collector.GetStorageUtilizationPerNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
