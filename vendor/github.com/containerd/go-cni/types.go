/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package cni

const (
	CNIPluginName        = "cni"
	DefaultNetDir        = "/etc/cni/net.d"
	DefaultCNIDir        = "/opt/cni/bin"
	VendorCNIDirTemplate = "%s/opt/%s/bin"
	DefaultPrefix        = "eth"

	// AnnotationPodFloatingIP is in pod annotation
	AnnotationPodFloatingIP = "wocloud.cn/floating-ip"
	// AnnotationPodSubnet is in pod annotation
	AnnotationPodSubnet = "wocloud.cn/floating-subnet"
	// AnnotationPodGateway is in pod annotation
	AnnotationPodGateway = "wocloud.cn/floating-gateway"
	// AnnotationPodVlan is in pod annotation
	AnnotationPodVlan = "wocloud.cn/floating-vlan"
)

type config struct {
	pluginDirs    []string
	pluginConfDir string
	prefix        string
}

type PortMapping struct {
	HostPort      int32
	ContainerPort int32
	Protocol      string
	HostIP        string
}

type IPRanges struct {
	Subnet     string
	RangeStart string
	RangeEnd   string
	Gateway    string
}

type FloatingIP struct {
	Ip      string
	Subnet  string
	Gateway string
	// Vlan is -1 stand for no vlan
	Vlan string
	Flag bool
}

func CreateFloatingIP(annotation map[string]string) *FloatingIP {
	f := new(FloatingIP)

	if v, ok := annotation[AnnotationPodFloatingIP]; !ok {
		return nil
	} else {
		f.Ip = v
	}

	if v, ok := annotation[AnnotationPodSubnet]; !ok {
		return nil
	} else {
		f.Subnet = v
	}

	if v, ok := annotation[AnnotationPodGateway]; !ok {
		return nil
	} else {
		f.Gateway = v
	}

	if v, ok := annotation[AnnotationPodVlan]; !ok {
		f.Vlan = "-1"
	} else {
		f.Vlan = v
	}
	f.Flag = true
	return f
}
