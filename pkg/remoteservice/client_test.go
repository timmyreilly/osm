package remoteservice

import (
	"testing"

	tassert "github.com/stretchr/testify/assert"
	fakeRemoteServiceClient "github.com/openservicemesh/osm/pkg/gen/client/config/clientset/versioned/fake"
)

func TestNewRemoteServiceClient(t *testing.T) {
	assert := tassert.New(t)

	client, err := newRemoteServiceClient(fakeRemoteServiceClient.NewSimpleClientset(), nil, nil)
	assert.Nil(err)
	assert.NotNil(client)
}
