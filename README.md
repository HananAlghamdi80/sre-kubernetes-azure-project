
---

# SRE Project on Azure Kubernetes Service (AKS)

## 📌 Overview

This project showcases how to deploy, secure, scale, and monitor a microservices-based system on **Azure Kubernetes Service (AKS)** using core **SRE best practices**.

The system includes **three microservices**, each written in a different programming language:

* **API Service** – Node.js
* **Auth Service** – Go
* **Images Service** – Python

Each service runs in its own Kubernetes Deployment and communicates internally using ClusterIP services.

The project focuses on:

* Securing traffic with **NGINX Ingress + Self-Signed TLS**
* Implementing **Network Policies** (Zero Trust Model)
* Adding reliability features: **HPA**, **Liveness/Readiness/Startup Probes**, **PDB**
* Full observability with **Prometheus + Grafana**
* Real-time alerts using **Alertmanager + Discord Webhook**

The result is a reliable and production-ready architecture running on AKS.

---

## 🏗️ Architecture Diagram

<p align="center">
  <img src="https://raw.githubusercontent.com/HananAlghamdi80/sre-kubernetes-project/main/hananDig.png" width="100%">
</p>

---

## 🧱 Technologies Used

### ☁️ Cloud & Kubernetes

* Azure Kubernetes Service (AKS)
* Kubernetes Deployments
* ClusterIP Services
* NGINX Ingress Controller
* cert-manager (SelfSigned Issuer)
* Secrets & Configurations

### 🐳 Containers

* Docker
* Azure Container Registry (ACR)

### 🧩 Microservices

* Node.js (API Service)
* Go (Auth Service)
* Python (Images Service)

### 🔐 Security & Networking

* Network Policies (Default Deny + Allow Rules)
* Self-Signed TLS Certificates
* Ingress TLS Termination

### 📈 Reliability & Scaling

* Horizontal Pod Autoscaler (HPA)
* Liveness, Readiness, Startup Probes
* Pod Disruption Budgets (PDB)

### 🔍 Monitoring & Alerting

* Prometheus
* Grafana
* ServiceMonitors
* Alertmanager
* Discord Webhook Integration

### 🗂️ Storage

* `emptyDir` volume (for image uploads)

---

## 🚀 Project Components

* API Service
* Auth Service
* Images Service
* Ingress & TLS
* Network Policies
* HPA
* Probes
* PDB
* Prometheus + Grafana
* Alertmanager + Discord

---

## 🛠️ Building & Pushing Docker Images

Each microservice is containerized and pushed to ACR:

```bash
docker build -t api-service .
docker tag api-service hananacr.azurecr.io/api-service:v1
docker push hananacr.azurecr.io/api-service:v1
```

Repeat for `auth-service` and `images-service`.

---

## ☸️ Kubernetes Deployment

All Kubernetes manifests were applied using:

```bash
kubectl apply -f api-deployment.yaml
kubectl apply -f auth-deployment.yaml
kubectl apply -f images-deployment.yaml
```

---

## 🔐 TLS & Ingress

Using cert-manager with SelfSigned issuer:

* Generate certificate
* Store it as a TLS secret
* Apply it in Ingress for HTTPS termination

---

## 🔒 Network Policies

Zero-trust architecture:

* `default-deny-ingress.yaml` blocks everything
* Allow policies enable:

  * API → Auth
  * API → Images

Ensures restricted east-west traffic inside the cluster.

---

## 📈 Autoscaling (HPA)

Each microservice scales based on CPU usage:

```bash
kubectl apply -f hpa-api.yaml
kubectl apply -f hpa-auth.yaml
kubectl apply -f hpa-images.yaml
```

---

## ❤️ Health Probes

Every service includes:

* **Startup Probe** – ensures app initializes correctly
* **Readiness Probe** – controls when the pod receives traffic
* **Liveness Probe** – restarts the pod if it becomes unhealthy

These ensure high reliability and smooth rollouts.

---

## 🧱 Pod Disruption Budgets (PDB)

Used to maintain minimum availability, especially during:

* Node upgrades
* Maintenance events
* Voluntary disruptions

Example:

```bash
kubectl apply -f api-pdb.yaml
```

---

## 🔍 Monitoring Setup

Prometheus scrapes metrics from:

* API
* Auth
* Images

Grafana visualizes metrics through dashboards such as:

* CPU / Memory usage
* Pod restarts
* API latency
* Error rate
* Custom service metrics

