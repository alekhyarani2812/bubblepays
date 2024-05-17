pipeline {
    agent any

    stages {
        stage('Hello') {
            steps {
                echo 'Hello, World!'
                echo 'Thius is Jenkins 1st build'
                echo '**************************'
            }
        }
        stage('Build') {
            steps {
                echo 'build go application'
                //sh 'go build -o bubblepays'
                sh 'make build'
         }
    }
    }
}
