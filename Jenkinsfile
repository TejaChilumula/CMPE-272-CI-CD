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
        stage('Check Backend') {
            steps {
                script {
                    dir('Backend') {
                        bat 'start go run main.go &'
                        // Add any other commands to check the backend
                    }
                }
            }
        }

        stage('Start Frontend') {
            steps {
                script {
                    dir('Frontend/chatapp') {
                        bat 'npm install'
                        bat 'npm start &'
                        //bat 'export REACT_APP_BACKEND_URL="http://localhost:8080"' // Replace 8080 with your backend port
                        bat 'sleep 20s' // Wait for the frontend to start
                        bat 'echo Frontend is running at http://localhost:3000' // Print the URL
                        // Add any other commands to check the frontend
                    }
                }
            }
        }

        // Add more stages as needed
    }

    post {
        always {
            echo 'This will always run'
        }
        success {
            echo 'This will run only if successful'
        }
        failure {
            echo 'This will run only if failed'
        }
    }
}
