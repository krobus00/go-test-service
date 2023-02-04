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
				checkout([
					$class: 'GitSCM',
					branches: [[name: '*/${GIT_BRANCH}']],
					doGenerateSubmoduleConfigurations: false,
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
        stage("unit-test") {
            steps {
                sh 'make test'
            }
        }
    }
}