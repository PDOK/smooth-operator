<p align="center">
    <a href="https://youtu.be/4TYv2PhG89A?si=X_a7zpjaD_esgUAO&t=74">
        <img src="docs/mascotte.png" alt="Smooth-operator" title="Smooth-operator" width="300" />
    </a>
</p>

[![Build](https://github.com/PDOK/smooth-operator/actions/workflows/test.yml/badge.svg)](https://github.com/PDOK/smooth-operator/actions/workflows/test.yml)
[![Lint](https://github.com/PDOK/smooth-operator/actions/workflows/lint.yml/badge.svg)](https://github.com/PDOK/smooth-operator/actions/workflows/lint-go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/PDOK/smooth-operator)](https://goreportcard.com/report/github.com/PDOK/smooth-operator)
[![GitHub license](https://img.shields.io/github/license/PDOK/smooth-operator)](https://github.com/PDOK/smooth-operator/blob/master/LICENSE)

# Smooth Operator Library
This repository contains code shared by the Kubernetes operators developed by [PDOK](https://www.pdok.nl/). There are no executable programs present in this codebase. 
The goal of this repository is to support the development work of PDOK, but it can also be used as a starter to develop custom map related operators. At the moment the following operators have this repository as a basis: 
- [Atom operator](https://github.com/PDOK/atom-operator), an operator for deploying atomfeed instances
- [Mapserver operator](https://github.com/PDOK/mapserver-operator), an operator for deploying mapserver instances

# Usage
This project does not contain any executable files and is only used as a basis for Kubernetes operators.

# Project structure
The code is separated based on context of code.

### Global Kinds
Directory `api` contains Kinds that are used globally, for example `OwnerInfo` which is used by WMS, WFS and Atom kind.
These kinds only provide information for other kinds.

### Shared code
Directory `pkg` contains the common codebase for operator logic that is reusable or generic.

### Shared structure
Directory `model` contains the structs that are reusable or generic.

# Contributing

### How to contribute
Smooth-operator is solely developed by PDOK. Contributions are however always welcome. If you have any questions or suggestions you can create an issue in the issue tracker.

### Contact
The maintainers can be contacted through the issue tracker.

# Authors
This project is developed by [PDOK](https://www.pdok.nl/), a platform for publication of geographic datasets of Dutch governmental institutions.
