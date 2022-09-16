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

// SubscriptionNotify represents the values of the 'subscription_notify' type.
//
// This struct is a request to send a templated email to a user related to this
// subscription.
type SubscriptionNotify struct {
	bitmap_                 uint32
	bccAddress              string
	clusterID               string
	clusterUUID             string
	subject                 string
	subscriptionID          string
	templateName            string
	templateParameters      []*TemplateParameter
	includeRedHatAssociates bool
	internalOnly            bool
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *SubscriptionNotify) Empty() bool {
	return o == nil || o.bitmap_ == 0
}

// BccAddress returns the value of the 'bcc_address' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The BCC address to be included on the email that is sent
func (o *SubscriptionNotify) BccAddress() string {
	if o != nil && o.bitmap_&1 != 0 {
		return o.bccAddress
	}
	return ""
}

// GetBccAddress returns the value of the 'bcc_address' attribute and
// a flag indicating if the attribute has a value.
//
// The BCC address to be included on the email that is sent
func (o *SubscriptionNotify) GetBccAddress() (value string, ok bool) {
	ok = o != nil && o.bitmap_&1 != 0
	if ok {
		value = o.bccAddress
	}
	return
}

// ClusterID returns the value of the 'cluster_ID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Indicates which Cluster (internal id) the resource type belongs to
func (o *SubscriptionNotify) ClusterID() string {
	if o != nil && o.bitmap_&2 != 0 {
		return o.clusterID
	}
	return ""
}

// GetClusterID returns the value of the 'cluster_ID' attribute and
// a flag indicating if the attribute has a value.
//
// Indicates which Cluster (internal id) the resource type belongs to
func (o *SubscriptionNotify) GetClusterID() (value string, ok bool) {
	ok = o != nil && o.bitmap_&2 != 0
	if ok {
		value = o.clusterID
	}
	return
}

// ClusterUUID returns the value of the 'cluster_UUID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Indicates which Cluster (external id) the resource type belongs to
func (o *SubscriptionNotify) ClusterUUID() string {
	if o != nil && o.bitmap_&4 != 0 {
		return o.clusterUUID
	}
	return ""
}

// GetClusterUUID returns the value of the 'cluster_UUID' attribute and
// a flag indicating if the attribute has a value.
//
// Indicates which Cluster (external id) the resource type belongs to
func (o *SubscriptionNotify) GetClusterUUID() (value string, ok bool) {
	ok = o != nil && o.bitmap_&4 != 0
	if ok {
		value = o.clusterUUID
	}
	return
}

// IncludeRedHatAssociates returns the value of the 'include_red_hat_associates' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Indicates whether to include red hat associates in the email notification
func (o *SubscriptionNotify) IncludeRedHatAssociates() bool {
	if o != nil && o.bitmap_&8 != 0 {
		return o.includeRedHatAssociates
	}
	return false
}

// GetIncludeRedHatAssociates returns the value of the 'include_red_hat_associates' attribute and
// a flag indicating if the attribute has a value.
//
// Indicates whether to include red hat associates in the email notification
func (o *SubscriptionNotify) GetIncludeRedHatAssociates() (value bool, ok bool) {
	ok = o != nil && o.bitmap_&8 != 0
	if ok {
		value = o.includeRedHatAssociates
	}
	return
}

// InternalOnly returns the value of the 'internal_only' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Indicates whether the service log is internal only to RH
func (o *SubscriptionNotify) InternalOnly() bool {
	if o != nil && o.bitmap_&16 != 0 {
		return o.internalOnly
	}
	return false
}

// GetInternalOnly returns the value of the 'internal_only' attribute and
// a flag indicating if the attribute has a value.
//
// Indicates whether the service log is internal only to RH
func (o *SubscriptionNotify) GetInternalOnly() (value bool, ok bool) {
	ok = o != nil && o.bitmap_&16 != 0
	if ok {
		value = o.internalOnly
	}
	return
}

// Subject returns the value of the 'subject' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The email subject
func (o *SubscriptionNotify) Subject() string {
	if o != nil && o.bitmap_&32 != 0 {
		return o.subject
	}
	return ""
}

// GetSubject returns the value of the 'subject' attribute and
// a flag indicating if the attribute has a value.
//
// The email subject
func (o *SubscriptionNotify) GetSubject() (value string, ok bool) {
	ok = o != nil && o.bitmap_&32 != 0
	if ok {
		value = o.subject
	}
	return
}

// SubscriptionID returns the value of the 'subscription_ID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Indicates which Subscription the resource type belongs to
func (o *SubscriptionNotify) SubscriptionID() string {
	if o != nil && o.bitmap_&64 != 0 {
		return o.subscriptionID
	}
	return ""
}

// GetSubscriptionID returns the value of the 'subscription_ID' attribute and
// a flag indicating if the attribute has a value.
//
// Indicates which Subscription the resource type belongs to
func (o *SubscriptionNotify) GetSubscriptionID() (value string, ok bool) {
	ok = o != nil && o.bitmap_&64 != 0
	if ok {
		value = o.subscriptionID
	}
	return
}

// TemplateName returns the value of the 'template_name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The name of the template used to construct the email contents
func (o *SubscriptionNotify) TemplateName() string {
	if o != nil && o.bitmap_&128 != 0 {
		return o.templateName
	}
	return ""
}

// GetTemplateName returns the value of the 'template_name' attribute and
// a flag indicating if the attribute has a value.
//
// The name of the template used to construct the email contents
func (o *SubscriptionNotify) GetTemplateName() (value string, ok bool) {
	ok = o != nil && o.bitmap_&128 != 0
	if ok {
		value = o.templateName
	}
	return
}

// TemplateParameters returns the value of the 'template_parameters' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The values which will be substituted into the templated email
func (o *SubscriptionNotify) TemplateParameters() []*TemplateParameter {
	if o != nil && o.bitmap_&256 != 0 {
		return o.templateParameters
	}
	return nil
}

// GetTemplateParameters returns the value of the 'template_parameters' attribute and
// a flag indicating if the attribute has a value.
//
// The values which will be substituted into the templated email
func (o *SubscriptionNotify) GetTemplateParameters() (value []*TemplateParameter, ok bool) {
	ok = o != nil && o.bitmap_&256 != 0
	if ok {
		value = o.templateParameters
	}
	return
}

// SubscriptionNotifyListKind is the name of the type used to represent list of objects of
// type 'subscription_notify'.
const SubscriptionNotifyListKind = "SubscriptionNotifyList"

// SubscriptionNotifyListLinkKind is the name of the type used to represent links to list
// of objects of type 'subscription_notify'.
const SubscriptionNotifyListLinkKind = "SubscriptionNotifyListLink"

// SubscriptionNotifyNilKind is the name of the type used to nil lists of objects of
// type 'subscription_notify'.
const SubscriptionNotifyListNilKind = "SubscriptionNotifyListNil"

// SubscriptionNotifyList is a list of values of the 'subscription_notify' type.
type SubscriptionNotifyList struct {
	href  string
	link  bool
	items []*SubscriptionNotify
}

// Len returns the length of the list.
func (l *SubscriptionNotifyList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *SubscriptionNotifyList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *SubscriptionNotifyList) Get(i int) *SubscriptionNotify {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *SubscriptionNotifyList) Slice() []*SubscriptionNotify {
	var slice []*SubscriptionNotify
	if l == nil {
		slice = make([]*SubscriptionNotify, 0)
	} else {
		slice = make([]*SubscriptionNotify, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *SubscriptionNotifyList) Each(f func(item *SubscriptionNotify) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *SubscriptionNotifyList) Range(f func(index int, item *SubscriptionNotify) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
