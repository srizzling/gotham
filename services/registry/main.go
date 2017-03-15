package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/srizzling/gotham/libs"
)

var devices map[string]libs.Device

// getDevice from an alias and return a pointer to a device obj
func getDevice(alias string) (*libs.Device, error) {
	aliasTo := devices[alias]
	if aliasTo.Alias == "" {
		return nil, errors.New("Alias doesn't exist in table")
	}
	return &aliasTo, nil
}

func listDevices() (*[]libs.Device, error) {
	keys := make([]libs.Device, 0, len(devices))
	for _, v := range devices {
		keys = append(keys, v)
	}
	return &keys, nil
}

func viewDeviceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filter := vars["device"]
	device, err := getDevice(filter)
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "An error has occured")
	}
	b, err := json.Marshal(device)
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "An error has occured")
	}
	fmt.Fprintf(w, string(b))
}

func viewServiceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filter := vars["service"]
	devicePrint, err := filterDevices(filter)
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "An error has occured")
	}
	printJSON, err := printJSON(*devicePrint)
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "An error has occured")
	}
	fmt.Fprintf(w, printJSON)

}

func viewIndexHandler(w http.ResponseWriter, r *http.Request) {
	devicePrint, err := listDevices()
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "An error has occured")
	}
	printJSON, err := printJSON(*devicePrint)
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "An error has occured")
	}
	fmt.Fprintf(w, printJSON)
}

func filterDevices(filter string) (*[]libs.Device, error) {
	// Currently only supports listing a single service for now
	// TODO: expand filter to be more consise
	filteredDevices := make([]libs.Device, 0, len(devices))
	for _, e := range devices {
		for _, f := range e.BoundServices {
			if f == filter {
				filteredDevices = append(filteredDevices, e)
			}
		}
	}
	return &filteredDevices, nil
}

// TODO: Make this dynamic with the ability to read the config file
func loadDevices(file string) error {
	devices = make(map[string]libs.Device)
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	fileErr := json.Unmarshal(content, &devices)
	if fileErr != nil {
		return fileErr
	}
	return nil
}

func printJSON(jsonPrint []libs.Device) (string, error) {
	b, err := json.Marshal(jsonPrint)
	if err != nil {
		return "", err
	}
	return string(b), err
}

func main() {
	err := loadDevices("test.json")
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/view", viewIndexHandler)
	router.HandleFunc("/view/{device}", viewDeviceHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
