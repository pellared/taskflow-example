# taskflow-example

>Example usage of [taskflow](https://github.com/pellared/taskflow)

[![Build Status](https://img.shields.io/github/workflow/status/pellared/taskflow-example/build)](https://github.com/pellared/taskflow-example/actions?query=workflow%3Abuild+branch%3Amain)

Running:

```sh
go run ./build
```

Notable files:
- [build/build.go](build/build.go) - taskflow build pipeline,
- [build/build_test.go](build/build_test.go) - sample unit test for a taskflow command,
- [.github/workflows/build.yml](.github/workflows/build.yml) - GitHub Actions workflow,
- [.vscode/launch.json](.vscode/launch.json) - Visual Studio Code configuration for debugging,
- [.vscode/tasks.json](.vscode/tasks.json) - Visual Studio Code task definition for runing taskflow using `Tasks: Run Build Task`.
