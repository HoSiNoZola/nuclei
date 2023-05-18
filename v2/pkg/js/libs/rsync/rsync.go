package rsync

import (
	"net"
	"strconv"
	"time"

	"github.com/praetorian-inc/fingerprintx/pkg/plugins"
	"github.com/praetorian-inc/fingerprintx/pkg/plugins/services/rsync"
)

// Client is a minimal Rsync client for nuclei scripts.
type Client struct{}

// IsRsyncResponse is the response from the IsRsync function.
type IsRsyncResponse struct {
	IsRsync bool
	Banner  string
}

// IsRsync checks if a host is running a Rsync server.
func (c *Client) IsRsync(host string, port int) (IsRsyncResponse, error) {
	resp := IsRsyncResponse{}

	timeout := 5 * time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, strconv.Itoa(port)), timeout)
	if err != nil {
		return resp, err
	}
	defer conn.Close()

	rsyncPlugin := rsync.RSYNCPlugin{}
	service, err := rsyncPlugin.Run(conn, timeout, plugins.Target{Host: host})
	if err != nil {
		return resp, err
	}
	resp.Banner = service.Version
	resp.IsRsync = true
	return resp, nil
}
