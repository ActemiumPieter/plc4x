package org.apache.plc4x;

import org.apache.nifi.components.ValidationContext;
import org.apache.nifi.components.ValidationResult;
import org.apache.nifi.components.Validator;
import org.apache.plc4x.java.DefaultPlcDriverManager;
import org.apache.plc4x.java.api.PlcDriver;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;


public class Plc4xConnectionStringValidator implements Validator {
    @Override
    public ValidationResult validate(String subject, String input, ValidationContext context) {
        DefaultPlcDriverManager manager = new DefaultPlcDriverManager();

        if (context.isExpressionLanguageSupported(subject) && context.isExpressionLanguagePresent(input)) {
            return new ValidationResult.Builder().subject(subject).input(input).explanation("Expression Language Present").valid(true).build();
        }
        try {
            PlcDriver driver =  manager.getDriverForUrl(input);
            driver.getConnection(input);
        } catch (PlcConnectionException e) {
            return new ValidationResult.Builder().subject(subject)
                .explanation(e.getMessage())
                .valid(false)
                .build();
        }
        return new ValidationResult.Builder().subject(subject)
            .explanation("")
            .valid(true)
            .build();
    }
}