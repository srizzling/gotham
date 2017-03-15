package main

import (
	"os"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("alfred-control")

// The idea of this service will progromattically
// create an accesory that will bridge the device with
// homekit

func setup() {
	// Log setup here
	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)

	loggingFile := logging.NewLogBackend(os.Stderr, "", 0)
	loggingFileFormatter := logging.NewBackendFormatter(loggingFile, format)
	logging.SetBackend(loggingFileFormatter)
}

func main() {
	// init setup here
	setup()

	aliases := wol.GetAliasesArray()
	accs := make([]*accessory.Accessory, 0, len(*aliases))

	for _, e := range *aliases {
		name := e
		switchInfo := accessory.Info{
			Name: name,
		}

		acc := accessory.NewSwitch(switchInfo)

		// Log to console when client (e.g. iOS app) changes the value of the on characteristis
		acc.Switch.On.OnValueRemoteUpdate(func(on bool) {
			if on == true {
				name := acc.Info.Name.GetValue()
				macObj, err := wol.GetMacAddress(name)
				if err != nil {
					log.Errorf("%s", err)
				}

				wakeupmethod := macObj.WakeUpMethod
				log.Infof("Attempting to turn on %s via the following WakeUpMethod %s", name, wakeupmethod)

				var e error
				if wakeupmethod == "wol" {
					e = wol.SendMagicPacket(name)
				}

				if e != nil {
					log.Criticalf("Cannot turn on %s due to the %e", name, e)
				}

			} else {

			}
		})

		accs = append(accs, acc.Accessory)

	}

	// TODO: figure out how to get generate this
	config := hc.Config{Pin: "12344321", Port: "12345", StoragePath: "./db"}
	t, err := hc.NewIPTransport(config, nil, accs...)

	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}
