# John Crickett Coding Challenges

Late to the party (weekend build) but we roll. Implementing [codingchallenges.fyi](https://codingchallenges.fyi/challenges/intro) in Go.

**Note: Doing this in addition to the day work to practice Go pretty much. Okay, maybe not. UNIX design principles are top-notch, so can learn a thing or two.**

Read: [The Art of Unix Programming](http://www.catb.org/~esr/writings/taoup/html/)

## Challenges

### Building My Own Unix's wc
- Implemented a version of Unix's `wc` command in Go.
- Optimized for performance and error handling.

### JSON Parser with Array Support
- Added functionality to a JSON parser to fully support JSON arrays.
- Handled array elements and nested arrays.

### Huffman Coding Compression Tool
- Developed a CLI tool for lossless data compression using Priority Queue (Min Heap) and Depth-First Search (DFS) algorithms.
- Implemented compression and decompression functions.

### CLI Cut Tool
- Implemented efficient field extraction using Go's string manipulation capabilities.
- Utilized flag package to parse command-line arguments and provide flexible field selection options.

### Load Balancer
- Load balancer distributes requests using a round-robin algorithm to evenly distribute load among active servers.
- Health check mechanism periodically checks the health status of backend servers and removes unhealthy servers from the pool to prevent service disruptions.
