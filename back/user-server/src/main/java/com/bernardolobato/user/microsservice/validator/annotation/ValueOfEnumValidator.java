package com.bernardolobato.user.microsservice.validator.annotation;

import java.lang.annotation.Documented;
import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;
import javax.validation.Constraint;
import javax.validation.Payload;
import com.bernardolobato.user.microsservice.validator.ValueOfEnumValidatorImpl;

@Target({ElementType.METHOD, ElementType.FIELD, ElementType.ANNOTATION_TYPE, ElementType.CONSTRUCTOR, ElementType.PARAMETER, ElementType.TYPE_USE})
@Retention(RetentionPolicy.RUNTIME)
@Documented
@Constraint(validatedBy = ValueOfEnumValidatorImpl.class)
public @interface ValueOfEnumValidator {
    Class<? extends Enum<?>> enumClass();
    String message() default "Deve ser um dos valores {enumClass}";
    Class<?>[] groups() default {};
    Class<? extends Payload>[] payload() default {};
}