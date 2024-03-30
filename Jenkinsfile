pipeline {
    agent any
    
    stages {
        stage('Checkout') {
            steps {
                // Checkout the specific branch
                checkout([$class: 'GitSCM',
                          branches: [[name: 'master']],
                          userRemoteConfigs: [[url: 'git@github.com:alekhyarani2812/bubblepays.git']]])
            }
        }
        stage('Run tests') {
            steps {
                // Run go tests
                sh 'echo "Run go tests"'
                sh 'make rununittests'
            }
        } 
        stage('Build') {
            steps {
                // Run make target
                sh 'make build'
            }
        }
    }
}
