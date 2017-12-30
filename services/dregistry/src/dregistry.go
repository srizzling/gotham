package dregistry

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"golang.org/x/net/context"

	proto "github.com/srizzling/gotham/services/dregistry/proto"
	pdevice "github.com/srizzling/gotham/shared/device/proto"
)

// DRegistry stands for DeviceRegistry and how devices register to Gotham.
type DRegistry struct {
	Devices map[string]*pdevice.Device
}

// GetDevice is a way to to return a single device based on Alias
func (g *DRegistry) GetDevice(ctx context.Context, req *proto.GetDeviceRequest, rsp *proto.GetDeviceResponse) error {
	alias := req.Alias
	device := g.Devices[alias]
	if device == nil {
		return errors.New("Alias doesn't exist")
	}
	rsp.Device = device
	return nil
}

// RegisterDevice (for now this method is kind of lame will make it a discoverable thing)
func (g *DRegistry) RegisterDevice(ctx context.Context, req *proto.RegisterDeviceRequest, rsp *proto.RegisterDeviceResponse) error {
	g.Devices[req.Device.Alias] = req.Device
	rsp.Success = true
	return nil
}

// LoadData is a function that will loaddata into the map
func LoadData(path string) map[string]*pdevice.Device {
	devices := make(map[string]*pdevice.Device)
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	fileErr := json.Unmarshal(content, &devices)

	if fileErr != nil {
		log.Fatal(err)
	}

	return devices
}
