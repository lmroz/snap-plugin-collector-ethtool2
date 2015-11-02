# IXGBE Collector Plugin

## Features
### Current version
This plugin allows data collection from Intel's NICs. In current version metrics exposed to `ethtool -S` are available.

### Future
Parsing RAW data from NIC's registry dump.

## Usage
This plugin uses *ethtool* to read data from NIC. Because *ethtool* is core Linux utility no further actions are required to start using this plugin.
Metrics are available under name space `/intel/net/[IFNAME]/[METRIC]` where `[IFNAME]` is interface name (ex. `eth0)` and `[METRIC]` is metric of interest (ex. `rx_packets`).
