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


<h2>Open Source .sln file in Visual Studio</h2>

![keylogger source](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/03f667a5-4926-4ed2-b8f9-758f0e7bcd2e)


<h2>Press run in Visual Studio to begin logging</h2>

![logging](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/96ec1009-68bb-4533-85d3-1fbf77d996b7)



<br />
<br />
<h2>The logger sends an email with the string and a file attachment also containing that info. Here the string is set to 300 but can be adjusted</h2>

![logger sending email automatically](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/ee75896a-c5d4-4e8b-b8ef-988845424364)

<h2>To conceal the application, in "Properties" on the project, change "Console Application" to "Windows Appliation"</h2>

![in properties change from console to windows to conceal the process](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/d3ace3e5-e45a-4376-8645-47029220a90b)

![switched](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/dc1f5679-66f2-4d46-84b7-1d2bd1128077)

<br />
<br />
<h2>Emails continue to come in while running the application as a Windows Application</h2>

![still sending logged emails while running concealed](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/6b097ca1-b9d7-4caa-a2b5-8543e66fa65d)


<h2>Using Task Scheduler we can set this application to run on startup creating a deeper level of persistence</h2>

![task scheduler to set task to run on startup](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/692aa34a-923a-44fe-83a7-c66d0343d6b2)


<br />
<br />
<h2>In order to allow gmail to send emails like this, you will need to activate 2-factor auth and get an application password</h2>





<!--
 ```diff
- text in red
+ text in green
! text in orange
# text in gray
@@ text in purple (and bold)@@
```
--!>
