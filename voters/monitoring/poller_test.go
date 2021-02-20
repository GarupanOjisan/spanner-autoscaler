package monitoring

import (
	"context"
	"os"
	"testing"

	"github.com/garupanojisan/spanner-autoscaler/metrics"
	"github.com/garupanojisan/spanner-autoscaler/metrics/monitoring"
)

func TestPoller_Poll(t *testing.T) {
	ctx := context.Background()
	c, err := monitoring.NewCollector(ctx)
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		collector metrics.Collector
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
		want    int64
		wantErr bool
	}{
		{
			name:   "ok",
			fields: fields{collector: c},
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
			p := &Voter{
				collector: tt.fields.collector,
			}
			got, err := p.Poll(tt.args.ctx, tt.args.projectID, tt.args.instanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Poller.Poll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Poller.Poll() = %v, want %v", got, tt.want)
			}
		})
	}
}
