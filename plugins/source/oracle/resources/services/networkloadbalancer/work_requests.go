package networkloadbalancer

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/networkloadbalancer"
)

func WorkRequests() *schema.Table {
	return &schema.Table{
		Name:      "oracle_networkloadbalancer_work_requests",
		Resolver:  fetchWorkRequests,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&networkloadbalancer.WorkRequestSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}

func fetchWorkRequests(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	var page *string
	for {
		request := networkloadbalancer.ListWorkRequestsRequest{
			CompartmentId: common.String(cqClient.CompartmentOcid),
			Page:          page,
		}

		response, err := cqClient.OracleClients[cqClient.Region].NetworkloadbalancerNetworkloadbalancerClient.ListWorkRequests(ctx, request)

		if err != nil {
			return err
		}

		res <- response.Items

		if response.OpcNextPage == nil {
			break
		}

		page = response.OpcNextPage
	}

	return nil
}
