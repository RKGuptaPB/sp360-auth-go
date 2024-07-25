# \GetAccessTokenAPI

All URIs are relative to *https://api-dev.sendpro360.pitneycloud.com/auth*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccessToken**](GetAccessTokenAPI.md#GetAccessToken) | **Post** /api/v1/token | Get access token



## GetAccessToken

> map[string]interface{} GetAccessToken(ctx).ContentType(contentType).GrantType(grantType).Scope(scope).Execute()

Get access token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RKGuptaPB/sp360-auth-go"
)

func main() {
	contentType := "contentType_example" // string | Set this to- <b>'application/x-www-form-urlencoded'
	grantType := "grantType_example" // string |  (optional)
	scope := "scope_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetAccessTokenAPI.GetAccessToken(context.Background()).ContentType(contentType).GrantType(grantType).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetAccessTokenAPI.GetAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessToken`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GetAccessTokenAPI.GetAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** | Set this to- &lt;b&gt;&#39;application/x-www-form-urlencoded&#39; | 
 **grantType** | **string** |  | 
 **scope** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

