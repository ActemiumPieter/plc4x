//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"encoding/xml"
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

type DeviceDescriptorMediumType uint8

type IDeviceDescriptorMediumType interface {
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

const (
	DeviceDescriptorMediumType_TP1    DeviceDescriptorMediumType = 0x0
	DeviceDescriptorMediumType_PL110  DeviceDescriptorMediumType = 0x1
	DeviceDescriptorMediumType_RF     DeviceDescriptorMediumType = 0x2
	DeviceDescriptorMediumType_TP0    DeviceDescriptorMediumType = 0x3
	DeviceDescriptorMediumType_PL132  DeviceDescriptorMediumType = 0x4
	DeviceDescriptorMediumType_KNX_IP DeviceDescriptorMediumType = 0x5
)

var DeviceDescriptorMediumTypeValues []DeviceDescriptorMediumType

func init() {
	DeviceDescriptorMediumTypeValues = []DeviceDescriptorMediumType{
		DeviceDescriptorMediumType_TP1,
		DeviceDescriptorMediumType_PL110,
		DeviceDescriptorMediumType_RF,
		DeviceDescriptorMediumType_TP0,
		DeviceDescriptorMediumType_PL132,
		DeviceDescriptorMediumType_KNX_IP,
	}
}

func DeviceDescriptorMediumTypeByValue(value uint8) DeviceDescriptorMediumType {
	switch value {
	case 0x0:
		return DeviceDescriptorMediumType_TP1
	case 0x1:
		return DeviceDescriptorMediumType_PL110
	case 0x2:
		return DeviceDescriptorMediumType_RF
	case 0x3:
		return DeviceDescriptorMediumType_TP0
	case 0x4:
		return DeviceDescriptorMediumType_PL132
	case 0x5:
		return DeviceDescriptorMediumType_KNX_IP
	}
	return 0
}

func DeviceDescriptorMediumTypeByName(value string) DeviceDescriptorMediumType {
	switch value {
	case "TP1":
		return DeviceDescriptorMediumType_TP1
	case "PL110":
		return DeviceDescriptorMediumType_PL110
	case "RF":
		return DeviceDescriptorMediumType_RF
	case "TP0":
		return DeviceDescriptorMediumType_TP0
	case "PL132":
		return DeviceDescriptorMediumType_PL132
	case "KNX_IP":
		return DeviceDescriptorMediumType_KNX_IP
	}
	return 0
}

func CastDeviceDescriptorMediumType(structType interface{}) DeviceDescriptorMediumType {
	castFunc := func(typ interface{}) DeviceDescriptorMediumType {
		if sDeviceDescriptorMediumType, ok := typ.(DeviceDescriptorMediumType); ok {
			return sDeviceDescriptorMediumType
		}
		return 0
	}
	return castFunc(structType)
}

func (m DeviceDescriptorMediumType) LengthInBits() uint16 {
	return 4
}

func (m DeviceDescriptorMediumType) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DeviceDescriptorMediumTypeParse(io utils.ReadBuffer) (DeviceDescriptorMediumType, error) {
	val, err := io.ReadUint8("DeviceDescriptorMediumType", 4)
	if err != nil {
		return 0, nil
	}
	return DeviceDescriptorMediumTypeByValue(val), nil
}

func (e DeviceDescriptorMediumType) Serialize(io utils.WriteBuffer) error {
	err := io.WriteUint8("DeviceDescriptorMediumType", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
	return err
}

func (m *DeviceDescriptorMediumType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.CharData:
			tok := token.(xml.CharData)
			*m = DeviceDescriptorMediumTypeByName(string(tok))
		}
	}
}

func (m DeviceDescriptorMediumType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.String(), start); err != nil {
		return err
	}
	return nil
}

func (e DeviceDescriptorMediumType) name() string {
	switch e {
	case DeviceDescriptorMediumType_TP1:
		return "TP1"
	case DeviceDescriptorMediumType_PL110:
		return "PL110"
	case DeviceDescriptorMediumType_RF:
		return "RF"
	case DeviceDescriptorMediumType_TP0:
		return "TP0"
	case DeviceDescriptorMediumType_PL132:
		return "PL132"
	case DeviceDescriptorMediumType_KNX_IP:
		return "KNX_IP"
	}
	return ""
}

func (e DeviceDescriptorMediumType) String() string {
	return e.name()
}

func (m DeviceDescriptorMediumType) Box(s string, i int) utils.AsciiBox {
	boxName := "DeviceDescriptorMediumType"
	if s != "" {
		boxName += "/" + s
	}
	return utils.BoxString(boxName, fmt.Sprintf("%#0*x %s", 1, uint8(m), m.name()), -1)
}
