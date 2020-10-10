pipeline {
    agent any

    stages {
        stage('Build Image') {
            steps {
                sh 'make docker'
            }
        }
        stage('Upload Image') {
            steps {
                sh 'docker images'
            }
        }
        stage('Test Fabcar') {
            steps {
                echo 'Test Fabcar'
            }
        }
    }
}
