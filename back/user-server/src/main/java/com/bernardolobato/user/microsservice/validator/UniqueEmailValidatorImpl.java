 package com.bernardolobato.user.microsservice.validator;

import javax.validation.ConstraintValidator;
import javax.validation.ConstraintValidatorContext;
import com.bernardolobato.user.microsservice.model.User;
import com.bernardolobato.user.microsservice.repository.UserRepository;
import com.bernardolobato.user.microsservice.validator.annotation.UniqueEmailValidator;
import org.springframework.beans.factory.annotation.Autowired;


public class UniqueEmailValidatorImpl implements ConstraintValidator<UniqueEmailValidator, String> {

	
	@Autowired
	private UserRepository userRepository;
	
	@Override
	public boolean isValid(String email, ConstraintValidatorContext cxt) {
		
		User u = userRepository.findFirstByEmail(email);

		if(u==null)
		{
			return true;
		}
		return false;		
	}
}