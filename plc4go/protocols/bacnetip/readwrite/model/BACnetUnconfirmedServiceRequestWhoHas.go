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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetUnconfirmedServiceRequestWhoHas is the corresponding interface of BACnetUnconfirmedServiceRequestWhoHas
type BACnetUnconfirmedServiceRequestWhoHas interface {
	utils.LengthAware
	utils.Serializable
	BACnetUnconfirmedServiceRequest
	// GetDeviceInstanceRangeLowLimit returns DeviceInstanceRangeLowLimit (property field)
	GetDeviceInstanceRangeLowLimit() BACnetContextTagUnsignedInteger
	// GetDeviceInstanceRangeHighLimit returns DeviceInstanceRangeHighLimit (property field)
	GetDeviceInstanceRangeHighLimit() BACnetContextTagUnsignedInteger
	// GetObject returns Object (property field)
	GetObject() BACnetUnconfirmedServiceRequestWhoHasObject
}

// BACnetUnconfirmedServiceRequestWhoHasExactly can be used when we want exactly this type and not a type which fulfills BACnetUnconfirmedServiceRequestWhoHas.
// This is useful for switch cases.
type BACnetUnconfirmedServiceRequestWhoHasExactly interface {
	BACnetUnconfirmedServiceRequestWhoHas
	isBACnetUnconfirmedServiceRequestWhoHas() bool
}

// _BACnetUnconfirmedServiceRequestWhoHas is the data-structure of this message
type _BACnetUnconfirmedServiceRequestWhoHas struct {
	*_BACnetUnconfirmedServiceRequest
	DeviceInstanceRangeLowLimit  BACnetContextTagUnsignedInteger
	DeviceInstanceRangeHighLimit BACnetContextTagUnsignedInteger
	Object                       BACnetUnconfirmedServiceRequestWhoHasObject
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestWhoHas) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_WHO_HAS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestWhoHas) InitializeParent(parent BACnetUnconfirmedServiceRequest) {
}

func (m *_BACnetUnconfirmedServiceRequestWhoHas) GetParent() BACnetUnconfirmedServiceRequest {
	return m._BACnetUnconfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestWhoHas) GetDeviceInstanceRangeLowLimit() BACnetContextTagUnsignedInteger {
	return m.DeviceInstanceRangeLowLimit
}

func (m *_BACnetUnconfirmedServiceRequestWhoHas) GetDeviceInstanceRangeHighLimit() BACnetContextTagUnsignedInteger {
	return m.DeviceInstanceRangeHighLimit
}

