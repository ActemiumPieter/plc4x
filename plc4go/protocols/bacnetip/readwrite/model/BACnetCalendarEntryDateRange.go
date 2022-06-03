/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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

// BACnetCalendarEntryDateRange is the data-structure of this message
type BACnetCalendarEntryDateRange struct {
	*BACnetCalendarEntry
	DateRange *BACnetDateRangeEnclosed

	// Arguments.
	TagNumber uint8
}

// IBACnetCalendarEntryDateRange is the corresponding interface of BACnetCalendarEntryDateRange
type IBACnetCalendarEntryDateRange interface {
	IBACnetCalendarEntry
	// GetDateRange returns DateRange (property field)
	GetDateRange() *BACnetDateRangeEnclosed
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

func (m *BACnetCalendarEntryDateRange) InitializeParent(parent *BACnetCalendarEntry, peekedTagHeader *BACnetTagHeader) {
	m.BACnetCalendarEntry.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetCalendarEntryDateRange) GetParent() *BACnetCalendarEntry {
	return m.BACnetCalendarEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetCalendarEntryDateRange) GetDateRange() *BACnetDateRangeEnclosed {
	return m.DateRange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCalendarEntryDateRange factory function for BACnetCalendarEntryDateRange
func NewBACnetCalendarEntryDateRange(dateRange *BACnetDateRangeEnclosed, peekedTagHeader *BACnetTagHeader, tagNumber uint8) *BACnetCalendarEntryDateRange {
	_result := &BACnetCalendarEntryDateRange{
		DateRange:           dateRange,
		BACnetCalendarEntry: NewBACnetCalendarEntry(peekedTagHeader, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetCalendarEntryDateRange(structType interface{}) *BACnetCalendarEntryDateRange {
	if casted, ok := structType.(BACnetCalendarEntryDateRange); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetCalendarEntryDateRange); ok {
		return casted
	}
	if casted, ok := structType.(BACnetCalendarEntry); ok {
		return CastBACnetCalendarEntryDateRange(casted.Child)
	}
	if casted, ok := structType.(*BACnetCalendarEntry); ok {
		return CastBACnetCalendarEntryDateRange(casted.Child)
	}
	return nil
}

func (m *BACnetCalendarEntryDateRange) GetTypeName() string {
	return "BACnetCalendarEntryDateRange"
}

func (m *BACnetCalendarEntryDateRange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetCalendarEntryDateRange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dateRange)
	lengthInBits += m.DateRange.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetCalendarEntryDateRange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetCalendarEntryDateRangeParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetCalendarEntryDateRange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCalendarEntryDateRange"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dateRange)
	if pullErr := readBuffer.PullContext("dateRange"); pullErr != nil {
		return nil, pullErr
	}
	_dateRange, _dateRangeErr := BACnetDateRangeEnclosedParse(readBuffer, uint8(uint8(1)))
	if _dateRangeErr != nil {
		return nil, errors.Wrap(_dateRangeErr, "Error parsing 'dateRange' field")
	}
	dateRange := CastBACnetDateRangeEnclosed(_dateRange)
	if closeErr := readBuffer.CloseContext("dateRange"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetCalendarEntryDateRange"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetCalendarEntryDateRange{
		DateRange:           CastBACnetDateRangeEnclosed(dateRange),
		BACnetCalendarEntry: &BACnetCalendarEntry{},
	}
	_child.BACnetCalendarEntry.Child = _child
	return _child, nil
}

func (m *BACnetCalendarEntryDateRange) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetCalendarEntryDateRange"); pushErr != nil {
			return pushErr
		}

		// Simple Field (dateRange)
		if pushErr := writeBuffer.PushContext("dateRange"); pushErr != nil {
			return pushErr
		}
		_dateRangeErr := m.DateRange.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("dateRange"); popErr != nil {
			return popErr
		}
		if _dateRangeErr != nil {
			return errors.Wrap(_dateRangeErr, "Error serializing 'dateRange' field")
		}

		if popErr := writeBuffer.PopContext("BACnetCalendarEntryDateRange"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetCalendarEntryDateRange) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
