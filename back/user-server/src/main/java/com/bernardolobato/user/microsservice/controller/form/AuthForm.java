 package com.bernardolobato.user.microsservice.controller.form;

import javax.validation.constraints.NotBlank;
import lombok.AllArgsConstructor;
import lombok.Getter;

@Getter
@AllArgsConstructor
public class AuthForm {
    @NotBlank
    String email;

    @NotBlank
    String password;

    public AuthForm() {}
}
