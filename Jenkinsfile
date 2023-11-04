pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                script {
                    git 'https://github.com/TejaChilumula/CMPE-272-CI-CD.git'
                }
            }
        }

        stage('Build Frontend') {
            steps {
                dir('frontend/chatapp') {
                    sh 'npm install'
                    sh 'npm run build'
                }
            }
        }

        stage('Build Backend') {
            steps {
                dir('backend') {
                    sh 'go build'
                }
            }
        }

        stage('Test') {
            steps {
                dir('frontend/chatapp') {
                    sh 'npm test'
                }
                dir('backend') {
                    sh 'go test ./...'
                }
            }
        }

        stage('Deploy') {
            steps {
                // Add your deployment commands here
            }
        }
    }

    post {
        always {
            echo 'This will always run'
        }

        success {
            echo 'This will run only if successful - Teja'
        }

        failure {
            echo 'This will run only if failed - Teja'
        }
    }
}
