pipeline {
    agent any

    stages {
        stage('Hello') {
            steps {
                echo 'Hello, World!'
            }
        }
    }
    stage('Build') {
            steps {
                echo 'build go application'
                //sh 'go build -o bubblepays'
                make build
         }
    }
}

