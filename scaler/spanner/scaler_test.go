package spanner

import (
	"context"
	"os"
	"testing"

	instance "cloud.google.com/go/spanner/admin/instance/apiv1"
)

func TestScaler_Scale(t *testing.T) {
	ctx := context.Background()
	adminClient, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		t.Fatal(err)
	}
	type fields struct {
		client *instance.InstanceAdminClient
	}
	type args struct {
		ctx        context.Context
		projectID  string
		instanceID string
		numNode    int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				client: adminClient,
			},
			args: args{
				ctx:        ctx,
				projectID:  os.Getenv("GCP_PROJECT_ID"),
				instanceID: os.Getenv("SPANNER_INSTANCE_ID"),
				numNode:    1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scaler{
				client: tt.fields.client,
			}
			if err := s.Scale(tt.args.ctx, tt.args.projectID, tt.args.instanceID, tt.args.numNode); (err != nil) != tt.wantErr {
				t.Errorf("Scaler.Scale() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
