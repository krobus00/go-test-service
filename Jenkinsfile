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
					url: '${ghprbAuthorRepoGitUrl}'
					]]
				])
			}
		}
		stage('Code Analysis') {            
			steps {				
                withEnv(["PATH+GO=${HOME}/go/bin"]){
                    sh 'go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2'
                    sh 'make lint'
                }
            }
		}
        stage("Unit Test") {
            steps {
                sh 'make test'
            }
        }
		stage("Build Image") {
            steps {
                sh 'docker build . -t krobus00/go-test-service test'
            }
        }
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