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

package v1 // github.com/renan-campos/ocm-sdk-go/accountsmgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/renan-campos/ocm-sdk-go/helpers"
)

// MarshalCloudAccount writes a value of the 'cloud_account' type to the given writer.
func MarshalCloudAccount(object *CloudAccount, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeCloudAccount(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// writeCloudAccount writes a value of the 'cloud_account' type to the given stream.
func writeCloudAccount(object *CloudAccount, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = object.bitmap_&1 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("cloud_account_id")
		stream.WriteString(object.cloudAccountID)
		count++
	}
	present_ = object.bitmap_&2 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("cloud_provider_id")
		stream.WriteString(object.cloudProviderID)
	}
	stream.WriteObjectEnd()
}

// UnmarshalCloudAccount reads a value of the 'cloud_account' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalCloudAccount(source interface{}) (object *CloudAccount, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readCloudAccount(iterator)
	err = iterator.Error
	return
}

// readCloudAccount reads a value of the 'cloud_account' type from the given iterator.
func readCloudAccount(iterator *jsoniter.Iterator) *CloudAccount {
	object := &CloudAccount{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "cloud_account_id":
			value := iterator.ReadString()
			object.cloudAccountID = value
			object.bitmap_ |= 1
		case "cloud_provider_id":
			value := iterator.ReadString()
			object.cloudProviderID = value
			object.bitmap_ |= 2
		default:
			iterator.ReadAny()
		}
	}
	return object
}
