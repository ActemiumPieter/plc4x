package org.apache.plc4x;

import org.apache.nifi.annotation.documentation.CapabilityDescription;
import org.apache.nifi.annotation.documentation.Tags;
import org.apache.nifi.controller.ControllerService;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;

import java.util.concurrent.ExecutionException;


@Tags({"Service API"})
@CapabilityDescription("Example Service API.")
public interface MyService extends ControllerService {

    PlcConnection getConnection()throws ExecutionException, InterruptedException, PlcConnectionException;

    public String getPlc4xConnecionString();

    void refreshConnectionManager();

    void closeConnection();
}
