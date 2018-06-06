package common

const (
	EndPoint        	= "myhuaweicloud.com"
	ServiceCCE      	= "cce"
	ServiceECS      	= "ecs"
	ServiceVPC      	= "vpc"
	DefaultCidr     	= "192.168.0.0/24"
	DefaultGateway  	= "192.168.0.1"
	DefaultVpcName 		= "default-vpc"
	DefaultSubnetName   = "default-subnet"
)

const (
	VirtualMachine = "VirtualMachine"
	BareMetal      = "BareMetal"
	Windows        = "Windows"
)

const (
	Available      = "Available"
	Unavailable    = "Unavailable"
	ScalingUp      = "ScalingUp"
	ScalingDown    = "ScalingDown"
	Creating       = "Creating"
	Deleting       = "Deleting"
	Upgrading      = "Upgrading"
	Resizing       = "Resizing"
	Empty          = "Empty"
)

// Error message
type ErrorInfo struct {
	Code        string `json:"code"`
	Description string `json:"message"`
}

type VpcSt struct {
	Name string `json:"name,omitempty"`
	Cidr string `json:"cidr,omitempty"`
}

type VpcRequest struct {
	Vpc VpcSt `json:"vpc"`
}

type SubnetSt struct {
	Name      string `json:"name"`
	Cidr      string `json:"cidr"`
	GatewayIP string `json:"gateway_ip"`
	VpcID     string `json:"vpc_id"`
}

type SubnetRequest struct {
	Subnet SubnetSt `json:"subnet"`
}

// Used to parse the Vpc response
type VpcInfo struct {
	Vpc VpcResp `json:"vpc"`
}

// Vpc response fields
type VpcResp struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Cidr string `json:"cidr"`
	Status string `json:"status"`
	NoSecurityGroup bool `json:"noSecurityGroup"`
}

// Used to parse the Subnet response
type SubnetInfo struct {
	Subnet Subnet `json:"subnet"`
}

// Subnet response fields
type Subnet struct {
	Id                	string `json:"id"`                // Specifies a resource ID in UUID format.
	Name              	string `json:"name"`              // Specifies the name of the subnet.
	Cidr              	string `json:"cidr"`              // Specifies the network segment of the subnet.
	GatewayIp        	string `json:"gateway_ip"`        // Specifies the gateway address of the subnet.
	DhcpEnable       	bool   `json:"dhcp_enable"`       // Specifies whether the DHCP function is enabled for the subnet.
	PrimaryDns       	string `json:"primary_dns"`       // Specifies the primary IP address of the DNS server on the subnet.
	SecondaryDns     	string `json:"secondary_dns"`     // Specifies the secondary IP address of the DNS server on the subnet.
	AvailabilityZone 	string `json:"availability_zone"` // Specifies the ID of the AZ to which the subnet belongs.
	VpcId            	string `json:"vpc_id"`            // Specifies the ID of the VPC to which the subnet belongs.
	Status            	string `json:"status"`            // Specifies the status of the subnet.The value can be ACTIVE, DOWN, BUILD, ERROR, or DELETE.
}

type NodeConfig struct {
	NodeFlavor      string
	AvailableZone   string
	SSHName         string
	RootVolumeSize  int64
	RootVolumeType  string
	DataVolumeSize  int64
	DataVolumeType  string
	BillingMode     int64
	NodeCount       int64
	PublicIP        PublicIP
	NodeLabels      map[string]string
}

//cluster struct
type MetaInfo struct {
	Name                 string `json:"name"`
	Uid                  string `json:"uid,omitempty"`
	CreationTimestamp    string `json:"creationTimestamp,omitempty"`
	UpdateTimestamp      string `json:"updateTimestamp,omitempty"`
}

type NetworkInfo struct {
	Vpc            string `json:"vpc,omitempty"`
	Subnet         string `json:"subnet,omitempty"`
	HighwaySubnet  string `json:"highwaySubnet,omitempty"`
}

type ContainerNetworkInfo struct {
	Mode string  `json:"mode,omitempty"`
	Cidr string  `json:"cidr,omitempty"`
}

type SpecInfo struct {
	ClusterType string      `json:"type,omitempty"`
	Flavor      string      `json:"flavor,omitempty"`
	K8sVersion  string      `json:"version,omitempty"`
	Description string		`json:"description,omitempty"`
	HostNetwork *NetworkInfo `json:"hostNetwork,omitempty"`
	ContainerNetwork *ContainerNetworkInfo `json:"containerNetwork,omitempty"`
}

