# Todo App Infrastructure Practice Project

## Overview

This project is designed to practice building and deploying a small production-like application using modern backend, frontend, and infrastructure technologies.

The application itself is a simple Todo App, but the primary focus is on the architecture, deployment pipeline, and infrastructure setup rather than application complexity.

## Goals

### Infrastructure

* Configure an **Nginx Load Balancer** to distribute incoming requests across multiple API servers.
* Deploy and manage **three Go API instances** behind the load balancer.
* Simulate a highly available backend architecture.

### Backend

* Build a REST API in **Go**.
* Follow **Clean Architecture** principles:

  * Separation of domain, application, infrastructure, and delivery layers.
  * Dependency inversion.
  * Testable and maintainable codebase.
* Implement CRUD operations for Todo items.

### Frontend

* Build a **Client-Side Rendered (CSR)** web application.
* Consume backend APIs through the load balancer.
* Provide a simple interface for:

  * Creating todos
  * Viewing todos
  * Updating todos
  * Deleting todos

### CI/CD

* Set up an automated CI/CD pipeline.
* Automatically:

  * Build the frontend application.
  * Upload static assets to a CDN.
  * Deploy backend services when changes are merged.

## Architecture

```text
                    ┌───────────────┐
                    │     User      │
                    └───────┬───────┘
                            │
                            ▼
                   ┌─────────────────┐
                   │ Nginx Load      │
                   │ Balancer        │
                   └───────┬─────────┘
                           │
         ┌─────────────────┼─────────────────┐
         ▼                 ▼                 ▼
 ┌─────────────┐   ┌─────────────┐   ┌─────────────┐
 │ API Server 1│   │ API Server 2│   │ API Server 3│
 │     Go      │   │     Go      │   │     Go      │
 └──────┬──────┘   └──────┬──────┘   └──────┬──────┘
        └─────────────────┼─────────────────┘
                          ▼
                    ┌─────────┐
                    │Database │
                    └─────────┘


Frontend (CSR)
       │
       ▼
      CDN
```

## Features

### Todo Management

* Create a todo
* List all todos
* Mark a todo as completed
* Update a todo
* Delete a todo

### Load Balancing

* Round-robin request distribution using Nginx
* Ability to verify requests are served by different API instances

### Automated Deployment

* Frontend build and upload to CDN
* Backend deployment pipeline
* Automated checks before deployment

## Tech Stack

### Backend

* Go
* Clean Architecture
* REST API

### Frontend

* React
* TypeScript

### Infrastructure

* Nginx
* CDN
* CI/CD Pipeline

## Learning Objectives

By completing this project, you will gain practical experience with:

* Load balancing using Nginx
* Designing maintainable Go applications
* Clean Architecture principles
* Frontend deployment to a CDN
* CI/CD automation
* Multi-instance backend deployments
* End-to-end application delivery

## Expected Outcome

A fully functional Todo application where:

* Requests are distributed across three Go API servers by Nginx.
* The frontend is served from a CDN.
* Changes are automatically deployed through CI/CD pipelines.
* The codebase follows clean architecture principles and production-oriented practices.
