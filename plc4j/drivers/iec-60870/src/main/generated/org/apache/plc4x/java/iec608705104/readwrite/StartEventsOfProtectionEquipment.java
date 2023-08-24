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

public class StartEventsOfProtectionEquipment implements Message {

  // Properties.
  protected final boolean startOfOperationInReverseDirection;
  protected final boolean startOfOperationIE;
  protected final boolean stateOfOperationPhaseL3;
  protected final boolean stateOfOperationPhaseL2;
  protected final boolean stateOfOperationPhaseL1;
  protected final boolean generalStartOfOperation;

  public StartEventsOfProtectionEquipment(
      boolean startOfOperationInReverseDirection,
      boolean startOfOperationIE,
      boolean stateOfOperationPhaseL3,
      boolean stateOfOperationPhaseL2,
      boolean stateOfOperationPhaseL1,
      boolean generalStartOfOperation) {
    super();
    this.startOfOperationInReverseDirection = startOfOperationInReverseDirection;
    this.startOfOperationIE = startOfOperationIE;
    this.stateOfOperationPhaseL3 = stateOfOperationPhaseL3;
    this.stateOfOperationPhaseL2 = stateOfOperationPhaseL2;
    this.stateOfOperationPhaseL1 = stateOfOperationPhaseL1;
    this.generalStartOfOperation = generalStartOfOperation;
  }

  public boolean getStartOfOperationInReverseDirection() {
    return startOfOperationInReverseDirection;
  }

  public boolean getStartOfOperationIE() {
    return startOfOperationIE;
  }

  public boolean getStateOfOperationPhaseL3() {
    return stateOfOperationPhaseL3;
  }

  public boolean getStateOfOperationPhaseL2() {
    return stateOfOperationPhaseL2;
  }

  public boolean getStateOfOperationPhaseL1() {
    return stateOfOperationPhaseL1;
  }

  public boolean getGeneralStartOfOperation() {
    return generalStartOfOperation;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("StartEventsOfProtectionEquipment");

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0, writeUnsignedByte(writeBuffer, 2));

    // Simple Field (startOfOperationInReverseDirection)
    writeSimpleField(
        "startOfOperationInReverseDirection",
        startOfOperationInReverseDirection,
        writeBoolean(writeBuffer));

    // Simple Field (startOfOperationIE)
    writeSimpleField("startOfOperationIE", startOfOperationIE, writeBoolean(writeBuffer));

    // Simple Field (stateOfOperationPhaseL3)
    writeSimpleField("stateOfOperationPhaseL3", stateOfOperationPhaseL3, writeBoolean(writeBuffer));

    // Simple Field (stateOfOperationPhaseL2)
    writeSimpleField("stateOfOperationPhaseL2", stateOfOperationPhaseL2, writeBoolean(writeBuffer));

    // Simple Field (stateOfOperationPhaseL1)
    writeSimpleField("stateOfOperationPhaseL1", stateOfOperationPhaseL1, writeBoolean(writeBuffer));

    // Simple Field (generalStartOfOperation)
    writeSimpleField("generalStartOfOperation", generalStartOfOperation, writeBoolean(writeBuffer));

    writeBuffer.popContext("StartEventsOfProtectionEquipment");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    StartEventsOfProtectionEquipment _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 2;

    // Simple field (startOfOperationInReverseDirection)
    lengthInBits += 1;

    // Simple field (startOfOperationIE)
    lengthInBits += 1;

    // Simple field (stateOfOperationPhaseL3)
    lengthInBits += 1;

    // Simple field (stateOfOperationPhaseL2)
    lengthInBits += 1;

    // Simple field (stateOfOperationPhaseL1)
    lengthInBits += 1;

    // Simple field (generalStartOfOperation)
    lengthInBits += 1;

    return lengthInBits;
  }

  public static StartEventsOfProtectionEquipment staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static StartEventsOfProtectionEquipment staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("StartEventsOfProtectionEquipment");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Byte reservedField0 = readReservedField("reserved", readUnsignedByte(readBuffer, 2), (byte) 0);

    boolean startOfOperationInReverseDirection =
        readSimpleField("startOfOperationInReverseDirection", readBoolean(readBuffer));

    boolean startOfOperationIE = readSimpleField("startOfOperationIE", readBoolean(readBuffer));

    boolean stateOfOperationPhaseL3 =
        readSimpleField("stateOfOperationPhaseL3", readBoolean(readBuffer));

    boolean stateOfOperationPhaseL2 =
        readSimpleField("stateOfOperationPhaseL2", readBoolean(readBuffer));

    boolean stateOfOperationPhaseL1 =
        readSimpleField("stateOfOperationPhaseL1", readBoolean(readBuffer));

    boolean generalStartOfOperation =
        readSimpleField("generalStartOfOperation", readBoolean(readBuffer));

    readBuffer.closeContext("StartEventsOfProtectionEquipment");
    // Create the instance
    StartEventsOfProtectionEquipment _startEventsOfProtectionEquipment;
    _startEventsOfProtectionEquipment =
        new StartEventsOfProtectionEquipment(
            startOfOperationInReverseDirection,
            startOfOperationIE,
            stateOfOperationPhaseL3,
            stateOfOperationPhaseL2,
            stateOfOperationPhaseL1,
            generalStartOfOperation);
    return _startEventsOfProtectionEquipment;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof StartEventsOfProtectionEquipment)) {
      return false;
    }
    StartEventsOfProtectionEquipment that = (StartEventsOfProtectionEquipment) o;
    return (getStartOfOperationInReverseDirection() == that.getStartOfOperationInReverseDirection())
        && (getStartOfOperationIE() == that.getStartOfOperationIE())
        && (getStateOfOperationPhaseL3() == that.getStateOfOperationPhaseL3())
        && (getStateOfOperationPhaseL2() == that.getStateOfOperationPhaseL2())
        && (getStateOfOperationPhaseL1() == that.getStateOfOperationPhaseL1())
        && (getGeneralStartOfOperation() == that.getGeneralStartOfOperation())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getStartOfOperationInReverseDirection(),
        getStartOfOperationIE(),
        getStateOfOperationPhaseL3(),
        getStateOfOperationPhaseL2(),
        getStateOfOperationPhaseL1(),
        getGeneralStartOfOperation());
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
