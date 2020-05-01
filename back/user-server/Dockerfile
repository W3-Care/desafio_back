FROM openjdk:8

# api
COPY ./target/source2it-api-1.0.jar /opt/source2it/source2it-api.jar

# Ports we will expose
EXPOSE 8080

CMD ["java", "-jar", "/opt/source2it/source2it-api.jar"]
