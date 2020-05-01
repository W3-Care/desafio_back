 package com.bernardolobato.user.microsservice.security;

import java.io.IOException;
import java.util.Date;
import javax.servlet.FilterChain;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import com.bernardolobato.user.microsservice.controller.form.UserForm;
import com.bernardolobato.user.microsservice.model.User;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import io.jsonwebtoken.security.Keys;

public class JwtAuthenticationFilter extends UsernamePasswordAuthenticationFilter {

    private final AuthenticationManager authenticationManager;

    public JwtAuthenticationFilter(AuthenticationManager authenticationManager) {
        this.authenticationManager = authenticationManager;

        setFilterProcessesUrl(SecurityConstants.AUTH_LOGIN_URL);
    }

    @Override
    public Authentication attemptAuthentication(HttpServletRequest request,
            HttpServletResponse response) {
                UserForm uf = null;
        try {
            uf = new ObjectMapper().readValue(request.getInputStream(), UserForm.class);
        } catch (IOException e) {
            e.printStackTrace();
        }
        Authentication authenticationToken = new UsernamePasswordAuthenticationToken(uf.getEmail(), uf.getPassword());
        return authenticationManager.authenticate(authenticationToken);
    }

    @Override
    protected void successfulAuthentication(HttpServletRequest request,
            HttpServletResponse response, FilterChain filterChain, Authentication authentication) {
        byte[] signingKey = SecurityConstants.JWT_SECRET.getBytes();
        String details = "";
        try {
            details = (new ObjectMapper()).writeValueAsString(authentication.getPrincipal());
        } catch (Exception e) {
            this.logger.error(e.getMessage());
        }
        String token =
                Jwts.builder().signWith(Keys.hmacShaKeyFor(signingKey), SignatureAlgorithm.HS512)
                        .setHeaderParam("typ", SecurityConstants.TOKEN_TYPE)
                        .setSubject(((User) authentication.getPrincipal()).getEmail())
                        .setIssuer(SecurityConstants.TOKEN_ISSUER)
                        .setExpiration(new Date(System.currentTimeMillis() + 864000000))
                        .claim("details", details)
                        .compact();
                        response.addHeader("Access-Control-Expose-Headers", "Authorization");
                        response.addHeader("Access-Control-Allow-Headers", "Authorization, X-PINGOTHER, Origin, X-Requested-With, Content-Type, Accept, X-Custom-header");                          
                        response.addHeader(SecurityConstants.TOKEN_HEADER, SecurityConstants.TOKEN_PREFIX + token);
    }

}
