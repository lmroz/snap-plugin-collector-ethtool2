# snap collector plugin - ethtool		
		

### Digital optical monitoring

Diagnostics data and alarms for optical transceivers (SFP, SFP+ or XFP).

Namespace: `/intel/net/<driver name>/<device name>/dom/<metric name>`


Metric Name| Description
------------ | -------------    
|alarm_warning_flags_implemented                    |	Alarm warning flags implemented (yes/no).
|br_nominal                                         |	The nominal bit rate  is specified in units of 100 Megabits per second.
|connector                                          |	The connector value indicates the external connector provided on the interface.																	A value 1 means SC, a value 7 means LC.
|encoding                                           |	The encoding value indicates the serial encoding mechanism. Values: 0 - Unspecified; 															1 - 8B10B; 2 - 4B5B; 3 - NRZ; 4 - Manchester; 5 - SONET Scrambled;
|extended_identifier                                |	This object describes the extended identifier value specifies the physical device 																described by the serial information. A value 4 means GBIC/SFP defined by 2-wire 																interface ID.
|identifier                     		            |	This object describes the identifier value specifies the physical device described by 														the serial information. A value 3 means SFP.
| |
|laser_bias_current                                 |	Displays the magnitude of the laser bias power current, in milliamperes (mA). 																	The laser bias provides direct modulation of laser diodes and modulates currents.
|laser_bias_current_high_alarm                      |	Displays whether the laser bias current high alarm is On or Off.
|laser_bias_current_high_alarm_threshold            |	Displays the vendor-specified threshold for the laser bias current high alarm.
|laser_bias_current_high_warning                    |	Displays whether the laser bias current high warning is On or Off.
|laser_bias_current_high_warning_threshold          |	Displays the vendor-specified threshold for the laser bias current high warning.
|laser_bias_current_low_alarm                       |	Displays whether the laser bias current low alarm is On or Off.
|laser_bias_current_low_alarm_threshold             |	Displays the vendor-specified threshold for the laser bias current low alarm.
|laser_bias_current_low_warning                     |	Displays whether the laser bias current low warning is On or Off.
|laser_bias_current_low_warning_threshold           |	Displays the vendor-specified threshold for the laser bias current low warning.
| |
|laser_output_power                                 |	Displays the laser output power, in milliwatts (mW) and decibels referred to 1.0 mW (dBm).
|laser_output_power_high_alarm                      |	Displays whether the laser output power high alarm is On or Off.
|laser_output_power_high_alarm_threshold            |	Displays the vendor-specified threshold for the laser output power high alarm.
|laser_output_power_high_warning                    |	Displays whether the laser output power high warning is On or Off.
|laser_output_power_high_warning_threshold          |	Displays the vendor-specified threshold for the laser output power high warning.
|laser_output_power_low_alarm                       |	Displays whether the laser output power low alarm is On or Off.
|laser_output_power_low_alarm_threshold             |	Displays the vendor-specified threshold for the laser output power low alarm.
|laser_output_power_low_warning                     |	Displays whether the laser output power low warning is On or Off.
|laser_output_power_low_warning_threshold			|	Displays the vendor-specified threshold for the laser output power low warning.
| 													|
|laser_rx_power<sup>(2)</sup>                   	|	Displays the receive laser optical power, in milliwatts (mW) and decibels referred to 														1.0 mW (dBm).
|laser_rx_power_high_alarm                          |	Displays whether the receive laser power high alarm is On or Off.
|laser_rx_power_high_alarm_threshold                |	Displays the vendor-specified threshold for  the receive laser power high alarm.
|laser_rx_power_high_warning                        |	Displays whether the receive laser power high warning is On or Off.
|laser_rx_power_high_warning_threshold              |	Displays the vendor-specified threshold the receive laser power high warning.
|laser_rx_power_low_alarm                           |	Displays whether the receive laser power low alarm is On or Off.
|laser_rx_power_low_alarm_threshold                 |	Displays the vendor-specified threshold forthe receive laser power low alarm.
|laser_rx_power_low_warning                         |	Displays whether the receive laser power low warning is On or Off.
|laser_rx_power_low_warning_threshold               |	Displays the vendor-specified threshold for the receive laser power low warning.
| |
|laser_wavelength                                   |	The fibre channel transmitter wave laser length (nm).
|length_50um                                        |	This value specifies the link length that is supported while operating in 																		compliance with the applicable standards using 50 micron multi-mode fiber. The value 															is in units of 10 meters. A value of 255 means that a link length greater than 2.54 															km is supported. A value of zero means that the 50 micron multi-mode fiber is not 																supported or that the length information must be determined from the transceiver 																technology.
|length_62.5um                                      |	This value specifies the link length that is supported while operating in 																		compliance with the applicable standards using 62.5 micron multi-mode fiber. The 																value is in units of 10 meters. A value of 255 means that a link length information 															must determined from the transceiver technology. It is common to support both 50 																micron and 62.5 micron fiber. 
|length_copper                                      |	This value specifies the minimum link length that is supported while operating in 																compliance with the applicable standards using copper cable. The value is in units of 														1 meter. A value of 255 means that a link length greater than 254 meters is 																	supported. A value of zero means that the copper cables are not supported or that the 														length information must be determined from the transceiver technology.
|length_om3                                         |	This value specifies the minimum link length that is supported while operating in 																compliance with the applicable standards using OM3 multimode fibers.
|length_smf                                         |	This value specifies the minimum link length that is supported while operating in 																compliance with the applicable standards using single-mode (SFM) optical fibers. The 															value is in units of meters. A value of zero means that the SFM fibers are not 																	supported or that the length information must be determined from the transceiver 																technology.
|length_smf_km                                      |	This value specifies the minimum link length that is supported while operating in 																compliance with the applicable standards using single-mode (SFM) optical fibers. The 															value is in units of kilometers. A value of zero means that the SFM fibers are not 																supported or that the length information must be determined from the transceiver 																technology.
| |
|module_not_ready_alarm<sup>(2)</sup>				| 	Displays whether the module not ready alarm is On or Off.
|module_power_down_alarm<sup>(2)</sup>				| 	Displays whether the module power down alarm is On or Off. When the output is On, 																module is in a limited power mode, low for normal operation.
| |
|module_temperature                                 |	Displays the temperature, in Celsius and Fahrenheit.
|module_temperature_high_alarm                      |	Displays whether the module temperature high alarm is On or Off.
|module_temperature_high_alarm_threshold            |	Displays the vendor-specified threshold for the module temperature high alarm.
|module_temperature_high_warning                    |	Displays whether the module temperature high warning is On or Off.
|module_temperature_high_warning_threshold          |	Displays the vendor-specified threshold for the module temperature high warning.
|module_temperature_low_alarm                       |	Displays whether the module temperature low alarm is On or Off.
|module_temperature_low_alarm_threshold             |	Displays the vendor-specified threshold for the module temperature low alarm.
|module_temperature_low_warning                     |	Displays whether the module temperature low warning is On or Off.
|module_temperature_low_warning_threshold           |	Displays the vendor-specified threshold for the module temperature low warning.
| |
|module_voltage<sup>(1)</sup>                       |	Displays the voltage in Volts.
|module_voltage_high_alarm<sup>(1)</sup>            |	Displays whether the module voltage high alarm is On or Off.
|module_voltage_high_alarm_threshold<sup>(1)</sup>  |	Displays the vendor-specified threshold for the module voltage high alarm.
|module_voltage_high_warning_threshold<sup>(1)</sup>|	Displays the vendor-specified threshold for the module voltage high warning.
|module_voltage_low_alarm<sup>(1)</sup>             |	Displays whether the module voltage low alarm is On or Off.
|module_voltage_low_alarm_threshold<sup>(1)</sup>   |	Displays the vendor-specified threshold for the module voltage low alarm.
|module_voltage_low_warning<sup>(1)</sup>           |	Displays whether the module voltage low warning is On or Off.
|module_voltage_low_warning_threshold<sup>(1)</sup> |	Displays the vendor-specified threshold for the module voltage low warning.
| |
|optical_diagnostics_support                        |	Displays whether the optical diagnostics is supported (yes/no).
|rate_identifier                                    |
|receiver_signal_average_optical_power<sup>(1)</sup>|	Displays the receiver signal average optical power, in milliwatts (mW) and decibels 															referred to 1.0 mW (dBm).
| |
|rx_cdr_loss_of_lock_alarm<sup>(2)</sup>			|	Displays whether the Rx CDR (receive clock and data recovery) loss of lock alarm is 															On or Off. Triggered by loss of lock on the receive side of the CDR.
|rx_loss_of_signal<sup>(2)</sup>		   			|	Displays whether the Rx loss of signal alarm is On or Off. When on, indicates 																	insufficient optical input power to the module.
|rx_not_ready_alarm<sup>(2)</sup>			   		|	Displays whether the Rx not ready alarm is On or Off.  Triggered by any condition 																leading to invalid data on the receive path.
| |
|transceiver_codes                                  |	This describes the transceiver compliance codes. 
|transceiver_type                                   |	This describes the transceiver type information. 
| |
|tx_cdr_loss_of_lock_alarm<sup>(2)</sup>			|	Displays whether the Tx CDR (transmit clock and data recovery) loss of lock alarm is 															On or Off. Triggered by loss of lock on the transmit side of the CDR.
|tx_data_not_ready_alarm<sup>(2)</sup>				|	Displays whether the Tx data not ready alarm is On or Off. Triggered by any condition 														leading to invalid data on the transmit path.
|tx_laser_fault_alarm<sup>(2)</sup>			   		|	Displays whether the Tx laser fault alarm is On or Off. Triggered by laser fault 																condition.
|tx_not_ready_alarm<sup>(2)</sup>			   		|	Displays whether the Tx not ready alarm is On or Off.  Triggered by any condition 																leading to invalid data on the transmit path.
| |
|vendor_name                                        |	The vendor name contains the name of the corporation.
|vendor_oui                                         |	The vendor organizationally unique identifier field contains the IEEE Company 																	Identifier for the vendor or zero if unspecified.
|vendor_pn                                          |	The vendor part number or product name, equals zero if unspecified.
|vendor_rev                                         |	The vendor revision number.
			
Legend: <br />																																		<sup>(1)</sup> Not available for XFP transceivers <br />																						<sup>(2)</sup> Not available for SFP and SFP+ transceivers <br />

<br />																																			Notice:	<br />																																	Thresholds that trigger a high alarm, low alarm, high warning, or low warning are set by the transponder vendors. Generally, a high alarm or low alarm indicates that the optics module is not operating properly. This information can be used to diagnose why a transceiver is not working.

