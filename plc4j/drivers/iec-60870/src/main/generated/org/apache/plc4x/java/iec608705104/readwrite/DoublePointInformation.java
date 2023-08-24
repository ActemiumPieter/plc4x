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

public class DoublePointInformation implements Message {

  // Properties.
  protected final boolean invalid;
  protected final boolean notTopical;
  protected final boolean substituted;
  protected final boolean blocked;
  protected final byte dpiCode;

  public DoublePointInformation(
      boolean invalid, boolean notTopical, boolean substituted, boolean blocked, byte dpiCode) {
    super();
    this.invalid = invalid;
    this.notTopical = notTopical;
    this.substituted = substituted;
    this.blocked = blocked;
    this.dpiCode = dpiCode;
  }

  public boolean getInvalid() {
    return invalid;
  }

  public boolean getNotTopical() {
    return notTopical;
  }

  public boolean getSubstituted() {
    return substituted;
  }

  public boolean getBlocked() {
    return blocked;
  }

  public byte getDpiCode() {
    return dpiCode;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("DoublePointInformation");

    // Simple Field (invalid)
    writeSimpleField("invalid", invalid, writeBoolean(writeBuffer));

    // Simple Field (notTopical)
    writeSimpleField("notTopical", notTopical, writeBoolean(writeBuffer));

    // Simple Field (substituted)
    writeSimpleField("substituted", substituted, writeBoolean(writeBuffer));

    // Simple Field (blocked)
    writeSimpleField("blocked", blocked, writeBoolean(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0, writeUnsignedByte(writeBuffer, 2));

    // Simple Field (dpiCode)
    writeSimpleField("dpiCode", dpiCode, writeUnsignedByte(writeBuffer, 2));

    writeBuffer.popContext("DoublePointInformation");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    DoublePointInformation _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (invalid)
    lengthInBits += 1;

    // Simple field (notTopical)
    lengthInBits += 1;

    // Simple field (substituted)
    lengthInBits += 1;

    // Simple field (blocked)
    lengthInBits += 1;

    // Reserved Field (reserved)
    lengthInBits += 2;

    // Simple field (dpiCode)
    lengthInBits += 2;

    return lengthInBits;
  }

  public static DoublePointInformation staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static DoublePointInformation staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("DoublePointInformation");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    boolean invalid = readSimpleField("invalid", readBoolean(readBuffer));

    boolean notTopical = readSimpleField("notTopical", readBoolean(readBuffer));

    boolean substituted = readSimpleField("substituted", readBoolean(readBuffer));

    boolean blocked = readSimpleField("blocked", readBoolean(readBuffer));

    Byte reservedField0 = readReservedField("reserved", readUnsignedByte(readBuffer, 2), (byte) 0);

    byte dpiCode = readSimpleField("dpiCode", readUnsignedByte(readBuffer, 2));

    readBuffer.closeContext("DoublePointInformation");
    // Create the instance
    DoublePointInformation _doublePointInformation;
    _doublePointInformation =
        new DoublePointInformation(invalid, notTopical, substituted, blocked, dpiCode);
    return _doublePointInformation;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DoublePointInformation)) {
      return false;
    }
    DoublePointInformation that = (DoublePointInformation) o;
    return (getInvalid() == that.getInvalid())
        && (getNotTopical() == that.getNotTopical())
        && (getSubstituted() == that.getSubstituted())
        && (getBlocked() == that.getBlocked())
        && (getDpiCode() == that.getDpiCode())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getInvalid(), getNotTopical(), getSubstituted(), getBlocked(), getDpiCode());
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