type EndPoints struct {
	Internal string `json:"internal,omitempty"`
	External string `json:"external,omitempty"`
}

type Conditions struct {
	Type                string  `json:"type,omitempty"`
	Status              string  `json:"status,omitempty"`
	Reason              string  `json:"reason,omitempty"`
	Message             string  `json:"message,omitempty"`
	LastProbeTime       string  `json:"lastProbeTime,omitempty"`
	LastTransitionTime  string  `json:"lastTransitionTime,omitempty"`
}

type StatusInfo struct {
	Phase      	string    		`json:"phase,omitempty"`
	JobID      	string    		`json:"jobID,omitempty"`
	Reason     	string    		`json:"reason,omitempty"`
	Message     string    		`json:"message,omitempty"`
	Conditions  *Conditions 	`json:"conditions,omitempty"`
	Endpoints  	*EndPoints  	`json:"endpoints,omitempty"`
}

type ClusterInfo struct {
	Kind       string    	`json:"kind"`
	ApiVersion string    	`json:"apiVersion"`
	MetaData   MetaInfo  	`json:"metadata"`
	Spec       SpecInfo  	`json:"spec"`
	Status	   *StatusInfo	`json:"status,omitempty"`
}

//update cluster struct
type UpdateInfo struct {
	Description string `json:"description"`
}
type UpdateCluster struct {
	Spec UpdateInfo `json:"spec"`
}

//node struct
type NodeMetaInfo struct {
	Name                 string `json:"name"`
	Uid                  string `json:"uid"`
	CreationTimestamp    string `json:"creationTimestamp,omitempty"`
	UpdateTimestamp      string `json:"updateTimestamp,omitempty"`
}

type NodeLogin struct {
	SSHKey string `json:"sshKey"`
}

type NodeVolume struct {
	Size       int64     `json:"size"`
	VolumeType string  	`json:"volumetype"`
}

type Bandwidth struct {
	ChargeMode     string 	`json:"chargemode"`
	Size           int64    `json:"size"`
	ShareType      string 	`json:"sharetype"`
}

type Eip struct {
	Iptype    string    `json:"iptype"`
	Bandwidth Bandwidth `json:"bandwidth"`
}

type PublicIP struct {
	Ids    []string   	`json:"ids"`
	Count  int64        `json:"count"`
	Eip    Eip        	`json:"eip"`
}

type NodeSpecInfo struct {
	ClusterType    string       	`json:"type"`
	Flavor         string       	`json:"flavor"`
	AvailableZone  string       	`json:"az"`
	Login          NodeLogin    	`json:"login"`
	RootVolume     NodeVolume   	`json:"rootVolume"`
	DataVolumes    []NodeVolume 	`json:"dataVolumes"`
	PublicIP       PublicIP   		`json:"publicIP"`
	Count          int64          	`json:"count"`
	BillingMode    int64          	`json:"billingMode"`
}

type NodeStatusInfo struct {
	JobID string `json:"jobID"`
}

type NodeInfo struct {
	Kind       string    		`json:"kind"`
	ApiVersion string    		`json:"apiversion"`
	MetaData   NodeMetaInfo  	`json:"metadata"`
	Spec       NodeSpecInfo  	`json:"spec"`
	Status	   NodeStatusInfo	`json:"status,omitempty"`
}

//Cluster cert info
type Cluster struct{
	Server                     string   `json:"server,omitempty"`
	CertificateAuthorityData   string   `json:"certificate-authority-data,omitempty"`
}
type ClusterConfig struct {
	Name     string     `json:"name,omitempty"`
	Cluster  Cluster	`json:"clusters,omitempty"`
}

type User struct {
	ClientCertificateData string    `json:"client-certificate-data,omitempty"`
	ClientKeyData         string    `json:"client-key-data,omitempty"`
}
type UserConfig struct {
	Name     string     `json:"name,omitempty"`
	User     User       `json:"user,omitempty"`
}

type Context struct {
	Cluster string  `json:"context,omitempty"`
	User    string  `json:"user,omitempty"`
}
type ContextConfig struct {
	Name      string  `json:"name,omitempty"`
	Context   Context `json:"context,omitempty"`
}

type ClusterCert struct {
	Kind       	string       		`json:"kind,omitempty"`
	ApiVersion 	string       		`json:"apiVersion,omitempty"`
	Clusters    []ClusterConfig		`json:"clusters,omitempty"`
	Users       []UserConfig    	`json:"users,omitempty"`
	Contexts    []ContextConfig     `json:"contexts,omitempty"`
}
