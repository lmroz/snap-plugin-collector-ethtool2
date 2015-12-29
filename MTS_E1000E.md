# snap collector plugin - ethtool

## Metrics exposed via driver e1000e:

**a) NIC statistics**

Namespace: `/intel/net/e1000e/<device name>/nic/<metric name>`

Metric Name|
------------ |
|alloc_rx_buff_failed |                 
|collisions |            
|corr_ecc_errors |                      
|dropped_smbus |                        
|multicast |                        
|rx_align_errors |             
|rx_broadcast |                        
|**rx_bytes** |                            
|rx_crc_errors |                      
|rx_csum_offload_errors |               
|rx_csum_offload_good |                
|rx_dma_failed |                        
|**rx_errors** |                           
|rx_flow_control_xoff |            
|rx_flow_control_xon |                 
|rx_frame_errors |                
|rx_header_split |                    
|rx_hwtstamp_cleared |                 
|rx_length_errors |          
|rx_long_length_errors |               
|rx_missed_errors |                  
|rx_multicast |                 
|rx_no_buffer_count |                 
|rx_over_errors |                     
|**rx_packets** | 
|rx_short_length_errors |          
|rx_smbus |
 |                            
|tx_abort_late_coll |                 
|tx_aborted_errors |                
|tx_broadcast |                     
|**tx_bytes** |                    
|tx_carrier_errors |                    
|tx_deferred_ok |                      
|tx_dma_failed |                       
|tx_dropped |                           
|**tx_errors** |                      
|tx_fifo_errors |                      
|tx_flow_control_xoff | 
|tx_flow_control_xon |              
|tx_heartbeat_errors |             
|tx_hwtstamp_timeouts |                   
|tx_multi_coll_ok |               
|tx_multicast |                  
|**tx_packets** |                         
|tx_restart_queue |                     
|tx_single_coll_ok |                    
|tx_smbus |                            
|tx_tcp_seg_failed |                    
|tx_tcp_seg_good |                  
|tx_timeout_count |                     
|tx_window_errors |                    
|uncorr_ecc_errors |  


**b) Register dump statistics**

Namespace: `/intel/net/e1000e/<device name>/reg/<metric name>`

Metric Name| Description
------------ | -------------                
|CTRL | Device control register status
|RCTL | Receive control register
|RDH | Receive desc head
|RDLEN | Receive desc length
|RDT | Receive desc tail
|RDTR | Receive delay timer
|STATUS | Device status register
|TCTL | Transmit ctrl register
|TDH | Transmit desc head
|TDLEN | Transmit desc length
|TDT | Transmit desc tail
|TIDV | Transmit delay timer
|auto_speed_detect | Auto speed detection status (enabled/disabled)                    
|broadcast_accept_mode |  Broadcast accept mode           
|bus_speed | Bus speed [Hz]                       
|bus_type | Bus type (eg. PCI)                         
|bus_width | Bus width (bit)                           
|canonical_form_indicator | Canonical form indicator status (enabled/disabled)
|descriptor_minimum_threshold_size |   Minimum threshold size of descriptor
|discard_pause_frames | Discard pause frames mode
|duplex | Duplex mode (half/full)
|endian_mode | Endian mode
|force_duplex | Force duplex status
|force_speed | Force speed status                      
|invert_loss-of-signal | Invert loss of signal
|link_reset | Link reset condition                        
|link_speed | Link speed [Mb/s]                       
|link_up |  Link up configuration                          
|long_packet | Long packet status (disabled/enabled)                        
|multicast_promiscuous | Multicast promiscuous status (enabled/disabled)            
|pad_short_packets | Pad short packets status (enabled/disabled)                
|pass_mac_control_frames | Pass MAC control frames (pass/don't pass)              
|phy_type | Type of PHY                            
|re-transmit_on_late_collision | Re-transmission on late collision (enabled/disabled)      
|receive_buffer_size | Receive buffer size [B]              
|receive_flow_control | Receive flow control status (enabled/disabled)
|receiver | Receiver status (enabled/disabled)                           
|set_link_up | Set link up status                        
|software_xoff_transmission | Software xoff transmission status (enabled/disabled)
|speed_select | Speed select [Mb/s]
|store_bad_packets | Store bad packets (enabled/disabled)                  
|tbi_mode | TBI mode (enabled/disabled)                          
|transmit_flow_control | Transmit flow control status (enabled/disabled)
|transmitter | Transmitter status (enabled/disabled)
|unicast_promiscuous | Unicast promiscuous status (enabled/disabled)                 
|vlan_filter | VLAN filter status (enabled/disabled)                         
|vlan_mode | VLAN mode (enabled/disabled)                          