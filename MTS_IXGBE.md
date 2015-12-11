# snap collector plugin - ethtool		
		
## Metrics exposed via driver ixgbe:		
		

**a) NIC statistics**

Namespace: `/intel/net/ixgbe/<device name>/nic/<metric name>`
		
Metric Name | 		
------------			
|alloc_rx_buff_failed                          	|	
|alloc_rx_page_failed                          	|	
|broadcast                                     	|	
|collisions                                    	|	
|fcoe_bad_fccrc                                	|	
|fcoe_noddp                                    	|	
|fcoe_noddp_ext_buff                           	|	
|fdir_match                                    	|	
|fdir_miss                                     	|	
|fdir_overflow                                 	|	
|hw_rsc_aggregated                             	|	
|hw_rsc_flushed                                	|	
|lsc_int                                       	|	
|multicast                                     	|	
|non_eop_descs                                 	|	
|os2bmc_rx_by_bmc                              	|	
|os2bmc_rx_by_host                             	|	
|os2bmc_tx_by_bmc                              	|	
|os2bmc_tx_by_host                             	|	
|rx_bytes                                      	|	
|rx_bytes_nic                                  	|	
|rx_crc_errors                                 	|	
|rx_csum_offload_errors                        	|	
|rx_dropped                                    	|	
|rx_errors                                     	|	
|rx_fcoe_dropped                               	|	
|rx_fcoe_dwords                                	|	
|rx_fcoe_packets                               	|	
|rx_fifo_errors                                	|	
|rx_flow_control_xoff                          	|	
|rx_flow_control_xon                           	|	
|rx_frame_errors                               	|	
|rx_long_length_errors                         	|	
|rx_missed_errors                              	|	
|rx_no_buffer_count                            	|	
|rx_no_dma_resources                           	|	
|rx_over_errors                                	|	
|rx_packets                                    	|	
|rx_pb_<num>_pxoff                              |	
|rx_pb_<num>_pxon                               |	
|rx_pkts_nic                                   	|	
|rx_queue_<num>_bp_cleaned                      |	
|rx_queue_<num>_bp_misses                       |	
|rx_queue_<num>_bp_poll_yield                   |	
|rx_queue_<num>_bytes                           |	
|rx_queue_<num>_packets                         |	
|rx_short_length_errors                        	|	
|tx_aborted_errors                             	|	
|tx_busy                                       	|	
|tx_bytes                                      	|	
|tx_bytes_nic                                  	|	
|tx_carrier_errors                             	|	
|tx_dropped                                    	|	
|tx_errors                                     	|	
|tx_fcoe_dwords                                	|	
|tx_fcoe_packets                               	|	
|tx_fifo_errors                                	|	
|tx_flow_control_xoff                          	|	
|tx_flow_control_xon                           	|	
|tx_heartbeat_errors                           	|	
|tx_packets                                    	|	
|tx_pb_<num>_pxoff                             	|	
|tx_pb_<num>_pxon                             	|	
|tx_pkts_nic                                   	|	
|tx_queue_<num>_bp_cleaned                      |	
|tx_queue_<num>_bp_misses                       |
|tx_queue_<num>_bp_napi_yield                   |	
|tx_queue_<num>_bytes                           |	
|tx_queue_<num>_packets                       	|	
|tx_restart_queue                              	|	
|tx_timeout_count                              	|	



**b) Register dump statistics**

Namespace: `/intel/net/ixgbe/<device name>/reg/<metric name>`

