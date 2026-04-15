package plugin

type RegisterFileSystemPayload struct {
	MapData []RegisterFileSystemMapData `json:"map_data" yaml:"map_data"`
}
type RegisterFileSystemMapData struct {
	SystemPath  string `json:"system_path" yaml:"system_path"`
	SandboxPath string `json:"sandbox_path" yaml:"sandbox_path"`
}

type RegisterWebAccessPayload struct {
	AllowedHosts  []string `json:"allowed_hosts" yaml:"allowed_hosts"`
	AllowAllHosts bool     `json:"allow_all_hosts" yaml:"allow_all_hosts"` // overwrite allowed hosts
}

type RegisterNetworkPayload struct {
	AllowedHosts  []RegisterNetworkHostPayload `json:"allowed_hosts,omitempty" yaml:"allowed_hosts"`     // host:port
	AllowAllHosts bool                         `json:"allow_all_hosts,omitempty" yaml:"allow_all_hosts"` // overwrite allowed hosts, if you apply this true, user will be prompted to a warning.
}

type RegisterNetworkHostPayload struct {
	Host     string
	Port     uint16
	Protocol RegisterNetworkHostProtocol
}

type RegisterNetworkHostProtocol int

const (
	RegisterNetworkHostProtocolTCP RegisterNetworkHostProtocol = 1 << iota
	RegisterNetworkHostProtocolUDP
	RegisterNetworkHostProtocolICMP
	RegisterNetworkHostProtocolIGMP
	RegisterNetworkHostProtocolUnix // available in macOS and linux
)

type RegisterResourcePayload struct {
	MaxCPU           int // Max mCPUs
	MaxMem           int // ? MiB
	MaxTimeoutSecond int // max timeout
}
