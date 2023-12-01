package org.apache.plc4x;

import org.apache.nifi.annotation.documentation.CapabilityDescription;
import org.apache.nifi.annotation.documentation.Tags;
import org.apache.nifi.controller.ControllerService;

import java.util.concurrent.ExecutionException;


@Tags({"Service API"})
@CapabilityDescription("Example Service API.")
public interface MyService extends ControllerService {


    public void execute() throws ExecutionException;

    String getProperty();
}
