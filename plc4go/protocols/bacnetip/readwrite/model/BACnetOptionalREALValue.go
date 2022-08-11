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

// BACnetOptionalREALValue is the corresponding interface of BACnetOptionalREALValue
type BACnetOptionalREALValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetOptionalREAL
	// GetRealValue returns RealValue (property field)
	GetRealValue() BACnetApplicationTagReal
}

// BACnetOptionalREALValueExactly can be used when we want exactly this type and not a type which fulfills BACnetOptionalREALValue.
// This is useful for switch cases.
type BACnetOptionalREALValueExactly interface {
	BACnetOptionalREALValue
	isBACnetOptionalREALValue() bool
}

// _BACnetOptionalREALValue is the data-structure of this message
type _BACnetOptionalREALValue struct {
	*_BACnetOptionalREAL
	RealValue BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetOptionalREALValue) InitializeParent(parent BACnetOptionalREAL, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetOptionalREALValue) GetParent() BACnetOptionalREAL {
	return m._BACnetOptionalREAL
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalREALValue) GetRealValue() BACnetApplicationTagReal {
	return m.RealValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetOptionalREALValue factory function for _BACnetOptionalREALValue
func NewBACnetOptionalREALValue(realValue BACnetApplicationTagReal, peekedTagHeader BACnetTagHeader) *_BACnetOptionalREALValue {
	_result := &_BACnetOptionalREALValue{
		RealValue:           realValue,
		_BACnetOptionalREAL: NewBACnetOptionalREAL(peekedTagHeader),
	}
	_result._BACnetOptionalREAL._BACnetOptionalREALChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetOptionalREALValue(structType interface{}) BACnetOptionalREALValue {
	if casted, ok := structType.(BACnetOptionalREALValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalREALValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalREALValue) GetTypeName() string {
	return "BACnetOptionalREALValue"
}

func (m *_BACnetOptionalREALValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetOptionalREALValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (realValue)
	lengthInBits += m.RealValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetOptionalREALValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetOptionalREALValueParse(readBuffer utils.ReadBuffer) (BACnetOptionalREALValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalREALValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalREALValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (realValue)
	if pullErr := readBuffer.PullContext("realValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for realValue")
	}
	_realValue, _realValueErr := BACnetApplicationTagParse(readBuffer)
	if _realValueErr != nil {
		return nil, errors.Wrap(_realValueErr, "Error parsing 'realValue' field of BACnetOptionalREALValue")
	}
	realValue := _realValue.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("realValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for realValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetOptionalREALValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalREALValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetOptionalREALValue{
		_BACnetOptionalREAL: &_BACnetOptionalREAL{},
		RealValue:           realValue,
	}
	_child._BACnetOptionalREAL._BACnetOptionalREALChildRequirements = _child
	return _child, nil
}

func (m *_BACnetOptionalREALValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetOptionalREALValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetOptionalREALValue")
		}

		// Simple Field (realValue)
		if pushErr := writeBuffer.PushContext("realValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for realValue")
		}
		_realValueErr := writeBuffer.WriteSerializable(m.GetRealValue())
		if popErr := writeBuffer.PopContext("realValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for realValue")
		}
		if _realValueErr != nil {
			return errors.Wrap(_realValueErr, "Error serializing 'realValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetOptionalREALValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetOptionalREALValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetOptionalREALValue) isBACnetOptionalREALValue() bool {
	return true
}

func (m *_BACnetOptionalREALValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
