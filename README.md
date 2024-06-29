<h1 align="center">
  <img alt="ignore" src="https://github.com/neptship/ignore/blob/main/assets/ignoregithub.jpg">
</h1>

<p align="center">
  <a href="https://pkg.go.dev/github.com/neptship/ignore">
    <img alt="Reference" src="https://pkg.go.dev/badge/github.com/neptship/ignore.svg"
  /></a>
  <a href="https://goreportcard.com/report/github.com/neptship/ignore">
    <img alt="goreportcard" src="https://goreportcard.com/badge/github.com/neptship/ignore"
  /></a>
  <a href="https://github.com/neptship/ignore/actions/workflows/test.yml">
    <img alt="test workflow" src="https://github.com/neptship/ignore/actions/workflows/test.yml/badge.svg"
  /></a>
  <a href="https://github.com/neptship/ignore/blob/main/LICENSE">
    <img alt="license" src="https://img.shields.io/github/license/neptship/ignore"
  /></a>
  <a href="https://github.com/neptship/ignore/releases">
    <img alt="latest release" src="https://img.shields.io/github/release/neptship/ignore.svg"
  /></a>
</p>
<h2>Features</h2>

- Fast
- Cross-Platform - Linux, macOS, Windows
- Easy installation
- More than 600 templates for different technologies
- Friendly error messages in case something goes wrong

<h2>Installation</h2>

<h3>Golang (Windows, Linux, MacOS)</h3>

Install using [Golang Packages](https://pkg.go.dev/github.com/neptship/ignore)

```shell
go install github.com/neptship/ignore@latest
```

This script will automatically detect OS & Distro and use the best option available.

<h3> From source </h3>

Clone the repo
```shell
git clone https://github.com/neptship/ignore.git
cd ignore
```

GNU Make **(Recommended)**
```shell
make setup # if you want to compile and install ignore cli to path
```

<details>
<summary>If you don't have GNU Make use this</summary>


```shell
# To build
go build

# To install
go install
```

</details>

<h2>Usage</h2>

![Usage](https://github.com/neptship/ignore/blob/main/assets/ignore.gif?raw=true)

<h3>Other</h3>

See `ignore help` for more information

<details>
<summary>Commands</summary>

| Name         | Description                           |
|--------------|---------------------------------------|
| create       | create .ignore file                   |
| add          | add a template to .ignore file        |
| list         | available templates for .ignore files |
</details>

<h2> Built With </h2>

* [Cobra](https://cobra.dev/) - The modern CLI framework used

<h2> Contributing </h2>

Please read [CONTRIBUTING.md](https://github.com/neptship/ignore/blob/main/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

<h2> Authors </h2>

* **Mikhail Chikankov** - *Creator project* - [neptship](https://github.com/neptship)


<h2>License</h2>

Sample and its code provided under MIT license, please see [LICENSE](/LICENSE). All third-party source code provided
under their own respective and MIT-compatible Open Source licenses.

Copyright (C) 2023, Mikhail Chikankov


<h2> Acknowledgments </h2>

* Hat tip to anyone whose code was used
* Inspiration
* etc

[![Stargazers repo roster for @neptship/ignore](https://reporoster.com/stars/neptship/ignore)](https://github.com/neptship/ignore/stargazers)

[![Forkers repo roster for @neptship/ignore](https://reporoster.com/forks/neptship/ignore)](https://github.com/neptship/ignore/network/members)
