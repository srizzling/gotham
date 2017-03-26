package dregistry

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	proto "github.com/srizzling/gotham/services/dregistry/proto"

	"strings"

	"golang.org/x/net/context"
)

// DRegistry stands for DeviceRegistry and how devices register to Gotham.
type DRegistry struct {
	Devices map[string]*proto.Device
}

//var devices map[string]*proto.Device

// View is an API call to return a list of devices provided a filter. If filter is empty, it will return all
// attributes
func (g *DRegistry) View(ctx context.Context, req *proto.ViewRequest, rsp *proto.DeviceRegistry) error {
	var devices map[string]*proto.Device
	var err error

	service := req.Service
	devices, err = g.filterDevicesByService(service)

	// TODO: Think about this, and make it a bit more efficient
	if len(req.Filter) > 0 {
		for _, d := range devices {
			miniDevice(d, req.Filter)
		}
	}

	// if filter is empty all attributes will be returned
	rsp.Devices = devices
	return err
}

func (g *DRegistry) filterDevicesByService(filter string) (map[string]*proto.Device, error) {
	// Currently only supports listing a single service for now
	// TODO: expand filter to be more consise
	if filter == "" {
		return g.devices, nil
	}

	// TODO: This is probably an expensive call
	filteredDevices := make(map[string]*proto.Device)
	for _, d := range g.devices {
		for _, f := range d.BoundServices {
			if f == filter {
				filteredDevices[d.Alias] = d
			}
		}
	}

	//Check if the filter returns nothing, return an error (no need to exit the call)
	if !(len(filteredDevices) > 0) {
		return nil, errors.New("Service " + filter + string(len(filteredDevices)) + " is not contained in registered devices")
	}

	return filteredDevices, nil
}

// Function will reduce the resulting struct to a small struct assuming the fields you return to it.
func miniDevice(d *proto.Device, filters []string) {
	device := new(proto.Device)
	for _, f := range filters {
		switch strings.ToLower(f) {
		case "alias":
			device.Alias = d.Alias
		case "manufacturer":
			device.Manufacturer = d.Manufacturer
		case "model":
			device.Model = d.Model
		case "serialnumber":
			device.SerialNumber = d.SerialNumber
		case "hwaddress":
			device.HWAddress = d.HWAddress
		case "wakeupmethod":
			device.WakeUpMethod = d.WakeUpMethod
		case "hk_accessory":
			device.HK_Accessory = d.HK_Accessory
		case "boundservices":
			device.BoundServices = d.BoundServices
		}
	}
	*d = *device
}

// LoadData is a function that will loaddata into the map
func loadData(path string) map[string]*proto.Device {
	devices := make(map[string]*proto.Device)
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
