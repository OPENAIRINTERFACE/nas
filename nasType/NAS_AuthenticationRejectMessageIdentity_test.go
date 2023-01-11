// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasType_test

import (
	"testing"

	"github.com/omec-project/nas/nasMessage"
	"github.com/omec-project/nas/nasType"

	"github.com/stretchr/testify/assert"
)

type nasTypeRejectMessageIdentityData struct {
	in  uint8
	out uint8
}

var nasTypeRejectMessageIdentityTable = []nasTypeRejectMessageIdentityData{
	{
		nasMessage.PDUSessionEstablishmentRejectEAPMessageType,
		nasMessage.PDUSessionEstablishmentRejectEAPMessageType,
	},
}

func TestNasTypeNewAuthenticationRejectMessageIdentity(t *testing.T) {
	a := nasType.NewAuthenticationRejectMessageIdentity()
	assert.NotNil(t, a)
}

func TestNasTypeGetSetAuthenticationRejectMessageIdentity(t *testing.T) {
	a := nasType.NewAuthenticationRejectMessageIdentity()
	for _, table := range nasTypeRejectMessageIdentityTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
