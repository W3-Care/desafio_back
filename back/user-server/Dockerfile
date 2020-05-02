FROM openjdk:8

# api
COPY ./target/w3-care-user-api-1.0.jar /opt/w3-care-user-api-1.0.jar

# Ports we will expose
EXPOSE 8080

CMD ["java", "-jar", "/opt/w3-care-user-api-1.0.jar"]
