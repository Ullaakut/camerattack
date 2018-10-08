# camerattack

<p align="center">
    <img src="https://raw.githubusercontent.com/Ullaakut/camerattack/master/images/camerattack.gif" width="100%"/>
</p>

<p align="center">
    <a href="#license">
        <img src="https://img.shields.io/badge/license-Apache-blue.svg?style=flat" />
    </a>
    <a href="https://golangci.com/r/github.com/Ullaakut/camerattack">
        <img src="https://golangci.com/badges/github.com/Ullaakut/camerattack.svg" />
    </a>
    <a href="https://goreportcard.com/report/github.com/Ullaakut/camerattack">
        <img src="https://goreportcard.com/badge/github.com/Ullaakut/camerattack" />
    </a>
    <a href="https://github.com/Ullaakut/camerattack/releases/latest">
        <img src="https://img.shields.io/github/release/Ullaakut/camerattack.svg?style=flat" />
    </a>
    <a href="https://godoc.org/github.com/Ullaakut/camerattack">
        <img src="https://godoc.org/github.com/Ullaakut/camerattack?status.svg" />
    </a>
</p>

## An attack tool designed to remotely shut down CCTV cameras.

<p align="center"><img src="images/Camerattack.png" width="250"/></p>

## Usage

Simply run it and provide it the RTSP stream URL as a parameter. Depending on your network speed and the camera you're trying to shutdown, it might work sooner or later, or might not work.

## How does it work?

It simply does lots of requests really quickly on the device, hoping that it will crash because of a memory leak or some mistakes in its firmware's code.
