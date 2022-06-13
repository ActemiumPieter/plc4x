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

// BACnetPropertyAccessResultAccessResultPropertyValue is the data-structure of this message
type BACnetPropertyAccessResultAccessResultPropertyValue struct {
	*BACnetPropertyAccessResultAccessResult
	PropertyValue *BACnetConstructedData

	// Arguments.
	ObjectTypeArgument         BACnetObjectType
	PropertyIdentifierArgument BACnetPropertyIdentifier
	PropertyArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetPropertyAccessResultAccessResultPropertyValue is the corresponding interface of BACnetPropertyAccessResultAccessResultPropertyValue
type IBACnetPropertyAccessResultAccessResultPropertyValue interface {
	IBACnetPropertyAccessResultAccessResult
	// GetPropertyValue returns PropertyValue (property field)
	GetPropertyValue() *BACnetConstructedData
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetPropertyAccessResultAccessResultPropertyValue) InitializeParent(parent *BACnetPropertyAccessResultAccessResult, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyAccessResultAccessResult.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyAccessResultAccessResultPropertyValue) GetParent() *BACnetPropertyAccessResultAccessResult {
	return m.BACnetPropertyAccessResultAccessResult
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyAccessResultAccessResultPropertyValue) GetPropertyValue() *BACnetConstructedData {
	return m.PropertyValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyAccessResultAccessResultPropertyValue factory function for BACnetPropertyAccessResultAccessResultPropertyValue
func NewBACnetPropertyAccessResultAccessResultPropertyValue(propertyValue *BACnetConstructedData, peekedTagHeader *BACnetTagHeader, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, propertyArrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetPropertyAccessResultAccessResultPropertyValue {
	_result := &BACnetPropertyAccessResultAccessResultPropertyValue{
		PropertyValue:                          propertyValue,
		BACnetPropertyAccessResultAccessResult: NewBACnetPropertyAccessResultAccessResult(peekedTagHeader, objectTypeArgument, propertyIdentifierArgument, propertyArrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyAccessResultAccessResultPropertyValue(structType interface{}) *BACnetPropertyAccessResultAccessResultPropertyValue {
	if casted, ok := structType.(BACnetPropertyAccessResultAccessResultPropertyValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyAccessResultAccessResultPropertyValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyAccessResultAccessResult); ok {
		return CastBACnetPropertyAccessResultAccessResultPropertyValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyAccessResultAccessResult); ok {
		return CastBACnetPropertyAccessResultAccessResultPropertyValue(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyAccessResultAccessResultPropertyValue) GetTypeName() string {
	return "BACnetPropertyAccessResultAccessResultPropertyValue"
}

func (m *BACnetPropertyAccessResultAccessResultPropertyValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyAccessResultAccessResultPropertyValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (propertyValue)
	lengthInBits += m.PropertyValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyAccessResultAccessResultPropertyValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyAccessResultAccessResultPropertyValueParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, propertyArrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetPropertyAccessResultAccessResultPropertyValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyAccessResultAccessResultPropertyValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyAccessResultAccessResultPropertyValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (propertyValue)
	if pullErr := readBuffer.PullContext("propertyValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyValue")
	}
	_propertyValue, _propertyValueErr := BACnetConstructedDataParse(readBuffer, uint8(uint8(4)), BACnetObjectType(objectTypeArgument), BACnetPropertyIdentifier(propertyIdentifierArgument), propertyArrayIndexArgument)
	if _propertyValueErr != nil {
		return nil, errors.Wrap(_propertyValueErr, "Error parsing 'propertyValue' field")
	}
	propertyValue := CastBACnetConstructedData(_propertyValue)
	if closeErr := readBuffer.CloseContext("propertyValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyAccessResultAccessResultPropertyValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyAccessResultAccessResultPropertyValue")
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyAccessResultAccessResultPropertyValue{
		PropertyValue:                          CastBACnetConstructedData(propertyValue),
		BACnetPropertyAccessResultAccessResult: &BACnetPropertyAccessResultAccessResult{},
	}
	_child.BACnetPropertyAccessResultAccessResult.Child = _child
	return _child, nil
}

func (m *BACnetPropertyAccessResultAccessResultPropertyValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyAccessResultAccessResultPropertyValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyAccessResultAccessResultPropertyValue")
		}

		// Simple Field (propertyValue)
		if pushErr := writeBuffer.PushContext("propertyValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for propertyValue")
		}
		_propertyValueErr := writeBuffer.WriteSerializable(m.PropertyValue)
		if popErr := writeBuffer.PopContext("propertyValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for propertyValue")
		}
		if _propertyValueErr != nil {
			return errors.Wrap(_propertyValueErr, "Error serializing 'propertyValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyAccessResultAccessResultPropertyValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyAccessResultAccessResultPropertyValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyAccessResultAccessResultPropertyValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
