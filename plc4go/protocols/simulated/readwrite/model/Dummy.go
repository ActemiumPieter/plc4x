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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Dummy is the data-structure of this message
type Dummy struct {
	Dummy uint16
}

// IDummy is the corresponding interface of Dummy
type IDummy interface {
	// GetDummy returns Dummy (property field)
	GetDummy() uint16
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *Dummy) GetDummy() uint16 {
	return m.Dummy
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDummy factory function for Dummy
func NewDummy(dummy uint16) *Dummy {
	return &Dummy{Dummy: dummy}
}

func CastDummy(structType interface{}) *Dummy {
	if casted, ok := structType.(Dummy); ok {
		return &casted
	}
	if casted, ok := structType.(*Dummy); ok {
		return casted
	}
	return nil
}

func (m *Dummy) GetTypeName() string {
	return "Dummy"
}

func (m *Dummy) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *Dummy) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (dummy)
	lengthInBits += 16

	return lengthInBits
}

func (m *Dummy) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DummyParse(readBuffer utils.ReadBuffer) (*Dummy, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Dummy"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Dummy")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dummy)
	_dummy, _dummyErr := readBuffer.ReadUint16("dummy", 16)
	if _dummyErr != nil {
		return nil, errors.Wrap(_dummyErr, "Error parsing 'dummy' field")
	}
	dummy := _dummy

	if closeErr := readBuffer.CloseContext("Dummy"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Dummy")
	}

	// Create the instance
	return NewDummy(dummy), nil
}

func (m *Dummy) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("Dummy"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Dummy")
	}

	// Simple Field (dummy)
	dummy := uint16(m.Dummy)
	_dummyErr := writeBuffer.WriteUint16("dummy", 16, (dummy))
	if _dummyErr != nil {
		return errors.Wrap(_dummyErr, "Error serializing 'dummy' field")
	}

	if popErr := writeBuffer.PopContext("Dummy"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Dummy")
	}
	return nil
}

func (m *Dummy) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
