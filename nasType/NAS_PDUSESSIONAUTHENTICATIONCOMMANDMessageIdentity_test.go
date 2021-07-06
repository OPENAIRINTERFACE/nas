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

func TestNasTypeNewPDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypePDUSESSIONAUTHENTICATIONCOMMANDMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONAUTHENTICATIONCOMMANDMessageIdentityMessageTypeTable = []nasTypePDUSESSIONAUTHENTICATIONCOMMANDMessageIdentityMessageType{
	{nas.MsgTypePDUSessionAuthenticationCommand, nas.MsgTypePDUSessionAuthenticationCommand},
}

func TestNasTypeGetSetPDUSESSIONAUTHENTICATIONCOMMANDMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity()
	for _, table := range nasTypePDUSESSIONAUTHENTICATIONCOMMANDMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
