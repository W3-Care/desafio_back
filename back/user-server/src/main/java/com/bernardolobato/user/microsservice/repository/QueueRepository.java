 package com.bernardolobato.user.microsservice.repository;

import com.bernardolobato.user.microsservice.model.ChatQueue;
import com.bernardolobato.user.microsservice.model.QueueStatus;
import org.springframework.data.jpa.repository.JpaRepository;

public interface QueueRepository extends JpaRepository<ChatQueue, Long> {


    public ChatQueue findFirstByStatusOrderById(QueueStatus status);

    public ChatQueue findFirstByPatientIdAndStatus(Long patientId, QueueStatus status);

}
