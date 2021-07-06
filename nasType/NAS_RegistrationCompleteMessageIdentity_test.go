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

func TestNasTypeNewRegistrationCompleteMessageIdentity(t *testing.T) {
	a := nasType.NewRegistrationCompleteMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypeRegistrationCompleteMessageIdentity struct {
	in  uint8
	out uint8
}

var nasTypeRegistrationCompleteMessageIdentityTable = []nasTypeRegistrationCompleteMessageIdentity{
	{0x03, 0x03},
}

func TestNasTypeRegistrationCompleteMessageIdentityGetSetMessageType(t *testing.T) {
	a := nasType.NewRegistrationCompleteMessageIdentity()
	for _, table := range nasTypeRegistrationCompleteMessageIdentityTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}

type RegistrationCompleteMessageIdentityTestDataTemplate struct {
	in  nasType.RegistrationCompleteMessageIdentity
	out nasType.RegistrationCompleteMessageIdentity
}

var RegistrationCompleteMessageIdentityTestData = []nasType.RegistrationCompleteMessageIdentity{
	{0x03},
}

var RegistrationCompleteMessageIdentityExpectedTestData = []nasType.RegistrationCompleteMessageIdentity{
	{0x03},
}

var RegistrationCompleteMessageIdentityTable = []RegistrationCompleteMessageIdentityTestDataTemplate{
	{RegistrationCompleteMessageIdentityTestData[0], RegistrationCompleteMessageIdentityExpectedTestData[0]},
}

func TestNasTypeRegistrationCompleteMessageIdentity(t *testing.T) {

	for _, table := range RegistrationCompleteMessageIdentityTable {

		a := nasType.NewRegistrationCompleteMessageIdentity()

		a.SetMessageType(table.in.GetMessageType())
		assert.Equal(t, table.out.GetMessageType(), a.GetMessageType())
	}
}
