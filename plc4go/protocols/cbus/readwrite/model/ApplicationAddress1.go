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
)

// Code generated by code-generation. DO NOT EDIT.

// ApplicationAddress1 is the corresponding interface of ApplicationAddress1
type ApplicationAddress1 interface {
	utils.LengthAware
	utils.Serializable
	// GetAddress returns Address (property field)
	GetAddress() byte
	// GetIsWildcard returns IsWildcard (virtual field)
	GetIsWildcard() bool
}

// ApplicationAddress1Exactly can be used when we want exactly this type and not a type which fulfills ApplicationAddress1.
// This is useful for switch cases.
type ApplicationAddress1Exactly interface {
	ApplicationAddress1
	isApplicationAddress1() bool
}

// _ApplicationAddress1 is the data-structure of this message
type _ApplicationAddress1 struct {
	Address byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApplicationAddress1) GetAddress() byte {
	return m.Address
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_ApplicationAddress1) GetIsWildcard() bool {
	return bool(bool((m.GetAddress()) == (0xFF)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApplicationAddress1 factory function for _ApplicationAddress1
func NewApplicationAddress1(address byte) *_ApplicationAddress1 {
	return &_ApplicationAddress1{Address: address}
}

// Deprecated: use the interface for direct cast
func CastApplicationAddress1(structType interface{}) ApplicationAddress1 {
	if casted, ok := structType.(ApplicationAddress1); ok {
		return casted
	}
	if casted, ok := structType.(*ApplicationAddress1); ok {
		return *casted
	}
	return nil
}

func (m *_ApplicationAddress1) GetTypeName() string {
	return "ApplicationAddress1"
}

func (m *_ApplicationAddress1) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApplicationAddress1) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (address)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_ApplicationAddress1) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApplicationAddress1Parse(readBuffer utils.ReadBuffer) (ApplicationAddress1, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApplicationAddress1"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApplicationAddress1")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (address)
	_address, _addressErr := readBuffer.ReadByte("address")
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field of ApplicationAddress1")
	}
	address := _address

	// Virtual field
	_isWildcard := bool((address) == (0xFF))
	isWildcard := bool(_isWildcard)
	_ = isWildcard

	if closeErr := readBuffer.CloseContext("ApplicationAddress1"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApplicationAddress1")
	}

	// Create the instance
	return &_ApplicationAddress1{
		Address: address,
	}, nil
}

func (m *_ApplicationAddress1) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ApplicationAddress1"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ApplicationAddress1")
	}

	// Simple Field (address)
	address := byte(m.GetAddress())
	_addressErr := writeBuffer.WriteByte("address", (address))
	if _addressErr != nil {
		return errors.Wrap(_addressErr, "Error serializing 'address' field")
	}
	// Virtual field
	if _isWildcardErr := writeBuffer.WriteVirtual("isWildcard", m.GetIsWildcard()); _isWildcardErr != nil {
		return errors.Wrap(_isWildcardErr, "Error serializing 'isWildcard' field")
	}

	if popErr := writeBuffer.PopContext("ApplicationAddress1"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ApplicationAddress1")
	}
	return nil
}

func (m *_ApplicationAddress1) isApplicationAddress1() bool {
	return true
}

func (m *_ApplicationAddress1) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
