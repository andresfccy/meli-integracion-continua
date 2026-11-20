FROM jenkins/jenkins:lts

USER jenkins
COPY --chown=jenkins:jenkins plugins.txt /usr/share/jenkins/ref/plugins.txt
RUN jenkins-plugin-cli --plugin-file /usr/share/jenkins/ref/plugins.txt

# Keep builds off the controller without bypassing the setup wizard or security.
RUN mkdir -p /usr/share/jenkins/ref/init.groovy.d && \
    printf '%s\n' \
    'import jenkins.model.Jenkins' \
    'Jenkins.get().setNumExecutors(0)' \
    'Jenkins.get().save()' \
    > /usr/share/jenkins/ref/init.groovy.d/00-controller-executors.groovy
