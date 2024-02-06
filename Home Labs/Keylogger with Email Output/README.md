<h1>FOR EDUCATIONAL USE ONLY</h1>
<h1>Keylogger with Email Output</h1>
<h2>Description</h2>
<b> This is an example of a keylogger written in C#. It's primary function is to:<b><b />
 
<br />
<br />
<b>- log keystrokes</b>

<br />
<br />
- convert that data into strings

<br />
<br />
- create log files 

<br />
<br />
- access an email account through SMTP 

<br />
<br />
- create a draft

<br />
<br />
- add the string to that draft

<br />
<br />
- add the attachment to that draft

<br />
<br />
- send to a target email

<br />
<br />
The length of the strings logged, as well as the sending and receiving email addresses are all fully customizable utilizing this process. We also will cover how to run the process as a windows application to make it difficult to detect, as well as an option to set this application to run on startup creating stronger persistence.
</b>

<br />
<br />


<h2>Languages Used</h2>

<h2>C#</h2> 


<h2>Open the source .sln file in Visual Studio</h2>

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
