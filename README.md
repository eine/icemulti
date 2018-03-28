# [icemulti] Go(lang) tools to exploit warm/cold boot in iCE40 FPGAs

## Repository (directory) organization

- `api`: go(lang) sources of the AJAX service and the file server (backend).
- `app`: HTML + CSS + JavaScript sources of a web-based GUI to provide a responsive and customizable layout for visualization (frontend).
- `assets`: common logos, images...
- `build`: development and release resources (mostly shell scripts).
- `cli`: go(lang) sources of the CLI tool which provides a clean and documented access to functions/procedures implemented in `lib`.
- `doc`: markdown sources of the documentation (a [Hugo](https://github.com/gohugoio/hugo) site).
- `lib`: go(lang) sources that are used by `api` and `cli`, and are expected to be used as a library by external projects.
- `test`: example files in multiple formats.

## Development

Script `build/init_devenv.sh` ensures that all the dependencies of `api`, `app`, `cli` and `lib` are up to date. However, a few dependencies are required to execute it, and to later work with the sources: yarn (nodejs), golang, git...

In order to make development easier, a docker image is provided. Script `build/docker_devenv.sh sh` runs a temporal container, executes `init_devenv.sh` and provides a shell.

Note that, by default, port `8080` inside the container is bind to `5555` in the host. As a result, `localhost:5555` should be browsed/used when the `app` and/or `api` are being tested, even though the services are configured to `8080` inside the container.

Proceed to the READMEs in each of the subdirs for further info.

## Releases

There are multiple options to release icemulti.

- icemulti-slim: `cli`; command-line interface (CLI) tool to (un)pack or reconfigure.
- icemulti-app: `api` [+ `app`]; service running in the background that accepts AJAX requests, and optionally serves a local directory.
- icemulti: `cli` + `api` [+ `app`]: 'icemulti-slim' and 'icemulti-service' in a single binary.

While 'icemulti-slim' is distributed as a single binary file, 'icemulti-app' and 'icemulti' release files are tarballs. This is because `app` (frontend) sources are given uncompressed in the same tarball, and its content is expected to be served by `api` (backend). To disable this behaviour, use flag `--nofs | -n` or set the directory to `none` with `--dir none | -d none`. The result will be an AJAX-only service.

Moreover, static files, such as the sources of `app`, can be embedded in a go binary file (see [gobuffalo/packr](https://github.com/gobuffalo/packr)). Therefore, 'icemulti-pack' and 'icemulti-app-pack' variants can be produced. These provide the same functionality, but a single binary file is distributed, instead of a tarball.

---

In order to generate release tarballs, first set the devenv (either `build/docker_devenv.sh sh` -docker- or `build/init_devenv.sh` -local-). Then, execute `build/release.sh`, which will build all the pieces and generate tarballs in the root dir. Last, and optionally, execute `build/release_img.sh` to generate docker images from scratch.

Note: `release_img.sh` must be executed in the host, not inside a container.

## Refs

- [Concise guide to golang/dep ](https://gist.github.com/subfuzion/12342599e26f5094e4e2d08e9d4ad50d)
