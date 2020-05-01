 package com.bernardolobato.user.microsservice.controller.form;

import lombok.AllArgsConstructor;
import lombok.Getter;

@Getter
@AllArgsConstructor
public class QueueForm {
    Long id;
    Long patientId;
    Long doctorId;
    
    public QueueForm() {
    }
}
