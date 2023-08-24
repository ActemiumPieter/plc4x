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

public class APDUUFormatStopDataTransferActivation extends APDU implements Message {

  // Accessors for discriminator values.
  public Integer getCommand() {
    return (int) 0x13;
  }

  public APDUUFormatStopDataTransferActivation() {
    super();
  }

  @Override
  protected void serializeAPDUChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("APDUUFormatStopDataTransferActivation");

    // Padding Field (padding)
    writePaddingField(
        "padding",
        (int) (2),
        (short) 0x00,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("APDUUFormatStopDataTransferActivation");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    APDUUFormatStopDataTransferActivation _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Padding Field (padding)
    int _timesPadding = (int) (2);
    while (_timesPadding-- > 0) {
      lengthInBits += 8;
    }

    return lengthInBits;
  }

  public static APDUBuilder staticParseAPDUBuilder(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("APDUUFormatStopDataTransferActivation");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    readPaddingField(
        readUnsignedShort(readBuffer, 8),
        (int) (2),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("APDUUFormatStopDataTransferActivation");
    // Create the instance
    return new APDUUFormatStopDataTransferActivationBuilderImpl();
  }

  public static class APDUUFormatStopDataTransferActivationBuilderImpl implements APDU.APDUBuilder {

    public APDUUFormatStopDataTransferActivationBuilderImpl() {}

    public APDUUFormatStopDataTransferActivation build() {
      APDUUFormatStopDataTransferActivation aPDUUFormatStopDataTransferActivation =
          new APDUUFormatStopDataTransferActivation();
      return aPDUUFormatStopDataTransferActivation;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof APDUUFormatStopDataTransferActivation)) {
      return false;
    }
    APDUUFormatStopDataTransferActivation that = (APDUUFormatStopDataTransferActivation) o;
    return super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode());
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
