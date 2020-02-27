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

import (
	"io"
	"net/http"
)

func readAccessReviewPostRequest(request *AccessReviewPostServerRequest, r *http.Request) error {
	var err error
	request.request, err = UnmarshalAccessReviewRequest(r)
	return err
}
func writeAccessReviewPostRequest(request *AccessReviewPostRequest, writer io.Writer) error {
	return MarshalAccessReviewRequest(request.request, writer)
}
func readAccessReviewPostResponse(response *AccessReviewPostResponse, reader io.Reader) error {
	var err error
	response.response, err = UnmarshalAccessReviewResponse(reader)
	return err
}
func writeAccessReviewPostResponse(response *AccessReviewPostServerResponse, w http.ResponseWriter) error {
	return MarshalAccessReviewResponse(response.response, w)
}
