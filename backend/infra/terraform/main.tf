provider "kubernetes" {
  config_path = "~/.kube/config"
}

resource "kubernetes_deployment" "postgres" {
  metadata {
    name = "postgres"
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "postgres"
      }
    }
    template {
      metadata {
        labels = {
          app = "postgres"
        }
      }
      spec {
        container {
          image = "postgres:15-alpine"
          name  = "postgres"
          env {
            name  = "POSTGRES_USER"
            value = "user"
          }
          env {
            name  = "POSTGRES_PASSWORD"
            value = "pass"
          }
          env {
            name  = "POSTGRES_DB"
            value = "shopnest"
          }
          port {
            container_port = 5432
          }
        }
      }
    }
  }
}

# Redis Deployment
resource "kubernetes_deployment" "redis" {
  metadata {
    name = "redis"
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "redis"
      }
    }
    template {
      metadata {
        labels = {
          app = "redis"
        }
      }
      spec {
        container {
          image = "redis:alpine"
          name  = "redis"
          port {
            container_port = 6379
          }
        }
      }
    }
  }
}

# RabbitMQ Deployment
resource "kubernetes_deployment" "rabbitmq" {
  metadata {
    name = "rabbitmq"
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "rabbitmq"
      }
    }
    template {
      metadata {
        labels = {
          app = "rabbitmq"
        }
      }
      spec {
        container {
          image = "rabbitmq:management"
          name  = "rabbitmq"
          env {
            name  = "RABBITMQ_DEFAULT_USER"
            value = "user"
          }
          env {
            name  = "RABBITMQ_DEFAULT_PASS"
            value = "pass"
          }
          port {
            container_port = 5672
          }
          port {
            container_port = 15672
          }
        }
      }
    }
  }
}

# User Service Deployment
resource "kubernetes_deployment" "user_service" {
  metadata {
    name = "user-service"
  }
  spec {
    replicas = 2
    selector {
      match_labels = {
        app = "user-service"
      }
    }
    template {
      metadata {
        labels = {
          app = "user-service"
        }
      }
      spec {
        container {
          image = "your-registry/user-service:latest"
          name  = "user-service"
          env {
            name  = "DB_CONNECTION_STRING"
            value = "postgres://user:pass@postgres:5432/shopnest"
          }
          env {
            name  = "REDIS_ADDRESS"
            value = "redis:6379"
          }
          env {
            name  = "RABBITMQ_URL"
            value = "amqp://user:pass@rabbitmq:5672"
          }
          port {
            container_port = 8080
          }
        }
      }
    }
  }
}

# Email Worker Deployment
resource "kubernetes_deployment" "email_worker" {
  metadata {
    name = "email-worker"
  }
  spec {
    replicas = 2
    selector {
      match_labels = {
        app = "email-worker"
      }
    }
    template {
      metadata {
        labels = {
          app = "email-worker"
        }
      }
      spec {
        container {
          image = "your-registry/email-worker:latest"
          name  = "email-worker"
          env {
            name  = "RABBITMQ_URL"
            value = "amqp://user:pass@rabbitmq:5672"
          }
          env {
            name  = "SMTP_HOST"
            value = "your-smtp-server.com"
          }
        }
      }
    }
  }
}

# Kong API Gateway Deployment
resource "kubernetes_deployment" "kong" {
  metadata {
    name = "kong"
  }
  spec {
    replicas = 2
    selector {
      match_labels = {
        app = "kong"
      }
    }
    template {
      metadata {
        labels = {
          app = "kong"
        }
      }
      spec {
        container {
          image = "kong:latest"
          name  = "kong"
          env {
            name  = "KONG_DATABASE"
            value = "postgres"
          }
          env {
            name  = "KONG_PG_HOST"
            value = "postgres"
          }
          env {
            name  = "KONG_PG_USER"
            value = "user"
          }
          env {
            name  = "KONG_PG_PASSWORD"
            value = "pass"
          }
          port {
            container_port = 8000
          }
          port {
            container_port = 8443
          }
          port {
            container_port = 8001
          }
        }
      }
    }
  }
}

# Services for all components
resource "kubernetes_service" "postgres" {
  metadata {
    name = "postgres"
  }
  spec {
    selector = {
      app = "postgres"
    }
    port {
      port        = 5432
      target_port = 5432
    }
  }
}

resource "kubernetes_service" "redis" {
  metadata {
    name = "redis"
  }
  spec {
    selector = {
      app = "redis"
    }
    port {
      port        = 6379
      target_port = 6379
    }
  }
}

resource "kubernetes_service" "rabbitmq" {
  metadata {
    name = "rabbitmq"
  }
  spec {
    selector = {
      app = "rabbitmq"
    }
    port {
      name       = "amqp"
      port       = 5672
      target_port = 5672
    }
    port {
      name       = "management"
      port       = 15672
      target_port = 15672
    }
  }
}

resource "kubernetes_service" "user_service" {
  metadata {
    name = "user-service"
  }
  spec {
    selector = {
      app = "user-service"
    }
    port {
      port        = 8080
      target_port = 8080
    }
  }
}

resource "kubernetes_service" "kong" {
  metadata {
    name = "kong"
  }
  spec {
    type = "LoadBalancer"
    selector = {
      app = "kong"
    }
    port {
      name       = "http"
      port       = 80
      target_port = 8000
    }
    port {
      name       = "https"
      port       = 443
      target_port = 8443
    }
    port {
      name       = "admin"
      port       = 8001
      target_port = 8001
    }
  }
}