Metric Name| Description
------------ | -------------     
|AIS                                           	|	Auto-Scan Interrupt Status
|ANLP1                                         	|	Auto Neg Lnk Part. Ctrl Word 1
|ANLP2                                         	|	Auto Neg Lnk Part. Ctrl Word 2
|APAE                                          	|	Auto-Scan PHY Address Enable
|ARD                                           	|	Auto-Scan Read Data
|ATLASCTL                                      	|	Atlas Analog Configuration
|AUTOC                                         	|	Auto Negotiation Control
|AUTOC2                                        	|	Auto Negotiation Control 2
|AUTOC3                                        	|	Auto Negotiation Control 3
|CTRL                                          	|	Device Control
|CTRL_EXT                                      	|	Extended Device Control
|DCA_RXCTRL<num>                             	|	Rx DCA Control <num>
|DCA_TXCTRL<num>                              	|	Tx DCA Control <num>
|DPMCS                                         	|	Desc. Plan Music Ctrl Status
|DROPEN                                        	|	Drop Enable Control
|DTXCTL                                        	|	DMA Tx Control
|EEC                                           	|	EEPROM/Flash Control
|EEMNGCTL                                      	|	Manageability EEPROM Control
|EEMNGDATA                                     	|	Manageability EEPROM R/W Data
|EERD                                          	|	EEPROM Read
|EIAC                                          	|	Extended Interrupt Auto Clear
|EIAM                                          	|	Extended Interr. Auto Mask EN
|EICR                                          	|	Extended Interrupt Cause
|EICS                                          	|	Extended Interrupt Cause Set
|EIMC                                          	|	Extended Interrupt Mask Clear
|EIMS                                          	|	Extended Interr. Mask Set/Read
|EITR0                                         	|	Extended Interrupt Throttle 0
|EODSDP                                        	|	Extended OD SDP Control
|ESDP                                          	|	Extended SDP Control
|FCCTV<0-3>                                      	|	Flow Ctrl Tx Timer Value <0-3>
|FCRTH<num>                                    	|	Flow Ctrl Rx Threshold High <num>
|FCRTL<num>                                    	|	Flow Ctrl Rx Threshold low <num>
|FCRTV                                         	|	Flow Control Refresh Threshold
|FCTRL                                         	|	Filter Control register
|FHFT                                          	|	Flexible Host Filter Table
|FLA                                           	|	Flash Access
|FLMNGCNT                                      	|	Manageability Flash Read Count
|FLMNGCTL                                      	|	Manageability Flash Control
|FLMNGDATA                                     	|	Manageability Flash Read Data
|FLOP                                          	|	Flash Opcode
|FRTIMER                                       	|	Free Running Timer
|GPIE                                          	|	General Purpose Interrupt EN
|GRC                                           	|	General Receive Control
|HLREG0                                        	|	Highlander Control 0 register
|HLREG1                                        	|	Highlander Status 1
|IMIR<num>                                     	|	Immediate Interrupt Rx <num>
|IMIREXT<num>                                  	|	Immed. Interr. Rx Extended <num>
|IMIRVP                                        	|	Immed. Interr. Rx VLAN Prior.
|IP4AT                                         	|	IPv4 Address Table
|IP6AT                                         	|	IPv6 Address Table
|IPAV                                          	|	IP Address Valid
|IVAR0                                         	|	Interrupt Vector Allocation 0
|LEDCTL                                        	|	LED Control
|LINKS                                         	|	Link Status register
|MACA                                          	|	MDI Auto-Scan Command and Addr
|MACS                                          	|	FIFO Status/CNTL Report
|MCSTCTRL                                      	|	Multicast Control
|MDFTC1                                        	|	MAC DFT Control 1
|MDFTC2                                        	|	MAC DFT Control
|MDFTFIFO1                                     	|	MAC DFT FIFO 1
|MDFTFIFO2                                     	|	MAC DFT FIFO 2
|MDFTS                                         	|	MAC DFT Status
|MFLCN                                         	|	TabMAC Flow Control register
|MHADD                                         	|	MAC Addr High/Max Frame size
|MLADD                                         	|	MAC Address Low
|MNGTXMAP                                      	|	Manageability Tx TC Mapping
|MRQC                                          	|	Multiple Rx Queues Command
|MSCA                                          	|	MDI Single Command and Addr
|MSIXPBA                                       	|	MSI-X Pending Bit Array
|MSIXT                                         	|	MSI-X Table
|MSRWD                                         	|	MDI Single Read and Write Data
|PAP                                           	|	Pause and Pace
|PBACL                                         	|	MSI-X PBA Clear
|PBRXECC                                       	|	Packet Buffer Rx ECC
|PBTXECC                                       	|	Packet Buffer Tx ECC
|PCIEECCCTL                                    	|	PCIe ECC Control
|PCIE_DIAG<num>                                   |   PCIe Diagnostic <num>	
|PCS1GANA                                      	|	PCS-1G Auto Neg. Adv.
|PCS1GANLP                                     	|	PCS-1G AN LP Ability
|PCS1GANLPNP                                   	|	PCS_1G Auto Neg LPs Next Page
|PCS1GANNP                                     	|	PCS_1G Auto Neg Next Page Tx
|PCS1GCFIG                                     	|	PCS_1G Gloabal Config 1
|PCS1GDBG0                                     	|	PCS_1G Debug 0
|PCS1GDBG1                                     	|	PCS_1G Debug 1
|PCS1GLCTL                                     	|	PCS_1G Link Control
|PCS1GLSTA                                     	|	PCS_1G Link Status
|PCSS1                                         	|	XGXS Status 1
|PCSS2                                         	|	XGXS Status 2
|PDPMCS                                        	|	Pkt Data Plan Music ctrl Stat
|PFCTOP                                        	|	Priority Flow Ctrl Type Opcode
|PSRTYPE                                       	|	Packet Split Receive Type
|RAH<num>                                       |	Receive Address High <num>
|RAL<num>                                       |	Receive Address Low <num>
|RDBAH<num>                                  	|	Rx Desc Base Addr High <num>
|RDBAL<00-63                                   	|	Rx Desc Base Addr Low <num>	
|RDH<num>                                    	|	Receive Descriptor Head <num>	
|RDHMPN                                        	|	Rx Desc Handler Mem Page num
|RDLEN<num>                                  	|	Receive Descriptor Length <num>	
|RDPROBE                                       	|	Rx Probe Mode Status
|RDRXCTL                                       	|	Receive DMA Control
|RDSTAT<num>                                   	|	Rx DMA Statistics <num>	
|RDSTATCTL                                     	|	Rx DMA Statistic Control
|RDT<num>                                    	|	Receive Descriptor Tail <num>	
|RFCTL                                         	|	Receive Filter Control
|RFVAL                                         	|	Receive Filter Validation
|RIC_DW0                                       	|	Rx Desc Hand. Mem Read Data 0
|RIC_DW1                                       	|	Rx Desc Hand. Mem Read Data 1
|RIC_DW2                                       	|	Rx Desc Hand. Mem Read Data
|RIC_DW3                                       	|	Rx Desc Hand. Mem Read Data 3
|RMCS                                          	|	Receive Music Control register
|RT2CR<num>                                    	|	Receive T2 Configure <num>	
|RT2SR<num>                                    	|	Recieve T2 Status <num>	
|RUPPBMR                                       	|	Rx User Prior to Pkt Buff Map
|RXBUFCTRL                                     	|	RX Buffer Access Control
|RXBUFDATA0                                    	|	RX Buffer DATA 0
|RXBUFDATA1                                    	|	RX Buffer DATA 1
|RXBUFDATA2                                    	|	RX Buffer DATA 3
|RXBUFDATA3                                    	|	RX Buffer DATA 4
|RXCSUM                                        	|	Receive Checksum Control
|RXCTRL                                        	|	Receive Control
|RXDCTL<num>                                 	|	Receive Descriptor Control <num>	
|RXPBSIZE<num>                                 	|	Receive Packet Buffer Size <num>	
|SERDESC                                       	|	SERDES Interface Control
|SRRCTL0                                       	|	Split and Replic Rx Control 0
|SRRCTL<num>                                 	|	Split and Replic Rx Control <num>	
|STATUS                                        	|	Device Status
|TCPTIMER                                      	|	TCP Timer
|TDBAH<num>                                  	|	Tx Desc Base Addr High <num>	
|TDBAL<num>                                  	|	Tx Desc Base Addr Low <num>	
|TDH<num>                                    	|	Transmit Descriptor Head <num>
|TDHMPN                                        	|	Tx Desc Handler Mem Page Num
|TDLEN<num>                                  	|	Tx Descriptor Length <num>	
|TDPROBE                                       	|	Tx Probe Mode Status
|TDPT2TCCR<num>                                	|	Tx Data Plane T2 TC Config <num>
|TDPT2TCSR<num>                                    |	Tx Data Plane T2 TC Status <num>	
|TDSTAT<num>                                       |	Tx DMA Statistics 0	
|TDSTATCTL                                     	|	Tx DMA Statistic Control
|TDT<num>                                        |	Transmit Descriptor Tail <num>	
|TDTQ2TCCR<num>                                    |	Tx Desc TQ2 TC Config <num>
|TDTQ2TCSR<num>                                    |	Tx Desc TQ2 TC Status <num>	
|TDWBAH<num>                                     |	Tx Desc Compl. WB Addr High <num>	
|TDWBAL<num>                                     |	Tx Desc Compl. WB Addr low <num>	
|TFCS                                          	|	Transmit Flow Control Status
|TIC_DW0                                       	|	Tx Desc Hand. Mem Read Data 0
|TIC_DW1                                       	|	Tx Desc Hand. Mem Read Data 1
|TIC_DW2                                       	|	Tx Desc Hand. Mem Read Data
|TIC_DW3                                       	|	Tx Desc Hand. Mem Read Data 3
|TIPG                                          	|	Transmit IPG Control
|TREG                                          	|	Test Register
|TXBUFCTRL                                     	|	TX Buffer Access Control
|TXBUFDATA0                                    	|	TX Buffer DATA 0
|TXBUFDATA1                                    	|	TX Buffer DATA 1
|TXBUFDATA2                                    	|	TX Buffer DATA 2
|TXBUFDATA3                                    	|	TX Buffer DATA 3
|TXDCTL<num>                                     |	Tx Descriptor Control <num>	
|TXPBSIZE<num>                                 	|	Transmit Packet Buffer Size <num>	
|VLNCTRL                                       	|	VLAN Control register
|VMD_CTL                                       	|	VMDq Control
|WUC                                           	|	Wake up Control
|WUFC                                          	|	Wake Up Filter Control
|WUPL                                          	|	Wake Up Packet Length
|WUPM                                          	|	Wake Up Packet Memory
|WUS                                           	|	Wake Up Status
|XPCSS                                         	|	10GBASE-X PCS Status
|bprc                                          	|	Broadcast Packets Rx Count
|bptc                                          	|	Broadcast Packets Tx Count
|broadcast_accept                              	|	broadcast acceptance (enabled/disabled)
|crcerrs                                       	|	CRC Error Count
|discard_pause_frames                          	|	Pause frames discards (enabled/disabled)
|errbc                                         	|	Error Byte Count
 |
