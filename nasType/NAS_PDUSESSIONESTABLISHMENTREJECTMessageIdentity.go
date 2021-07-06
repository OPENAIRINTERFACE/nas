// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType

// PDUSESSIONESTABLISHMENTREJECTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONESTABLISHMENTREJECTMessageIdentity struct {
	Octet uint8
}

func NewPDUSESSIONESTABLISHMENTREJECTMessageIdentity() (pDUSESSIONESTABLISHMENTREJECTMessageIdentity *PDUSESSIONESTABLISHMENTREJECTMessageIdentity) {
	pDUSESSIONESTABLISHMENTREJECTMessageIdentity = &PDUSESSIONESTABLISHMENTREJECTMessageIdentity{}
	return pDUSESSIONESTABLISHMENTREJECTMessageIdentity
}

// PDUSESSIONESTABLISHMENTREJECTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *PDUSESSIONESTABLISHMENTREJECTMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// PDUSESSIONESTABLISHMENTREJECTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *PDUSESSIONESTABLISHMENTREJECTMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
