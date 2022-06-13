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

// BACnetConstructedDataProportionalConstantUnits is the data-structure of this message
type BACnetConstructedDataProportionalConstantUnits struct {
	*BACnetConstructedData
	Units *BACnetEngineeringUnitsTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataProportionalConstantUnits is the corresponding interface of BACnetConstructedDataProportionalConstantUnits
type IBACnetConstructedDataProportionalConstantUnits interface {
	IBACnetConstructedData
	// GetUnits returns Units (property field)
	GetUnits() *BACnetEngineeringUnitsTagged
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

func (m *BACnetConstructedDataProportionalConstantUnits) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataProportionalConstantUnits) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROPORTIONAL_CONSTANT_UNITS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataProportionalConstantUnits) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataProportionalConstantUnits) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataProportionalConstantUnits) GetUnits() *BACnetEngineeringUnitsTagged {
	return m.Units
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataProportionalConstantUnits factory function for BACnetConstructedDataProportionalConstantUnits
func NewBACnetConstructedDataProportionalConstantUnits(units *BACnetEngineeringUnitsTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataProportionalConstantUnits {
	_result := &BACnetConstructedDataProportionalConstantUnits{
		Units:                 units,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataProportionalConstantUnits(structType interface{}) *BACnetConstructedDataProportionalConstantUnits {
	if casted, ok := structType.(BACnetConstructedDataProportionalConstantUnits); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProportionalConstantUnits); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataProportionalConstantUnits(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataProportionalConstantUnits(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataProportionalConstantUnits) GetTypeName() string {
	return "BACnetConstructedDataProportionalConstantUnits"
}

func (m *BACnetConstructedDataProportionalConstantUnits) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataProportionalConstantUnits) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (units)
	lengthInBits += m.Units.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataProportionalConstantUnits) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataProportionalConstantUnitsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataProportionalConstantUnits, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProportionalConstantUnits"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProportionalConstantUnits")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (units)
	if pullErr := readBuffer.PullContext("units"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for units")
	}
	_units, _unitsErr := BACnetEngineeringUnitsTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _unitsErr != nil {
		return nil, errors.Wrap(_unitsErr, "Error parsing 'units' field")
	}
	units := CastBACnetEngineeringUnitsTagged(_units)
	if closeErr := readBuffer.CloseContext("units"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for units")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProportionalConstantUnits"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProportionalConstantUnits")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataProportionalConstantUnits{
		Units:                 CastBACnetEngineeringUnitsTagged(units),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataProportionalConstantUnits) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProportionalConstantUnits"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProportionalConstantUnits")
		}

		// Simple Field (units)
		if pushErr := writeBuffer.PushContext("units"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for units")
		}
		_unitsErr := writeBuffer.WriteSerializable(m.Units)
		if popErr := writeBuffer.PopContext("units"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for units")
		}
		if _unitsErr != nil {
			return errors.Wrap(_unitsErr, "Error serializing 'units' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProportionalConstantUnits"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProportionalConstantUnits")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataProportionalConstantUnits) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
