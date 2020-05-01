 package com.bernardolobato.user.microsservice.validator.annotation;

import java.lang.annotation.Documented;
import java.lang.annotation.Retention;
import java.lang.annotation.Target;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.ElementType;
import javax.validation.Constraint;
import javax.validation.Payload;
import com.bernardolobato.user.microsservice.validator.UniqueEmailValidatorImpl;

@Documented
@Constraint(validatedBy = UniqueEmailValidatorImpl.class)
@Target( { ElementType.METHOD, ElementType.FIELD })
@Retention(RetentionPolicy.RUNTIME)

public @interface UniqueEmailValidator {
	   String message() default "Este e-mail já está cadastrado";
       Class<?>[] groups() default {};
	   Class<? extends Payload>[] payload() default {};
	
}