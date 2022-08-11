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

// BACnetServiceAckAtomicReadFile is the corresponding interface of BACnetServiceAckAtomicReadFile
type BACnetServiceAckAtomicReadFile interface {
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetEndOfFile returns EndOfFile (property field)
	GetEndOfFile() BACnetApplicationTagBoolean
	// GetAccessMethod returns AccessMethod (property field)
	GetAccessMethod() BACnetServiceAckAtomicReadFileStreamOrRecord
}

// BACnetServiceAckAtomicReadFileExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckAtomicReadFile.
// This is useful for switch cases.
type BACnetServiceAckAtomicReadFileExactly interface {
	BACnetServiceAckAtomicReadFile
	isBACnetServiceAckAtomicReadFile() bool
}

// _BACnetServiceAckAtomicReadFile is the data-structure of this message
type _BACnetServiceAckAtomicReadFile struct {
	*_BACnetServiceAck
	EndOfFile    BACnetApplicationTagBoolean
	AccessMethod BACnetServiceAckAtomicReadFileStreamOrRecord
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckAtomicReadFile) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_ATOMIC_READ_FILE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckAtomicReadFile) InitializeParent(parent BACnetServiceAck) {}

func (m *_BACnetServiceAckAtomicReadFile) GetParent() BACnetServiceAck {
	return m._BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckAtomicReadFile) GetEndOfFile() BACnetApplicationTagBoolean {
	return m.EndOfFile
}

func (m *_BACnetServiceAckAtomicReadFile) GetAccessMethod() BACnetServiceAckAtomicReadFileStreamOrRecord {
	return m.AccessMethod
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckAtomicReadFile factory function for _BACnetServiceAckAtomicReadFile
func NewBACnetServiceAckAtomicReadFile(endOfFile BACnetApplicationTagBoolean, accessMethod BACnetServiceAckAtomicReadFileStreamOrRecord, serviceAckLength uint16) *_BACnetServiceAckAtomicReadFile {
	_result := &_BACnetServiceAckAtomicReadFile{
		EndOfFile:         endOfFile,
		AccessMethod:      accessMethod,
		_BACnetServiceAck: NewBACnetServiceAck(serviceAckLength),
	}
	_result._BACnetServiceAck._BACnetServiceAckChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckAtomicReadFile(structType interface{}) BACnetServiceAckAtomicReadFile {
	if casted, ok := structType.(BACnetServiceAckAtomicReadFile); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckAtomicReadFile); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckAtomicReadFile) GetTypeName() string {
	return "BACnetServiceAckAtomicReadFile"
}

func (m *_BACnetServiceAckAtomicReadFile) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetServiceAckAtomicReadFile) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (endOfFile)
	lengthInBits += m.EndOfFile.GetLengthInBits()

	// Simple field (accessMethod)
	lengthInBits += m.AccessMethod.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetServiceAckAtomicReadFile) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckAtomicReadFileParse(readBuffer utils.ReadBuffer, serviceAckLength uint16) (BACnetServiceAckAtomicReadFile, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckAtomicReadFile"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckAtomicReadFile")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (endOfFile)
	if pullErr := readBuffer.PullContext("endOfFile"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for endOfFile")
	}
	_endOfFile, _endOfFileErr := BACnetApplicationTagParse(readBuffer)
	if _endOfFileErr != nil {
		return nil, errors.Wrap(_endOfFileErr, "Error parsing 'endOfFile' field of BACnetServiceAckAtomicReadFile")
	}
	endOfFile := _endOfFile.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("endOfFile"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for endOfFile")
	}

	// Simple Field (accessMethod)
	if pullErr := readBuffer.PullContext("accessMethod"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessMethod")
	}
	_accessMethod, _accessMethodErr := BACnetServiceAckAtomicReadFileStreamOrRecordParse(readBuffer)
	if _accessMethodErr != nil {
		return nil, errors.Wrap(_accessMethodErr, "Error parsing 'accessMethod' field of BACnetServiceAckAtomicReadFile")
	}
	accessMethod := _accessMethod.(BACnetServiceAckAtomicReadFileStreamOrRecord)
	if closeErr := readBuffer.CloseContext("accessMethod"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessMethod")
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckAtomicReadFile"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckAtomicReadFile")
	}

	// Create a partially initialized instance
	_child := &_BACnetServiceAckAtomicReadFile{
		_BACnetServiceAck: &_BACnetServiceAck{
			ServiceAckLength: serviceAckLength,
		},
		EndOfFile:    endOfFile,
		AccessMethod: accessMethod,
	}
	_child._BACnetServiceAck._BACnetServiceAckChildRequirements = _child
	return _child, nil
}

func (m *_BACnetServiceAckAtomicReadFile) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckAtomicReadFile"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckAtomicReadFile")
		}

		// Simple Field (endOfFile)
		if pushErr := writeBuffer.PushContext("endOfFile"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for endOfFile")
		}
		_endOfFileErr := writeBuffer.WriteSerializable(m.GetEndOfFile())
		if popErr := writeBuffer.PopContext("endOfFile"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for endOfFile")
		}
		if _endOfFileErr != nil {
			return errors.Wrap(_endOfFileErr, "Error serializing 'endOfFile' field")
		}

		// Simple Field (accessMethod)
		if pushErr := writeBuffer.PushContext("accessMethod"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accessMethod")
		}
		_accessMethodErr := writeBuffer.WriteSerializable(m.GetAccessMethod())
		if popErr := writeBuffer.PopContext("accessMethod"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accessMethod")
		}
		if _accessMethodErr != nil {
			return errors.Wrap(_accessMethodErr, "Error serializing 'accessMethod' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckAtomicReadFile"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckAtomicReadFile")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetServiceAckAtomicReadFile) isBACnetServiceAckAtomicReadFile() bool {
	return true
}

func (m *_BACnetServiceAckAtomicReadFile) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
