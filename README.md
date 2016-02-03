# A Go package for generating and manipulating UUIDs

Generate, encode, and decode UUIDs v1, as defined in [RFC 4122](http://www.ietf.org/rfc/rfc4122.txt), in [Go](http://golang.org).

## Maturity

[![Build Status](https://travis-ci.org/agext/uuid.svg?branch=master)](https://travis-ci.org/agext/uuid)

v1.0 Stable: Guaranteed no breaking changes to the API in future v1.x releases. No known bugs or performance issues. Probably safe to use in production, though provided on "AS IS" basis.

## Overview

[![GoDoc](https://godoc.org/github.com/agext/uuid?status.png)](https://godoc.org/github.com/agext/uuid)

Package uuid implements generation and manipulation of UUIDs (v1 defined in RFC 4122).

Version 1 UUIDs are time-based and include a node identifier that can be a MAC address or a random 48-bit value.

This package uses the random approach for the node identifier, setting both the 'multicast' and 'local' bits to make sure the value cannot be confused with a real IEEE 802 address (see section 4.5 of RFC 4122). The initial node identifier is a cryptographic-quality random 46-bit value. The first 30 bits can be set and retrieved with the `SetNodeId` and `NodeId` functions and method, so that they can be used as a hard-coded instance id. The remaining 16 bits are reserved for increasing the randomness of the UUIDs and to avoid collisions on clock sequence rollovers.

The basic generator `New` increments the clock sequence on every call and when the counter rolls over the last 16 bits of the node identifier are regenerated using a PRNG seeded at init()-time with the initial node identifier. This approach sacrifices cryptographic quality for speed and for avoiding depletion of the OS entropy pool (yes, it can and does happen).

The `NewCrypto` generator replaces the clock sequence and last 16 bits of the node identifier on each call with cryptographic-quality random values.

## Installation

```
go get github.com/agext/uuid
```

## License

Package log is released under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.
