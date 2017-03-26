package dregistry

import "testing"

func stringInSlice(a string, b []string) bool {
	for _, s := range b {
		if s == a {
			return true
		}
	}
	return false
}

func TestMiniDevice(t *testing.T) {
	// This test is a way drift between the proto defination, and switch case for minimizling the array

}

func TestFilterDevicesByService(t *testing.T) {
	err := loadData()
	if err != nil {
		t.Error(err)
	}

	// Test that provided "wol" service 3/3 device are returned
	service := "wol"
	devices, err := filterDevicesByService(service)

	if err != nil {
		t.Error(err)
	}

	// Verify that indeed all three devices are turned, and have the wol service
	// in the bound service
	n := len(devices)
	if n != 3 {
		t.Errorf("Expected: length of device map 3 \n got: %d", n)
	}
	var found bool
	for _, d := range devices {
		found = stringInSlice(service, d.BoundServices)
	}
	if !found {
		t.Errorf("Service %s was not found in one of the devices, it should.", service)
	}

	// Test that provided "light" service that 1/3 devices are returned
	service = "light"
	devices, err = filterDevicesByService(service)

	if err != nil {
		t.Error(err)
	}

	n = len(devices)
	if n != 1 {
		t.Errorf("Expected: length of device map 1 \n got: %d", n)
	}

	for _, d := range devices {
		found = stringInSlice(service, d.BoundServices)
	}
	if !found {
		t.Errorf("Service %s was not found in one of the devices, it should.", service)
	}

	// Test provided no service all are returned
	service = ""
	devices, err = filterDevicesByService(service)

	if err != nil {
		t.Error(err)
	}

	n = len(devices)
	if n != 3 {
		t.Errorf("Expected: length of device map 3\n got: %d", n)
	}

	// Test provided a service that doesn't exist, an error is returned?
	service = "nonexist"
	devices, err = filterDevicesByService(service)

	if err == nil {
		t.Error("Filter ( " + service + " ) doesn't exist, but function didn't return error")
	}
}
