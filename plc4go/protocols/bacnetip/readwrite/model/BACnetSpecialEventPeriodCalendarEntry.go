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

// BACnetSpecialEventPeriodCalendarEntry is the data-structure of this message
type BACnetSpecialEventPeriodCalendarEntry struct {
	*BACnetSpecialEventPeriod
	CalendarEntry *BACnetCalendarEntryEnclosed
}

// IBACnetSpecialEventPeriodCalendarEntry is the corresponding interface of BACnetSpecialEventPeriodCalendarEntry
type IBACnetSpecialEventPeriodCalendarEntry interface {
	IBACnetSpecialEventPeriod
	// GetCalendarEntry returns CalendarEntry (property field)
	GetCalendarEntry() *BACnetCalendarEntryEnclosed
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetSpecialEventPeriodCalendarEntry) InitializeParent(parent *BACnetSpecialEventPeriod, peekedTagHeader *BACnetTagHeader) {
	m.BACnetSpecialEventPeriod.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetSpecialEventPeriodCalendarEntry) GetParent() *BACnetSpecialEventPeriod {
	return m.BACnetSpecialEventPeriod
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetSpecialEventPeriodCalendarEntry) GetCalendarEntry() *BACnetCalendarEntryEnclosed {
	return m.CalendarEntry
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetSpecialEventPeriodCalendarEntry factory function for BACnetSpecialEventPeriodCalendarEntry
func NewBACnetSpecialEventPeriodCalendarEntry(calendarEntry *BACnetCalendarEntryEnclosed, peekedTagHeader *BACnetTagHeader) *BACnetSpecialEventPeriodCalendarEntry {
	_result := &BACnetSpecialEventPeriodCalendarEntry{
		CalendarEntry:            calendarEntry,
		BACnetSpecialEventPeriod: NewBACnetSpecialEventPeriod(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetSpecialEventPeriodCalendarEntry(structType interface{}) *BACnetSpecialEventPeriodCalendarEntry {
	if casted, ok := structType.(BACnetSpecialEventPeriodCalendarEntry); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetSpecialEventPeriodCalendarEntry); ok {
		return casted
	}
	if casted, ok := structType.(BACnetSpecialEventPeriod); ok {
		return CastBACnetSpecialEventPeriodCalendarEntry(casted.Child)
	}
	if casted, ok := structType.(*BACnetSpecialEventPeriod); ok {
		return CastBACnetSpecialEventPeriodCalendarEntry(casted.Child)
	}
	return nil
}

func (m *BACnetSpecialEventPeriodCalendarEntry) GetTypeName() string {
	return "BACnetSpecialEventPeriodCalendarEntry"
}

func (m *BACnetSpecialEventPeriodCalendarEntry) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetSpecialEventPeriodCalendarEntry) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (calendarEntry)
	lengthInBits += m.CalendarEntry.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetSpecialEventPeriodCalendarEntry) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetSpecialEventPeriodCalendarEntryParse(readBuffer utils.ReadBuffer) (*BACnetSpecialEventPeriodCalendarEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetSpecialEventPeriodCalendarEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetSpecialEventPeriodCalendarEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (calendarEntry)
	if pullErr := readBuffer.PullContext("calendarEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for calendarEntry")
	}
	_calendarEntry, _calendarEntryErr := BACnetCalendarEntryEnclosedParse(readBuffer, uint8(uint8(0)))
	if _calendarEntryErr != nil {
		return nil, errors.Wrap(_calendarEntryErr, "Error parsing 'calendarEntry' field")
	}
	calendarEntry := CastBACnetCalendarEntryEnclosed(_calendarEntry)
	if closeErr := readBuffer.CloseContext("calendarEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for calendarEntry")
	}

	if closeErr := readBuffer.CloseContext("BACnetSpecialEventPeriodCalendarEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetSpecialEventPeriodCalendarEntry")
	}

	// Create a partially initialized instance
	_child := &BACnetSpecialEventPeriodCalendarEntry{
		CalendarEntry:            CastBACnetCalendarEntryEnclosed(calendarEntry),
		BACnetSpecialEventPeriod: &BACnetSpecialEventPeriod{},
	}
	_child.BACnetSpecialEventPeriod.Child = _child
	return _child, nil
}

func (m *BACnetSpecialEventPeriodCalendarEntry) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetSpecialEventPeriodCalendarEntry"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetSpecialEventPeriodCalendarEntry")
		}

		// Simple Field (calendarEntry)
		if pushErr := writeBuffer.PushContext("calendarEntry"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for calendarEntry")
		}
		_calendarEntryErr := writeBuffer.WriteSerializable(m.CalendarEntry)
		if popErr := writeBuffer.PopContext("calendarEntry"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for calendarEntry")
		}
		if _calendarEntryErr != nil {
			return errors.Wrap(_calendarEntryErr, "Error serializing 'calendarEntry' field")
		}

		if popErr := writeBuffer.PopContext("BACnetSpecialEventPeriodCalendarEntry"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetSpecialEventPeriodCalendarEntry")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetSpecialEventPeriodCalendarEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
