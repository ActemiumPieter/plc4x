package org.actemium.plc4x;

import org.apache.nifi.reporting.InitializationException;
import org.apache.nifi.util.TestRunner;
import org.apache.nifi.util.TestRunners;
import org.apache.plc4x.PLCConnectionService;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.apache.plc4x.java.api.PlcConnection;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;


public class Plc4xPLCConnectionServiceTest {

    @BeforeEach
    public void init(){

    }

    @Test
    public void testService() throws InitializationException, PlcConnectionException {
        final TestRunner runner = TestRunners.newTestRunner(TestProcessor.class);
        final PLCConnectionService service = new PLCConnectionService();
        runner.addControllerService("test", service);

        runner.setProperty(service, PLCConnectionService.PLC_CONNECTION_STRING, "simulated://127.0.0.1");
        runner.enableControllerService(service);
        runner.assertValid(service);

        // Attempt to get a connection from the service
        PlcConnection plcConnection = service.getConnection();
        try {
            // Assert that the connection is not null (indicating successful connection)
            assertTrue(plcConnection != null && plcConnection.isConnected());
        } catch (Exception e) {
            e.printStackTrace();
        }
    }


    @Test
    public void testGetPlc4xConnectionString() throws InitializationException {
        final TestRunner runner = TestRunners.newTestRunner(TestProcessor.class);
        final PLCConnectionService service = new PLCConnectionService();
        runner.addControllerService("test", service);
        runner.setProperty(service, PLCConnectionService.PLC_CONNECTION_STRING, "simulated://127.0.0.1");
        runner.enableControllerService(service);
        runner.assertValid(service);

        // Attempt to get a connection from the service
        try {
            String connectionString = service.getPlc4xConnecionString();
            // Assert that the connection is not null (indicating successful connection)
            assertEquals(connectionString, "simulated://127.0.0.1");
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
