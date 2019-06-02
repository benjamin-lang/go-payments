#!groovy​

pipeline {
    agent {
        node 'docker'
    }

    environment {
        SERVICE_VERSION = VersionNumber(projectStartDate: '2019-01-01', worstResultForIncrement: 'SUCCESS'
                , versionNumberString: '${BUILD_DATE_FORMATTED,"yyyyMMdd"}-r${BUILDS_TODAY, XX}', versionPrefix: '')
        SERVICE_NAME = 'payment-service-go'
        SCM_URL = 'git@github.com:benjamin-lang/go-payments.git'
        DOCKER_REGISTRY = 'springfield:5000'
    }

    options {
        buildDiscarder(logRotator(numToKeepStr: '7'))
    }

    stages {
        stage('init') {
            steps {
                echo sh(returnStdout: true, script: 'env')
            }
        }
        stage('display version') {
            when {
                branch 'master'
            }
            steps {
                script {
                    currentBuild.displayName = SERVICE_VERSION
                }
            }
        }
        stage('ssh checkout') {
            steps {
                git(url: SCM_URL, credentialsId: 'jenkins_github_credentials', branch: '$BRANCH_NAME')
            }
        }
        stage('docker') {
            steps {
                script {
                    docker.build("$DOCKER_REGISTRY/$SERVICE_NAME:$SERVICE_VERSION", "--build-arg BUILD_ID=$BUILD_ID -f ./build/docker/Dockerfile .")
                    sh 'docker image prune -f --filter label=stage=builder --filter label=build=$BUILD_ID'
                }
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}