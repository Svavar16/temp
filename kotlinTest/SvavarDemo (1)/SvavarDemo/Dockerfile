#FROM amazoncorretto:17-alpine3.18-jdk as BUILD
#LABEL authors="Svavar Páll Guðgeirsson"
#
#COPY . /src
#WORKDIR /src
#RUN ./gradlew clean build --no-daemon
#
## Should try to get this smaller, is 341mb right now
#FROM amazoncorretto:17-alpine3.18-jdk as runtime
#
#COPY --from=BUILD /src/build/libs/SvavarDemo-0.0.1-SNAPSHOT.jar run.jar
#
#ENTRYPOINT ["java", "-jar", "run.jar"]
# Build stage
#FROM openjdk:17-alpine3.18 as BUILD
#LABEL authors="Svavar Páll Guðgeirsson"
#
#WORKDIR /src
#COPY . .
#
#RUN ./gradlew clean build --no-daemon
#
## Final image
#FROM openjdk:17-alpine3.18
#
#COPY --from=BUILD /src/build/libs/SvavarDemo-0.0.1-SNAPSHOT.jar run.jar
#
#ENTRYPOINT ["java", "-jar", "run.jar"]
#FROM openjdk:17-jdk-alpine3.14 as BUILD
#LABEL authors="Svavar Páll Guðgeirsson"
#
#COPY . /src
#WORKDIR /src
#RUN ./gradlew clean build --no-daemon
#
## Should try to get this smaller, is 341mb right now
#FROM openjdk:17-jdk-alpine3.14
#
#COPY --from=BUILD /src/build/libs/SvavarDemo-0.0.1-SNAPSHOT.jar run.jar
#
#ENTRYPOINT ["java", "-jar", "run.jar"]
FROM amazoncorretto:17-alpine3.18-jdk as BUILD
LABEL authors="Svavar Páll Guðgeirsson"

COPY . /src
WORKDIR /src
RUN ./gradlew clean build --no-daemon

# Should try to get this smaller, is 341mb right now
FROM amazoncorretto:17-alpine3.18-jdk

COPY --from=BUILD /src/build/libs/SvavarDemo-0.0.1-SNAPSHOT.jar run.jar

ENTRYPOINT ["java", "-jar", "run.jar"]