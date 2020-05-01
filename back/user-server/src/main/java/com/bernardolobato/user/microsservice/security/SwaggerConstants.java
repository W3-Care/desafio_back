package com.bernardolobato.user.microsservice.security;
 

public final class SwaggerConstants {

    public static final String URL_UI = "/swagger-ui.html";
    public static final String URL_ASSETS = "/webjars/**";
    public static final String URL_RESOURCES = "/swagger-resources/**";
    public static final String API_DOCS = "/v2/api-docs";
    

    private SwaggerConstants() {
        throw new IllegalStateException("Cannot create instance of static util class");
    }
}