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

package v1 // github.com/openshift-online/ocm-sdk-go/authorizations/v1

// ExportControlReviewResponseBuilder contains the data and logic needed to build 'export_control_review_response' objects.
//
//
type ExportControlReviewResponseBuilder struct {
	restricted *bool
}

// NewExportControlReviewResponse creates a new builder of 'export_control_review_response' objects.
func NewExportControlReviewResponse() *ExportControlReviewResponseBuilder {
	return new(ExportControlReviewResponseBuilder)
}

// Restricted sets the value of the 'restricted' attribute to the given value.
//
//
func (b *ExportControlReviewResponseBuilder) Restricted(value bool) *ExportControlReviewResponseBuilder {
	b.restricted = &value
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ExportControlReviewResponseBuilder) Copy(object *ExportControlReviewResponse) *ExportControlReviewResponseBuilder {
	if object == nil {
		return b
	}
	b.restricted = object.restricted
	return b
}

// Build creates a 'export_control_review_response' object using the configuration stored in the builder.
func (b *ExportControlReviewResponseBuilder) Build() (object *ExportControlReviewResponse, err error) {
	object = new(ExportControlReviewResponse)
	object.restricted = b.restricted
	return
}