|**gprc                                         |	Good Packets Received Count
|**gorch**                                      |	Good Octets Rx Count High
|**gorcl**                                      |	Good Octets Rx Count Low
|**gptc                                         |	Good Packets Transmitted Count
|**gotch**                                     	|	Good Octets Tx Count High
|**gotcl                                      	|	Good Octets Tx Count Low
 |
|illerrc                                       	|	Illegal Byte Error Count
|jumbo_frames                                  	|	Jumbo frames (enabled/disabled)
|link_speed                                    	|	Speed of link
|link_status                                   	|	Status of link (down/up)
|loopback                                      	|	Loopback (enabled/disabled)
|lxoffrxc                                      	|	Link XOFF Received Count
|lxofftxc                                      	|	Link XOFF Transmitted Count
|lxonrxc                                       	|	Link XON Received Count
|lxontxc                                       	|	Link XON Transmitted Count
|mlfc                                          	|	MAC Local Fault Count
|mngpdc                                        	|	Management Pkts Dropped Count
|mngprc                                        	|	Management Packets Rx Count
|mngptc                                        	|	Management Packets Tx Count
|mpc<num>                                      	|	Missed Packets Count <num>
|mprc                                          	|	Multicast Packets Rx Count
|mptc                                          	|	Multicast Packets Tx Count
|mrfc                                          	|	MAC Remote Fault Count
|mspdc                                         	|	MAC Short Packet Discard Count
|multicast_promiscuous                         	|	Multicast promiscuous(enabled/disabled)
|pad_short_frames                              	|	Pad short frames (enabled/disabled)
|pass_mac_control_frames                       	|	Pass MAC control frames (enable/disabled)
 |			