---

## 🚨 Alerting Setup

Alertmanager is configured to send alerts to **Discord Webhook**, such as:

* High CPU usage
* Pod crashes / restarts
* Service downtime

This enables fast detection and response to system issues.


🚀 SRE Kubernetes Project — Monitoring, Scaling & Failure Simulation
This project demonstrates a full SRE-style Kubernetes setup on Azure, including workload deployment, secure ingress, autoscaling, network policies, monitoring, alerting, and failure-simulation scenarios.

📦 1. Cluster Workloads Overview
🔹 Pods Status
Shows all running workloads inside the sre namespace.
![Pods Status](https://raw.githubusercontent.com/HananAlghamdi80/sre-kubernetes-azure-project/main/%D9%80%20Pods.png)


🔹 Deployments
Ensures each service has the required number of replicas and is up-to-date.
![Deployments](https://raw.githubusercontent.com/HananAlghamdi80/sre-kubernetes-azure-project/main/D.png)


🔹 Services
ClusterIP services exposing the API, Auth, and Images microservices internally.
![Services](https://raw.githubusercontent.com/HananAlghamdi80/sre-kubernetes-azure-project/main/L.png)


🔹 TLS Certificate (Let's Encrypt)
Valid HTTPS certificate issued via cert-manager using ACME challenges.
![TLS Certificate](https://raw.githubusercontent.com/HananAlghamdi80/sre-kubernetes-azure-project/main/CER.png)


🔹 Network Policies
Zero-trust networking implemented to restrict service-to-service communication.
![Network Policies](https://raw.githubusercontent.com/HananAlghamdi80/sre-kubernetes-azure-project/main/net.png)


⚠️ 2. Failure Simulation Scenarios (SRE Testing)
Below are real SRE simulations performed to test system reliability.

🚨 Scenario 1 — Ingress Failure
Simulating ingress failure when internal services are unreachable or blocked.
![Ingress Failure](https://raw.githubusercontent.com/HananAlghamdi80/sre-kubernetes-azure-project/main/dis.png)


♻️ Scenario 2 — Self-Healing (Pod Restart Recovery)
Deleting a running pod and observing Kubernetes automatically recreate it.
![Self-Healing](https://raw.githubusercontent.com/HananAlghamdi80/sre-kubernetes-azure-project/main/Self-Healing.png)


📈 Scenario 3 — HPA Autoscaling Under Load
Testing Horizontal Pod Autoscaler scaling new replicas when CPU exceeds threshold.
![HPA Autoscaling](https://raw.githubusercontent.com/HananAlghamdi80/sre-kubernetes-azure-project/main/hpa.png)


🔔 Grafana CPU Alert
Alert triggered and sent to Discord when API pod CPU spikes.
![Grafana Alert](https://raw.githubusercontent.com/HananAlghamdi80/sre-kubernetes-azure-project/main/grav.png)


##  Summary This project covers:


Summary
This project covers:


Multi-service Kubernetes deployment


NGINX Ingress with HTTPS & Let's Encrypt


Network Policies (Zero Trust)


Prometheus, Grafana & Alertmanager


Autoscaling (HPA)


Real failure simulation:


Ingress failure


Pod self-healing


High CPU alert & autoscaling

## 🔧 Future Improvements


This project currently uses a self-signed TLS certificate for demonstration purposes, and it will be upgraded to a production-grade certificate (Let’s Encrypt or a dedicated CA) in future iterations.
I also plan to integrate Loki for centralized log aggregation and deeper observability across all services.

With strong hands-on experience in CI/CD automation, I can fully automate the delivery workflow using GitHub Actions, ArgoCD (GitOps), blue-green deployments, automated testing, secure secret handling, and container image scanning. These enhancements will streamline deployments, improve reliability, and bring the platform even closer to production-level standards.

## 🔄 CI/CD Automation (Planned)
I have strong hands-on experience building CI/CD pipelines using GitHub Actions and ArgoCD.  
I can fully automate this platform with:

- OIDC authentication for secure deployments  
- Blue/Green & Canary strategies  
- Automated testing & linting  
- Security scans with Trivy  
- Automatic image builds & pushes to ACR  
- GitOps sync via ArgoCD  

This will make the entire system fully automated, reliable, and production-ready.

I have delivered several advanced DevOps/SRE projects that follow industry best practices, and I’m continuously improving the platform to make it more secure, scalable, and fully automated.



️
