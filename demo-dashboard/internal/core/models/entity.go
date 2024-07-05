package models

import (
	"database/sql/driver"
	"demo-dashboard/internal/log"
	"encoding/json"

	"github.com/lib/pq"
)

type JSON json.RawMessage

func (j *JSON) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		log.Logger.Error("Failed to unmarshall JSON value: ", value)
	}
	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

type BaseInfo struct {
	ID         any   `json:"id" gorm:"type:varchar(64);not null"`
	CreateTime int64 `json:"create_time,omitempty" gorm:"type:varchar(64)"`
	UpdateTime int64 `json:"update_time,omitempty" gorm:"type:varchar(64)"`
}

type Status uint8

// swagger:model Route
type Route struct {
	BaseInfo
	URI             string            `json:"uri,omitempty" gorm:"type:varchar(64)"`
	Uris            pq.StringArray    `json:"uris,omitempty" gorm:"type:varchar(64)[]"`
	Name            string            `json:"name" gorm:"type:varchar(128)"`
	Desc            string            `json:"desc,omitempty" gorm:"type:text"`
	Priority        int               `json:"priority,omitempty" gorm:"type:tinyint;unsigned"`
	Methods         pq.StringArray    `json:"methods,omitempty" gorm:"type:varchar(32)[]"`
	Host            string            `json:"host,omitempty" gorm:"type:varchar(64)"`
	Hosts           pq.StringArray    `json:"hosts,omitempty" gorm:"type:varchar(64)[]"`
	RemoteAddr      string            `json:"remote_addr,omitempty" gorm:"type:varchar(64)"`
	RemoteAddrs     pq.StringArray    `json:"remote_addrs,omitempty" gorm:"type:varchar(64)[]"`
	Vars            pq.GenericArray   `json:"vars,omitempty" gorm:"type:varchar(64)[]"`
	FilterFunc      string            `json:"filter_func,omitempty" gorm:"type:varchar(64)"`
	Script          any               `json:"script,omitempty" gorm:"type:varchar(128)"`
	ScriptID        any               `json:"script_id,omitempty" gorm:"type:varchar(128)"` // For debug and optimization(cache), currently same as Route's ID
	Plugins         map[string]any    `json:"plugins,omitempty" gorm:"type:json"`
	PluginConfigID  any               `json:"plugin_config_id,omitempty" gorm:"type:varchar(64)"`
	Upstream        *UpstreamDef      `json:"upstream,omitempty"`
	ServiceID       any               `json:"service_id,omitempty"`
	UpstreamID      any               `json:"upstream_id,omitempty"`
	ServiceProtocol string            `json:"service_protocol,omitempty" gorm:"type:varchar(64)"`
	Labels          map[string]string `json:"labels,omitempty" gorm:"type:json"`
	EnableWebsocket bool              `json:"enable_websocket,omitempty" gorm:"type:bool"`
	Status          Status            `json:"status" gorm:"type:tinyint;unsigned"`
}

func (Route) TableName() string {
	return "tb_routes"
}

// --- structures for upstream start  ---
type TimeoutValue float32

func (t *TimeoutValue) Scan(value any) error {
	float, ok := value.(float32)
	if !ok {
		log.Logger.Errorf("failed to get float value: %v", value)
	}
	*t = TimeoutValue(float)
	return nil
}
func (t TimeoutValue) Value() (driver.Value, error) {
	return t, nil
}

type Timeout struct {
	UpstreamID uint
	Connect    float32 `json:"connect,omitempty"`
	Send       float32 `json:"send,omitempty"`
	Read       float32 `json:"read,omitempty"`
}

func (Timeout) TableName() string {
	return "tb_upstream_timeout"
}

