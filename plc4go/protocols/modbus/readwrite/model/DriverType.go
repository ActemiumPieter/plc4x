/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// DriverType is an enum
type DriverType uint32

type IDriverType interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	DriverType_MODBUS_TCP   DriverType = 0x01
	DriverType_MODBUS_RTU   DriverType = 0x02
	DriverType_MODBUS_ASCII DriverType = 0x03
)

var DriverTypeValues []DriverType

func init() {
	_ = errors.New
	DriverTypeValues = []DriverType{
		DriverType_MODBUS_TCP,
		DriverType_MODBUS_RTU,
		DriverType_MODBUS_ASCII,
	}
}

func DriverTypeByValue(value uint32) (enum DriverType, ok bool) {
	switch value {
	case 0x01:
		return DriverType_MODBUS_TCP, true
	case 0x02:
		return DriverType_MODBUS_RTU, true
	case 0x03:
		return DriverType_MODBUS_ASCII, true
	}
	return 0, false
}

func DriverTypeByName(value string) (enum DriverType, ok bool) {
	switch value {
	case "MODBUS_TCP":
		return DriverType_MODBUS_TCP, true
	case "MODBUS_RTU":
		return DriverType_MODBUS_RTU, true
	case "MODBUS_ASCII":
		return DriverType_MODBUS_ASCII, true
	}
	return 0, false
}

func DriverTypeKnows(value uint32) bool {
	for _, typeValue := range DriverTypeValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastDriverType(structType interface{}) DriverType {
	castFunc := func(typ interface{}) DriverType {
		if sDriverType, ok := typ.(DriverType); ok {
			return sDriverType
		}
		return 0
	}
	return castFunc(structType)
}

func (m DriverType) GetLengthInBits() uint16 {
	return 32
}

func (m DriverType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DriverTypeParse(readBuffer utils.ReadBuffer) (DriverType, error) {
	val, err := readBuffer.ReadUint32("DriverType", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading DriverType")
	}
	if enum, ok := DriverTypeByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return DriverType(val), nil
	} else {
		return enum, nil
	}
}

func (e DriverType) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint32("DriverType", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e DriverType) PLC4XEnumName() string {
	switch e {
	case DriverType_MODBUS_TCP:
		return "MODBUS_TCP"
	case DriverType_MODBUS_RTU:
		return "MODBUS_RTU"
	case DriverType_MODBUS_ASCII:
		return "MODBUS_ASCII"
	}
	return ""
}

func (e DriverType) String() string {
	return e.PLC4XEnumName()
}
