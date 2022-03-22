/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/v2/jobqueue/v1

import (
	"net/http"

	"github.com/openshift-online/ocm-sdk-go/v2/errors"
)

// JobsServer represents the interface the manages the 'jobs' resource.
type JobsServer interface {

	// Job returns the target 'job' server for the given identifier.
	//
	// jobs' operations (success, failure)
	Job(id string) JobServer
}

// dispatchJobs navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchJobs(w http.ResponseWriter, r *http.Request, server JobsServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	}
	switch segments[0] {
	default:
		target := server.Job(segments[0])
		if target == nil {
			errors.SendNotFound(w, r)
			return
		}
		dispatchJob(w, r, target, segments[1:])
	}
}
