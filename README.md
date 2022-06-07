<div align="center" id="top"> 
  <img src="go.jpeg" alt="Todos Api" />
</div>
<h1 align="center">Todos List API</h1>
<p align="center">
  <img alt="Tool" src="https://badges.aleen42.com/src/docker.svg">
  <img alt="Tool" src="https://badges.aleen42.com/src/golang.svg">
  <img alt="Github top language" src="https://img.shields.io/github/languages/top/devenes/todos-api?color=56BEB8">
  <img alt="Github language count" src="https://img.shields.io/github/languages/count/devenes/todos-api?color=purple">
  <img alt="Repository size" src="https://img.shields.io/github/repo-size/devenes/todos-api?color=orange">
  <img alt="License" src="https://img.shields.io/github/license/devenes/todos-api?color=yellow">
  <!-- <img alt="Github issues" src="https://img.shields.io/github/issues/devenes/todos-api?color=56BEB8" /> -->
  <!-- <img alt="Github forks" src="https://img.shields.io/github/forks/devenes/todos-api?color=56BEB8" /> -->
  <!-- <img alt="Github stars" src="https://img.shields.io/github/stars/devenes/todos-api?color=56BEB8" /> -->
</p>

<p align="center">
  <a href="#dart-about">About</a> &#xa0; | &#xa0; 
  <a href="#sparkles-features">Features</a> &#xa0; | &#xa0;
  <a href="#rocket-technologies">Technologies</a> &#xa0; | &#xa0;
  <a href="#white_check_mark-requirements">Requirements</a> &#xa0; | &#xa0;
  <a href="#checkered_flag-starting">Starting</a> &#xa0; | &#xa0;
  <a href="#memo-license">License</a> &#xa0; | &#xa0;
  <a href="https://github.com/devenes" target="_blank">Author</a>
</p>

<br>

## :dart: About ##

This API is a simple todo list. It is a RESTful API that allows you to create, read, update and delete todos. 

## :sparkles: Features ##

:heavy_check_mark: List todos\
:heavy_check_mark: Update todos

## :rocket: Technologies ##

The following tools were used in this project:

- [Go](https://golang.org/): The language used in this project.

## :white_check_mark: Requirements ##

Before starting :checkered_flag:, you need to have [Git](https://git-scm.com) and [Go](https://golang.org) installed.

## :checkered_flag: Starting ##

- Clone the project
```bash
git clone https://github.com/devenes/todos-api
```

- Go to the project directory:
```bash
cd todos-api
```

- Run the following command to start the server:
```bash
go run main.go
```

- List todos with the following command:
```bash
curl http://localhost:8080/todos
```

- Get the first todo with the following command:
```bash
curl http://localhost:8080/todos/1
```

- Create a new todo with the following command:
```bash
curl -X POST -H 'content-type: application/json' --data '{"id": "4", "name": "Buy milk"}' http://localhost:8080/todos
```

- Update the first todo with the following command:
```bash
curl -X PUT -H 'content-type: application/json' --data '{"id": "1", "name": "Check the mailbox"}' http://localhost:8080/todos
```

- Get the second todo with the following command:
```bash
curl http://localhost:8080/todos/2
```

- The server will initialize in the `http://localhost:8080` URL.

## :memo: License ##

This project is under license from Apache 2.0. For more details, see the [LICENSE](LICENSE) file.


Made with :heart: by <a href="https://github.com/devenes" target="_blank">devenes</a>

&#xa0;

<a href="#top">Back to top</a>