|**prc64**                                      |	Packets Received (64B) Count
|**prc127**                                     |	Packets Rx (65-127B) Count
|**prc255**                                     |	Packets Rx (128-255B) Count
|**prc511**                                     |	Packets Rx (256-511B) Count
|**prc1023**                                    |	Packets Rx (512-1023B) Count
|**prc1522**                                    |	Packets Rx (1024-Max) Count	
|**ptc64**                                      |	Packets Transmitted (64B) Count
|**ptc127**                                     |	Packets Tx (65-127B) Count
|**ptc255**                                     |	Packets Tx (128-255B) Count
|**ptc511**                                     |	Packets Tx (256-511B) Count
|**ptc1023**                                    |	Packets Tx (512-1023B) Count
|**ptc1522**                                    |	Packets Tx (1024-Max) Count
 |
|priority_flow_control                         	|	Priority flow control status (enabled/disabled)	
|pxoffrxc<num>                                 	|	Priority XOFF Received Count <num>
|pxofftxc<num>                                 	|	Priority XOFF Transmitted Count <num>	
|pxonrxc<num>                                  	|	Priority XON Received Count <num>
|pxontxc<num>                                  	|	Priority XON Transmitted Count <num>	
|qbrc<num>                                   	|	Queue Bytes Rx Count <num>
|qbtc<num>                                   	|	Queue Bytes Tx Count <num>	
|qprc<num>                                   	|	Queue Packets Rx Count <num>
|qptc<num>                                   	|	Queue Packets Tx Count <num>	
|receive_buffer_size                           	|	Receive buffer size [KB]
|receive_crc_strip                             	|	Receiving of CRC strip (enabled/disabled)
|receive_flow_control_packets                  	|	Receiving of flow control packets (enabled/disabled)
|receive_priority_flow_control_packets         	|	Receiving of flow control packets priority (enabled/disabled)
|rfc                                           	|	Receive Filter Control
|rjc                                           	|	Receive Jabber Count
|rlec                                          	|	Receive Length Error Count
|rnbc<num>                                     	|	Receive No Buffers Count <num>
|roc                                           	|	Receive Oversize Count
|ruc                                           	|	Receive Undersize count
|store_bad_packets                             	|	Store bad packets status (enabled/disabled)
 |	
|**torh**                                       |	Total Octets Rx Count High
|**torl**                                       |	Total Octets Rx Count Low
 |
|**tpr**                                        |	Total Packets Received
|**tpt**										|	Total Packets Transmitted
 |
|transmit_crc                                  	|	Transmit CRC (enabled/disabled)
|transmit_flow_control                         	|	Transmit flow control (enabled/disabled)
|unicast_promiscuous                           	|	Unicast promiscuous (enabled/disabled)
|vlan_filter                                   	|	VLAN filter status (enabled/disabled)
|vlan_mode                                     	|	VLAN mode (enabled/disabled)
|xec                                           	|	XSUM Error Count	