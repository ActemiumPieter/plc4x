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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type GroupObjectDescriptorRealisationType1 struct {
	DataPointer           uint8
	TransmitEnable        bool
	SegmentSelectorEnable bool
	WriteEnable           bool
	ReadEnable            bool
	CommunicationEnable   bool
	Priority              CEMIPriority
	ValueType             ComObjectValueType
}

// The corresponding interface
type IGroupObjectDescriptorRealisationType1 interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

func NewGroupObjectDescriptorRealisationType1(dataPointer uint8, transmitEnable bool, segmentSelectorEnable bool, writeEnable bool, readEnable bool, communicationEnable bool, priority CEMIPriority, valueType ComObjectValueType) *GroupObjectDescriptorRealisationType1 {
	return &GroupObjectDescriptorRealisationType1{DataPointer: dataPointer, TransmitEnable: transmitEnable, SegmentSelectorEnable: segmentSelectorEnable, WriteEnable: writeEnable, ReadEnable: readEnable, CommunicationEnable: communicationEnable, Priority: priority, ValueType: valueType}
}

func CastGroupObjectDescriptorRealisationType1(structType interface{}) *GroupObjectDescriptorRealisationType1 {
	castFunc := func(typ interface{}) *GroupObjectDescriptorRealisationType1 {
		if casted, ok := typ.(GroupObjectDescriptorRealisationType1); ok {
			return &casted
		}
		if casted, ok := typ.(*GroupObjectDescriptorRealisationType1); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *GroupObjectDescriptorRealisationType1) GetTypeName() string {
	return "GroupObjectDescriptorRealisationType1"
}

func (m *GroupObjectDescriptorRealisationType1) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *GroupObjectDescriptorRealisationType1) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (dataPointer)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (transmitEnable)
	lengthInBits += 1

	// Simple field (segmentSelectorEnable)
	lengthInBits += 1

	// Simple field (writeEnable)
	lengthInBits += 1

	// Simple field (readEnable)
	lengthInBits += 1

	// Simple field (communicationEnable)
	lengthInBits += 1

	// Simple field (priority)
	lengthInBits += 2

	// Simple field (valueType)
	lengthInBits += 8

	return lengthInBits
}

func (m *GroupObjectDescriptorRealisationType1) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func GroupObjectDescriptorRealisationType1Parse(io utils.ReadBuffer) (*GroupObjectDescriptorRealisationType1, error) {
	io.PullContext("GroupObjectDescriptorRealisationType1")

	// Simple Field (dataPointer)
	dataPointer, _dataPointerErr := io.ReadUint8("dataPointer", 8)
	if _dataPointerErr != nil {
		return nil, errors.Wrap(_dataPointerErr, "Error parsing 'dataPointer' field")
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := io.ReadUint8("reserved", 1)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x1) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x1),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (transmitEnable)
	transmitEnable, _transmitEnableErr := io.ReadBit("transmitEnable")
	if _transmitEnableErr != nil {
		return nil, errors.Wrap(_transmitEnableErr, "Error parsing 'transmitEnable' field")
	}

	// Simple Field (segmentSelectorEnable)
	segmentSelectorEnable, _segmentSelectorEnableErr := io.ReadBit("segmentSelectorEnable")
	if _segmentSelectorEnableErr != nil {
		return nil, errors.Wrap(_segmentSelectorEnableErr, "Error parsing 'segmentSelectorEnable' field")
	}

	// Simple Field (writeEnable)
	writeEnable, _writeEnableErr := io.ReadBit("writeEnable")
	if _writeEnableErr != nil {
		return nil, errors.Wrap(_writeEnableErr, "Error parsing 'writeEnable' field")
	}

	// Simple Field (readEnable)
	readEnable, _readEnableErr := io.ReadBit("readEnable")
	if _readEnableErr != nil {
		return nil, errors.Wrap(_readEnableErr, "Error parsing 'readEnable' field")
	}

	// Simple Field (communicationEnable)
	communicationEnable, _communicationEnableErr := io.ReadBit("communicationEnable")
	if _communicationEnableErr != nil {
		return nil, errors.Wrap(_communicationEnableErr, "Error parsing 'communicationEnable' field")
	}

	// Simple Field (priority)
	priority, _priorityErr := CEMIPriorityParse(io)
	if _priorityErr != nil {
		return nil, errors.Wrap(_priorityErr, "Error parsing 'priority' field")
	}

	// Simple Field (valueType)
	valueType, _valueTypeErr := ComObjectValueTypeParse(io)
	if _valueTypeErr != nil {
		return nil, errors.Wrap(_valueTypeErr, "Error parsing 'valueType' field")
	}

	io.CloseContext("GroupObjectDescriptorRealisationType1")

	// Create the instance
	return NewGroupObjectDescriptorRealisationType1(dataPointer, transmitEnable, segmentSelectorEnable, writeEnable, readEnable, communicationEnable, priority, valueType), nil
}

