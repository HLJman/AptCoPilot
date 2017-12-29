# AptCopilot

This project contains two parts, a server written in [Go](https://golang.org) and frontend web app written in [Polymer](https://www.polymer-project.org). 

The frontend app is contained in the `assets` directory with the custom components in the `src` folder. Routing is client-side using the Polymer element [app-route](https://www.polymer-project.org/blog/routing).

The server is in the root of the project and serves both the web app and REST api. The REST api can be found under path `/api`.

## Dependencies
- mysql database

## Run

- With data

        ./dev/run.sh

- UI only
        
        cd assets
        polymer serve --open

## Publish to AWS

        make publish

## SSH to AWS Server

        make ssh