<h1>FOR EDUCATIONAL USE ONLY</h1>
<h1>Keylogger with Email Output</h1>
<h2>Description</h2>
<b> Using a third party API, I'm able to run a PowerShell script to send real-time failed RDP info (event 4625) to collect geographic data that is then given to a custom log in Sentinel. From there we set up a workbook that displays that custom log in a live (updates every 5min) heat map of the world.
</b>
<br />
<br />
This demo is a way to observe live Brute Force attacks from across the globe. The script does need to be running on the target machine in order to relay the information to our SIEM, so this isn't a practical tool for real world situations, but it is a very interesting way to collect and display information on attackers. 

<br />
<br />


<h2>Languages Used</h2>

- <b>C#:</b> Used to: Log keystrokes, create log files, log into gmail account, create draft, add string to draft, attach log to draft, and send to target email


<h2>Tests coming in; Custom logs being output with geodata</h2>




<h2>World map of incoming attacks after 1 hour</h2>



<br />
<br />
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