func (m *GroupObjectDescriptorRealisationType1) Serialize(io utils.WriteBuffer) error {
	io.PushContext("GroupObjectDescriptorRealisationType1")

	// Simple Field (dataPointer)
	dataPointer := uint8(m.DataPointer)
	_dataPointerErr := io.WriteUint8("dataPointer", 8, (dataPointer))
	if _dataPointerErr != nil {
		return errors.Wrap(_dataPointerErr, "Error serializing 'dataPointer' field")
	}

	// Reserved Field (reserved)
	{
		_err := io.WriteUint8("reserved", 1, uint8(0x1))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (transmitEnable)
	transmitEnable := bool(m.TransmitEnable)
	_transmitEnableErr := io.WriteBit("transmitEnable", (transmitEnable))
	if _transmitEnableErr != nil {
		return errors.Wrap(_transmitEnableErr, "Error serializing 'transmitEnable' field")
	}

	// Simple Field (segmentSelectorEnable)
	segmentSelectorEnable := bool(m.SegmentSelectorEnable)
	_segmentSelectorEnableErr := io.WriteBit("segmentSelectorEnable", (segmentSelectorEnable))
	if _segmentSelectorEnableErr != nil {
		return errors.Wrap(_segmentSelectorEnableErr, "Error serializing 'segmentSelectorEnable' field")
	}

	// Simple Field (writeEnable)
	writeEnable := bool(m.WriteEnable)
	_writeEnableErr := io.WriteBit("writeEnable", (writeEnable))
	if _writeEnableErr != nil {
		return errors.Wrap(_writeEnableErr, "Error serializing 'writeEnable' field")
	}

	// Simple Field (readEnable)
	readEnable := bool(m.ReadEnable)
	_readEnableErr := io.WriteBit("readEnable", (readEnable))
	if _readEnableErr != nil {
		return errors.Wrap(_readEnableErr, "Error serializing 'readEnable' field")
	}

	// Simple Field (communicationEnable)
	communicationEnable := bool(m.CommunicationEnable)
	_communicationEnableErr := io.WriteBit("communicationEnable", (communicationEnable))
	if _communicationEnableErr != nil {
		return errors.Wrap(_communicationEnableErr, "Error serializing 'communicationEnable' field")
	}

	// Simple Field (priority)
	_priorityErr := m.Priority.Serialize(io)
	if _priorityErr != nil {
		return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
	}

	// Simple Field (valueType)
	_valueTypeErr := m.ValueType.Serialize(io)
	if _valueTypeErr != nil {
		return errors.Wrap(_valueTypeErr, "Error serializing 'valueType' field")
	}

	io.PopContext("GroupObjectDescriptorRealisationType1")
	return nil
}

func (m *GroupObjectDescriptorRealisationType1) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "dataPointer":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.DataPointer = data
			case "transmitEnable":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.TransmitEnable = data
			case "segmentSelectorEnable":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.SegmentSelectorEnable = data
			case "writeEnable":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.WriteEnable = data
			case "readEnable":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ReadEnable = data
			case "communicationEnable":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.CommunicationEnable = data
			case "priority":
				var data CEMIPriority
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Priority = data
			case "valueType":
				var data ComObjectValueType
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ValueType = data
			}
		}
	}
}

func (m *GroupObjectDescriptorRealisationType1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := "org.apache.plc4x.java.knxnetip.readwrite.GroupObjectDescriptorRealisationType1"
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.DataPointer, xml.StartElement{Name: xml.Name{Local: "dataPointer"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.TransmitEnable, xml.StartElement{Name: xml.Name{Local: "transmitEnable"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.SegmentSelectorEnable, xml.StartElement{Name: xml.Name{Local: "segmentSelectorEnable"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.WriteEnable, xml.StartElement{Name: xml.Name{Local: "writeEnable"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ReadEnable, xml.StartElement{Name: xml.Name{Local: "readEnable"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.CommunicationEnable, xml.StartElement{Name: xml.Name{Local: "communicationEnable"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Priority, xml.StartElement{Name: xml.Name{Local: "priority"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ValueType, xml.StartElement{Name: xml.Name{Local: "valueType"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m GroupObjectDescriptorRealisationType1) String() string {
	return string(m.Box("", 120))
}

func (m GroupObjectDescriptorRealisationType1) Box(name string, width int) utils.AsciiBox {
	boxName := "GroupObjectDescriptorRealisationType1"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Simple field (case simple)
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("DataPointer", m.DataPointer, -1))
	// Reserved Field (reserved)
	// reserved field can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("reserved", uint8(0x1), -1))
	// Simple field (case simple)
	// bool can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("TransmitEnable", m.TransmitEnable, -1))
	// Simple field (case simple)
	// bool can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("SegmentSelectorEnable", m.SegmentSelectorEnable, -1))
	// Simple field (case simple)
	// bool can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("WriteEnable", m.WriteEnable, -1))
	// Simple field (case simple)
	// bool can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("ReadEnable", m.ReadEnable, -1))
	// Simple field (case simple)
	// bool can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("CommunicationEnable", m.CommunicationEnable, -1))
	// Complex field (case complex)
	boxes = append(boxes, m.Priority.Box("priority", width-2))
	// Complex field (case complex)
	boxes = append(boxes, m.ValueType.Box("valueType", width-2))
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}
