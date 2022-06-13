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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// The data-structure of this message
type MQTT_Property_WILL_DELAY_INTERVAL struct {
	*MQTT_Property
	Value uint32
}

// The corresponding interface
type IMQTT_Property_WILL_DELAY_INTERVAL interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////

func (m *MQTT_Property_WILL_DELAY_INTERVAL) InitializeParent(parent *MQTT_Property, propertyType MQTT_PropertyType) {
	m.PropertyType = propertyType
}

func NewMQTT_Property_WILL_DELAY_INTERVAL(value uint32, propertyType MQTT_PropertyType) *MQTT_Property {
	child := &MQTT_Property_WILL_DELAY_INTERVAL{
		Value: value,
    	MQTT_Property: NewMQTT_Property(propertyType),
	}
	child.Child = child
	return child.MQTT_Property
}

func CastMQTT_Property_WILL_DELAY_INTERVAL(structType interface{}) *MQTT_Property_WILL_DELAY_INTERVAL {
	castFunc := func(typ interface{}) *MQTT_Property_WILL_DELAY_INTERVAL {
		if casted, ok := typ.(MQTT_Property_WILL_DELAY_INTERVAL); ok {
			return &casted
		}
		if casted, ok := typ.(*MQTT_Property_WILL_DELAY_INTERVAL); ok {
			return casted
		}
		if casted, ok := typ.(MQTT_Property); ok {
			return CastMQTT_Property_WILL_DELAY_INTERVAL(casted.Child)
		}
		if casted, ok := typ.(*MQTT_Property); ok {
			return CastMQTT_Property_WILL_DELAY_INTERVAL(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *MQTT_Property_WILL_DELAY_INTERVAL) GetTypeName() string {
	return "MQTT_Property_WILL_DELAY_INTERVAL"
}

func (m *MQTT_Property_WILL_DELAY_INTERVAL) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *MQTT_Property_WILL_DELAY_INTERVAL) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (value)
	lengthInBits += 32;

	return lengthInBits
}


func (m *MQTT_Property_WILL_DELAY_INTERVAL) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func MQTT_Property_WILL_DELAY_INTERVALParse(readBuffer utils.ReadBuffer) (*MQTT_Property, error) {
	if pullErr := readBuffer.PullContext("MQTT_Property_WILL_DELAY_INTERVAL"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (value)
_value, _valueErr := readBuffer.ReadUint32("value", 32)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("MQTT_Property_WILL_DELAY_INTERVAL"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &MQTT_Property_WILL_DELAY_INTERVAL{
		Value: value,
        MQTT_Property: &MQTT_Property{},
	}
	_child.MQTT_Property.Child = _child
	return _child.MQTT_Property, nil
}

func (m *MQTT_Property_WILL_DELAY_INTERVAL) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MQTT_Property_WILL_DELAY_INTERVAL"); pushErr != nil {
			return pushErr
		}

	// Simple Field (value)
	value := uint32(m.Value)
	_valueErr := writeBuffer.WriteUint32("value", 32, (value))
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

		if popErr := writeBuffer.PopContext("MQTT_Property_WILL_DELAY_INTERVAL"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *MQTT_Property_WILL_DELAY_INTERVAL) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}



