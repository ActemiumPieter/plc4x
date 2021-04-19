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
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// Constant values.
const BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_VENDORIDHEADER uint8 = 0x09
const BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_SERVICENUMBERHEADER uint8 = 0x1A
const BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_LISTOFVALUESOPENINGTAG uint8 = 0x2E
const BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_LISTOFVALUESCLOSINGTAG uint8 = 0x2F

// The data-structure of this message
type BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer struct {
	VendorId      uint8
	ServiceNumber uint16
	Values        []int8
	Parent        *BACnetUnconfirmedServiceRequest
}

// The corresponding interface
type IBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) ServiceChoice() uint8 {
	return 0x04
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func NewBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer(vendorId uint8, serviceNumber uint16, values []int8) *BACnetUnconfirmedServiceRequest {
	child := &BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer{
		VendorId:      vendorId,
		ServiceNumber: serviceNumber,
		Values:        values,
		Parent:        NewBACnetUnconfirmedServiceRequest(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer(structType interface{}) *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer {
	castFunc := func(typ interface{}) *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer {
		if casted, ok := typ.(BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer); ok {
			return casted
		}
		if casted, ok := typ.(BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer(casted.Child)
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Const Field (vendorIdHeader)
	lengthInBits += 8

	// Simple field (vendorId)
	lengthInBits += 8

	// Const Field (serviceNumberHeader)
	lengthInBits += 8

	// Simple field (serviceNumber)
	lengthInBits += 16

	// Const Field (listOfValuesOpeningTag)
	lengthInBits += 8

	// Array field
	if len(m.Values) > 0 {
		lengthInBits += 8 * uint16(len(m.Values))
	}

	// Const Field (listOfValuesClosingTag)
	lengthInBits += 8

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransferParse(io utils.ReadBuffer, len uint16) (*BACnetUnconfirmedServiceRequest, error) {
	io.PullContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer")

	// Const Field (vendorIdHeader)
	vendorIdHeader, _vendorIdHeaderErr := io.ReadUint8("vendorIdHeader", 8)
	if _vendorIdHeaderErr != nil {
		return nil, errors.Wrap(_vendorIdHeaderErr, "Error parsing 'vendorIdHeader' field")
	}
	if vendorIdHeader != BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_VENDORIDHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_VENDORIDHEADER) + " but got " + fmt.Sprintf("%d", vendorIdHeader))
	}

	// Simple Field (vendorId)
	vendorId, _vendorIdErr := io.ReadUint8("vendorId", 8)
	if _vendorIdErr != nil {
		return nil, errors.Wrap(_vendorIdErr, "Error parsing 'vendorId' field")
	}

	// Const Field (serviceNumberHeader)
	serviceNumberHeader, _serviceNumberHeaderErr := io.ReadUint8("serviceNumberHeader", 8)
	if _serviceNumberHeaderErr != nil {
		return nil, errors.Wrap(_serviceNumberHeaderErr, "Error parsing 'serviceNumberHeader' field")
	}
	if serviceNumberHeader != BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_SERVICENUMBERHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_SERVICENUMBERHEADER) + " but got " + fmt.Sprintf("%d", serviceNumberHeader))
	}

	// Simple Field (serviceNumber)
	serviceNumber, _serviceNumberErr := io.ReadUint16("serviceNumber", 16)
	if _serviceNumberErr != nil {
		return nil, errors.Wrap(_serviceNumberErr, "Error parsing 'serviceNumber' field")
	}

	// Const Field (listOfValuesOpeningTag)
	listOfValuesOpeningTag, _listOfValuesOpeningTagErr := io.ReadUint8("listOfValuesOpeningTag", 8)
	if _listOfValuesOpeningTagErr != nil {
		return nil, errors.Wrap(_listOfValuesOpeningTagErr, "Error parsing 'listOfValuesOpeningTag' field")
	}
	if listOfValuesOpeningTag != BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_LISTOFVALUESOPENINGTAG {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_LISTOFVALUESOPENINGTAG) + " but got " + fmt.Sprintf("%d", listOfValuesOpeningTag))
	}
	io.PullContext("values")

	// Array field (values)
	// Length array
	values := make([]int8, 0)
	_valuesLength := uint16(len) - uint16(uint16(8))
	_valuesEndPos := io.GetPos() + uint16(_valuesLength)
	for io.GetPos() < _valuesEndPos {
		_item, _err := io.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'values' field")
		}
		values = append(values, _item)
	}

	// Const Field (listOfValuesClosingTag)
	listOfValuesClosingTag, _listOfValuesClosingTagErr := io.ReadUint8("listOfValuesClosingTag", 8)
	if _listOfValuesClosingTagErr != nil {
		return nil, errors.Wrap(_listOfValuesClosingTagErr, "Error parsing 'listOfValuesClosingTag' field")
	}
	if listOfValuesClosingTag != BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_LISTOFVALUESCLOSINGTAG {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_LISTOFVALUESCLOSINGTAG) + " but got " + fmt.Sprintf("%d", listOfValuesClosingTag))
	}

	io.CloseContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer")

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer{
		VendorId:      vendorId,
		ServiceNumber: serviceNumber,
		Values:        values,
		Parent:        &BACnetUnconfirmedServiceRequest{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer")

		// Const Field (vendorIdHeader)
		_vendorIdHeaderErr := io.WriteUint8("vendorIdHeader", 8, 0x09)
		if _vendorIdHeaderErr != nil {
			return errors.Wrap(_vendorIdHeaderErr, "Error serializing 'vendorIdHeader' field")
		}

		// Simple Field (vendorId)
		vendorId := uint8(m.VendorId)
		_vendorIdErr := io.WriteUint8("vendorId", 8, (vendorId))
		if _vendorIdErr != nil {
			return errors.Wrap(_vendorIdErr, "Error serializing 'vendorId' field")
		}

		// Const Field (serviceNumberHeader)
		_serviceNumberHeaderErr := io.WriteUint8("serviceNumberHeader", 8, 0x1A)
		if _serviceNumberHeaderErr != nil {
			return errors.Wrap(_serviceNumberHeaderErr, "Error serializing 'serviceNumberHeader' field")
		}

		// Simple Field (serviceNumber)
		serviceNumber := uint16(m.ServiceNumber)
		_serviceNumberErr := io.WriteUint16("serviceNumber", 16, (serviceNumber))
		if _serviceNumberErr != nil {
			return errors.Wrap(_serviceNumberErr, "Error serializing 'serviceNumber' field")
		}

		// Const Field (listOfValuesOpeningTag)
		_listOfValuesOpeningTagErr := io.WriteUint8("listOfValuesOpeningTag", 8, 0x2E)
		if _listOfValuesOpeningTagErr != nil {
			return errors.Wrap(_listOfValuesOpeningTagErr, "Error serializing 'listOfValuesOpeningTag' field")
		}

		// Array Field (values)
		if m.Values != nil {
			io.PushContext("values")
			for _, _element := range m.Values {
				_elementErr := io.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'values' field")
				}
			}
			io.PopContext("values")
		}

		// Const Field (listOfValuesClosingTag)
		_listOfValuesClosingTagErr := io.WriteUint8("listOfValuesClosingTag", 8, 0x2F)
		if _listOfValuesClosingTagErr != nil {
			return errors.Wrap(_listOfValuesClosingTagErr, "Error serializing 'listOfValuesClosingTag' field")
		}

		io.PopContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "vendorId":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.VendorId = data
			case "serviceNumber":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ServiceNumber = data
			case "values":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.Values = utils.ByteArrayToInt8Array(_decoded[0:_len])
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

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.VendorId, xml.StartElement{Name: xml.Name{Local: "vendorId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ServiceNumber, xml.StartElement{Name: xml.Name{Local: "serviceNumber"}}); err != nil {
		return err
	}
	_encodedValues := hex.EncodeToString(utils.Int8ArrayToByteArray(m.Values))
	_encodedValues = strings.ToUpper(_encodedValues)
	if err := e.EncodeElement(_encodedValues, xml.StartElement{Name: xml.Name{Local: "values"}}); err != nil {
		return err
	}
	return nil
}

func (m BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) String() string {
	return string(m.Box("", 120))
}

func (m BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) Box(name string, width int) utils.AsciiBox {
	boxName := "BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Const Field (vendorIdHeader)
		boxes = append(boxes, utils.BoxAnything("VendorIdHeader", uint8(0x09), -1))
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("VendorId", m.VendorId, -1))
		// Const Field (serviceNumberHeader)
		boxes = append(boxes, utils.BoxAnything("ServiceNumberHeader", uint8(0x1A), -1))
		// Simple field (case simple)
		// uint16 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("ServiceNumber", m.ServiceNumber, -1))
		// Const Field (listOfValuesOpeningTag)
		boxes = append(boxes, utils.BoxAnything("ListOfValuesOpeningTag", uint8(0x2E), -1))
		// Array Field (values)
		if m.Values != nil {
			// Simple array base type int8 will be rendered one by one
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.Values {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("Values", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		// Const Field (listOfValuesClosingTag)
		boxes = append(boxes, utils.BoxAnything("ListOfValuesClosingTag", uint8(0x2F), -1))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
