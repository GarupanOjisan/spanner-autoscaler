package spanner

import (
	"context"
	"fmt"

	instance "cloud.google.com/go/spanner/admin/instance/apiv1"
	instancepb "google.golang.org/genproto/googleapis/spanner/admin/instance/v1"
	"google.golang.org/genproto/protobuf/field_mask"

	"github.com/garupanojisan/spanner-autoscaler/scaler"
)

// Scaler is implement of Cloud Spanner
type Scaler struct {
	client *instance.InstanceAdminClient
}

// NewSpannerScaler create a new Scaler
func NewSpannerScaler(ctx context.Context) (scaler.Scaler, error) {
	adminClient, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		return nil, err
	}
	return &Scaler{client: adminClient}, nil
}

// Scale scales number of node
func (s *Scaler) Scale(ctx context.Context, projectID, instanceID string, numNode int32) error {
	op, err := s.client.UpdateInstance(ctx, &instancepb.UpdateInstanceRequest{
		Instance: &instancepb.Instance{
			Name:      fmt.Sprintf("projects/%s/instances/%s", projectID, instanceID),
			NodeCount: numNode,
		},
		FieldMask: &field_mask.FieldMask{
			Paths: []string{"node_count"},
		},
	})
	if err != nil {
		return err
	}
	if _, err := op.Wait(ctx); err != nil {
		return err
	}
	return nil
}
