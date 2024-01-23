# UniBlox Ecommerce Backend

This project implements the backend functionality for a simple ecommerce store as part of the UniBlox assignment.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Technology Stack](#technology-stack)
- [Getting Started](#getting-started)
- [API Documentation](#api-documentation)

## Introduction

The UniBlox Ecommerce Backend is designed to simulate the backend processes of an ecommerce store. It includes functionalities such as adding items to the cart, processing checkouts, generating discount codes, and providing purchase-related information.

## Features

- Add items to the cart
- Checkout functionality with discount code validation
- Admin APIs for generating discount codes and retrieving purchase information
- Simulated in-memory store

## Technology Stack

- Golang
- gin-gonic

## Getting Started

**Clone the Repository:**

```bash
git clone https://github.com/harshvsinghme/uniblox-assmt-backend.git
```

Get .env file and save that at the same directory level of the main.go file.

And, for this project .env file can be empty

```bash
go mod tidy
```

### Run Server

```bash
go run main.go
```

## Documentation

#### https://documenter.getpostman.com/view/16741530/2s9YymHkPj
