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
    "errors"
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
    "io"
)

// The data-structure of this message
type AdsWriteResponse struct {
    Result ReturnCode
    Parent *AdsData
    IAdsWriteResponse
}

// The corresponding interface
type IAdsWriteResponse interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsWriteResponse) CommandId() CommandId {
    return CommandId_ADS_WRITE
}

func (m *AdsWriteResponse) Response() bool {
    return true
}


func (m *AdsWriteResponse) InitializeParent(parent *AdsData) {
}

func NewAdsWriteResponse(result ReturnCode) *AdsData {
    child := &AdsWriteResponse{
        Result: result,
        Parent: NewAdsData(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastAdsWriteResponse(structType interface{}) *AdsWriteResponse {
    castFunc := func(typ interface{}) *AdsWriteResponse {
        if casted, ok := typ.(AdsWriteResponse); ok {
            return &casted
        }
        if casted, ok := typ.(*AdsWriteResponse); ok {
            return casted
        }
        if casted, ok := typ.(AdsData); ok {
            return CastAdsWriteResponse(casted.Child)
        }
        if casted, ok := typ.(*AdsData); ok {
            return CastAdsWriteResponse(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *AdsWriteResponse) GetTypeName() string {
    return "AdsWriteResponse"
}

func (m *AdsWriteResponse) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Enum Field (result)
    lengthInBits += 32

    return lengthInBits
}

func (m *AdsWriteResponse) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func AdsWriteResponseParse(io *utils.ReadBuffer) (*AdsData, error) {

    // Enum field (result)
    result, _resultErr := ReturnCodeParse(io)
    if _resultErr != nil {
        return nil, errors.New("Error parsing 'result' field " + _resultErr.Error())
    }

    // Create a partially initialized instance
    _child := &AdsWriteResponse{
        Result: result,
        Parent: &AdsData{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *AdsWriteResponse) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Enum field (result)
    result := CastReturnCode(m.Result)
    _resultErr := result.Serialize(io)
    if _resultErr != nil {
        return errors.New("Error serializing 'result' field " + _resultErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *AdsWriteResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "result":
                var data ReturnCode
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Result = data
            }
        }
        token, err = d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
    }
}

func (m *AdsWriteResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.Result, xml.StartElement{Name: xml.Name{Local: "result"}}); err != nil {
        return err
    }
    return nil
}

