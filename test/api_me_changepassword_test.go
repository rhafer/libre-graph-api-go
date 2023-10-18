/*
Libre Graph API

Testing MeChangepasswordAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package libregraph

import (
	"context"
	openapiclient "github.com/owncloud/libre-graph-api-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_libregraph_MeChangepasswordAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MeChangepasswordAPIService ChangeOwnPassword", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.MeChangepasswordAPI.ChangeOwnPassword(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
