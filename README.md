<p align="center">
    <a href="https://youtu.be/4TYv2PhG89A?si=X_a7zpjaD_esgUAO&t=74">
        <img src="docs/mascotte.png" alt="Smooth-operator" title="Smooth-operator" width="300" />
    </a>
</p>

# Smooth Operator
This repository contains code shared by the Kubernetes operators developed by [PDOK](https://www.pdok.nl/). There are no executable programs present in this codebase. 
The goal of this repository is to support the development work of PDOK, but it can also be used as a starter to develop custom map related operators. At the moment the following operators have this repository as a basis: 
- [Atom operator](https://github.com/PDOK/atom-operator), an operator for deploying atomfeed instances
- [Mapserver operator](https://github.com/PDOK/mapserver-operator), an operator for deploying mapserver instances

Although this repository is oriented towards geoinformation related operators, it can be used as a base for non-geoinformation operators.

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

# License

```
MIT License

Copyright (c) 2025 Publieke Dienstverlening op de Kaart

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
