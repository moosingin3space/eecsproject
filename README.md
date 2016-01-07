# eecsproject

`eecsproject` is a simple project generator for University of Michigan
intro EECS courses. It will generate a skeleton project with a Vagrantfile and
a template Makefile for either C or C++, based on parameters passed in.

To download, go to the [releases](https://github.com/moosingin3space/eecsproject/releases).
The *snapshot* release is built from the latest git master, and the other
versions are tagged releases.

## Installation

- Install [VirtualBox](https://www.virtualbox.org)
- Install [Vagrant](https://vagrantup.com)
- Download a [release](https://github.com/moosingin3space/eecsproject/releases)
  and copy it into a directory on your `PATH`.

## Usage

- To generate a C project: `eecsproject project-name`
- To generate a C++ project: `eecsproject -cpp project-name`
- To run the CAEN-compatible VM: `vagrant up`, then `vagrant ssh`.
  You will then be presented with a Linux shell.

## List of software on VM

- `gcc`
- `g++`
- `gdb`
- `valgrind`
