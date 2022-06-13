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

// BACnetConstructedDataBelongsTo is the data-structure of this message
type BACnetConstructedDataBelongsTo struct {
	*BACnetConstructedData
	BelongsTo *BACnetDeviceObjectReference

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataBelongsTo is the corresponding interface of BACnetConstructedDataBelongsTo
type IBACnetConstructedDataBelongsTo interface {
	IBACnetConstructedData
	// GetBelongsTo returns BelongsTo (property field)
	GetBelongsTo() *BACnetDeviceObjectReference
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

func (m *BACnetConstructedDataBelongsTo) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataBelongsTo) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BELONGS_TO
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataBelongsTo) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataBelongsTo) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataBelongsTo) GetBelongsTo() *BACnetDeviceObjectReference {
	return m.BelongsTo
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBelongsTo factory function for BACnetConstructedDataBelongsTo
func NewBACnetConstructedDataBelongsTo(belongsTo *BACnetDeviceObjectReference, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataBelongsTo {
	_result := &BACnetConstructedDataBelongsTo{
		BelongsTo:             belongsTo,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataBelongsTo(structType interface{}) *BACnetConstructedDataBelongsTo {
	if casted, ok := structType.(BACnetConstructedDataBelongsTo); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBelongsTo); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataBelongsTo(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataBelongsTo(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataBelongsTo) GetTypeName() string {
	return "BACnetConstructedDataBelongsTo"
}

func (m *BACnetConstructedDataBelongsTo) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataBelongsTo) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (belongsTo)
	lengthInBits += m.BelongsTo.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataBelongsTo) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBelongsToParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataBelongsTo, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBelongsTo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBelongsTo")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (belongsTo)
	if pullErr := readBuffer.PullContext("belongsTo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for belongsTo")
	}
	_belongsTo, _belongsToErr := BACnetDeviceObjectReferenceParse(readBuffer)
	if _belongsToErr != nil {
		return nil, errors.Wrap(_belongsToErr, "Error parsing 'belongsTo' field")
	}
	belongsTo := CastBACnetDeviceObjectReference(_belongsTo)
	if closeErr := readBuffer.CloseContext("belongsTo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for belongsTo")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBelongsTo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBelongsTo")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataBelongsTo{
		BelongsTo:             CastBACnetDeviceObjectReference(belongsTo),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataBelongsTo) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBelongsTo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBelongsTo")
		}

		// Simple Field (belongsTo)
		if pushErr := writeBuffer.PushContext("belongsTo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for belongsTo")
		}
		_belongsToErr := writeBuffer.WriteSerializable(m.BelongsTo)
		if popErr := writeBuffer.PopContext("belongsTo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for belongsTo")
		}
		if _belongsToErr != nil {
			return errors.Wrap(_belongsToErr, "Error serializing 'belongsTo' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBelongsTo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBelongsTo")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataBelongsTo) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
