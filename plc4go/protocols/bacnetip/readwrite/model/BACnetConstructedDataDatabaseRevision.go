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

// BACnetConstructedDataDatabaseRevision is the data-structure of this message
type BACnetConstructedDataDatabaseRevision struct {
	*BACnetConstructedData
	DatabaseRevision *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataDatabaseRevision is the corresponding interface of BACnetConstructedDataDatabaseRevision
type IBACnetConstructedDataDatabaseRevision interface {
	IBACnetConstructedData
	// GetDatabaseRevision returns DatabaseRevision (property field)
	GetDatabaseRevision() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataDatabaseRevision) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataDatabaseRevision) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DATABASE_REVISION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataDatabaseRevision) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataDatabaseRevision) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataDatabaseRevision) GetDatabaseRevision() *BACnetApplicationTagUnsignedInteger {
	return m.DatabaseRevision
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDatabaseRevision factory function for BACnetConstructedDataDatabaseRevision
func NewBACnetConstructedDataDatabaseRevision(databaseRevision *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataDatabaseRevision {
	_result := &BACnetConstructedDataDatabaseRevision{
		DatabaseRevision:      databaseRevision,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataDatabaseRevision(structType interface{}) *BACnetConstructedDataDatabaseRevision {
	if casted, ok := structType.(BACnetConstructedDataDatabaseRevision); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDatabaseRevision); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataDatabaseRevision(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataDatabaseRevision(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataDatabaseRevision) GetTypeName() string {
	return "BACnetConstructedDataDatabaseRevision"
}

func (m *BACnetConstructedDataDatabaseRevision) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataDatabaseRevision) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (databaseRevision)
	lengthInBits += m.DatabaseRevision.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataDatabaseRevision) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDatabaseRevisionParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataDatabaseRevision, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDatabaseRevision"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDatabaseRevision")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (databaseRevision)
	if pullErr := readBuffer.PullContext("databaseRevision"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for databaseRevision")
	}
	_databaseRevision, _databaseRevisionErr := BACnetApplicationTagParse(readBuffer)
	if _databaseRevisionErr != nil {
		return nil, errors.Wrap(_databaseRevisionErr, "Error parsing 'databaseRevision' field")
	}
	databaseRevision := CastBACnetApplicationTagUnsignedInteger(_databaseRevision)
	if closeErr := readBuffer.CloseContext("databaseRevision"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for databaseRevision")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDatabaseRevision"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDatabaseRevision")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataDatabaseRevision{
		DatabaseRevision:      CastBACnetApplicationTagUnsignedInteger(databaseRevision),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataDatabaseRevision) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDatabaseRevision"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDatabaseRevision")
		}

		// Simple Field (databaseRevision)
		if pushErr := writeBuffer.PushContext("databaseRevision"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for databaseRevision")
		}
		_databaseRevisionErr := writeBuffer.WriteSerializable(m.DatabaseRevision)
		if popErr := writeBuffer.PopContext("databaseRevision"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for databaseRevision")
		}
		if _databaseRevisionErr != nil {
			return errors.Wrap(_databaseRevisionErr, "Error serializing 'databaseRevision' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDatabaseRevision"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDatabaseRevision")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataDatabaseRevision) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
