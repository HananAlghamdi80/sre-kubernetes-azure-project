
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


️
