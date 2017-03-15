package libs

// Device Structure that will be loaded by the registery
type Device struct {
	Alias         string
	HWAddress     string
	WakeUpMethod  string
	BoundServices []string
}
