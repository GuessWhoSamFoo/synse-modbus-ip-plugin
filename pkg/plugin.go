package pkg

import (
	log "github.com/sirupsen/logrus"
	"github.com/vapor-ware/synse-modbus-ip-plugin/pkg/devices"
	"github.com/vapor-ware/synse-modbus-ip-plugin/pkg/outputs"
	"github.com/vapor-ware/synse-sdk/sdk"
)

// MakePlugin creates a new instance of the Synse Modbus TCP/IP Plugin.
func MakePlugin() *sdk.Plugin {
	plugin, err := sdk.NewPlugin()
	if err != nil {
		log.Fatal(err)
	}

	// Register output types
	err = plugin.RegisterOutputs(
		&outputs.GallonsPerMin,
		&outputs.InchesWaterColumn,
		&outputs.MacAddressWide,
		&outputs.PowerFactor,
		&outputs.VoltAmp,
		&outputs.VAR,
		&outputs.WattHour,
		&outputs.Identity,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Register device handlers
	err = plugin.RegisterDeviceHandlers(
		&devices.CoilsHandler,
		&devices.ReadOnlyCoilsHandler,
		&devices.HoldingRegisterHandler,
		&devices.ReadOnlyHoldingRegisterHandler,
		&devices.InputRegisterHandler,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Register setup actions
	err = plugin.RegisterDeviceSetupActions(
		&devices.OnModbusDeviceLoad,
	)
	if err != nil {
		log.Fatal(err)
	}

	return plugin
}
