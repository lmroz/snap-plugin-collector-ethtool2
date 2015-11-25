# Snap Collector Plugin - Ethtool

This plugin uses ethtool to gather interface statistics. Current version exposes stats available using `ethtool -S` command.

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](configuration-and-usage)
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

- Linux system
- ethtool available under `$PATH` or `/sbin/`

### Configuration and Usage
####Tips:
-Adding more metrics to monitor per NIC is very cheap in terms of CPU time.


## Documentation

You can learn about some of exposed metrics [here](https://www.myricom.com/software/myri10ge/397-could-you-explain-the-meanings-of-the-myri10ge-counters-reported-in-the-output-of-ethtool.html)

### Collected Metrics
List of metrics for each device is dependent on it's driver.
Metrics are available in namespace `/intel/net/<device name>/<metric name>` and all are `int`'s.

### Roadmap
As we launch this plugin, we have a few items in mind for the next release:

* Metrics from device's registry dump when using IXGBE driver

## Community Support
This repository is one of **many** plugins in the **Snap Framework**: a powerful telemetry agent framework. To reach out on other use cases, visit:

* Snap Gitter channel (@TODO Link)
* Our Google Group (@TODO Link)

The full project is at http://github.com:intelsdi-x/snap.

## Contributing
We love contributions! :heart_eyes:

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
Snap, along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
List authors, co-authors and anyone you'd like to mention

* Author: [Lukasz Mroz](https://github.com/lmroz)

**Thank you!** Your contribution is incredibly important to us.
