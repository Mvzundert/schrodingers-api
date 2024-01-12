# Schrodinger's API

## Introduction
This guide explains how to set up and run the Schrodinger's API, which interfaces a Go server with Rockstar scripts via the `chirp` interpreter.

## Prerequisites
- Git
- Go environment
- Compiled `chirp` binary (separate setup instructions provided)

## Setup
1. **Clone the Repository:**
   ```
   git clone https://github.com/Mvzundert/schrodingers-api
   cd schrodingers-api
   ```

2. **Place the `chirp` Binary:**
   - Copy the compiled `chirp` binary to the `bin` directory in the project.

3. **Run the Server:**
   ```
   go run server.go
   ```

4. **Access the API:**
   - The API is now accessible. For example, use `curl http://localhost:8080/ping` to test.
