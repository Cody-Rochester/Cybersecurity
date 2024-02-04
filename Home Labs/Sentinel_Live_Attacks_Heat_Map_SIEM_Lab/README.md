<h1>Failed RDP to IP Geolocation Information</h1>

<h2>Description</h2>
<b> Using a third party API, I'm able to run a PowerShell script to send real-time failed RDP info (event 4625) to collect geographic data that is then given to a custom log in Sentinel. From there we set up a workbook that displays that custom log in a live (updates every 5min) heat map of the world.
</b>
<br />
<br />
This demo is a way to observe live Brute Force attacks from across the globe. The script does need to be running on the target machine in order to relay the information to our SIEM, so this isn't a practical tool for real world situations, but it is a very interesting way to collect and display information on attackers. 

<br />
<br />


<h2>Languages Used</h2>

- <b>PowerShell:</b> Extract RDP failed logon logs from Windows Event Viewer (EventID 4625)

<h2>Utilities Used</h2>

- <b>ipgeolocation.io:</b> IP Address to Geolocation API
- <b>Azure Virtual Machines:</b> Creating Honeypot VM to expose to attackers
- <b>Azure Logging and Analytics Workspace:</b> Linking PowerShell script to custom log to collect and parse data
- <b>Azure Sentinel:</b> Organize and display information gathered from attackers in a map format

<h2>Attacks from China coming in; Custom logs being output with geodata</h2>



<h2>World map of incoming attacks after 24 hours (built custom logs including geodata)</h2>



<!--
 ```diff
- text in red
+ text in green
! text in orange
# text in gray
@@ text in purple (and bold)@@
```
--!>