package vo

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-09 10:03
**/

type RegisterServiceInstanceParam struct {
	Ip          string            `param:"ip"`
	Port        uint64            `param:"port"`
	NamespaceId string            `param:"namespaceId"`
	Weight      float64           `param:"weight"`
	Enable      bool              `param:"enable"`
	Healthy     bool              `param:"healthy"`
	Metadata    map[string]string `param:"metadata"`
	ClusterName string            `param:"clusterName"`
	ServiceName string            `param:"serviceName"`
	GroupName   string            `param:"groupName"`
	Ephemeral   bool              `param:"ephemeral"`
}

type LogoutServiceInstanceParam struct {
	Ip          string `param:"ip"`
	Port        uint64 `param:"port"`
	NamespaceId string `param:"namespaceId"`
	ClusterName string `param:"clusterName"`
	ServiceName string `param:"serviceName"`
	GroupName   string `param:"groupName"`
	Ephemeral   bool   `param:"ephemeral"`
}

type ModifyServiceInstanceParam struct {
	Ip          string            `param:"ip"`
	Port        uint64            `param:"port"`
	NamespaceId string            `param:"namespaceId"`
	Weight      float64           `param:"weight"`
	Enable      bool              `param:"enable"`
	Healthy     bool              `param:"healthy"`
	Metadata    map[string]string `param:"metadata"`
	ClusterName string            `param:"clusterName"`
	ServiceName string            `param:"serviceName"`
	GroupName   string            `param:"groupName"`
	Ephemeral   bool              `param:"ephemeral"`
}

type GetServiceParam struct {
	ServiceName string   `param:"serviceName"`
	GroupName   string   `param:"groupName"`
	NamespaceId string   `param:"namespaceId"`
	Clusters    []string `param:"clusters"`
	HealthyOnly bool     `param:"healthyOnly"`
}

type GetServiceListParam struct {
	StartPage   uint32 `param:"pageNo"`
	PageSize    uint32 `param:"pageSize"`
	GroupName   string `param:"groupName"`
	NamespaceId string `param:"namespaceId"`
}

type GetServiceInstanceParam struct {
	ServiceName string `param:"serviceName"`
	GroupName   string `param:"groupName"`
	NamespaceId string `param:"namespaceId"`
	Ip          string `param:"ip"`
	Port        uint64 `param:"port"`
	Cluster     string `param:"cluster"`
	HealthyOnly bool   `param:"healthyOnly"`
	Ephemeral   bool   `param:"ephemeral"`
}

type BeatTaskParam struct {
	Ip          string            `json:"ip"`
	Port        uint64            `json:"port"`
	Weight      float64           `json:"weight"`
	ServiceName string            `json:"serviceName"`
	GroupName   string            `json:"groupName"`
	Ephemeral   bool              `json:"ephemeral"`
	Cluster     string            `json:"cluster"`
	Metadata    map[string]string `json:"metadata"`
}

//type BeatHealthyParam struct {
//	ServiceName string `json:"serviceName"`
//	GroupName   string `json:"groupName"`
//	NamespaceId string `json:"namespaceId"`
//	ClusterName string `json:"clusterName"`
//	Ip          string `json:"ip"`
//	Port        uint64 `json:"port"`
//	Healthy     bool   `json:"healthy"`
//}

type GetServiceDetailParam struct {
	ServiceName string `param:"serviceName"`
	GroupName   string `param:"groupName"`
	NamespaceId string `param:"namespaceId"`
}

type SubscribeParam struct {
	ServiceName       string   `param:"serviceName"`
	GroupName         string   `param:"groupName"`
	NamespaceId       string   `param:"namespaceId"`
	Clusters          []string `param:"clusters"`
	HealthyOnly       bool     `param:"healthyOnly"`
	SubscribeCallback func([]SubscribeService, error)
}
