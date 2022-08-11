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

// BACnetConstructedDataVendorName is the corresponding interface of BACnetConstructedDataVendorName
type BACnetConstructedDataVendorName interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetVendorName returns VendorName (property field)
	GetVendorName() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataVendorNameExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataVendorName.
// This is useful for switch cases.
type BACnetConstructedDataVendorNameExactly interface {
	BACnetConstructedDataVendorName
	isBACnetConstructedDataVendorName() bool
}

// _BACnetConstructedDataVendorName is the data-structure of this message
type _BACnetConstructedDataVendorName struct {
	*_BACnetConstructedData
	VendorName BACnetApplicationTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataVendorName) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataVendorName) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VENDOR_NAME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataVendorName) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataVendorName) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataVendorName) GetVendorName() BACnetApplicationTagCharacterString {
	return m.VendorName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataVendorName) GetActualValue() BACnetApplicationTagCharacterString {
	return CastBACnetApplicationTagCharacterString(m.GetVendorName())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataVendorName factory function for _BACnetConstructedDataVendorName
func NewBACnetConstructedDataVendorName(vendorName BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataVendorName {
	_result := &_BACnetConstructedDataVendorName{
		VendorName:             vendorName,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataVendorName(structType interface{}) BACnetConstructedDataVendorName {
	if casted, ok := structType.(BACnetConstructedDataVendorName); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataVendorName); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataVendorName) GetTypeName() string {
	return "BACnetConstructedDataVendorName"
}

func (m *_BACnetConstructedDataVendorName) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataVendorName) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (vendorName)
	lengthInBits += m.VendorName.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataVendorName) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataVendorNameParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataVendorName, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataVendorName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataVendorName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (vendorName)
	if pullErr := readBuffer.PullContext("vendorName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vendorName")
	}
	_vendorName, _vendorNameErr := BACnetApplicationTagParse(readBuffer)
	if _vendorNameErr != nil {
		return nil, errors.Wrap(_vendorNameErr, "Error parsing 'vendorName' field of BACnetConstructedDataVendorName")
	}
	vendorName := _vendorName.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("vendorName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vendorName")
	}

	// Virtual field
	_actualValue := vendorName
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataVendorName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataVendorName")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataVendorName{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		VendorName: vendorName,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataVendorName) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataVendorName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataVendorName")
		}

		// Simple Field (vendorName)
		if pushErr := writeBuffer.PushContext("vendorName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vendorName")
		}
		_vendorNameErr := writeBuffer.WriteSerializable(m.GetVendorName())
		if popErr := writeBuffer.PopContext("vendorName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vendorName")
		}
		if _vendorNameErr != nil {
			return errors.Wrap(_vendorNameErr, "Error serializing 'vendorName' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataVendorName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataVendorName")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataVendorName) isBACnetConstructedDataVendorName() bool {
	return true
}

func (m *_BACnetConstructedDataVendorName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
