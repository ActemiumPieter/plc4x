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

// BACnetConstructedDataFileAccessMethod is the data-structure of this message
type BACnetConstructedDataFileAccessMethod struct {
	*BACnetConstructedData
	FileAccessMethod *BACnetFileAccessMethodTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataFileAccessMethod is the corresponding interface of BACnetConstructedDataFileAccessMethod
type IBACnetConstructedDataFileAccessMethod interface {
	IBACnetConstructedData
	// GetFileAccessMethod returns FileAccessMethod (property field)
	GetFileAccessMethod() *BACnetFileAccessMethodTagged
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

func (m *BACnetConstructedDataFileAccessMethod) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataFileAccessMethod) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FILE_ACCESS_METHOD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataFileAccessMethod) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataFileAccessMethod) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataFileAccessMethod) GetFileAccessMethod() *BACnetFileAccessMethodTagged {
	return m.FileAccessMethod
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataFileAccessMethod factory function for BACnetConstructedDataFileAccessMethod
func NewBACnetConstructedDataFileAccessMethod(fileAccessMethod *BACnetFileAccessMethodTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataFileAccessMethod {
	_result := &BACnetConstructedDataFileAccessMethod{
		FileAccessMethod:      fileAccessMethod,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataFileAccessMethod(structType interface{}) *BACnetConstructedDataFileAccessMethod {
	if casted, ok := structType.(BACnetConstructedDataFileAccessMethod); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFileAccessMethod); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataFileAccessMethod(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataFileAccessMethod(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataFileAccessMethod) GetTypeName() string {
	return "BACnetConstructedDataFileAccessMethod"
}

func (m *BACnetConstructedDataFileAccessMethod) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataFileAccessMethod) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (fileAccessMethod)
	lengthInBits += m.FileAccessMethod.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataFileAccessMethod) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataFileAccessMethodParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataFileAccessMethod, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFileAccessMethod"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFileAccessMethod")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (fileAccessMethod)
	if pullErr := readBuffer.PullContext("fileAccessMethod"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for fileAccessMethod")
	}
	_fileAccessMethod, _fileAccessMethodErr := BACnetFileAccessMethodTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _fileAccessMethodErr != nil {
		return nil, errors.Wrap(_fileAccessMethodErr, "Error parsing 'fileAccessMethod' field")
	}
	fileAccessMethod := CastBACnetFileAccessMethodTagged(_fileAccessMethod)
	if closeErr := readBuffer.CloseContext("fileAccessMethod"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for fileAccessMethod")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFileAccessMethod"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFileAccessMethod")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataFileAccessMethod{
		FileAccessMethod:      CastBACnetFileAccessMethodTagged(fileAccessMethod),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataFileAccessMethod) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFileAccessMethod"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFileAccessMethod")
		}

		// Simple Field (fileAccessMethod)
		if pushErr := writeBuffer.PushContext("fileAccessMethod"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for fileAccessMethod")
		}
		_fileAccessMethodErr := writeBuffer.WriteSerializable(m.FileAccessMethod)
		if popErr := writeBuffer.PopContext("fileAccessMethod"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for fileAccessMethod")
		}
		if _fileAccessMethodErr != nil {
			return errors.Wrap(_fileAccessMethodErr, "Error serializing 'fileAccessMethod' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFileAccessMethod"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFileAccessMethod")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataFileAccessMethod) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
