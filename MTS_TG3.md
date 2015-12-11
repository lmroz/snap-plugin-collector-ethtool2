# snap collector plugin - ethtool		
		
## Metrics exposed via driver tg3:		
		
**a) NIC statistics**

Namespace: `/intel/net/tg3/<device name>/nic/<metric name>`
			
Metrics |	
------------			
|dma_read_prioq_full                             	
|dma_readq_full                                  	
|dma_write_prioq_full                            	
|dma_writeq_full                                 	
|mbuf_lwm_thresh_hit                             	
|nic_avoided_irqs                                	
|nic_irqs                                        	
|nic_tx_threshold_hit                            	
|ring_set_send_prod_index                        	
|ring_status_update                              	
 |
|rx_64_or_less_octet_packets                     	
|rx_65_to_127_octet_packets                      	
|rx_128_to_255_octet_packets                     	
|rx_256_to_511_octet_packets                     	
|rx_512_to_1023_octet_packets                    	
|rx_1024_to_1522_octet_packets                   	
|rx_1523_to_2047_octet_packets                   	
|rx_2048_to_4095_octet_packets                   	
|rx_4096_to_8191_octet_packets                   	
|rx_8192_to_9022_octet_packets   
 |
|rx_align_errors                                 	
|rx_bcast_packets                                	
|rx_discards                                     	
|rx_errors                                       	
|rx_fcs_errors                                   	
|rx_fragments                                    	
|rx_frame_too_long_errors                        	
|rx_in_length_errors                             	
|rx_jabbers                                      	
|rx_mac_ctrl_rcvd                                	
|rx_mcast_packets                                	
|rx_octets                                       	
|rx_out_length_errors                            	
|rx_threshold_hit                                	
|rx_ucast_packets                                	
|rx_undersize_packets                            	
|rx_xoff_entered                                 	
|rx_xoff_pause_rcvd                              	
|rx_xon_pause_rcvd                               	
|rxbds_empty                                     	
|tx_bcast_packets                                	
|tx_carrier_sense_errors                         	
|tx_collide_<num>times                              
|tx_collisions                                   	
|tx_comp_queue_full                              	
|tx_deferred                                     	
|tx_discards                                     	
|tx_errors                                       	
|tx_excessive_collisions                         	
|tx_flow_control                                 	
|tx_late_collisions                              	
|tx_mac_errors                                   	
|tx_mcast_packets                                	
|tx_mult_collisions                              	
|tx_octets                                       	
|tx_single_collisions                            	
|tx_ucast_packets                                	
|tx_xoff_sent                                    	
|tx_xon_sent                                     	


Notice:	No metric exposed reg statistics for tg3 driver, only raw register dump values available.
