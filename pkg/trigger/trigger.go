//
// Copyright (c) 2019 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package trigger

import (
	"github.com/edgexfoundry/app-functions-sdk-go/pkg/common"
	logger "github.com/edgexfoundry/go-mod-core-contracts/clients/logging"
)

// ITrigger interface is used to hold event data and allow function to
type ITrigger interface {
	// Initialize performs post creation initializations
	Initialize(logger.LoggingClient) error

	// function to call to get current configuration for function
	GetConfiguration() common.ConfigurationStruct
	// function to call to retrieve data from input
	GetData() interface{}

	Complete(string)
}