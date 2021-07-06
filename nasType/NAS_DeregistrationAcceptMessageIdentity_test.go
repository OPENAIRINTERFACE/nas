// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType_test

import (
	"testing"

	"github.com/free5gc/nas"
	"github.com/free5gc/nas/nasType"

	"github.com/stretchr/testify/assert"
)

type nasTypeDeregistrationAcceptMessageIdentityData struct {
	in  uint8
	out uint8
}

var nasTypeDeregistrationAcceptMessageIdentityTable = []nasTypeDeregistrationAcceptMessageIdentityData{
	{nas.MsgTypeDeregistrationAcceptUETerminatedDeregistration, nas.MsgTypeDeregistrationAcceptUETerminatedDeregistration},
}

func TestNasTypeNewDeregistrationAcceptMessageIdentity(t *testing.T) {
	a := nasType.NewDeregistrationAcceptMessageIdentity()
	assert.NotNil(t, a)
}

func TestNasTypeGetSetDeregistrationAcceptMessageIdentity(t *testing.T) {
	a := nasType.NewDeregistrationAcceptMessageIdentity()
	for _, table := range nasTypeDeregistrationAcceptMessageIdentityTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
