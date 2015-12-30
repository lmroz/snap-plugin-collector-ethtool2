# snap collector plugin - ethtool		
		
## Metrics exposed via driver FM10k:		
				
### NIC statistics

Namespace: `/intel/net/fm10k/<device name>/nic/<metric name>`
		
Metrics |	
------------ |			
|ca                                            	|
|loopback_drop                                 	|
|mac_rules_avail                               	|
|mac_rules_used                                	|
|mbx_rx_dwords                                 	|
|mbx_rx_mbmem_pushed                           	|
|mbx_rx_messages                               	|
|mbx_rx_parse_err                              	|
|mbx_tx_busy                                   	|
|mbx_tx_dropped                                	|
|mbx_tx_dwords                                 	|
|mbx_tx_mbmem_pulled                           	|
|mbx_tx_messages                               	|
|nodesc_drop                                   	|
|rx_alloc_failed                               	|
|rx_bytes                                      	|
|rx_bytes_nic                                  	|
|rx_crc_errors                                 	|
|rx_csum_errors                                	|
|rx_dropped                                    	|
|rx_drops_nic                                  	|
|rx_errors                                     	|
|rx_fifo_errors                                	|
|rx_length_errors                              	|
|rx_overrun_pf                                 	|
|rx_overrun_vf                                 	|
|rx_packets                                    	|
|rx_packets_nic                                	|
|rx_queue_0_bytes                              	|
|rx_queue_0_packets                            	|
|rx_queue_100_bytes                            	|
|rx_queue_100_packets                          	|
|rx_queue_101_bytes                            	|
|rx_queue_101_packets                          	|
|rx_queue_102_bytes                            	|
|rx_queue_102_packets                          	|
|rx_queue_103_bytes                            	|
|rx_queue_103_packets                          	|
|rx_queue_104_bytes                            	|
|rx_queue_104_packets                          	|
|rx_queue_105_bytes                            	|
|rx_queue_105_packets                          	|
|rx_queue_106_bytes                            	|
|rx_queue_106_packets                          	|
|rx_queue_107_bytes                            	|
|rx_queue_107_packets                          	|
|rx_queue_108_bytes                            	|
|rx_queue_108_packets                          	|
|rx_queue_109_bytes                            	|
|rx_queue_109_packets                          	|
|rx_queue_10_bytes                             	|
|rx_queue_10_packets                           	|
|rx_queue_110_bytes                            	|
|rx_queue_110_packets                          	|
|rx_queue_111_bytes                            	|
|rx_queue_111_packets                          	|
|rx_queue_112_bytes                            	|
|rx_queue_112_packets                          	|
|rx_queue_113_bytes                            	|
|rx_queue_113_packets                          	|
|rx_queue_114_bytes                            	|
|rx_queue_114_packets                          	|
|rx_queue_115_bytes                            	|
|rx_queue_115_packets                          	|
|rx_queue_116_bytes                            	|

Notice: No metric exposed reg statistics for fm10k driver, only raw register dump values available.
