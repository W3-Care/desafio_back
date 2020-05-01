package com.bernardolobato.user.microsservice.validator;

import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;
import javax.validation.ConstraintValidator;
import javax.validation.ConstraintValidatorContext;
import com.bernardolobato.user.microsservice.validator.annotation.ValueOfEnumValidator;

public class ValueOfEnumValidatorImpl implements ConstraintValidator<ValueOfEnumValidator, CharSequence> {
    private List<String> acceptedValues;
 
    @Override
    public void initialize(ValueOfEnumValidator annotation) {
        acceptedValues = Stream.of(annotation.enumClass().getEnumConstants())
                .map(Enum::name)
                .collect(Collectors.toList());
    }
 
    @Override
    public boolean isValid(CharSequence value, ConstraintValidatorContext context) {
        if (value == null) {
            return true;
        }
 
        return acceptedValues.contains(value.toString());
    }
}