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
type ComObjectTableRealisationType2 struct {
	NumEntries           uint8
	RamFlagsTablePointer uint8
	ComObjectDescriptors []*GroupObjectDescriptorRealisationType2
	Parent               *ComObjectTable
}

// The corresponding interface
type IComObjectTableRealisationType2 interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ComObjectTableRealisationType2) FirmwareType() FirmwareType {
	return FirmwareType_SYSTEM_2
}

func (m *ComObjectTableRealisationType2) InitializeParent(parent *ComObjectTable) {
}

func NewComObjectTableRealisationType2(numEntries uint8, ramFlagsTablePointer uint8, comObjectDescriptors []*GroupObjectDescriptorRealisationType2) *ComObjectTable {
	child := &ComObjectTableRealisationType2{
		NumEntries:           numEntries,
		RamFlagsTablePointer: ramFlagsTablePointer,
		ComObjectDescriptors: comObjectDescriptors,
		Parent:               NewComObjectTable(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastComObjectTableRealisationType2(structType interface{}) *ComObjectTableRealisationType2 {
	castFunc := func(typ interface{}) *ComObjectTableRealisationType2 {
		if casted, ok := typ.(ComObjectTableRealisationType2); ok {
			return &casted
		}
		if casted, ok := typ.(*ComObjectTableRealisationType2); ok {
			return casted
		}
		if casted, ok := typ.(ComObjectTable); ok {
			return CastComObjectTableRealisationType2(casted.Child)
		}
		if casted, ok := typ.(*ComObjectTable); ok {
			return CastComObjectTableRealisationType2(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ComObjectTableRealisationType2) GetTypeName() string {
	return "ComObjectTableRealisationType2"
}

func (m *ComObjectTableRealisationType2) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ComObjectTableRealisationType2) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (numEntries)
	lengthInBits += 8

	// Simple field (ramFlagsTablePointer)
	lengthInBits += 8

	// Array field
	if len(m.ComObjectDescriptors) > 0 {
		for i, element := range m.ComObjectDescriptors {
			last := i == len(m.ComObjectDescriptors)-1
			lengthInBits += element.LengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *ComObjectTableRealisationType2) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ComObjectTableRealisationType2Parse(io utils.ReadBuffer) (*ComObjectTable, error) {
	io.PullContext("ComObjectTableRealisationType2")

	// Simple Field (numEntries)
	numEntries, _numEntriesErr := io.ReadUint8("numEntries", 8)
	if _numEntriesErr != nil {
		return nil, errors.Wrap(_numEntriesErr, "Error parsing 'numEntries' field")
	}

	// Simple Field (ramFlagsTablePointer)
	ramFlagsTablePointer, _ramFlagsTablePointerErr := io.ReadUint8("ramFlagsTablePointer", 8)
	if _ramFlagsTablePointerErr != nil {
		return nil, errors.Wrap(_ramFlagsTablePointerErr, "Error parsing 'ramFlagsTablePointer' field")
	}
	io.PullContext("comObjectDescriptors")

	// Array field (comObjectDescriptors)
	// Count array
	comObjectDescriptors := make([]*GroupObjectDescriptorRealisationType2, numEntries)
	for curItem := uint16(0); curItem < uint16(numEntries); curItem++ {
		_item, _err := GroupObjectDescriptorRealisationType2Parse(io)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'comObjectDescriptors' field")
		}
		comObjectDescriptors[curItem] = _item
	}

	io.CloseContext("ComObjectTableRealisationType2")

	// Create a partially initialized instance
	_child := &ComObjectTableRealisationType2{
		NumEntries:           numEntries,
		RamFlagsTablePointer: ramFlagsTablePointer,
		ComObjectDescriptors: comObjectDescriptors,
		Parent:               &ComObjectTable{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ComObjectTableRealisationType2) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("ComObjectTableRealisationType2")

		// Simple Field (numEntries)
		numEntries := uint8(m.NumEntries)
		_numEntriesErr := io.WriteUint8("numEntries", 8, (numEntries))
		if _numEntriesErr != nil {
			return errors.Wrap(_numEntriesErr, "Error serializing 'numEntries' field")
		}

		// Simple Field (ramFlagsTablePointer)
		ramFlagsTablePointer := uint8(m.RamFlagsTablePointer)
		_ramFlagsTablePointerErr := io.WriteUint8("ramFlagsTablePointer", 8, (ramFlagsTablePointer))
		if _ramFlagsTablePointerErr != nil {
			return errors.Wrap(_ramFlagsTablePointerErr, "Error serializing 'ramFlagsTablePointer' field")
		}

		// Array Field (comObjectDescriptors)
		if m.ComObjectDescriptors != nil {
			io.PushContext("comObjectDescriptors")
			for _, _element := range m.ComObjectDescriptors {
				_elementErr := _element.Serialize(io)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'comObjectDescriptors' field")
				}
			}
			io.PopContext("comObjectDescriptors")
		}

		io.PopContext("ComObjectTableRealisationType2")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ComObjectTableRealisationType2) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "numEntries":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.NumEntries = data
			case "ramFlagsTablePointer":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.RamFlagsTablePointer = data
			case "comObjectDescriptors":
				var data []*GroupObjectDescriptorRealisationType2
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ComObjectDescriptors = data
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

func (m *ComObjectTableRealisationType2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.NumEntries, xml.StartElement{Name: xml.Name{Local: "numEntries"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.RamFlagsTablePointer, xml.StartElement{Name: xml.Name{Local: "ramFlagsTablePointer"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "comObjectDescriptors"}}); err != nil {
		return err
	}
	for _, arrayElement := range m.ComObjectDescriptors {
		if err := e.EncodeElement(arrayElement, xml.StartElement{Name: xml.Name{Local: "comObjectDescriptors"}}); err != nil {
			return err
		}
	}
	if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "comObjectDescriptors"}}); err != nil {
		return err
	}
	return nil
}

func (m ComObjectTableRealisationType2) String() string {
	return string(m.Box("", 120))
}

func (m ComObjectTableRealisationType2) Box(name string, width int) utils.AsciiBox {
	boxName := "ComObjectTableRealisationType2"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("NumEntries", m.NumEntries, -1))
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("RamFlagsTablePointer", m.RamFlagsTablePointer, -1))
		// Array Field (comObjectDescriptors)
		if m.ComObjectDescriptors != nil {
			// Complex array base type
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.ComObjectDescriptors {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("ComObjectDescriptors", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
