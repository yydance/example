package entity

import "demo-dashboard/internal/conf"

var (
	AdminURL    = conf.ApisixConfig.AdminAPI + "/apisix/admin"
	UpstreamURL = AdminURL + "/upstreams"
	ServiceURL  = AdminURL + "/services"
	RouteURL    = AdminURL + "/routes"
)

type BaseInfo struct {
}

type Timeout struct {
	Connect float32 `json:"connect,omitempty"`
	Send    float32 `json:"send,omitempty"`
	Read    float32 `json:"read,omitempty"`
}

type UpstreamTLS struct {
	ClientCert string `json:"client_cert,omitempty"`
	ClientKey  string `json:"client_key,omitempty"`
}

type UpstreamKeepalivePool struct {
	Size        int `json:"size,omitempty"`
	IdleTimeout int `json:"idle_timeout,omitempty"`
	Requests    int `json:"requests,omitempty"`
}

type Upstream struct {
	Name          string                 `json:"name,omitempty"`
	Type          string                 `json:"type,omitempty"`
	Nodes         any                    `json:"nodes,omitempty"`
	Checks        any                    `json:"checks,omitempty"`
	ServiceName   string                 `json:"service_name,omitempty"`
	DiscoveryType string                 `json:"discovery_type,omitempty"`
	HashOn        string                 `json:"hash_on,omitempty"`
	Key           string                 `json:"key,omitempty"`
	PassHost      string                 `json:"pass_host,omitempty"`
	UpstreamHost  string                 `json:"upstream_host,omitempty"`
	Scheme        string                 `json:"scheme,omitempty"`
	Labels        map[string]string      `json:"labels,omitempty"`
	Retries       int                    `json:"retries,omitempty"`
	RetryTimeout  float32                `json:"retry_timeout,omitempty"`
	Timeout       *Timeout               `json:"timeout,omitempty"`
	Desc          string                 `json:"desc,omitempty"`
	TLS           *UpstreamTLS           `json:"tls,omitempty"`
	KeepalivePool *UpstreamKeepalivePool `json:"keepalive_pool,omitempty"`
}
