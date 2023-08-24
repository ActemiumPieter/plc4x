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
package org.apache.plc4x.java.iec608705104.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class ASDU implements Message {

  // Properties.
  protected final TypeIdentification typeIdentification;
  protected final boolean structureQualifier;
  protected final boolean test;
  protected final boolean negative;
  protected final CauseOfTransmission causeOfTransmission;
  protected final short originatorAddress;
  protected final int asduAddressField;
  protected final List<InformationObject> informationObjects;

  public ASDU(
      TypeIdentification typeIdentification,
      boolean structureQualifier,
      boolean test,
      boolean negative,
      CauseOfTransmission causeOfTransmission,
      short originatorAddress,
      int asduAddressField,
      List<InformationObject> informationObjects) {
    super();
    this.typeIdentification = typeIdentification;
    this.structureQualifier = structureQualifier;
    this.test = test;
    this.negative = negative;
    this.causeOfTransmission = causeOfTransmission;
    this.originatorAddress = originatorAddress;
    this.asduAddressField = asduAddressField;
    this.informationObjects = informationObjects;
  }

  public TypeIdentification getTypeIdentification() {
    return typeIdentification;
  }

  public boolean getStructureQualifier() {
    return structureQualifier;
  }

  public boolean getTest() {
    return test;
  }

  public boolean getNegative() {
    return negative;
  }

  public CauseOfTransmission getCauseOfTransmission() {
    return causeOfTransmission;
  }

  public short getOriginatorAddress() {
    return originatorAddress;
  }

  public int getAsduAddressField() {
    return asduAddressField;
  }

  public List<InformationObject> getInformationObjects() {
    return informationObjects;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ASDU");

    // Simple Field (typeIdentification)
    writeSimpleEnumField(
        "typeIdentification",
        "TypeIdentification",
        typeIdentification,
        new DataWriterEnumDefault<>(
            TypeIdentification::getValue,
            TypeIdentification::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (structureQualifier)
    writeSimpleField("structureQualifier", structureQualifier, writeBoolean(writeBuffer));

    // Implicit Field (numberOfObjects) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    byte numberOfObjects = (byte) (COUNT(getInformationObjects()));
    writeImplicitField("numberOfObjects", numberOfObjects, writeUnsignedByte(writeBuffer, 7));

    // Simple Field (test)
    writeSimpleField("test", test, writeBoolean(writeBuffer));

    // Simple Field (negative)
    writeSimpleField("negative", negative, writeBoolean(writeBuffer));

    // Simple Field (causeOfTransmission)
    writeSimpleEnumField(
        "causeOfTransmission",
        "CauseOfTransmission",
        causeOfTransmission,
        new DataWriterEnumDefault<>(
            CauseOfTransmission::getValue,
            CauseOfTransmission::name,
            writeUnsignedByte(writeBuffer, 6)));

    // Simple Field (originatorAddress)
    writeSimpleField("originatorAddress", originatorAddress, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (asduAddressField)
    writeSimpleField("asduAddressField", asduAddressField, writeUnsignedInt(writeBuffer, 16));

    // Array Field (informationObjects)
    writeComplexTypeArrayField("informationObjects", informationObjects, writeBuffer);

    writeBuffer.popContext("ASDU");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    ASDU _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (typeIdentification)
    lengthInBits += 8;

    // Simple field (structureQualifier)
    lengthInBits += 1;

    // Implicit Field (numberOfObjects)
    lengthInBits += 7;

    // Simple field (test)
    lengthInBits += 1;

    // Simple field (negative)
    lengthInBits += 1;

    // Simple field (causeOfTransmission)
    lengthInBits += 6;

    // Simple field (originatorAddress)
    lengthInBits += 8;

    // Simple field (asduAddressField)
    lengthInBits += 16;

    // Array field
    if (informationObjects != null) {
      int i = 0;
      for (InformationObject element : informationObjects) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= informationObjects.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ASDU staticParse(ReadBuffer readBuffer, Object... args) throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static ASDU staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("ASDU");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    TypeIdentification typeIdentification =
        readEnumField(
            "typeIdentification",
            "TypeIdentification",
            new DataReaderEnumDefault<>(
                TypeIdentification::enumForValue, readUnsignedShort(readBuffer, 8)));

    boolean structureQualifier = readSimpleField("structureQualifier", readBoolean(readBuffer));

    byte numberOfObjects = readImplicitField("numberOfObjects", readUnsignedByte(readBuffer, 7));

    boolean test = readSimpleField("test", readBoolean(readBuffer));

    boolean negative = readSimpleField("negative", readBoolean(readBuffer));

    CauseOfTransmission causeOfTransmission =
        readEnumField(
            "causeOfTransmission",
            "CauseOfTransmission",
            new DataReaderEnumDefault<>(
                CauseOfTransmission::enumForValue, readUnsignedByte(readBuffer, 6)));

    short originatorAddress =
        readSimpleField("originatorAddress", readUnsignedShort(readBuffer, 8));

    int asduAddressField = readSimpleField("asduAddressField", readUnsignedInt(readBuffer, 16));

    List<InformationObject> informationObjects =
        readCountArrayField(
            "informationObjects",
            new DataReaderComplexDefault<>(
                () -> InformationObject.staticParse(readBuffer), readBuffer),
            numberOfObjects);

    readBuffer.closeContext("ASDU");
    // Create the instance
    ASDU _aSDU;
    _aSDU =
        new ASDU(
            typeIdentification,
            structureQualifier,
            test,
            negative,
            causeOfTransmission,
            originatorAddress,
            asduAddressField,
            informationObjects);
    return _aSDU;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ASDU)) {
      return false;
    }
    ASDU that = (ASDU) o;
    return (getTypeIdentification() == that.getTypeIdentification())
        && (getStructureQualifier() == that.getStructureQualifier())
        && (getTest() == that.getTest())
        && (getNegative() == that.getNegative())
        && (getCauseOfTransmission() == that.getCauseOfTransmission())
        && (getOriginatorAddress() == that.getOriginatorAddress())
        && (getAsduAddressField() == that.getAsduAddressField())
        && (getInformationObjects() == that.getInformationObjects())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getTypeIdentification(),
        getStructureQualifier(),
        getTest(),
        getNegative(),
        getCauseOfTransmission(),
        getOriginatorAddress(),
        getAsduAddressField(),
        getInformationObjects());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
