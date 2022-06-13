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

// BACnetConstructedDataNetworkPortMaxMaster is the data-structure of this message
type BACnetConstructedDataNetworkPortMaxMaster struct {
	*BACnetConstructedData
	MaxMaster *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataNetworkPortMaxMaster is the corresponding interface of BACnetConstructedDataNetworkPortMaxMaster
type IBACnetConstructedDataNetworkPortMaxMaster interface {
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

func (m *BACnetConstructedDataNetworkPortMaxMaster) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_NETWORK_PORT
}

func (m *BACnetConstructedDataNetworkPortMaxMaster) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_MASTER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataNetworkPortMaxMaster) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataNetworkPortMaxMaster) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataNetworkPortMaxMaster) GetMaxMaster() *BACnetApplicationTagUnsignedInteger {
	return m.MaxMaster
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataNetworkPortMaxMaster factory function for BACnetConstructedDataNetworkPortMaxMaster
func NewBACnetConstructedDataNetworkPortMaxMaster(maxMaster *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataNetworkPortMaxMaster {
	_result := &BACnetConstructedDataNetworkPortMaxMaster{
		MaxMaster:             maxMaster,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataNetworkPortMaxMaster(structType interface{}) *BACnetConstructedDataNetworkPortMaxMaster {
	if casted, ok := structType.(BACnetConstructedDataNetworkPortMaxMaster); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNetworkPortMaxMaster); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataNetworkPortMaxMaster(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataNetworkPortMaxMaster(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataNetworkPortMaxMaster) GetTypeName() string {
	return "BACnetConstructedDataNetworkPortMaxMaster"
}

func (m *BACnetConstructedDataNetworkPortMaxMaster) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataNetworkPortMaxMaster) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maxMaster)
	lengthInBits += m.MaxMaster.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataNetworkPortMaxMaster) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataNetworkPortMaxMasterParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataNetworkPortMaxMaster, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNetworkPortMaxMaster"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNetworkPortMaxMaster")
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

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNetworkPortMaxMaster"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNetworkPortMaxMaster")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataNetworkPortMaxMaster{
		MaxMaster:             CastBACnetApplicationTagUnsignedInteger(maxMaster),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataNetworkPortMaxMaster) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNetworkPortMaxMaster"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNetworkPortMaxMaster")
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

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNetworkPortMaxMaster"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNetworkPortMaxMaster")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataNetworkPortMaxMaster) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
