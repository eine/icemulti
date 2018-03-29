# Tools to exploit dynamic reconfiguration in iCE40 FPGAs

This project is based on [IceStorm](http://www.clifford.at/icestorm) ([cliffordwolf/icestorm](https://github.com/cliffordwolf/icestorm)):

> Project IceStorm aims at documenting the bitstream format of Lattice iCE40
FPGAs and providing simple tools for analyzing and creating bitstream files.

> See http://www.clifford.at/icestorm/ for more information.

> The focus of the project is on the iCE40 LP/HX 1K/4K/8K chips. Most of the work was done on HX1K-TQ144 and HX8K-CT256 parts.

The list of available boards with such chips comprises, but is not limited to: [iCEstick Evaluation Kit](http://www.latticesemi.com/icestick), [iCE40-HX8K Breakout Board](http://www.latticesemi.com/Products/DevelopmentBoardsAndKits/iCE40HX8KBreakoutBoard.aspx), [IceZUM Alhambra](https://github.com/FPGAwars/icezum), [icoBOARD 1.0](http://icoboard.org/about-icoboard.html), [Kéfir I iCE40-HX4K](http://fpgalibre.sourceforge.net/Kefir/), [Nandland Go board](https://www.nandland.com/goboard/introduction.html), [eCow-Logic](http://www.ecowlogic.fr/), [myStorm](https://folknologylabs.wordpress.com/2016/07/21/a-perfect-storm/), [BlackIce](https://mystorm.uk/we-forecast-blackice-this-winter-2/), [TinyFPGA](http://tinyfpga.com/)...

Since these devices are SRAM-based, thus volatile, it is common practice to include a flash memory in any board design. This is used to automatically load a configuration on power-up, and SPI is one of the most used interfaces. As a result, it is common in FPGAs devices to find hard IP cores implementing SPI controllers.

Furthermore, according to [TN1248: iCE40 Programming and Configuration](http://www.latticesemi.com/dynamic/view_document.cfm?document_id=46502), the hard SPI cores in the iCE40 devices support not only the *master* mode required to load from flash. In *slave* mode,

>the iCE40 configuration data can be downloaded from an external processor, microcontroller, or DSP processor using the SPI interface.

On top of that, some devices support so-called *Cold Boot* and *Warm Boot* configuration options. This allows to write up to four images/bitstreams at a time to the flash memory, so that any of them can be loaded on reset, without requiring any additional external communication. This is known as *Dynamic Reconfiguration* in the FPGA community.

There is a fifth configuration mode: the one-time programmable NVCM (on-volatile Configuration Memory), which is, naturally, out of the scope of this project.

## Extending dynamic reconfiguration beyond four images/bitstreams

After diving into the existing documentation, and having performed some experiments, the contributors to this project realized that the *cold/warm boot* feature can be extended far beyond the limit of four (in)directly addressable images/bitstreams. E.g. up to 130 images can fit in the 4MB (32Mb) flash included in the iCEstick.
As a result, this projects aims to provide documentation and helper tools to make the best of *dynamic reconfiguration* possibilities with *HX* versions of the *iCE40* FPGA family.

Mimicking latest SoC designs which include programmable logic, such as Xilinx's Zynq or Intel/Altera's Arria/Cyclone, the main orchestrator in the system is expected to be a CPU (either a PC or a microcontroller). This is already true for most of the available boards. Nevertheless, embedded CPUs can be synthetised, in order to have software executed. However, this sets a quite large list of devices to choose from. Although not exclusively, work is focused on the following:

- USB-TTL adapter: FTDI, CH340, PL2303...
  - USB-SPI
  - USB-UART
- External uC: AVR Atmega/Attiny, ARM CortexM...
- Embedded uC: Lattuino (AVR), VexRiscv (RISC-V)...

Depending on the design of the PCB(s), multiple connection schemes might be required in order to achieve the same functional result. Indeed, some custom HDL modules should be designed. What's more, some systems might require a not negligible stop-the-world approach.

The ideal outcome would be a HDL bootloader, small enough to provide extended and (optionally) compressed *dynamic reconfiguration*, while enough space is left for user apps. It would provide an illusion of *partial reconfiguration*.

## Features

NOTE: to avoid mixing terms, *image* is used to reference the bitstream corresponding to a single design, and *pack* relates to multiple images packed in a single bitstream.

- Tools are released either as a go(lang) library, a CLI application and/or a web application.
- Manage, reorder, edit, merge, split, (un)compess... images and packs, up to the size of the target flash memory.
- Manage metadata of images to uniquely identify these and the implemented designs.
- Generate, edit, merge, split... memory maps from formatted packs of images.
- Select the connection scheme that best suits your needs.
- Given the connection scheme, generate test/example routines for orchestration CPUs.
- If the target CPU is connected to the PC (or if it is the PC itself), test the generated routines in the web app.

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

NOTE: it is out of the scope of this project how to generate bitstream images from HDL sources. Therefore, several bitstreams of individual designs are expected to be previously generated. Alternatively, some are provided in subdir `test/bitstreams`.

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
