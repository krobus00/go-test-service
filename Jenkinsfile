pipeline {
    agent any
    tools {
        go 'go1.18'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
		stage('Checkout') {
			steps {
				sh "ls -lar"
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
				sh "ls -lar"
			}
		}
		stage('Code Analysis') {
			steps {
				withEnv(["PATH+GO=${$PATH}/bin:${HOME}/go/bin"]) {
					sh "go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2"
					sh "golangci-lint --version"
				}
                sh 'go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2'
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