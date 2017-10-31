# camerattack

<p align="center"><img src="https://raw.githubusercontent.com/Ullaakut/camerattack/master/images/camerattack.gif" width="100%"/></p>

[![Latest release](https://img.shields.io/github/release/Ullaakut/camerattack.svg?style=flat)](https://github.com/Ullaakut/camerattack/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/Ullaakut/camerattack)](https://goreportcard.com/report/github.com/Ullaakut/camerattack)

## An attack tool designed to remotely shut down CCTV cameras.

## Usage

Simply run it and provide it the RTSP stream URL as a parameter. Depending on your network speed and the camera you're trying to shutdown, it might work sooner or later, or might not work.

## How does it work?

It simply does lots of requests really quickly on the device, hoping that it will crash because of a memory leak or some mistakes in its firmware's code.
