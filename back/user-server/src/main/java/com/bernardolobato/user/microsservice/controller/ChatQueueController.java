 package com.bernardolobato.user.microsservice.controller;

import java.util.Optional;
import javax.validation.Valid;
import com.bernardolobato.user.microsservice.controller.form.QueueForm;
import com.bernardolobato.user.microsservice.model.ChatQueue;
import com.bernardolobato.user.microsservice.model.QueueStatus;
import com.bernardolobato.user.microsservice.model.User;
import com.bernardolobato.user.microsservice.repository.QueueRepository;
import com.bernardolobato.user.microsservice.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/queues")
public class ChatQueueController {

    @Autowired
    QueueRepository queueRepository;

    @Autowired
    UserRepository userRepository;

    @PostMapping("/")
public ResponseEntity<?> salvar(@Valid @RequestBody QueueForm queueForm) {
        try {
            ChatQueue c = new ChatQueue();
            Optional<User> u = userRepository.findById(queueForm.getPatientId());
            if (u.isPresent()) {
                c.setPatient(u.get());
                c.setStatus(QueueStatus.WAITING);
                c = this.queueRepository.save(c);
                return ResponseEntity.ok().body(c);
            } else {
                return ResponseEntity.notFound().build();
            }
        } catch (Exception e) {
            System.err.println(e);
            return ResponseEntity.badRequest().body(e.getMessage());
        }

    }

    @PostMapping("/pull")
    public ResponseEntity<?> pull(@Valid @RequestBody QueueForm queueForm) {
            try {
                ChatQueue c = this.queueRepository.findFirstByStatus(QueueStatus.WAITING);
                Optional<User> u = userRepository.findById(queueForm.getDoctorId());
                if (u.isPresent()) {
                    c.setDoctor(u.get());
                    c.setStatus(QueueStatus.IN_EXECUTION);
                    c = this.queueRepository.save(c);
                    return ResponseEntity.ok().body(c);
                } else {
                    return ResponseEntity.notFound().build();
                }
            } catch (Exception e) {
                System.err.println(e);
                return ResponseEntity.badRequest().body(e.getMessage());
            }
    
        }
        @PostMapping("/finish")
    public ResponseEntity<?> finish(@Valid @RequestBody QueueForm queueForm) {
            try {
                ChatQueue c = this.queueRepository.findFirstByStatus(QueueStatus.WAITING);
                Optional<User> u = userRepository.findById(queueForm.getDoctorId());
                if (u.isPresent()) {
                    c.setDoctor(u.get());
                    c.setStatus(QueueStatus.IN_EXECUTION);
                    c = this.queueRepository.save(c);
                    return ResponseEntity.ok().body(c);
                } else {
                    return ResponseEntity.notFound().build();
                }
            } catch (Exception e) {
                System.err.println(e);
                return ResponseEntity.badRequest().body(e.getMessage());
            }
    
        }
}
