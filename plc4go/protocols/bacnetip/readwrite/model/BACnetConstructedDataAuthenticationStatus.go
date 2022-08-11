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

// BACnetConstructedDataAuthenticationStatus is the corresponding interface of BACnetConstructedDataAuthenticationStatus
type BACnetConstructedDataAuthenticationStatus interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAuthenticationStatus returns AuthenticationStatus (property field)
	GetAuthenticationStatus() BACnetAuthenticationStatusTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAuthenticationStatusTagged
}

// BACnetConstructedDataAuthenticationStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAuthenticationStatus.
// This is useful for switch cases.
type BACnetConstructedDataAuthenticationStatusExactly interface {
	BACnetConstructedDataAuthenticationStatus
	isBACnetConstructedDataAuthenticationStatus() bool
}

// _BACnetConstructedDataAuthenticationStatus is the data-structure of this message
type _BACnetConstructedDataAuthenticationStatus struct {
	*_BACnetConstructedData
	AuthenticationStatus BACnetAuthenticationStatusTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAuthenticationStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_AUTHENTICATION_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAuthenticationStatus) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAuthenticationStatus) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationStatus) GetAuthenticationStatus() BACnetAuthenticationStatusTagged {
	return m.AuthenticationStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationStatus) GetActualValue() BACnetAuthenticationStatusTagged {
	return CastBACnetAuthenticationStatusTagged(m.GetAuthenticationStatus())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAuthenticationStatus factory function for _BACnetConstructedDataAuthenticationStatus
func NewBACnetConstructedDataAuthenticationStatus(authenticationStatus BACnetAuthenticationStatusTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAuthenticationStatus {
	_result := &_BACnetConstructedDataAuthenticationStatus{
		AuthenticationStatus:   authenticationStatus,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAuthenticationStatus(structType interface{}) BACnetConstructedDataAuthenticationStatus {
	if casted, ok := structType.(BACnetConstructedDataAuthenticationStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAuthenticationStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAuthenticationStatus) GetTypeName() string {
	return "BACnetConstructedDataAuthenticationStatus"
}

func (m *_BACnetConstructedDataAuthenticationStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAuthenticationStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (authenticationStatus)
	lengthInBits += m.AuthenticationStatus.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAuthenticationStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAuthenticationStatusParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAuthenticationStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAuthenticationStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAuthenticationStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (authenticationStatus)
	if pullErr := readBuffer.PullContext("authenticationStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for authenticationStatus")
	}
	_authenticationStatus, _authenticationStatusErr := BACnetAuthenticationStatusTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _authenticationStatusErr != nil {
		return nil, errors.Wrap(_authenticationStatusErr, "Error parsing 'authenticationStatus' field of BACnetConstructedDataAuthenticationStatus")
	}
	authenticationStatus := _authenticationStatus.(BACnetAuthenticationStatusTagged)
	if closeErr := readBuffer.CloseContext("authenticationStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for authenticationStatus")
	}

	// Virtual field
	_actualValue := authenticationStatus
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAuthenticationStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAuthenticationStatus")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAuthenticationStatus{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		AuthenticationStatus: authenticationStatus,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAuthenticationStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAuthenticationStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAuthenticationStatus")
		}

		// Simple Field (authenticationStatus)
		if pushErr := writeBuffer.PushContext("authenticationStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for authenticationStatus")
		}
		_authenticationStatusErr := writeBuffer.WriteSerializable(m.GetAuthenticationStatus())
		if popErr := writeBuffer.PopContext("authenticationStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for authenticationStatus")
		}
		if _authenticationStatusErr != nil {
			return errors.Wrap(_authenticationStatusErr, "Error serializing 'authenticationStatus' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAuthenticationStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAuthenticationStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAuthenticationStatus) isBACnetConstructedDataAuthenticationStatus() bool {
	return true
}

func (m *_BACnetConstructedDataAuthenticationStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
