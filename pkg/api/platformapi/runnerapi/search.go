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

package runnerapi

import (
	"context"

	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_infrastructure"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// SearchParams contains parameters used to search runner's data using Query DSL
type SearchParams struct {
	*api.API
	Region  string
	Request models.SearchRequest
}

// Validate is the implementation for the ecctl.Validator interface
func (params SearchParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid runner search params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	merr = merr.Append(params.Request.Validate(strfmt.Default))

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// Search searches all the runners using Query DSL
func Search(params SearchParams) (*models.RunnerOverview, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.API.V1API.PlatformInfrastructure.SearchRunners(
		platform_infrastructure.NewSearchRunnersParams().
			WithContext(api.WithRegion(context.Background(), params.Region)).
			WithBody(&params.Request),
		params.AuthWriter,
	)
	if err != nil {
		return nil, api.UnwrapError(err)
	}
	return res.Payload, nil
}
