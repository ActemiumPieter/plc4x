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

// AdsAddDeviceNotificationResponse is the data-structure of this message
type AdsAddDeviceNotificationResponse struct {
	*AdsData
	Result             ReturnCode
	NotificationHandle uint32
}

// IAdsAddDeviceNotificationResponse is the corresponding interface of AdsAddDeviceNotificationResponse
type IAdsAddDeviceNotificationResponse interface {
	IAdsData
	// GetResult returns Result (property field)
	GetResult() ReturnCode
	// GetNotificationHandle returns NotificationHandle (property field)
	GetNotificationHandle() uint32
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

func (m *AdsAddDeviceNotificationResponse) GetCommandId() CommandId {
	return CommandId_ADS_ADD_DEVICE_NOTIFICATION
}

func (m *AdsAddDeviceNotificationResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *AdsAddDeviceNotificationResponse) InitializeParent(parent *AdsData) {}

func (m *AdsAddDeviceNotificationResponse) GetParent() *AdsData {
	return m.AdsData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *AdsAddDeviceNotificationResponse) GetResult() ReturnCode {
	return m.Result
}

func (m *AdsAddDeviceNotificationResponse) GetNotificationHandle() uint32 {
	return m.NotificationHandle
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsAddDeviceNotificationResponse factory function for AdsAddDeviceNotificationResponse
func NewAdsAddDeviceNotificationResponse(result ReturnCode, notificationHandle uint32) *AdsAddDeviceNotificationResponse {
	_result := &AdsAddDeviceNotificationResponse{
		Result:             result,
		NotificationHandle: notificationHandle,
		AdsData:            NewAdsData(),
	}
	_result.Child = _result
	return _result
}

func CastAdsAddDeviceNotificationResponse(structType interface{}) *AdsAddDeviceNotificationResponse {
	if casted, ok := structType.(AdsAddDeviceNotificationResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*AdsAddDeviceNotificationResponse); ok {
		return casted
	}
	if casted, ok := structType.(AdsData); ok {
		return CastAdsAddDeviceNotificationResponse(casted.Child)
	}
	if casted, ok := structType.(*AdsData); ok {
		return CastAdsAddDeviceNotificationResponse(casted.Child)
	}
	return nil
}

func (m *AdsAddDeviceNotificationResponse) GetTypeName() string {
	return "AdsAddDeviceNotificationResponse"
}

func (m *AdsAddDeviceNotificationResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AdsAddDeviceNotificationResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (result)
	lengthInBits += 32

	// Simple field (notificationHandle)
	lengthInBits += 32

	return lengthInBits
}

func (m *AdsAddDeviceNotificationResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsAddDeviceNotificationResponseParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (*AdsAddDeviceNotificationResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsAddDeviceNotificationResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsAddDeviceNotificationResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (result)
	if pullErr := readBuffer.PullContext("result"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for result")
	}
	_result, _resultErr := ReturnCodeParse(readBuffer)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field")
	}
	result := _result
	if closeErr := readBuffer.CloseContext("result"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for result")
	}

	// Simple Field (notificationHandle)
	_notificationHandle, _notificationHandleErr := readBuffer.ReadUint32("notificationHandle", 32)
	if _notificationHandleErr != nil {
		return nil, errors.Wrap(_notificationHandleErr, "Error parsing 'notificationHandle' field")
	}
	notificationHandle := _notificationHandle

	if closeErr := readBuffer.CloseContext("AdsAddDeviceNotificationResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsAddDeviceNotificationResponse")
	}

	// Create a partially initialized instance
	_child := &AdsAddDeviceNotificationResponse{
		Result:             result,
		NotificationHandle: notificationHandle,
		AdsData:            &AdsData{},
	}
	_child.AdsData.Child = _child
	return _child, nil
}

func (m *AdsAddDeviceNotificationResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsAddDeviceNotificationResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsAddDeviceNotificationResponse")
		}

		// Simple Field (result)
		if pushErr := writeBuffer.PushContext("result"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for result")
		}
		_resultErr := writeBuffer.WriteSerializable(m.Result)
		if popErr := writeBuffer.PopContext("result"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for result")
		}
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		// Simple Field (notificationHandle)
		notificationHandle := uint32(m.NotificationHandle)
		_notificationHandleErr := writeBuffer.WriteUint32("notificationHandle", 32, (notificationHandle))
		if _notificationHandleErr != nil {
			return errors.Wrap(_notificationHandleErr, "Error serializing 'notificationHandle' field")
		}

		if popErr := writeBuffer.PopContext("AdsAddDeviceNotificationResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsAddDeviceNotificationResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *AdsAddDeviceNotificationResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
