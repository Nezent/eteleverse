pipeline {
    agent any
    
    environment {
        APP_NAME = 'eteleverse'
        DOCKER_IMAGE = "${APP_NAME}:${BUILD_NUMBER}"
    }
    
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        
        stage('Build') {
            steps {
                script {
                    sh 'go mod download'
                    sh 'go build -o main .'
                }
            }
        }
        
        stage('Test') {
            steps {
                script {
                    sh 'go test -v ./test/...'
                }
            }
        }
        
        stage('Docker Build') {
            steps {
                script {
                    sh "docker build -t ${DOCKER_IMAGE} ."
                }
            }
        }
        
        stage('Docker Deploy') {
            steps {
                script {
                    sh 'docker stop eteleverse-api-1 || true'
                    sh 'docker rm eteleverse-api-1 || true'
                    sh "docker run -d --name eteleverse-api-1 --network eteleverse_app-network -p 8080:8080 -e PORT=8080 ${DOCKER_IMAGE}"
                }
            }
        }
        
        stage('Health Check') {
            steps {
                script {
                    sleep(time: 10, unit: 'SECONDS')
                    
                    def maxRetries = 5
                    def retryCount = 0
                    def healthCheckPassed = false
                    
                    while (retryCount < maxRetries && !healthCheckPassed) {
                        try {
                            sh '''
                                response=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/api/v1/health)
                                if [ "$response" -eq 200 ]; then
                                    echo "Health check passed"
                                    exit 0
                                else
                                    echo "Health check failed with status: $response"
                                    exit 1
                                fi
                            '''
                            healthCheckPassed = true
                        } catch (Exception e) {
                            retryCount++
                            if (retryCount < maxRetries) {
                                echo "Health check attempt ${retryCount} failed, retrying..."
                                sleep(time: 5, unit: 'SECONDS')
                            } else {
                                error "Health check failed after ${maxRetries} attempts"
                            }
                        }
                    }
                }
            }
        }
    }
    
    post {
        success {
            echo 'Pipeline completed successfully!'
        }
        failure {
            echo 'Pipeline failed!'
            sh 'docker stop eteleverse-api-1 || true'
            sh 'docker rm eteleverse-api-1 || true'
        }
        always {
            sh 'docker system prune -f || true'
        }
    }
}
