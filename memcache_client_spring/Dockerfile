FROM gradle:8.7.0-jdk17 as builder

# Set the working directory in the container
WORKDIR /app

# Copy the Gradle files
COPY build.gradle .
COPY settings.gradle .
COPY gradlew .
COPY gradle/ gradle/

# Build the application
RUN ./gradlew build --no-daemon

# Copy the application code
COPY . .

# Build the application again
RUN ./gradlew build

# Use adoptopenjdk/openjdk17 as the base image
FROM eclipse-temurin:17

# Set the working directory in the container
WORKDIR /app

# Copy the built JAR file from the builder stage
COPY --from=builder /app/build/libs/*.jar app.jar

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["java", "-jar", "app.jar"]
