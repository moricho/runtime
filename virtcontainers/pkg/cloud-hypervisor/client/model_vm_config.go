/*
 * Cloud Hypervisor API
 *
 * Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.
 *
 * API version: 0.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// VmConfig Virtual machine configuration
type VmConfig struct {
	Cpus         CpusConfig           `json:"cpus,omitempty"`
	Memory       MemoryConfig         `json:"memory,omitempty"`
	Kernel       KernelConfig         `json:"kernel"`
	Cmdline      CmdLineConfig        `json:"cmdline"`
	Disks        []DiskConfig         `json:"disks,omitempty"`
	Net          []NetConfig          `json:"net,omitempty"`
	Rng          RngConfig            `json:"rng,omitempty"`
	Fs           []FsConfig           `json:"fs,omitempty"`
	Pmem         []PmemConfig         `json:"pmem,omitempty"`
	Serial       ConsoleConfig        `json:"serial,omitempty"`
	Console      ConsoleConfig        `json:"console,omitempty"`
	Devices      []DeviceConfig       `json:"devices,omitempty"`
	VhostUserNet []VhostUserNetConfig `json:"vhost_user_net,omitempty"`
	VhostUserBlk []VhostUserBlkConfig `json:"vhost_user_blk,omitempty"`
	Vsock        []VsockConfig        `json:"vsock,omitempty"`
	Iommu        bool                 `json:"iommu,omitempty"`
}
