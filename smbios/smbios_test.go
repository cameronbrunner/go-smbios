// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/siderolabs/go-smbios/smbios"
)

func TestASRockSingleRyzen(t *testing.T) {
	DoTestDesktopManagementInterface(t, "ASRock-Single-Ryzen")
}

func TestDellPowerEdgeR630DualXeon(t *testing.T) {
	DoTestDesktopManagementInterface(t, "Dell-PowerEdge-R630-Dual-Xeon")
}

func TestDellSuperMicroDualXeon(t *testing.T) {
	DoTestDesktopManagementInterface(t, "SuperMicro-Dual-Xeon")
}

func TestDellSuperMicroQuadOpteron(t *testing.T) {
	DoTestDesktopManagementInterface(t, "SuperMicro-Quad-Opteron")
}

func DoTestDesktopManagementInterface(t *testing.T, name string) {
	t.Run(name, func(t *testing.T) {
		stream, err := os.Open("testdata/" + name + ".dmi")
		require.NoError(t, err)

		//nolint: errcheck
		defer stream.Close()

		version := smbios.Version{Major: 3, Minor: 3, Revision: 0} // dummy version
		actual, err := smbios.Decode(stream, version)
		require.NoError(t, err)

		expectedJSON, err := os.ReadFile("testdata/" + name + ".json")
		require.NoError(t, err)

		var expected smbios.SMBIOS

		require.NoError(t, json.Unmarshal(expectedJSON, &expected))

		// remove parsed structures as they are not in JSON
		actual.Structures = nil

		require.Equal(t, &expected, actual)
	})
}
