package vnc

import (
	"net"
	"strconv"
	"time"

	"github.com/praetorian-inc/fingerprintx/pkg/plugins"
	"github.com/praetorian-inc/fingerprintx/pkg/plugins/services/vnc"
)

// Client is a minimal VNC client for nuclei scripts.
type Client struct{}

// IsVNCResponse is the response from the IsVNC function.
type IsVNCResponse struct {
	IsVNC  bool
	Banner string
}

// IsVNC checks if a host is running a VNC server.
// It returns a boolean indicating if the host is running a VNC server
// and the banner of the VNC server.
func (c *Client) IsVNC(host string, port int) (IsVNCResponse, error) {
	resp := IsVNCResponse{}

	timeout := 5 * time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, strconv.Itoa(port)), timeout)
	if err != nil {
		return resp, err
	}
	defer conn.Close()

	vncPlugin := vnc.VNCPlugin{}
	service, err := vncPlugin.Run(conn, timeout, plugins.Target{Host: host})
	if err != nil {
		return resp, err
	}
	resp.Banner = service.Version
	resp.IsVNC = true
	return resp, nil
}