func (m *_BACnetUnconfirmedServiceRequestWhoHas) GetObject() BACnetUnconfirmedServiceRequestWhoHasObject {
	return m.Object
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestWhoHas factory function for _BACnetUnconfirmedServiceRequestWhoHas
func NewBACnetUnconfirmedServiceRequestWhoHas(deviceInstanceRangeLowLimit BACnetContextTagUnsignedInteger, deviceInstanceRangeHighLimit BACnetContextTagUnsignedInteger, object BACnetUnconfirmedServiceRequestWhoHasObject, serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequestWhoHas {
	_result := &_BACnetUnconfirmedServiceRequestWhoHas{
		DeviceInstanceRangeLowLimit:      deviceInstanceRangeLowLimit,
		DeviceInstanceRangeHighLimit:     deviceInstanceRangeHighLimit,
		Object:                           object,
		_BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestWhoHas(structType interface{}) BACnetUnconfirmedServiceRequestWhoHas {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestWhoHas); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestWhoHas); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoHas) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestWhoHas"
}

func (m *_BACnetUnconfirmedServiceRequestWhoHas) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetUnconfirmedServiceRequestWhoHas) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Optional Field (deviceInstanceRangeLowLimit)
	if m.DeviceInstanceRangeLowLimit != nil {
		lengthInBits += m.DeviceInstanceRangeLowLimit.GetLengthInBits()
	}

	// Optional Field (deviceInstanceRangeHighLimit)
	if m.DeviceInstanceRangeHighLimit != nil {
		lengthInBits += m.DeviceInstanceRangeHighLimit.GetLengthInBits()
	}

	// Simple field (object)
	lengthInBits += m.Object.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestWhoHas) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestWhoHasParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequestWhoHas, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestWhoHas"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestWhoHas")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Optional Field (deviceInstanceRangeLowLimit) (Can be skipped, if a given expression evaluates to false)
	var deviceInstanceRangeLowLimit BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("deviceInstanceRangeLowLimit"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for deviceInstanceRangeLowLimit")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(0), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'deviceInstanceRangeLowLimit' field of BACnetUnconfirmedServiceRequestWhoHas")
		default:
			deviceInstanceRangeLowLimit = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("deviceInstanceRangeLowLimit"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for deviceInstanceRangeLowLimit")
			}
		}
	}

	// Optional Field (deviceInstanceRangeHighLimit) (Can be skipped, if a given expression evaluates to false)
	var deviceInstanceRangeHighLimit BACnetContextTagUnsignedInteger = nil
	if bool((deviceInstanceRangeLowLimit) != (nil)) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("deviceInstanceRangeHighLimit"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for deviceInstanceRangeHighLimit")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'deviceInstanceRangeHighLimit' field of BACnetUnconfirmedServiceRequestWhoHas")
		default:
			deviceInstanceRangeHighLimit = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("deviceInstanceRangeHighLimit"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for deviceInstanceRangeHighLimit")
			}
		}
	}

	// Simple Field (object)
	if pullErr := readBuffer.PullContext("object"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for object")
	}
	_object, _objectErr := BACnetUnconfirmedServiceRequestWhoHasObjectParse(readBuffer)
	if _objectErr != nil {
		return nil, errors.Wrap(_objectErr, "Error parsing 'object' field of BACnetUnconfirmedServiceRequestWhoHas")
	}
	object := _object.(BACnetUnconfirmedServiceRequestWhoHasObject)
	if closeErr := readBuffer.CloseContext("object"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for object")
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestWhoHas"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestWhoHas")
	}

	// Create a partially initialized instance
	_child := &_BACnetUnconfirmedServiceRequestWhoHas{
		_BACnetUnconfirmedServiceRequest: &_BACnetUnconfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		DeviceInstanceRangeLowLimit:  deviceInstanceRangeLowLimit,
		DeviceInstanceRangeHighLimit: deviceInstanceRangeHighLimit,
		Object:                       object,
	}
	_child._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoHas) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestWhoHas"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestWhoHas")
		}

		// Optional Field (deviceInstanceRangeLowLimit) (Can be skipped, if the value is null)
		var deviceInstanceRangeLowLimit BACnetContextTagUnsignedInteger = nil
		if m.GetDeviceInstanceRangeLowLimit() != nil {
			if pushErr := writeBuffer.PushContext("deviceInstanceRangeLowLimit"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for deviceInstanceRangeLowLimit")
			}
			deviceInstanceRangeLowLimit = m.GetDeviceInstanceRangeLowLimit()
			_deviceInstanceRangeLowLimitErr := writeBuffer.WriteSerializable(deviceInstanceRangeLowLimit)
			if popErr := writeBuffer.PopContext("deviceInstanceRangeLowLimit"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for deviceInstanceRangeLowLimit")
			}
			if _deviceInstanceRangeLowLimitErr != nil {
				return errors.Wrap(_deviceInstanceRangeLowLimitErr, "Error serializing 'deviceInstanceRangeLowLimit' field")
			}
		}

		// Optional Field (deviceInstanceRangeHighLimit) (Can be skipped, if the value is null)
		var deviceInstanceRangeHighLimit BACnetContextTagUnsignedInteger = nil
		if m.GetDeviceInstanceRangeHighLimit() != nil {
			if pushErr := writeBuffer.PushContext("deviceInstanceRangeHighLimit"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for deviceInstanceRangeHighLimit")
			}
			deviceInstanceRangeHighLimit = m.GetDeviceInstanceRangeHighLimit()
			_deviceInstanceRangeHighLimitErr := writeBuffer.WriteSerializable(deviceInstanceRangeHighLimit)
			if popErr := writeBuffer.PopContext("deviceInstanceRangeHighLimit"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for deviceInstanceRangeHighLimit")
			}
			if _deviceInstanceRangeHighLimitErr != nil {
				return errors.Wrap(_deviceInstanceRangeHighLimitErr, "Error serializing 'deviceInstanceRangeHighLimit' field")
			}
		}

		// Simple Field (object)
		if pushErr := writeBuffer.PushContext("object"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for object")
		}
		_objectErr := writeBuffer.WriteSerializable(m.GetObject())
		if popErr := writeBuffer.PopContext("object"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for object")
		}
		if _objectErr != nil {
			return errors.Wrap(_objectErr, "Error serializing 'object' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestWhoHas"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestWhoHas")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestWhoHas) isBACnetUnconfirmedServiceRequestWhoHas() bool {
	return true
}

func (m *_BACnetUnconfirmedServiceRequestWhoHas) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
