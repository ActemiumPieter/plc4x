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

// ApduDataExtLinkWrite is the data-structure of this message
type ApduDataExtLinkWrite struct {
	*ApduDataExt

	// Arguments.
	Length uint8
}

// IApduDataExtLinkWrite is the corresponding interface of ApduDataExtLinkWrite
type IApduDataExtLinkWrite interface {
	IApduDataExt
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *ApduDataExtLinkWrite) GetExtApciType() uint8 {
	return 0x27
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataExtLinkWrite) InitializeParent(parent *ApduDataExt) {}

func (m *ApduDataExtLinkWrite) GetParent() *ApduDataExt {
	return m.ApduDataExt
}

// NewApduDataExtLinkWrite factory function for ApduDataExtLinkWrite
func NewApduDataExtLinkWrite(length uint8) *ApduDataExtLinkWrite {
	_result := &ApduDataExtLinkWrite{
		ApduDataExt: NewApduDataExt(length),
	}
	_result.Child = _result
	return _result
}

func CastApduDataExtLinkWrite(structType interface{}) *ApduDataExtLinkWrite {
	if casted, ok := structType.(ApduDataExtLinkWrite); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataExtLinkWrite); ok {
		return casted
	}
	if casted, ok := structType.(ApduDataExt); ok {
		return CastApduDataExtLinkWrite(casted.Child)
	}
	if casted, ok := structType.(*ApduDataExt); ok {
		return CastApduDataExtLinkWrite(casted.Child)
	}
	return nil
}

func (m *ApduDataExtLinkWrite) GetTypeName() string {
	return "ApduDataExtLinkWrite"
}

func (m *ApduDataExtLinkWrite) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataExtLinkWrite) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtLinkWrite) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtLinkWriteParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExtLinkWrite, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtLinkWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtLinkWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtLinkWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtLinkWrite")
	}

	// Create a partially initialized instance
	_child := &ApduDataExtLinkWrite{
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child, nil
}

func (m *ApduDataExtLinkWrite) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtLinkWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtLinkWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtLinkWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtLinkWrite")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtLinkWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
