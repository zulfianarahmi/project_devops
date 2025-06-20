# project_devops

## Overview

This repository contains the implementation of a proposed DevOps transformation strategy for a one-year-old startup specializing in business reimbursements. As the first DevOps Engineer, the primary goal of this project is to address critical pain points faced by the development team and clients, including slow feature delivery, frequent application downtime/crashes, long Mean Time To Recovery (MTTR), and lack of visibility into the production environment.

The solution aims to transition from a manual, bi-weekly deployment process to an automated, reliable, and observable CI/CD pipeline leveraging AWS cloud services and modern DevOps practices. This transformation is designed to enhance developer productivity, improve application stability, accelerate feature delivery, and provide clear insights into system health, all while staying within the allocated budget.

## Project Goals & Metrics for Success

The success of this transformation will be measured by key DevOps metrics:

* **Uptime/Availability:** Significant reduction in application downtime.
* **Mean Time To Recovery (MTTR):** Drastic decrease in time taken to detect and resolve production issues.
* **Deployment Frequency:** Increase in the frequency of code deployments to production (aiming for daily or multiple times a day).
* **Lead Time for Changes:** Reduction in the time from code commit to production deployment.
* **Change Failure Rate:** Decrease in the percentage of deployments causing new issues.
* **Engineer Satisfaction:** Improved morale and reduced stress for the development team.
* **Client Satisfaction:** Reduced client complaints regarding application stability and feature delivery.

## Architecture & Technology Stack

The proposed architecture leverages a modern containerized approach on AWS, orchestrated by Infrastructure as Code (IaC) and an automated CI/CD pipeline.

**Key Components:**

* **Frontend:** ReactJS application.
* **Backend:** Go (Fiber/FastAPI) application.
* **Version Control:** GitHub (Monorepo strategy).
* **CI/CD:** GitHub Actions.
* **Container Registry:** DockerHub.
* **Infrastructure as Code (IaC):** Terraform.
* **Cloud Provider:** Amazon Web Services (AWS).
* **Container Orchestration:** AWS ECS Fargate.
* **Load Balancing:** AWS Application Load Balancer (ALB).
* **Networking:** AWS VPC, Subnets, Security Groups.
* **Observability:** Prometheus (Metrics), Loki (Logs), Grafana (Visualization), AWS CloudWatch Logs.
* **Database (Future):** AWS RDS (PostgreSQL) and AWS ElastiCache (Redis).
