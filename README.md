# smooth-operator
Shared logic and structure library for PDOK operators. 
This repo contains global Kinds and shared code base for other operator.

<p align="center">
    <a href="https://youtu.be/4TYv2PhG89A?si=X_a7zpjaD_esgUAO&t=74">
        <img src="docs/mascotte.png" alt="Smooth-operator" title="Smooth-operator" width="300" />
    </a>
</p>

# Usage
We separate concerns based on context of code.

## Global Kinds
In `api/v1` we store global kinds that are used globally.
For example `OwnerInfo` which is used by WMS, WFS and Atom kind.
These kinds only provide information for other kinds and are not reconciled.

## Shared code
In `pkg` we store the common codebase for operator logic that is reusable/generic.
For example, the reconciling of the ingress route is generic for all kinds.
This makes our operators run smooth together.

## Shared structure
In `model` we store the structs that are reusable/generic.
For example, all operators use the same `OperatorStatus` subresource.

# License

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
