pipeline {
    agent { label 'docker' }
    options {
        skipDefaultCheckout(true)
        disableConcurrentBuilds()
        timeout(time: 30, unit: 'MINUTES')
        buildDiscarder(logRotator(numToKeepStr: '10'))
    }
    parameters {
        booleanParam(name: 'DEMONSTRATE_FAILURE', defaultValue: true,
            description: 'Fail the Test stage deliberately in this disposable educational build.')
    }
    stages {
        stage('Checkout') {
            steps {
                deleteDir()
                dir('github-scm') {
                    checkout(changelog: false, poll: false, scm: [$class: 'GitSCM', branches: [[name: '*/feature/two-containers']], userRemoteConfigs: [[url: 'https://github.com/andresfccy/meli-integracion-continua.git', credentialsId: 'github-meli-readonly']]])
                    sh 'git rev-parse HEAD'
                }
                sh '''
                    echo 'b98ffdb09c4e1e4e4749470eb60f7a53f973bfc58fff351b7b5c78a83b91502e  /tmp/meli-candidate-ac45ad99643aa39645c81eb939b13d5374eafd2a.tar' | sha256sum -c -
                    tar -xf /tmp/meli-candidate-ac45ad99643aa39645c81eb939b13d5374eafd2a.tar
                '''
                script {
                    env.CI_PROJECT = "meli-check-${env.BUILD_TAG}".toLowerCase().replaceAll('[^a-z0-9-]', '-')
                    env.EVIDENCE_DIR = "evidence/entrega2/build-${env.BUILD_NUMBER}"
                    env.WEB_PORT = '0'
                }
            }
        }
        stage('Build') {
            steps { sh 'sh scripts/ci.sh build' }
        }
        stage('Test') {
            steps {
                sh 'sh scripts/ci.sh test'
                script {
                    if (params.DEMONSTRATE_FAILURE) {
                        error('Controlled failure requested; no application data was changed.')
                    }
                }
                sh 'sh scripts/ci.sh smoke'
            }
        }
    }
    post {
        always {
            script {
                if (env.CI_PROJECT) { sh 'sh scripts/ci.sh cleanup' }
            }
        }
        cleanup {
            archiveArtifacts artifacts: 'evidence/entrega2/build-*/**', allowEmptyArchive: true
        }
        success { echo 'Application image builds and all required tests passed.' }
        failure { echo 'Build failed. Inspect the failing stage and archived evidence.' }
    }
}
