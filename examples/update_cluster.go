/*
Copyright (c) 2018 Red Hat, Inc.

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

// This example shows how to update the display name of a cluster.

package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go/v2"
	cmv1 "github.com/openshift-online/ocm-sdk-go/v2/clustersmgmt/v1"
	"github.com/openshift-online/ocm-sdk-go/v2/logging"
)

func main() {
	// Create a context:
	ctx := context.Background()

	// Create a logger that has the debug level enabled:
	logger, err := logging.NewGoLoggerBuilder().
		Debug(true).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build logger: %v\n", err)
		os.Exit(1)
	}

	// Create the connection, and remember to close it:
	token := os.Getenv("OCM_TOKEN")
	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Tokens(token).
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Get the client for the resource that manages the collection of clusters:
	collection := connection.ClustersMgmt().V1().Clusters()

	// Get the client for the resource that manages the cluster that we want to update. Note
	// that this will not yet send any request to the server, so it will succeed even if the
	// cluster doesn't exist.
	resource := collection.Cluster("1BDFg66jv2kDfBh6bBog3IsZWVH")

	// Prepare the patch to send:
	patch, err := cmv1.NewCluster().
		DisplayName("My cluster").
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create cluster patch: %v\n", err)
		os.Exit(1)
	}

	// Send the request to update the cluster:
	_, err = resource.Update().
		Body(patch).
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't update cluster: %v\n", err)
		os.Exit(1)
	}
}
