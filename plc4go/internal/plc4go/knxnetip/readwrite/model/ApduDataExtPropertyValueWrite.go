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
type ApduDataExtPropertyValueWrite struct {
	ObjectIndex uint8
	PropertyId  uint8
	Count       uint8
	Index       uint16
	Data        []uint8
	Parent      *ApduDataExt
}

// The corresponding interface
type IApduDataExtPropertyValueWrite interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtPropertyValueWrite) ExtApciType() uint8 {
	return 0x17
}

func (m *ApduDataExtPropertyValueWrite) InitializeParent(parent *ApduDataExt) {
}

func NewApduDataExtPropertyValueWrite(objectIndex uint8, propertyId uint8, count uint8, index uint16, data []uint8) *ApduDataExt {
	child := &ApduDataExtPropertyValueWrite{
		ObjectIndex: objectIndex,
		PropertyId:  propertyId,
		Count:       count,
		Index:       index,
		Data:        data,
		Parent:      NewApduDataExt(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduDataExtPropertyValueWrite(structType interface{}) *ApduDataExtPropertyValueWrite {
	castFunc := func(typ interface{}) *ApduDataExtPropertyValueWrite {
		if casted, ok := typ.(ApduDataExtPropertyValueWrite); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataExtPropertyValueWrite); ok {
			return casted
		}
		if casted, ok := typ.(ApduDataExt); ok {
			return CastApduDataExtPropertyValueWrite(casted.Child)
		}
		if casted, ok := typ.(*ApduDataExt); ok {
			return CastApduDataExtPropertyValueWrite(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataExtPropertyValueWrite) GetTypeName() string {
	return "ApduDataExtPropertyValueWrite"
}

func (m *ApduDataExtPropertyValueWrite) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataExtPropertyValueWrite) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (objectIndex)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (count)
	lengthInBits += 4

	// Simple field (index)
	lengthInBits += 12

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *ApduDataExtPropertyValueWrite) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataExtPropertyValueWriteParse(io utils.ReadBuffer, length uint8) (*ApduDataExt, error) {
	io.PullContext("ApduDataExtPropertyValueWrite")

	// Simple Field (objectIndex)
	objectIndex, _objectIndexErr := io.ReadUint8("objectIndex", 8)
	if _objectIndexErr != nil {
		return nil, errors.Wrap(_objectIndexErr, "Error parsing 'objectIndex' field")
	}

	// Simple Field (propertyId)
	propertyId, _propertyIdErr := io.ReadUint8("propertyId", 8)
	if _propertyIdErr != nil {
		return nil, errors.Wrap(_propertyIdErr, "Error parsing 'propertyId' field")
	}

	// Simple Field (count)
	count, _countErr := io.ReadUint8("count", 4)
	if _countErr != nil {
		return nil, errors.Wrap(_countErr, "Error parsing 'count' field")
	}

	// Simple Field (index)
	index, _indexErr := io.ReadUint16("index", 12)
	if _indexErr != nil {
		return nil, errors.Wrap(_indexErr, "Error parsing 'index' field")
	}
	io.PullContext("data")

	// Array field (data)
	// Count array
	data := make([]uint8, uint16(length)-uint16(uint16(5)))
	for curItem := uint16(0); curItem < uint16(uint16(length)-uint16(uint16(5))); curItem++ {
		_item, _err := io.ReadUint8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'data' field")
		}
		data[curItem] = _item
	}

	io.CloseContext("ApduDataExtPropertyValueWrite")

	// Create a partially initialized instance
	_child := &ApduDataExtPropertyValueWrite{
		ObjectIndex: objectIndex,
		PropertyId:  propertyId,
		Count:       count,
		Index:       index,
		Data:        data,
		Parent:      &ApduDataExt{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduDataExtPropertyValueWrite) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("ApduDataExtPropertyValueWrite")

		// Simple Field (objectIndex)
		objectIndex := uint8(m.ObjectIndex)
		_objectIndexErr := io.WriteUint8("objectIndex", 8, (objectIndex))
		if _objectIndexErr != nil {
			return errors.Wrap(_objectIndexErr, "Error serializing 'objectIndex' field")
		}

		// Simple Field (propertyId)
		propertyId := uint8(m.PropertyId)
		_propertyIdErr := io.WriteUint8("propertyId", 8, (propertyId))
		if _propertyIdErr != nil {
			return errors.Wrap(_propertyIdErr, "Error serializing 'propertyId' field")
		}

		// Simple Field (count)
		count := uint8(m.Count)
		_countErr := io.WriteUint8("count", 4, (count))
		if _countErr != nil {
			return errors.Wrap(_countErr, "Error serializing 'count' field")
		}

		// Simple Field (index)
		index := uint16(m.Index)
		_indexErr := io.WriteUint16("index", 12, (index))
		if _indexErr != nil {
			return errors.Wrap(_indexErr, "Error serializing 'index' field")
		}

		// Array Field (data)
		if m.Data != nil {
			io.PushContext("data")
			for _, _element := range m.Data {
				_elementErr := io.WriteUint8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
			io.PopContext("data")
		}

		io.PopContext("ApduDataExtPropertyValueWrite")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataExtPropertyValueWrite) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "objectIndex":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ObjectIndex = data
			case "propertyId":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.PropertyId = data
			case "count":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Count = data
			case "index":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Index = data
			case "data":
				var data []uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Data = data
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

func (m *ApduDataExtPropertyValueWrite) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.ObjectIndex, xml.StartElement{Name: xml.Name{Local: "objectIndex"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.PropertyId, xml.StartElement{Name: xml.Name{Local: "propertyId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Count, xml.StartElement{Name: xml.Name{Local: "count"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Index, xml.StartElement{Name: xml.Name{Local: "index"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Data, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
		return err
	}
	return nil
}

func (m ApduDataExtPropertyValueWrite) String() string {
	return string(m.Box("", 120))
}

func (m ApduDataExtPropertyValueWrite) Box(name string, width int) utils.AsciiBox {
	boxName := "ApduDataExtPropertyValueWrite"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("ObjectIndex", m.ObjectIndex, -1))
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("PropertyId", m.PropertyId, -1))
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("Count", m.Count, -1))
		// Simple field (case simple)
		// uint16 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("Index", m.Index, -1))
		// Array Field (data)
		if m.Data != nil {
			// Simple array base type uint8 will be hex dumped
			boxes = append(boxes, utils.BoxedDumpAnything("Data", m.Data))
			// Simple array base type uint8 will be rendered one by one
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.Data {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("Data", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
