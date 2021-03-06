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

package depresourceapi

import (
	"io"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

// ParseElasticsearchInputParams is consumed by ParseElasticsearchInput.
type ParseElasticsearchInputParams struct {
	*api.API
	NewElasticsearchParams

	Payload          *models.ElasticsearchPayload
	TopologyElements []string
	Size, ZoneCount  int32
	Writer           io.Writer
}

// Validate ensures the parameters are usable by the consuming function.
func (params *ParseElasticsearchInputParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid deployment resource params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	return merr.ErrorOrNil()
}

// ParseElasticsearchInput handles all but the API parameters as optional,
// providing a nicer API when it's used. The bulk of what it does is:
// * If a Payload is already specified, it returns it.
// * Tries to create an []ElasticsearchTopologyElement from a raw []string.
// * If the previous step returns an empty slice, it uses a default slice which
//   might override the values when Size or ZoneCount are set in the params.
// * Auto-discovers the latest Stack version if Version is not specified.
// When all of those steps are done, it finally calls NewElasticsearch building
// the resulting ElasticsearchPayload.
func ParseElasticsearchInput(params ParseElasticsearchInputParams) (*models.ElasticsearchPayload, error) {
	if params.Payload != nil {
		return params.Payload, nil
	}

	if err := params.Validate(); err != nil {
		return nil, err
	}

	topology, err := NewElasticsearchTopology(params.TopologyElements)
	if err != nil {
		return nil, err
	}

	// On empty topology, use the default one with the size & count specified
	if len(topology) == 0 {
		topology = append(topology, NewElasticsearchTopologyElement(
			params.Size, params.ZoneCount,
		))
	}

	// Version Discovery
	version, err := LatestStackVersion(LatestStackVersionParams{
		Writer:  params.Writer,
		API:     params.API,
		Version: params.Version,
		Region:  params.Region,
	})
	if err != nil {
		return nil, err
	}

	var NewEsparams = params.NewElasticsearchParams
	NewEsparams.Version = version
	NewEsparams.Topology = topology

	return NewElasticsearch(NewEsparams)
}
