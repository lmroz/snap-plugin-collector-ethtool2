# snap collector plugin - ethtool

This plugin uses ethtool to gather interface statistics. 																						Current version exposes stats available using ethtool given by below commands:
* `ethtool -S`, interface statistics
* `ethtool -d`, register dump
* `ethtool -m`, digital optical monitoring

It's used in the [snap framework](http://github.com:intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license-and-authors)
6. [Acknowledgements](#acknowledgements)

## Getting Started

The plugin is ready to use out of the box. You don't have to perform any configuration or installation steps.

### System Requirements

* ethtool available under `$PATH` or `/sbin/`
* [golang 1.5+](https://golang.org/dl/)
* Root privileges are required 

### Operating systems
All OSs currently supported by plugin:
* Linux/amd64

### Installation
#### Download ethtool plugin binary:
You can get the pre-built binaries for your OS and architecture at snap's [GitHub Releases](https://github.com/intelsdi-x/snap/releases) page.

#### To build the plugin binary:
Fork https://github.com/intelsdi-x/snap-plugin-collector-ethtool  
Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-collector-ethtool.git
```

Build the plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `/build/rootfs/`

### Configuration and Usage
* Set up the [snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Ensure `$SNAP_PATH` is exported  
`export SNAP_PATH=$GOPATH/src/github.com/intelsdi-x/snap/build`


## Documentation

You can learn about some of exposed metrics [here](https://www.myricom.com/software/myri10ge/397-could-you-explain-the-meanings-of-the-myri10ge-counters-reported-in-the-output-of-ethtool.html)

### Collected Metrics
This plugin allows to collect interface network statistics such like received|transmitted bytes|packets and more.                                                                                                    
List of metrics for each device is dependent on it's driver.

This plugin has the ability to gather the following metrics (driver specific):
* [for driver E1000E](METRICS_E1000E.md)
* [for driver IXGBE](METRICS_IXGBE.md)
* [for driver FM10K](METRICS_FM10K.md)
* [for driver TG3](METRICS_TG3.md)


A few drivers such as IXGBE support exposing optical transceivers (SFP, SFP+, or XFP) information too. The information is known as [digital optical monitoring (DOM)] (METRICS_DOM.md#digital-optical-monitoring) information.

Metrics are available in namespace: 
*	`/intel/net/<driver name>/<device name>/nic/<metric name>` (from cmd `ethtool -S`, interface statistics)
*	`/intel/net/<driver name>/<device name>/reg/<metric name>` (from cmd `ethtool -d`, register dump)
*	`/intel/net/<driver name>/<device name>/dom/<metric name>` (from cmd `ethtool -m`, digital optical monitoring)


### Roadmap
As we launch this plugin, we have a few items in mind for the next release:

- [ x ] Metrics from device's registry dump when using IXGBE driver
- [ x ] Expose pluggable optics (SFP & SFP+) information

If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-ethtool/issues).

## Community Support
This repository is one of **many** plugins in **snap**, a powerful telemetry framework. The full project is at http://github.com:intelsdi-x/snap.
To reach out on other use cases, visit:
* [Snap Gitter channel] (https://gitter.im/intelsdi-x/snap)

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[snap](http://github.com:intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements

* Author: 	[Lukasz Mroz](https://github.com/lmroz)
* Co-author:[Izabella Raulin](https://github.com/IzabellaRaulin)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.us.
