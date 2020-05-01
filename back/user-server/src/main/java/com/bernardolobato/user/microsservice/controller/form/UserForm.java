 package com.bernardolobato.user.microsservice.controller.form;

import java.util.Date;
import javax.validation.constraints.NotBlank;
import com.bernardolobato.user.microsservice.model.User;
import com.bernardolobato.user.microsservice.model.UserType;
import com.bernardolobato.user.microsservice.validator.annotation.UniqueEmailValidator;
import com.bernardolobato.user.microsservice.validator.annotation.ValueOfEnumValidator;
import com.fasterxml.jackson.annotation.JsonFormat;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

@Getter
@AllArgsConstructor
public class UserForm {
    String name;
    @JsonFormat
          (shape = JsonFormat.Shape.STRING, pattern = "dd/MM/yyyy")          
    Date birthday;

    @NotBlank
    @UniqueEmailValidator
    String email;

    @Setter
    String password;

    @NotBlank
    @ValueOfEnumValidator(enumClass = UserType.class)
    String type;

    public UserForm() {

    }

    public User converter() {
        return new User(null, this.name, this.birthday, this.email, this.password, UserType.valueOf(this.type));
    }

}
