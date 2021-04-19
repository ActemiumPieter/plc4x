//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type MPropReadReq struct {
	InterfaceObjectType uint16
	ObjectInstance      uint8
	PropertyId          uint8
	NumberOfElements    uint8
	StartIndex          uint16
	Parent              *CEMI
}

// The corresponding interface
type IMPropReadReq interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *MPropReadReq) MessageCode() uint8 {
	return 0xFC
}

func (m *MPropReadReq) InitializeParent(parent *CEMI) {
}

func NewMPropReadReq(interfaceObjectType uint16, objectInstance uint8, propertyId uint8, numberOfElements uint8, startIndex uint16) *CEMI {
	child := &MPropReadReq{
		InterfaceObjectType: interfaceObjectType,
		ObjectInstance:      objectInstance,
		PropertyId:          propertyId,
		NumberOfElements:    numberOfElements,
		StartIndex:          startIndex,
		Parent:              NewCEMI(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastMPropReadReq(structType interface{}) *MPropReadReq {
	castFunc := func(typ interface{}) *MPropReadReq {
		if casted, ok := typ.(MPropReadReq); ok {
			return &casted
		}
		if casted, ok := typ.(*MPropReadReq); ok {
			return casted
		}
		if casted, ok := typ.(CEMI); ok {
			return CastMPropReadReq(casted.Child)
		}
		if casted, ok := typ.(*CEMI); ok {
			return CastMPropReadReq(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *MPropReadReq) GetTypeName() string {
	return "MPropReadReq"
}

func (m *MPropReadReq) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *MPropReadReq) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (interfaceObjectType)
	lengthInBits += 16

	// Simple field (objectInstance)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (numberOfElements)
	lengthInBits += 4

	// Simple field (startIndex)
	lengthInBits += 12

	return lengthInBits
}

func (m *MPropReadReq) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func MPropReadReqParse(io utils.ReadBuffer) (*CEMI, error) {
	io.PullContext("MPropReadReq")

	// Simple Field (interfaceObjectType)
	interfaceObjectType, _interfaceObjectTypeErr := io.ReadUint16("interfaceObjectType", 16)
	if _interfaceObjectTypeErr != nil {
		return nil, errors.Wrap(_interfaceObjectTypeErr, "Error parsing 'interfaceObjectType' field")
	}

	// Simple Field (objectInstance)
	objectInstance, _objectInstanceErr := io.ReadUint8("objectInstance", 8)
	if _objectInstanceErr != nil {
		return nil, errors.Wrap(_objectInstanceErr, "Error parsing 'objectInstance' field")
	}

	// Simple Field (propertyId)
	propertyId, _propertyIdErr := io.ReadUint8("propertyId", 8)
	if _propertyIdErr != nil {
		return nil, errors.Wrap(_propertyIdErr, "Error parsing 'propertyId' field")
	}

	// Simple Field (numberOfElements)
	numberOfElements, _numberOfElementsErr := io.ReadUint8("numberOfElements", 4)
	if _numberOfElementsErr != nil {
		return nil, errors.Wrap(_numberOfElementsErr, "Error parsing 'numberOfElements' field")
	}

	// Simple Field (startIndex)
	startIndex, _startIndexErr := io.ReadUint16("startIndex", 12)
	if _startIndexErr != nil {
		return nil, errors.Wrap(_startIndexErr, "Error parsing 'startIndex' field")
	}

	io.CloseContext("MPropReadReq")

	// Create a partially initialized instance
	_child := &MPropReadReq{
		InterfaceObjectType: interfaceObjectType,
		ObjectInstance:      objectInstance,
		PropertyId:          propertyId,
		NumberOfElements:    numberOfElements,
		StartIndex:          startIndex,
		Parent:              &CEMI{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *MPropReadReq) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("MPropReadReq")

		// Simple Field (interfaceObjectType)
		interfaceObjectType := uint16(m.InterfaceObjectType)
		_interfaceObjectTypeErr := io.WriteUint16("interfaceObjectType", 16, (interfaceObjectType))
		if _interfaceObjectTypeErr != nil {
			return errors.Wrap(_interfaceObjectTypeErr, "Error serializing 'interfaceObjectType' field")
		}

		// Simple Field (objectInstance)
		objectInstance := uint8(m.ObjectInstance)
		_objectInstanceErr := io.WriteUint8("objectInstance", 8, (objectInstance))
		if _objectInstanceErr != nil {
			return errors.Wrap(_objectInstanceErr, "Error serializing 'objectInstance' field")
		}

		// Simple Field (propertyId)
		propertyId := uint8(m.PropertyId)
		_propertyIdErr := io.WriteUint8("propertyId", 8, (propertyId))
		if _propertyIdErr != nil {
			return errors.Wrap(_propertyIdErr, "Error serializing 'propertyId' field")
		}

		// Simple Field (numberOfElements)
		numberOfElements := uint8(m.NumberOfElements)
		_numberOfElementsErr := io.WriteUint8("numberOfElements", 4, (numberOfElements))
		if _numberOfElementsErr != nil {
			return errors.Wrap(_numberOfElementsErr, "Error serializing 'numberOfElements' field")
		}

		// Simple Field (startIndex)
		startIndex := uint16(m.StartIndex)
		_startIndexErr := io.WriteUint16("startIndex", 12, (startIndex))
		if _startIndexErr != nil {
			return errors.Wrap(_startIndexErr, "Error serializing 'startIndex' field")
		}

		io.PopContext("MPropReadReq")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *MPropReadReq) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "interfaceObjectType":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.InterfaceObjectType = data
			case "objectInstance":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ObjectInstance = data
			case "propertyId":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.PropertyId = data
			case "numberOfElements":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.NumberOfElements = data
			case "startIndex":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.StartIndex = data
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

func (m *MPropReadReq) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.InterfaceObjectType, xml.StartElement{Name: xml.Name{Local: "interfaceObjectType"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ObjectInstance, xml.StartElement{Name: xml.Name{Local: "objectInstance"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.PropertyId, xml.StartElement{Name: xml.Name{Local: "propertyId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.NumberOfElements, xml.StartElement{Name: xml.Name{Local: "numberOfElements"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.StartIndex, xml.StartElement{Name: xml.Name{Local: "startIndex"}}); err != nil {
		return err
	}
	return nil
}

func (m MPropReadReq) String() string {
	return string(m.Box("", 120))
}

func (m MPropReadReq) Box(name string, width int) utils.AsciiBox {
	boxName := "MPropReadReq"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint16 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("InterfaceObjectType", m.InterfaceObjectType, -1))
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("ObjectInstance", m.ObjectInstance, -1))
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("PropertyId", m.PropertyId, -1))
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("NumberOfElements", m.NumberOfElements, -1))
		// Simple field (case simple)
		// uint16 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("StartIndex", m.StartIndex, -1))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
