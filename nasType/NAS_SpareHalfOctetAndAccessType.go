// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType

// SpareHalfOctetAndAccessType 9.11.3.11 9.5
// AccessType Row, sBit, len = [0, 0], 2 , 2
type SpareHalfOctetAndAccessType struct {
	Octet uint8
}

func NewSpareHalfOctetAndAccessType() (spareHalfOctetAndAccessType *SpareHalfOctetAndAccessType) {
	spareHalfOctetAndAccessType = &SpareHalfOctetAndAccessType{}
	return spareHalfOctetAndAccessType
}

// SpareHalfOctetAndAccessType 9.11.3.11 9.5
// AccessType Row, sBit, len = [0, 0], 2 , 2
func (a *SpareHalfOctetAndAccessType) GetAccessType() (accessType uint8) {
	return a.Octet & GetBitMask(2, 0)
}

// SpareHalfOctetAndAccessType 9.11.3.11 9.5
// AccessType Row, sBit, len = [0, 0], 2 , 2
func (a *SpareHalfOctetAndAccessType) SetAccessType(accessType uint8) {
	a.Octet = (a.Octet & 252) + (accessType & 3)
}
