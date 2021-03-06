// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package deploymentapi

import (
	"errors"
	"reflect"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestShutdown(t *testing.T) {
	type args struct {
		params ShutdownParams
	}
	tests := []struct {
		name string
		args args
		want *models.DeploymentShutdownResponse
		err  error
	}{
		{
			name: "fails on parameter validation",
			err: multierror.NewPrefixed("deployment shutdown",
				apierror.ErrMissingAPI,
				apierror.ErrDeploymentID,
			),
		},
		{
			name: "fails on API error",
			args: args{params: ShutdownParams{
				API:          api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
				DeploymentID: mock.ValidClusterID,
			}},
			err: errors.New(`{"error": "some error"}`),
		},
		{
			name: "Succeeds",
			args: args{params: ShutdownParams{
				API: api.NewMock(mock.New200Response(mock.NewStructBody(models.DeploymentShutdownResponse{
					ID: ec.String(mock.ValidClusterID),
				}))),
				DeploymentID: mock.ValidClusterID,
			}},
			want: &models.DeploymentShutdownResponse{
				ID: ec.String(mock.ValidClusterID),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Shutdown(tt.args.params)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Shutdown() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shutdown() = %v, want %v", got, tt.want)
			}
		})
	}
}
