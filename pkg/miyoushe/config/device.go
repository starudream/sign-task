package config

import (
	"fmt"

	"github.com/starudream/sign-task/util"
)

type Device struct {
	Id string `json:"id,omitempty" yaml:"id,omitempty"`
	// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/other/authentication.md#x-rpc-client_type
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
	// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/other/authentication.md#x-rpc-device_name
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/other/authentication.md#x-rpc-device_model
	Model string `json:"model,omitempty" yaml:"model,omitempty"`
	// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/other/authentication.md#x-rpc-sys_version
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
	// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/other/authentication.md#x-rpc-channel
	Channel string `json:"channel,omitempty" yaml:"channel,omitempty"`
}

func (d Device) TableCellString() string {
	return fmt.Sprintf("%s (%s)", d.Id, d.Name)
}

var device = Device{
	Id:      "",
	Type:    "2",
	Name:    "Xiaomi 22011211C",
	Model:   "22011211C",
	Version: "13",
	Channel: "miyousheluodi",
}

func NewDevice() Device {
	d := device
	d.Id = util.UUID()
	return d
}
