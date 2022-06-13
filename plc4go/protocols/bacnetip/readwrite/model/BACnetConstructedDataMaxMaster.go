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

// BACnetConstructedDataMaxMaster is the data-structure of this message
type BACnetConstructedDataMaxMaster struct {
	*BACnetConstructedData
	MaxMaster *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataMaxMaster is the corresponding interface of BACnetConstructedDataMaxMaster
type IBACnetConstructedDataMaxMaster interface {
	IBACnetConstructedData
	// GetMaxMaster returns MaxMaster (property field)
	GetMaxMaster() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataMaxMaster) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataMaxMaster) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_MASTER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataMaxMaster) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataMaxMaster) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataMaxMaster) GetMaxMaster() *BACnetApplicationTagUnsignedInteger {
	return m.MaxMaster
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMaxMaster factory function for BACnetConstructedDataMaxMaster
func NewBACnetConstructedDataMaxMaster(maxMaster *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataMaxMaster {
	_result := &BACnetConstructedDataMaxMaster{
		MaxMaster:             maxMaster,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataMaxMaster(structType interface{}) *BACnetConstructedDataMaxMaster {
	if casted, ok := structType.(BACnetConstructedDataMaxMaster); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaxMaster); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataMaxMaster(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataMaxMaster(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataMaxMaster) GetTypeName() string {
	return "BACnetConstructedDataMaxMaster"
}

func (m *BACnetConstructedDataMaxMaster) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataMaxMaster) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maxMaster)
	lengthInBits += m.MaxMaster.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataMaxMaster) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMaxMasterParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataMaxMaster, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaxMaster"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaxMaster")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maxMaster)
	if pullErr := readBuffer.PullContext("maxMaster"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maxMaster")
	}
	_maxMaster, _maxMasterErr := BACnetApplicationTagParse(readBuffer)
	if _maxMasterErr != nil {
		return nil, errors.Wrap(_maxMasterErr, "Error parsing 'maxMaster' field")
	}
	maxMaster := CastBACnetApplicationTagUnsignedInteger(_maxMaster)
	if closeErr := readBuffer.CloseContext("maxMaster"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maxMaster")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaxMaster"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaxMaster")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataMaxMaster{
		MaxMaster:             CastBACnetApplicationTagUnsignedInteger(maxMaster),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataMaxMaster) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaxMaster"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaxMaster")
		}

		// Simple Field (maxMaster)
		if pushErr := writeBuffer.PushContext("maxMaster"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maxMaster")
		}
		_maxMasterErr := writeBuffer.WriteSerializable(m.MaxMaster)
		if popErr := writeBuffer.PopContext("maxMaster"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maxMaster")
		}
		if _maxMasterErr != nil {
			return errors.Wrap(_maxMasterErr, "Error serializing 'maxMaster' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaxMaster"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaxMaster")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataMaxMaster) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
