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

// BACnetConstructedDataLifeSafetyPointAlarmValues is the data-structure of this message
type BACnetConstructedDataLifeSafetyPointAlarmValues struct {
	*BACnetConstructedData
	AlarmValues []*BACnetLifeSafetyStateTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLifeSafetyPointAlarmValues is the corresponding interface of BACnetConstructedDataLifeSafetyPointAlarmValues
type IBACnetConstructedDataLifeSafetyPointAlarmValues interface {
	IBACnetConstructedData
	// GetAlarmValues returns AlarmValues (property field)
	GetAlarmValues() []*BACnetLifeSafetyStateTagged
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

func (m *BACnetConstructedDataLifeSafetyPointAlarmValues) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LIFE_SAFETY_POINT
}

func (m *BACnetConstructedDataLifeSafetyPointAlarmValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALARM_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLifeSafetyPointAlarmValues) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLifeSafetyPointAlarmValues) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLifeSafetyPointAlarmValues) GetAlarmValues() []*BACnetLifeSafetyStateTagged {
	return m.AlarmValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLifeSafetyPointAlarmValues factory function for BACnetConstructedDataLifeSafetyPointAlarmValues
func NewBACnetConstructedDataLifeSafetyPointAlarmValues(alarmValues []*BACnetLifeSafetyStateTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLifeSafetyPointAlarmValues {
	_result := &BACnetConstructedDataLifeSafetyPointAlarmValues{
		AlarmValues:           alarmValues,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLifeSafetyPointAlarmValues(structType interface{}) *BACnetConstructedDataLifeSafetyPointAlarmValues {
	if casted, ok := structType.(BACnetConstructedDataLifeSafetyPointAlarmValues); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLifeSafetyPointAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLifeSafetyPointAlarmValues(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLifeSafetyPointAlarmValues(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLifeSafetyPointAlarmValues) GetTypeName() string {
	return "BACnetConstructedDataLifeSafetyPointAlarmValues"
}

func (m *BACnetConstructedDataLifeSafetyPointAlarmValues) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLifeSafetyPointAlarmValues) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.AlarmValues) > 0 {
		for _, element := range m.AlarmValues {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataLifeSafetyPointAlarmValues) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLifeSafetyPointAlarmValuesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLifeSafetyPointAlarmValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLifeSafetyPointAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLifeSafetyPointAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (alarmValues)
	if pullErr := readBuffer.PullContext("alarmValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for alarmValues")
	}
	// Terminated array
	alarmValues := make([]*BACnetLifeSafetyStateTagged, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetLifeSafetyStateTaggedParse(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'alarmValues' field")
			}
			alarmValues = append(alarmValues, CastBACnetLifeSafetyStateTagged(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("alarmValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for alarmValues")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLifeSafetyPointAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLifeSafetyPointAlarmValues")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLifeSafetyPointAlarmValues{
		AlarmValues:           alarmValues,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLifeSafetyPointAlarmValues) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLifeSafetyPointAlarmValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLifeSafetyPointAlarmValues")
		}

		// Array Field (alarmValues)
		if m.AlarmValues != nil {
			if pushErr := writeBuffer.PushContext("alarmValues", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for alarmValues")
			}
			for _, _element := range m.AlarmValues {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'alarmValues' field")
				}
			}
			if popErr := writeBuffer.PopContext("alarmValues", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for alarmValues")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLifeSafetyPointAlarmValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLifeSafetyPointAlarmValues")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLifeSafetyPointAlarmValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
