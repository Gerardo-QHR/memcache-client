FROM gradle:8.7.0-jdk17 as builder

WORKDIR /app

COPY build.gradle .
COPY settings.gradle .
COPY gradlew .
COPY gradle/ gradle/

RUN ./gradlew build --no-daemon

COPY . .

RUN ./gradlew build

FROM eclipse-temurin:17

WORKDIR /app

COPY --from=builder /app/build/libs/*.jar app.jar

EXPOSE 8080

CMD ["java", "-jar", "app.jar"]
