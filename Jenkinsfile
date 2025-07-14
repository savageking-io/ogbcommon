pipeline {
    agent {
        node {
            label 'Go Builder'
        }
    }

    stages {
        stage('Test') {
            steps {
                echo 'Testing..'
                sh 'go test ./...'
            }
        }
    }
}