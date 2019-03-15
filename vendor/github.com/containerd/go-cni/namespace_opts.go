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

type NamespaceOpts func(s *Namespace) error

// Capabilities
func WithCapabilityPortMap(portMapping []PortMapping) NamespaceOpts {
	return func(c *Namespace) error {
		c.capabilityArgs["portMappings"] = portMapping
		return nil
	}
}

func WithCapabilityIPRanges(ipRanges []IPRanges) NamespaceOpts {
	return func(c *Namespace) error {
		c.capabilityArgs["ipRanges"] = ipRanges
		return nil
	}
}

func WithCapabilityFloatingIP(fip *FloatingIP) NamespaceOpts {
	return func(c *Namespace) error {
		if fip != nil {
			c.capabilityArgs["floatingip"] = *fip
		}
		return nil
	}
}

func WithCapability(name string, capability interface{}) NamespaceOpts {
	return func(c *Namespace) error {
		c.capabilityArgs[name] = capability
		return nil
	}
}

// Args
func WithLabels(labels map[string]string) NamespaceOpts {
	return func(c *Namespace) error {
		for k, v := range labels {
			c.args[k] = v
		}
		return nil
	}
}

func WithArgs(k, v string) NamespaceOpts {
	return func(c *Namespace) error {
		c.args[k] = v
		return nil
	}
}
