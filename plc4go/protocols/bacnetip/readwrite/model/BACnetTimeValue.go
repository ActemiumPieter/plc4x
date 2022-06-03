/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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

// BACnetTimeValue is the data-structure of this message
type BACnetTimeValue struct {
	TimeValue *BACnetApplicationTagTime
	Value     *BACnetConstructedDataElement
}

// IBACnetTimeValue is the corresponding interface of BACnetTimeValue
type IBACnetTimeValue interface {
	// GetTimeValue returns TimeValue (property field)
	GetTimeValue() *BACnetApplicationTagTime
	// GetValue returns Value (property field)
	GetValue() *BACnetConstructedDataElement
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

func (m *BACnetTimeValue) GetTimeValue() *BACnetApplicationTagTime {
	return m.TimeValue
}

func (m *BACnetTimeValue) GetValue() *BACnetConstructedDataElement {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimeValue factory function for BACnetTimeValue
func NewBACnetTimeValue(timeValue *BACnetApplicationTagTime, value *BACnetConstructedDataElement) *BACnetTimeValue {
	return &BACnetTimeValue{TimeValue: timeValue, Value: value}
}

func CastBACnetTimeValue(structType interface{}) *BACnetTimeValue {
	if casted, ok := structType.(BACnetTimeValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetTimeValue); ok {
		return casted
	}
	return nil
}

func (m *BACnetTimeValue) GetTypeName() string {
	return "BACnetTimeValue"
}

func (m *BACnetTimeValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetTimeValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timeValue)
	lengthInBits += m.TimeValue.GetLengthInBits()

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetTimeValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTimeValueParse(readBuffer utils.ReadBuffer) (*BACnetTimeValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimeValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeValue)
	if pullErr := readBuffer.PullContext("timeValue"); pullErr != nil {
		return nil, pullErr
	}
	_timeValue, _timeValueErr := BACnetApplicationTagParse(readBuffer)
	if _timeValueErr != nil {
		return nil, errors.Wrap(_timeValueErr, "Error parsing 'timeValue' field")
	}
	timeValue := CastBACnetApplicationTagTime(_timeValue)
	if closeErr := readBuffer.CloseContext("timeValue"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, pullErr
	}
	_value, _valueErr := BACnetConstructedDataElementParse(readBuffer, BACnetObjectType(BACnetObjectType_VENDOR_PROPRIETARY_VALUE), BACnetPropertyIdentifier(BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE))
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := CastBACnetConstructedDataElement(_value)
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetTimeValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetTimeValue(timeValue, value), nil
}

func (m *BACnetTimeValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetTimeValue"); pushErr != nil {
		return pushErr
	}

	// Simple Field (timeValue)
	if pushErr := writeBuffer.PushContext("timeValue"); pushErr != nil {
		return pushErr
	}
	_timeValueErr := m.TimeValue.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("timeValue"); popErr != nil {
		return popErr
	}
	if _timeValueErr != nil {
		return errors.Wrap(_timeValueErr, "Error serializing 'timeValue' field")
	}

	// Simple Field (value)
	if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
		return pushErr
	}
	_valueErr := m.Value.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("value"); popErr != nil {
		return popErr
	}
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTimeValue"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetTimeValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
