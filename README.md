# xproject
X Project

Start the app:
make up / make down - start/stop the app on port 8080 in dev mode with docker-compose
(no hot-reload yet)

Debug:
make debug / make debug-down - start/stop the app in debug mode

to attach from vscode add the following configuration to .vscode/launch.json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Docker debug",
            "type": "go",
            "request": "launch",
            "mode": "remote",
            "remotePath": "/go/src/github.com/pavlov-tony/xproject/cmd/xproject",
            "port": 2345,
            "host": "127.0.0.1",
            "program": "${workspaceRoot}/cmd/xproject",
            "env": {},
            "args": ["-v"],
            "showLog": true,
            "trace": "verbose"
        },
    ]
}

Various commands to be run on the base image:
make lint
make unit-test
make cover