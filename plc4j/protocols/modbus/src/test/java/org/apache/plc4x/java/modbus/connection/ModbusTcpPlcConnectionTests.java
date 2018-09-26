/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
 */

package org.apache.plc4x.java.modbus.connection;

import io.netty.channel.Channel;
import org.apache.commons.lang3.reflect.FieldUtils;
import org.junit.After;
import org.junit.Before;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.net.InetAddress;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

import static org.mockito.Mockito.RETURNS_DEEP_STUBS;
import static org.mockito.Mockito.mock;

public class ModbusTcpPlcConnectionTests {

    private static final Logger LOGGER = LoggerFactory.getLogger(ModbusTcpPlcConnectionTests.class);

    private ModbusTcpPlcConnection SUT;

    private Channel channelMock;

    private ExecutorService executorService;

    @Before
    public void setUp() throws Exception {
        SUT = ModbusTcpPlcConnection.of(InetAddress.getByName("localhost"), null);
        channelMock = mock(Channel.class, RETURNS_DEEP_STUBS);
        FieldUtils.writeField(SUT, "channel", channelMock, true);
        executorService = Executors.newFixedThreadPool(10);
    }

    @After
    public void tearDown() {
        executorService.shutdownNow();
        SUT = null;
    }

}