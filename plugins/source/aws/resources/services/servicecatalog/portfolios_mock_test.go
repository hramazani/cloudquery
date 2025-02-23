package servicecatalog

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildPortfolios(t *testing.T, ctrl *gomock.Controller) client.Services {
	mk := mocks.NewMockServicecatalogClient(ctrl)
	ma := mocks.NewMockServicecatalogappregistryClient(ctrl)

	o := servicecatalog.ListPortfoliosOutput{}
	require.NoError(t, faker.FakeObject(&o))

	o.NextPageToken = nil

	mk.EXPECT().ListPortfolios(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&o,
		nil,
	)

	to := servicecatalogappregistry.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&to))

	ma.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&to,
		nil,
	)

	return client.Services{
		Servicecatalog:            mk,
		Servicecatalogappregistry: ma,
	}
}

func TestPortfolios(t *testing.T) {
	client.AwsMockTestHelper(t, Portfolios(), buildPortfolios, client.TestOptions{})
}
