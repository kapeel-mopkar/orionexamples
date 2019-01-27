package service

import (
	"context"
	"reflect"
	"testing"

	proto "github.com/carousell/orionexamples/consignmentservice/pkg/api/v1"
)

func Test_svc_CreateConsignment(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.ConsignmentRequest
	}

	tests := []struct {
		name    string
		s       *svc
		args    args
		want    *proto.ConsignmentResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateConsignment(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("svc.CreateConsignment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("svc.CreateConsignment() = %v, want %v", got, tt.want)
			}
		})
	}
}
