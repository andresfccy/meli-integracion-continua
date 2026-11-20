FROM docker:29-cli AS docker_cli
FROM jenkins/inbound-agent:jdk21
USER root
COPY --from=docker_cli /usr/local/bin/docker /usr/local/bin/docker
COPY --from=docker_cli /usr/local/libexec/docker/cli-plugins/docker-compose /usr/local/libexec/docker/cli-plugins/docker-compose
COPY --from=docker_cli /usr/local/libexec/docker/cli-plugins/docker-buildx /usr/local/libexec/docker/cli-plugins/docker-buildx
COPY --chmod=0555 start.sh /usr/local/bin/start-docker-agent
USER jenkins
WORKDIR /home/jenkins/agent
ENTRYPOINT ["/usr/local/bin/start-docker-agent"]
