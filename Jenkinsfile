pipeline {
    agent any
    tools {
        go 'go1.18'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0
    }
    stages {
		stage('Checkout') {
			steps {
				checkout([
					$class: 'GitSCM',
					branches: [[name: '*/${GIT_BRANCH}']],
					clean: true,
					extensions: [],
					submoduleCfg: [],
					userRemoteConfigs: [[
					name: 'origin',
					refspec: '+refs/pull/${ghprbPullId}/*:refs/remotes/origin/pr/${ghprbPullId}/*',
					url: 'https://github.com/krobus00/go-test-service.git'
					]]
				])
			}
		}
		stage('Code Analysis') {
			steps {
				
                sh 'go version'
                sh 'curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin'
                sh 'make lint'
								
                
            }
		}
        stage("Unit Test") {
            steps {
                sh 'make test'
            }
        }
		// stage("build") {
        //     steps {
        //         echo 'BUILD EXECUTION STARTED'
        //         sh 'go version'
        //         sh 'go get ./...'
        //         sh 'docker build . -t krobus00/go-test-service'
        //     }
        // }
        // stage('deliver') {
        //     agent any
        //     steps {
        //         withCredentials([usernamePassword(credentialsId: 'dockerhub', passwordVariable: 'dockerhubPassword', usernameVariable: 'dockerhubUser')]) {
        //         sh "docker login -u ${env.dockerhubUser} -p ${env.dockerhubPassword}"
        //         sh 'docker push krobus00/go-test-service'
        //         }
        //     }
        // }
    }
	post {
        always {
            cleanWs()
        }
    }
}