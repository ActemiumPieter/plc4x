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

// BACnetConstructedDataBACnetIPMulticastAddress is the data-structure of this message
type BACnetConstructedDataBACnetIPMulticastAddress struct {
	*BACnetConstructedData
	IpMulticastAddress *BACnetApplicationTagOctetString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataBACnetIPMulticastAddress is the corresponding interface of BACnetConstructedDataBACnetIPMulticastAddress
type IBACnetConstructedDataBACnetIPMulticastAddress interface {
	IBACnetConstructedData
	// GetIpMulticastAddress returns IpMulticastAddress (property field)
	GetIpMulticastAddress() *BACnetApplicationTagOctetString
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

func (m *BACnetConstructedDataBACnetIPMulticastAddress) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataBACnetIPMulticastAddress) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACNET_IP_MULTICAST_ADDRESS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataBACnetIPMulticastAddress) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataBACnetIPMulticastAddress) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataBACnetIPMulticastAddress) GetIpMulticastAddress() *BACnetApplicationTagOctetString {
	return m.IpMulticastAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBACnetIPMulticastAddress factory function for BACnetConstructedDataBACnetIPMulticastAddress
func NewBACnetConstructedDataBACnetIPMulticastAddress(ipMulticastAddress *BACnetApplicationTagOctetString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataBACnetIPMulticastAddress {
	_result := &BACnetConstructedDataBACnetIPMulticastAddress{
		IpMulticastAddress:    ipMulticastAddress,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataBACnetIPMulticastAddress(structType interface{}) *BACnetConstructedDataBACnetIPMulticastAddress {
	if casted, ok := structType.(BACnetConstructedDataBACnetIPMulticastAddress); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBACnetIPMulticastAddress); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataBACnetIPMulticastAddress(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataBACnetIPMulticastAddress(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataBACnetIPMulticastAddress) GetTypeName() string {
	return "BACnetConstructedDataBACnetIPMulticastAddress"
}

func (m *BACnetConstructedDataBACnetIPMulticastAddress) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataBACnetIPMulticastAddress) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ipMulticastAddress)
	lengthInBits += m.IpMulticastAddress.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataBACnetIPMulticastAddress) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBACnetIPMulticastAddressParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataBACnetIPMulticastAddress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBACnetIPMulticastAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBACnetIPMulticastAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipMulticastAddress)
	if pullErr := readBuffer.PullContext("ipMulticastAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipMulticastAddress")
	}
	_ipMulticastAddress, _ipMulticastAddressErr := BACnetApplicationTagParse(readBuffer)
	if _ipMulticastAddressErr != nil {
		return nil, errors.Wrap(_ipMulticastAddressErr, "Error parsing 'ipMulticastAddress' field")
	}
	ipMulticastAddress := CastBACnetApplicationTagOctetString(_ipMulticastAddress)
	if closeErr := readBuffer.CloseContext("ipMulticastAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipMulticastAddress")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBACnetIPMulticastAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBACnetIPMulticastAddress")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataBACnetIPMulticastAddress{
		IpMulticastAddress:    CastBACnetApplicationTagOctetString(ipMulticastAddress),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataBACnetIPMulticastAddress) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBACnetIPMulticastAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBACnetIPMulticastAddress")
		}

		// Simple Field (ipMulticastAddress)
		if pushErr := writeBuffer.PushContext("ipMulticastAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ipMulticastAddress")
		}
		_ipMulticastAddressErr := writeBuffer.WriteSerializable(m.IpMulticastAddress)
		if popErr := writeBuffer.PopContext("ipMulticastAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ipMulticastAddress")
		}
		if _ipMulticastAddressErr != nil {
			return errors.Wrap(_ipMulticastAddressErr, "Error serializing 'ipMulticastAddress' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBACnetIPMulticastAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBACnetIPMulticastAddress")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataBACnetIPMulticastAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
