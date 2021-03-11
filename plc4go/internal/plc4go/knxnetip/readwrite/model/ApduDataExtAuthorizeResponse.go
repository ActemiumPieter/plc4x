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
type ApduDataExtAuthorizeResponse struct {
    Level uint8
    Parent *ApduDataExt
    IApduDataExtAuthorizeResponse
}

// The corresponding interface
type IApduDataExtAuthorizeResponse interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtAuthorizeResponse) ExtApciType() uint8 {
    return 0x12
}


func (m *ApduDataExtAuthorizeResponse) InitializeParent(parent *ApduDataExt) {
}

func NewApduDataExtAuthorizeResponse(level uint8) *ApduDataExt {
    child := &ApduDataExtAuthorizeResponse{
        Level: level,
        Parent: NewApduDataExt(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastApduDataExtAuthorizeResponse(structType interface{}) *ApduDataExtAuthorizeResponse {
    castFunc := func(typ interface{}) *ApduDataExtAuthorizeResponse {
        if casted, ok := typ.(ApduDataExtAuthorizeResponse); ok {
            return &casted
        }
        if casted, ok := typ.(*ApduDataExtAuthorizeResponse); ok {
            return casted
        }
        if casted, ok := typ.(ApduDataExt); ok {
            return CastApduDataExtAuthorizeResponse(casted.Child)
        }
        if casted, ok := typ.(*ApduDataExt); ok {
            return CastApduDataExtAuthorizeResponse(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *ApduDataExtAuthorizeResponse) GetTypeName() string {
    return "ApduDataExtAuthorizeResponse"
}

func (m *ApduDataExtAuthorizeResponse) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (level)
    lengthInBits += 8

    return lengthInBits
}

func (m *ApduDataExtAuthorizeResponse) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ApduDataExtAuthorizeResponseParse(io *utils.ReadBuffer) (*ApduDataExt, error) {

    // Simple Field (level)
    level, _levelErr := io.ReadUint8(8)
    if _levelErr != nil {
        return nil, errors.New("Error parsing 'level' field " + _levelErr.Error())
    }

    // Create a partially initialized instance
    _child := &ApduDataExtAuthorizeResponse{
        Level: level,
        Parent: &ApduDataExt{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ApduDataExtAuthorizeResponse) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (level)
    level := uint8(m.Level)
    _levelErr := io.WriteUint8(8, (level))
    if _levelErr != nil {
        return errors.New("Error serializing 'level' field " + _levelErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataExtAuthorizeResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "level":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Level = data
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

func (m *ApduDataExtAuthorizeResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.Level, xml.StartElement{Name: xml.Name{Local: "level"}}); err != nil {
        return err
    }
    return nil
}

