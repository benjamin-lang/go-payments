pipeline {
    agent none

    environment {
        SERVICE_VERSION = VersionNumber(projectStartDate: '2019-01-01', worstResultForIncrement: 'SUCCESS'
            , versionNumberString: '${BUILD_DATE_FORMATTED,"yyyyMMdd"}-r${BUILDS_TODAY, XX}', versionPrefix: '')
        SERVICE_NAME = 'go-payments'
        SCM_URL = 'git@github.com:benjamin-lang/go-payments.git'
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
                git(url: SCM_URL, credentialsId: 'jenkins_ssh_credentials', branch: '$BRANCH_NAME')
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}
