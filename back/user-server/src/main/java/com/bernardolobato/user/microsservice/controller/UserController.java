 package com.bernardolobato.user.microsservice.controller;

import java.util.List;
import java.util.Optional;
import javax.validation.Valid;
import com.bernardolobato.user.microsservice.controller.form.UserForm;
import com.bernardolobato.user.microsservice.model.User;
import com.bernardolobato.user.microsservice.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.dao.EmptyResultDataAccessException;
import org.springframework.http.ResponseEntity;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/usuarios")
public class UserController {

    @Autowired
    UserRepository userRepository;

    @Autowired
    private  PasswordEncoder passwordEncoder;
    
    @PostMapping("/")
    public ResponseEntity<?> save(@Valid @RequestBody UserForm userForm) {
        try {
            userForm.setPassword(passwordEncoder.encode(userForm.getPassword()));
            User u = this.userRepository.save(userForm.converter());
            return ResponseEntity.ok().body(u);
        } catch (Exception e) {
            System.err.println(e);
            return ResponseEntity.badRequest().body(e.getMessage());
        }

    }

    @GetMapping("/")
    public ResponseEntity<?> list() {
        try {
            final List<User> usuarios = this.userRepository.findAll();
            return ResponseEntity.ok().body(usuarios);
        } catch (final Exception e) {
            return ResponseEntity.badRequest().build();
        }

    }

    @GetMapping("/{id}")
    public ResponseEntity<?> find(@PathVariable final Long id) {
        final Optional<User> u = this.userRepository.findById(id);
        if (u.isPresent()) {
            return ResponseEntity.ok().body(u.get());
        } else {
            return ResponseEntity.notFound().build();
        }
    }

    @PutMapping("/{id}")
    public ResponseEntity<?> update(@RequestBody final UserForm user, @PathVariable final Long id) {

        final Optional<User> u = userRepository.findById(id).map(el -> {
            el.setName(user.getName());
            el.setBirthday(user.getBirthday());
            el.setEmail(user.getEmail());
            return userRepository.save(el);
        });
        if (u.isPresent()) {
            return ResponseEntity.ok().body(u.get());
        } else {
            return ResponseEntity.notFound().build();
        }
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<?> deleta(@PathVariable final Long id) {
        try {
            userRepository.deleteById(id);
            return ResponseEntity.ok().build();
        } catch (final EmptyResultDataAccessException e) {
            return ResponseEntity.notFound().build();
        } catch (final Exception e) {
        System.out.println(e);
        return ResponseEntity.badRequest().build();
    }
  }
}