type Node struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Weight   int    `json:"weight"`
	Metadata any    `json:"metadata,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

type K8sInfo struct {
	Namespace   string `json:"namespace,omitempty"`
	DeployName  string `json:"deploy_name,omitempty"`
	ServiceName string `json:"service_name,omitempty"`
	Port        int    `json:"port,omitempty"`
	BackendType string `json:"backend_type,omitempty"`
}

type Healthy struct {
	Interval     int   `json:"interval,omitempty"`
	HttpStatuses []int `json:"http_statuses,omitempty"`
	Successes    int   `json:"successes,omitempty"`
}

type UnHealthy struct {
	Interval     int   `json:"interval,omitempty"`
	HTTPStatuses []int `json:"http_statuses,omitempty"`
	TCPFailures  int   `json:"tcp_failures,omitempty"`
	Timeouts     int   `json:"timeouts,omitempty"`
	HTTPFailures int   `json:"http_failures,omitempty"`
}

type Active struct {
	Type                   string       `json:"type,omitempty"`
	Timeout                TimeoutValue `json:"timeout,omitempty"`
	Concurrency            int          `json:"concurrency,omitempty"`
	Host                   string       `json:"host,omitempty"`
	Port                   int          `json:"port,omitempty"`
	HTTPPath               string       `json:"http_path,omitempty"`
	HTTPSVerifyCertificate bool         `json:"https_verify_certificate,omitempty"`
	Healthy                Healthy      `json:"healthy,omitempty"`
	UnHealthy              UnHealthy    `json:"unhealthy,omitempty"`
	ReqHeaders             []string     `json:"req_headers,omitempty"`
}

type Passive struct {
	Type      string    `json:"type,omitempty"`
	Healthy   Healthy   `json:"healthy,omitempty"`
	UnHealthy UnHealthy `json:"unhealthy,omitempty"`
}

type HealthChecker struct {
	Active  Active  `json:"active,omitempty"`
	Passive Passive `json:"passive,omitempty"`
}

type UpstreamTLS struct {
	UpstreamID uint
	ClientCert string `json:"client_cert,omitempty"`
	ClientKey  string `json:"client_key,omitempty"`
}

func (UpstreamTLS) TableName() string {
	return "tb_upstreamtls"
}

type UpstreamKeepalivePool struct {
	UpstreamID  uint
	IdleTimeout *float32 `json:"idle_timeout,omitempty" gorm:"type:float"`
	Requests    int      `json:"requests,omitempty" gorm:"type:smallint;unsigned"`
	Size        int      `json:"size" gorm:"type:smallint;unsigned"`
}

func (UpstreamKeepalivePool) TableName() string {
	return "tb_upstream_keepalivepool"
}

type UpstreamDef struct {
	ServiceID     uint
	Nodes         any                    `json:"nodes,omitempty" gorm:"type:longtext"`
	Retries       *int                   `json:"retries,omitempty" gorm:"type:tinyint;unsigned"`
	Timeout       *Timeout               `json:"timeout,omitempty" gorm:"foreignKey:UpstreamID"`
	Type          string                 `json:"type,omitempty" gorm:"type:varchar(64)"`
	Checks        any                    `json:"checks,omitempty" gorm:"type:longtext"`
	HashOn        string                 `json:"hash_on,omitempty" gorm:"type:varchar(128)"`
	Key           string                 `json:"key,omitempty" gorm:"type:varchar(64)"`
	Scheme        string                 `json:"scheme,omitempty" gorm:"type:varchar(32)"`
	DiscoveryType string                 `json:"discovery_type,omitempty" gorm:"type:varchar(64)"`
	DiscoveryArgs map[string]any         `json:"discovery_args,omitempty" gorm:"type:json"`
	PassHost      string                 `json:"pass_host,omitempty" gorm:"type:varchar(128)"`
	UpstreamHost  string                 `json:"upstream_host,omitempty" gorm:"type:varchar(128)"`
	Name          string                 `json:"name,omitempty" gorm:"type:varchar(256)"`
	Desc          string                 `json:"desc,omitempty" gorm:"type:varchar(256)"`
	ServiceName   string                 `json:"service_name,omitempty" gorm:"type:varchar(64)"`
	Labels        map[string]string      `json:"labels,omitempty" gorm:"type:json"`
	TLS           *UpstreamTLS           `json:"tls,omitempty" gorm:"foreignKey:UpstreamID"`
	KeepalivePool *UpstreamKeepalivePool `json:"keepalive_pool,omitempty" gorm:"foreignKey:UpstreamID"`
	RetryTimeout  float32                `json:"retry_timeout,omitempty" gorm:"type:float"`
}

// swagger:model Upstream
type Upstream struct {
	BaseInfo
	UpstreamDef
}

func (Upstream) TableName() string {
	return "tb_upstream"
}

type UpstreamNameResponse struct {
	ID   any    `json:"id"`
	Name string `json:"name"`
}

// --- structures for upstream end  ---

// swagger:model Consumer
type Consumer struct {
	Username   string            `json:"username" gorm:"type:varchar(128)"`
	Desc       string            `json:"desc,omitempty" gorm:"type:varchar(256)"`
	Plugins    map[string]any    `json:"plugins,omitempty" gorm:"type:json"`
	Labels     map[string]string `json:"labels,omitempty" gorm:"type:json"`
	CreateTime int64             `json:"create_time,omitempty"`
	UpdateTime int64             `json:"update_time,omitempty"`
}

func (Consumer) TableName() string {
	return "tb_consumer"
}

type SSLClient struct {
	CA    string `json:"ca,omitempty"`
	Depth int    `json:"depth,omitempty"`
}

// swagger:model SSL
type SSL struct {
	BaseInfo
	Cert          string            `json:"cert,omitempty"`
	Key           string            `json:"key,omitempty"`
	Sni           string            `json:"sni,omitempty"`
	Snis          []string          `json:"snis,omitempty"`
	Certs         []string          `json:"certs,omitempty"`
	Keys          []string          `json:"keys,omitempty"`
	ExpTime       int64             `json:"exptime,omitempty"`
	Status        int               `json:"status"`
	ValidityStart int64             `json:"validity_start,omitempty"`
	ValidityEnd   int64             `json:"validity_end,omitempty"`
	Labels        map[string]string `json:"labels,omitempty"`
	Client        *SSLClient        `json:"client,omitempty"`
}

// swagger:model Service
type Service struct {
	BaseInfo
	Name            string            `json:"name,omitempty" gorm:"type:varchar(128)"`
	Desc            string            `json:"desc,omitempty" gorm:"type:varchar(256)"`
	Upstream        *UpstreamDef      `json:"upstream,omitempty" gorm:"foreignKey:ServiceID"`
	Plugins         map[string]any    `json:"plugins,omitempty" gorm:"type:json"`
	Script          string            `json:"script,omitempty" gorm:"type:varchar(256)"`
	Labels          map[string]string `json:"labels,omitempty" gorm:"type:json"`
	EnableWebsocket bool              `json:"enable_websocket,omitempty" gorm:"type:bool"`
	Hosts           []string          `json:"hosts,omitempty" gorm:"type:varchar(128)"`
}

type Script struct {
	ID     string `json:"id"`
	Script any    `json:"script,omitempty"`
}

type RequestValidation struct {
	Type       string   `json:"type,omitempty"`
	Required   []string `json:"required,omitempty"`
	Properties any      `json:"properties,omitempty"`
}

// swagger:model GlobalPlugins
type GlobalPlugins struct {
	BaseInfo
	Plugins map[string]any `json:"plugins" gorm:"type:json"`
}

func (GlobalPlugins) TableName() string {
	return "tb_global_plugins"
}

type ServerInfo struct {
	BaseInfo
	LastReportTime int64  `json:"last_report_time,omitempty"`
	UpTime         int64  `json:"up_time,omitempty"`
	BootTime       int64  `json:"boot_time,omitempty"`
	EtcdVersion    string `json:"etcd_version,omitempty" gorm:"type:varchar(32)"`
	Hostname       string `json:"hostname,omitempty" gorm:"type:varchar(64)"`
	Version        string `json:"version,omitempty" gorm:"type:varchar(64)"`
}

// swagger:model GlobalPlugins
type PluginConfig struct {
	BaseInfo
	Desc    string            `json:"desc,omitempty"`
	Plugins map[string]any    `json:"plugins"`
	Labels  map[string]string `json:"labels,omitempty"`
}

// swagger:model Proto
type Proto struct {
	BaseInfo
	Desc    string `json:"desc,omitempty"`
	Content string `json:"content"`
}

// swagger:model StreamRoute
type StreamRoute struct {
	BaseInfo
	Desc       string         `json:"desc,omitempty"`
	RemoteAddr string         `json:"remote_addr,omitempty"`
	ServerAddr string         `json:"server_addr,omitempty"`
	ServerPort int            `json:"server_port,omitempty"`
	SNI        string         `json:"sni,omitempty"`
	Upstream   *UpstreamDef   `json:"upstream,omitempty"`
	UpstreamID any            `json:"upstream_id,omitempty"`
	Plugins    map[string]any `json:"plugins,omitempty"`
}

// swagger:model SystemConfig
type SystemConfig struct {
	ConfigName string         `json:"config_name"`
	Desc       string         `json:"desc,omitempty"`
	Payload    map[string]any `json:"payload,omitempty"`
	CreateTime int64          `json:"create_time,omitempty"`
	UpdateTime int64          `json:"update_time,omitempty"`
}

type User struct {
	BaseInfo
	Name    string `json:"name,omitempty"`
	Status  bool   `json:"status" default:"true"`
	Type    string `json:"type,omitempty"`
	TeamsID []any  `json:"teams_id,omitempty"`
	RoleID  []any  `json:"role_id,omitempty"`
}

type Team struct {
	BaseInfo
	Name      string   `json:"name,omitempty"`
	UsersID   []any    `json:"users_id,omitempty"`
	TeamAdmin []string `json:"team_admin,omitempty"`
}

type Role struct {
	BaseInfo
	Name          string   `json:"name,omitempty"`
	Authorization string   `json:"authorization,omitempty"`
	Features      []string `json:"features,omitempty"`
}
