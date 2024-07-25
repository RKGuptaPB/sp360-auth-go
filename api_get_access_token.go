/*
Auth APIs

This is the API to manage credentials and generate token

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sp360auth

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// GetAccessTokenAPIService GetAccessTokenAPI service
type GetAccessTokenAPIService service

type ApiGetAccessTokenRequest struct {
	ctx context.Context
	ApiService *GetAccessTokenAPIService
	contentType *string
	grantType *string
	scope *string
}

// Set this to- &lt;b&gt;&#39;application/x-www-form-urlencoded&#39;
func (r ApiGetAccessTokenRequest) ContentType(contentType string) ApiGetAccessTokenRequest {
	r.contentType = &contentType
	return r
}

func (r ApiGetAccessTokenRequest) GrantType(grantType string) ApiGetAccessTokenRequest {
	r.grantType = &grantType
	return r
}

func (r ApiGetAccessTokenRequest) Scope(scope string) ApiGetAccessTokenRequest {
	r.scope = &scope
	return r
}

func (r ApiGetAccessTokenRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetAccessTokenExecute(r)
}

/*
GetAccessToken Get access token

Each request to the PB PitneyShip APIs requires authentication via an OAuth token. This API call generates the OAuth token based on the Base64-encoded value of the API key and secret associated with your Pitney Bowes developer account. The token expires after 10 hours,means you must create a new one.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAccessTokenRequest
*/
func (a *GetAccessTokenAPIService) GetAccessToken(ctx context.Context) ApiGetAccessTokenRequest {
	return ApiGetAccessTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *GetAccessTokenAPIService) GetAccessTokenExecute(r ApiGetAccessTokenRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GetAccessTokenAPIService.GetAccessToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contentType == nil {
		return localVarReturnValue, nil, reportError("contentType is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
	if r.grantType != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "grant_type", r.grantType, "")
	}
	if r.scope != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "scope", r.scope, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
