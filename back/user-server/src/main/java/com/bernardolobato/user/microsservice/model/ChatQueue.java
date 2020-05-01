 package com.bernardolobato.user.microsservice.model;

import java.util.Date;
import javax.persistence.CascadeType;
import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Entity
@Getter
@AllArgsConstructor
@Setter
@NoArgsConstructor
public class ChatQueue {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    Long id;

    @ManyToOne( cascade=CascadeType.MERGE )
    @JoinColumn(name="doctor_id")
    @Setter
    User doctor;
    
    @ManyToOne(cascade=CascadeType.MERGE )
    @JoinColumn(name="patient_id")
    @Setter
    User patient;

    Date startDate;
    Date endDate;

    @Setter
    @Enumerated(EnumType.STRING)
    private QueueStatus status;
}
