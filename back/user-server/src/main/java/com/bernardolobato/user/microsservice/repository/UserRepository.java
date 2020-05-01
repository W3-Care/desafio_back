 package com.bernardolobato.user.microsservice.repository;

import com.bernardolobato.user.microsservice.model.User;
import org.springframework.data.jpa.repository.JpaRepository;

public interface UserRepository extends JpaRepository<User, Long> {

    User findFirstByEmail(String email);

    
}
