# GoDashery

An extensible WebSocket driven dashboard with Golang backend and web frontend.

The frontend uses TypeScript and SASS, but not a whole lot else.

The system is built to be fast and easily extensible. It's meant to be used for various purposes, on even low performance devices, e.g. Raspberry Pi.


## Building frontend

```
cd frontend
npm install -g typescript gulp
gulp build
```

When developing you'll want `gulp` to watch for changes, which it does on the default task. Simply run `gulp` with no arguments.


## Running backend

```
cd cmd/godashery
go build
./godashery
```

The backend will serve the built frontend files if you don't want to set that up with your own web server. Simply open up `http://localhost:8080`.


## Adding widgets

Add your server-side code to `widgets/mywidget.go`. Frontend logic goes to `frontend/src/widgets/mywidget.ts`
 
Check existing ones for examples.
