# Code Flow
<div align="center">
  <img alt="Code Flow" src="https://user-images.githubusercontent.com/52092313/162116410-9a1d96fd-a876-4f51-83cd-94a58f844cf5.png">
  A simple web app to create simple programs
</div>

## About the project
Code Flow is a web application powered with amazing features that allows the user create simple programs without the need of programming knowledge:
### Node Editor
Drag and drop nodes and create simple programs!
### Code Visualizer
Transform your diagrams into programs and run them!

## Built With
* [Vue.js](https://vuejs.org/)
* [Drawflow](https://github.com/jerosoler/Drawflow)
* [Go](https://go.dev/)
* [DGraph\Dgo](https://github.com/dgraph-io/dgo)
* [Chi](https://github.com/go-chi/chi)
* [Element+](https://element-plus.org/)

## Getting Started
### Prerequisites
To run the project on your local system, you need the following:
* Node.js
* npm

### Installation
1. Clone the repo
   ```sh
   git clone https://github.com/MateoAltamirano/CodeFlow.git
   ```
2. Install NPM packages
   ```sh
   npm install
   ```

## Usage
### Frontend
1. Navigate to the workspace folder /code-flow
2. Run the following command
   ```sh
   npm run dev
   ```
3. Open http://localhost:3000/
### Backend
1. Navigate to the workspace folder /code-flow-api
2. Run the following command
   ```sh
   docker run --rm -it -p 8080:8080 -p 9080:9080 -p 8000:8000 -v ~/dgraph:/dgraph dgraph/standalone:master
   ```
3. Run the following command
   ```sh
   go run .
   ```

## Contact
Mateo Altamirano - mateo.altamirano.1998@gmail.com
