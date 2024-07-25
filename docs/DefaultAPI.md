# \DefaultAPI

All URIs are relative to *https://api-dev.sendpro360.pitneycloud.com/auth*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1TokenOptions**](DefaultAPI.md#ApiV1TokenOptions) | **Options** /api/v1/token | 



## ApiV1TokenOptions

> map[string]interface{} ApiV1TokenOptions(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1TokenOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1TokenOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TokenOptions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1TokenOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TokenOptionsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

