// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType_test

import (
	"testing"

	"github.com/free5gc/nas/nasType"

	"github.com/stretchr/testify/assert"
)

func TestNasTypeNewExtendedProtocolDiscriminatort(t *testing.T) {
	a := nasType.NewExtendedProtocolDiscriminator()
	assert.NotNil(t, a)
}

type nasTypeExtendedProtocolDiscriminatorData struct {
	in  uint8
	out uint8
}

var nasTypeExtendedProtocolDiscriminatorTable = []nasTypeExtendedProtocolDiscriminatorData{
	{2, 2},
}

func TestNasTypeGetSetExtendedProtocolDiscriminator(t *testing.T) {
	a := nasType.NewExtendedProtocolDiscriminator()
	for _, table := range nasTypeExtendedProtocolDiscriminatorTable {
		a.SetExtendedProtocolDiscriminator(table.in)
		assert.Equal(t, table.out, a.GetExtendedProtocolDiscriminator())
	}
}
