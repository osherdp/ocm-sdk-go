/*
Copyright (c) 2019 Red Hat, Inc.

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

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/clustersmgmt/v1

// GitlabIdentityProvider represents the values of the 'gitlab_identity_provider' type.
//
// Details for `gitlab` identity providers.
type GitlabIdentityProvider struct {
	ca           *string
	clientID     *string
	clientSecret *string
	url          *string
}

// CA returns the value of the 'CA' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Optional trusted certificate authority bundle to use when making requests tot he server.
func (o *GitlabIdentityProvider) CA() string {
	if o != nil && o.ca != nil {
		return *o.ca
	}
	return ""
}

// GetCA returns the value of the 'CA' attribute and
// a flag indicating if the attribute has a value.
//
// Optional trusted certificate authority bundle to use when making requests tot he server.
func (o *GitlabIdentityProvider) GetCA() (value string, ok bool) {
	ok = o != nil && o.ca != nil
	if ok {
		value = *o.ca
	}
	return
}

// ClientID returns the value of the 'client_ID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Client identifier of a registered _GitLab_ OAuth application.
func (o *GitlabIdentityProvider) ClientID() string {
	if o != nil && o.clientID != nil {
		return *o.clientID
	}
	return ""
}

// GetClientID returns the value of the 'client_ID' attribute and
// a flag indicating if the attribute has a value.
//
// Client identifier of a registered _GitLab_ OAuth application.
func (o *GitlabIdentityProvider) GetClientID() (value string, ok bool) {
	ok = o != nil && o.clientID != nil
	if ok {
		value = *o.clientID
	}
	return
}

// ClientSecret returns the value of the 'client_secret' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Client secret issued by _GitLab_.
func (o *GitlabIdentityProvider) ClientSecret() string {
	if o != nil && o.clientSecret != nil {
		return *o.clientSecret
	}
	return ""
}

// GetClientSecret returns the value of the 'client_secret' attribute and
// a flag indicating if the attribute has a value.
//
// Client secret issued by _GitLab_.
func (o *GitlabIdentityProvider) GetClientSecret() (value string, ok bool) {
	ok = o != nil && o.clientSecret != nil
	if ok {
		value = *o.clientSecret
	}
	return
}

// URL returns the value of the 'URL' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// URL of the _GitLab_ instance.
func (o *GitlabIdentityProvider) URL() string {
	if o != nil && o.url != nil {
		return *o.url
	}
	return ""
}

// GetURL returns the value of the 'URL' attribute and
// a flag indicating if the attribute has a value.
//
// URL of the _GitLab_ instance.
func (o *GitlabIdentityProvider) GetURL() (value string, ok bool) {
	ok = o != nil && o.url != nil
	if ok {
		value = *o.url
	}
	return
}

// GitlabIdentityProviderList is a list of values of the 'gitlab_identity_provider' type.
type GitlabIdentityProviderList struct {
	items []*GitlabIdentityProvider
}

// Len returns the length of the list.
func (l *GitlabIdentityProviderList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *GitlabIdentityProviderList) Slice() []*GitlabIdentityProvider {
	var slice []*GitlabIdentityProvider
	if l == nil {
		slice = make([]*GitlabIdentityProvider, 0)
	} else {
		slice = make([]*GitlabIdentityProvider, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *GitlabIdentityProviderList) Each(f func(item *GitlabIdentityProvider) bool) {
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
func (l *GitlabIdentityProviderList) Range(f func(index int, item *GitlabIdentityProvider) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
