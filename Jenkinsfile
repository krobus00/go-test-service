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
			agent {
                docker {
                    image 'golangci/golangci-lint:v1.46.2'
                }
            }
            steps {
                // Create our project directory.
                sh 'cd ${GOPATH}/src'
                sh 'mkdir -p ${GOPATH}/src/go-test-service'
                // Copy all files in our Jenkins workspace to our project directory.
                sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/go-test-service'
                sh 'chmod -R 777 ${GOPATH}/src/go-test-service'
                catchError {
                    sh 'make lint'
                }
            }
            post {
                success {
                    echo 'Static code analysis stage successful'
                }
                failure {
                    error('Build is aborted due to failure of static code analysis stage')
                }
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