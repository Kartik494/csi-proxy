// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"net"

	"github.com/Microsoft/go-winio"
	"github.com/kubernetes-csi/csi-proxy/client"
	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/api/dummy/v1alpha1"
	"google.golang.org/grpc"
)

const groupName = "dummy"

var version = apiversion.NewVersionOrPanic("v1alpha1")

type Client struct {
	client     v1alpha1.DummyClient
	connection *grpc.ClientConn
}

// NewClient returns a client to make calls to the dummy API group version v1alpha1.
// It's the caller's responsibility to Close the client when done.
func NewClient() (*Client, error) {
	pipePath := client.PipePath(groupName, version)

	connection, err := grpc.Dial(pipePath,
		grpc.WithContextDialer(func(context context.Context, s string) (net.Conn, error) {
			return winio.DialPipeContext(context, s)
		}),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := v1alpha1.NewDummyClient(connection)
	return &Client{
		client:     client,
		connection: connection,
	}, nil
}

// Close closes the client. It must be called before the client gets GC-ed.
func (w *Client) Close() error {
	return w.connection.Close()
}

// ensures we implement all the required methods
var _ v1alpha1.DummyClient = &Client{}

func (w *Client) ComputeDouble(context context.Context, request *v1alpha1.ComputeDoubleRequest, opts ...grpc.CallOption) (*v1alpha1.ComputeDoubleResponse, error) {
	return w.client.ComputeDouble(context, request, opts...)
}
