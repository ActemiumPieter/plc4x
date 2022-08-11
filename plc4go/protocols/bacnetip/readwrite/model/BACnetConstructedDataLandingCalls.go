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

// BACnetConstructedDataLandingCalls is the corresponding interface of BACnetConstructedDataLandingCalls
type BACnetConstructedDataLandingCalls interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLandingCallStatus returns LandingCallStatus (property field)
	GetLandingCallStatus() []BACnetLandingCallStatus
}

// BACnetConstructedDataLandingCallsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLandingCalls.
// This is useful for switch cases.
type BACnetConstructedDataLandingCallsExactly interface {
	BACnetConstructedDataLandingCalls
	isBACnetConstructedDataLandingCalls() bool
}

// _BACnetConstructedDataLandingCalls is the data-structure of this message
type _BACnetConstructedDataLandingCalls struct {
	*_BACnetConstructedData
	LandingCallStatus []BACnetLandingCallStatus
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLandingCalls) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLandingCalls) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LANDING_CALLS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLandingCalls) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLandingCalls) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLandingCalls) GetLandingCallStatus() []BACnetLandingCallStatus {
	return m.LandingCallStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLandingCalls factory function for _BACnetConstructedDataLandingCalls
func NewBACnetConstructedDataLandingCalls(landingCallStatus []BACnetLandingCallStatus, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLandingCalls {
	_result := &_BACnetConstructedDataLandingCalls{
		LandingCallStatus:      landingCallStatus,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLandingCalls(structType interface{}) BACnetConstructedDataLandingCalls {
	if casted, ok := structType.(BACnetConstructedDataLandingCalls); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLandingCalls); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLandingCalls) GetTypeName() string {
	return "BACnetConstructedDataLandingCalls"
}

func (m *_BACnetConstructedDataLandingCalls) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataLandingCalls) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.LandingCallStatus) > 0 {
		for _, element := range m.LandingCallStatus {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataLandingCalls) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLandingCallsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLandingCalls, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLandingCalls"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLandingCalls")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (landingCallStatus)
	if pullErr := readBuffer.PullContext("landingCallStatus", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for landingCallStatus")
	}
	// Terminated array
	var landingCallStatus []BACnetLandingCallStatus
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetLandingCallStatusParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'landingCallStatus' field of BACnetConstructedDataLandingCalls")
			}
			landingCallStatus = append(landingCallStatus, _item.(BACnetLandingCallStatus))

		}
	}
	if closeErr := readBuffer.CloseContext("landingCallStatus", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for landingCallStatus")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLandingCalls"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLandingCalls")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLandingCalls{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LandingCallStatus: landingCallStatus,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLandingCalls) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLandingCalls"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLandingCalls")
		}

		// Array Field (landingCallStatus)
		if pushErr := writeBuffer.PushContext("landingCallStatus", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for landingCallStatus")
		}
		for _, _element := range m.GetLandingCallStatus() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'landingCallStatus' field")
			}
		}
		if popErr := writeBuffer.PopContext("landingCallStatus", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for landingCallStatus")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLandingCalls"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLandingCalls")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLandingCalls) isBACnetConstructedDataLandingCalls() bool {
	return true
}

func (m *_BACnetConstructedDataLandingCalls) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
