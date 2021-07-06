// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType

// PDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity struct {
	Octet uint8
}

func NewPDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity() (pDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity *PDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity) {
	pDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity = &PDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity{}
	return pDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity
}

// PDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *PDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// PDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *PDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
