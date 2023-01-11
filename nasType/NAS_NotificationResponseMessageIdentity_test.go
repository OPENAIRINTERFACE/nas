// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasType_test

import (
	"testing"

	"github.com/omec-project/nas"
	"github.com/omec-project/nas/nasType"

	"github.com/stretchr/testify/assert"
)

func TestNasTypeNewNotificationResponseMessageIdentity(t *testing.T) {
	a := nasType.NewNotificationResponseMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypeNotificationResponseMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypeNotificationResponseMessageIdentityMessageTypeTable = []nasTypeNotificationResponseMessageIdentityMessageType{
	{nas.MsgTypeNotificationResponse, nas.MsgTypeNotificationResponse},
}

func TestNasTypeGetSetNotificationResponseMessageIdentityMessageType(
	t *testing.T,
) {
	a := nasType.NewNotificationResponseMessageIdentity()
	for _, table := range nasTypeNotificationResponseMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
