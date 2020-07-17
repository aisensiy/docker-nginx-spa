[![docker pull openbayes/openbayes-nginx-spa][image shield]][docker hub]

This is a Docker image used to serve a Single Page App (pure frontend javascript) using nginx, it support PushState, and includes a way to pass configuration at run time.

## Supported tags and `Dockerfile` links

-	[`latest` (*Dockerfile*)][latest]

## Included on top of [base][base image] nginx image

- [pushState][push state] support. Every request is routed to `/app/index.html`. Useful for the clean urls (no `!#`)
- [ENV-based Config](#env-config)

# App Setup

This docker image is built for `index.html` file being in the `/app` directory. `pushState` is enabled.

At a minimum, you will want this in your `Dockerfile`:

```Dockerfile
FROM openbayes/openbayes-nginx-spa

COPY build/ /app
COPY index.html /app/index.html
```

Then you can build & run your app in the docker container. It will be served by a nginx static server.

```bash
$ docker build -t your-app-image .
$ docker run \
  -e PREFIX_API_KEY=yourkey \
  -e PREFIX_API_URL=http://myapi.example.com \
  -e CONFIG_VARS=PREFIX_ \
  -p 8000:80 your-app-image
```

You can then go to `http://docker-ip:8000/` to see it in action.

## Env Config

Included is ability to pass `runtime` environmental variables to your app.

This is very useful in case your API is on a different domain, or if you want to configure central error logging. 

```bash
$ docker run \
  -e PREFIX_RAVEN_DSN=yourkey \
  -e PREFIX_API_URL=http://myapi.example.com \
  -e CONFIG_VARS=PREFIX_ -p 8000:80 socialengine/nginx-spa:latest

 ==> Writing /app/config.js with { 
   "RAVEN_DSN": "yourkey", 
   "API_URL": "http://myapi.example.com"
 }
```

All the env variables start with `CONFIG_VARS` will be filtered into the `config.js` file. And the `CONFIG_VARS` prefix will be removed. The `config.js` file, which you can then add to your index.html, or load asynchronously. The path can be controlled with `CONFIG_FILE_PATH` environmental variable.

### Config Override

> This is supported in `openbayea/openbayes-nginx-spa:v2`

If you wanna have a `config.js` file with default configs in it. The `config.js` should fulfill the following requirements:

1. The `json` part must be a valid `json format` thing without any `javascript` part. There is a `config.example.js` file in this repository as an example.
2. The `config.js` file should only have a format like `window.__env = {}` do not add more code after the end of `}`.
3. Read the code `main.go` if you wanna know all the detail.

[push state]: https://developer.mozilla.org/en-US/docs/Web/API/History_API
[latest]: https://github.com/aisensiy/docker-nginx-spa/blob/master/Dockerfile
[base image]: https://github.com/nginxinc/docker-nginx
[image shield]: https://img.shields.io/badge/dockerhub-openbayes%2Fopenbayes--nginx--spa-blue.svg
[docker hub]: https://registry.hub.docker.com/u/openbayes/openbayes-nginx-spa/
