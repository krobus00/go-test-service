pipeline {
    agent any
    tools {
        go 'go1.18'
        nodejs "nodejs18"
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
                    credentialsId: 'gh-krobus00',
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
        stage("DangerJS Code Review") {
            environment {
                DANGER_FAKE_CI="YEP"
                DANGER_TEST_REPO="${params.ghprbGhRepository}"
                GITHUB_PR_ID="${params.ghprbPullId}"
                DANGER_TEST_PR="${params.ghprbPullId}"
            }
            steps {
                withCredentials([string(credentialsId: 'api-gh-krobus00', variable: 'DANGER_GITHUB_API_TOKEN')]) {
                    echo ''
                    sh 'danger ci --id $ghprbActualCommit'
                }
            }
        }
    }
	post {
        always {
            cleanWs()
        }
    }
}