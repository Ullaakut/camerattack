# camerattack
An attack tool designed to remotely shut down CCTV cameras

## Current state of the repo
I am currently researching ways to do that. The golang file you can see was just a quick test tool I made to try to spam the RTSP servers with requests, send them random data on purpose, etc. As of right now, I didn't find any security leak that could be exploited on all camera models. I want this tool to be able to remotely close any RTSP stream. Maybe session bruteforcing could be the only answer to that problem, but it would take a very long time and hopefully there are simpler ways.

## Solutions

### Session hijacking + TEARDOWN request

RTSP session identifiers are generated like such:

> Session identifiers are opaque strings of arbitrary length. Linear white space must be URL-escaped. A session identifier MUST be chosen randomly and MUST be at least eight octets long to make guessing it more difficult. (See Section 16.)
> `session-id   =   1*( ALPHA | DIGIT | safe )`

**Example of generated session ID:**  `HtMdQMfjnH3DgnPr`

### Requesting RTSP paths with `..` path component

The RTSP RFC specifies that this needs to be handled by the RTSP server, but it's worth a try. The RFC also states that clients having a persistently suspicious behavior should be ignored in order not to let them guess the credentials of the camera, but as [cameradar](https://github.com/EtixLabs/cameradar/) proved it, it seems that none of the cameras nor RTSP servers implement this kind of security. This gives me the hope that there might be other security risks that were ignored by RTSP server implementers.

### Concentrated denial of service attack

Not related, but once a camera has been accessed with Cameradar, it can also be used to initiate traffic flows to one or more IP addresses by specifying them as the destination in SETUP requests. Once again, the RFC specifies that this can be stopped by not letting an IP initiate a connection to other clients, but I don't think that many RTSP servers would bother.

## Conclusion

From my little research, it seems that most RTSP servers are assumed to be inaccessible from the outside and thus their security is almost never considered. In many cases, simply connecting a machine in Ethernet in a building gives access to the camera's network, that can be probed easily using [cameradar](https://github.com/EtixLabs/cameradar/). Sometimes they're even accessible through WiFi, when the people who installed them really have no idea of what they're doing.