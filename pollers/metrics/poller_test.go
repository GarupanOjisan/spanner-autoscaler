package metrics

import (
	"context"
	"testing"
)

func TestPoller_Poll(t *testing.T) {
	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
	}

	tests := []struct {
		name    string
		p       *Poller
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Poller{}
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
