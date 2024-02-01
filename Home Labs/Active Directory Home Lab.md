<h1>Active Directory Home Lab</h1>

 ### youtube link

<h2>Description</h2>
This project is creating a simulated large scale Active Directory environment. We accomplish this by running two Windows virtual machines using Oracle Virtual Box. One of the machines will act as the domain controller and the other acts as an individual client accessing the network through the internal network.  In order to simulate the large environment, we will be using a Powershell script to create 1,000 users. This exercise includes Domain /AD DS, RAS/NAT, DHCP (1 Scope), Domain Controller setup, VMWare Network, and client access of the internet through domain controller on an AD environment.
<br />


<h2>Languages and Utilities Used</h2>

- <b>PowerShell</b> 
- <b>Active Directory</b>
- <b>Oracle Virtual Box</b>
- <b>RAS/NAT</b>
- <b>DHCP</b>
- <b>IAM</b>


<h2>Environments Used </h2>

- <b>Windows 10</b>
- <b>Windows 19</b>
<h2>Program walk-through:</h2>

<p align="center">
Download and Spin up the Windows 10 and Windows 2019 VMs: <br/>

<br />
<br />
Set up IP Addressing on Domain Controller (Windows 2019 vm). Two NICs are required. One will be connecting to the internet, and the other wil be our internal network. The internet IP will be dynamically assigned, and we will manually set up our internal network IP:  <br/>

<br />
<br />
Rename PC through Start Menu > System > Rename PC: <br/>

<br />
<br />
Install Active Directory Domain Services:  <br/>

<br />
<br />
Configure the Post Deployment Settings:  <br/>

<br />
<br />
Create a Dedicated Domain Admin Account:  <br/>

<br />
<br />
Create an Organizational Unit:  <br/>

<br />
<br />
Create a New User: <br/>

<br />
<br />
Make a Domain Admin:  <br/>

<br />
<br />
Configure Routing and Remote Access:  <br/>

<br />
<br />
Install NAT:  <br/>

<br />
<br />
Set Up a DHCP Server on the Domain Controller: <br/>

<br />
<br />
DHCP:  <br/>

<br />
<br />
Configure Lease Duration:  <br/>

<br />
<br />
Configure DHCP Options:  <br/>

<br />
<br />
Add an IP Address for a Router: <br/>

<br />
<br />
Source Code for the Powershell Script:  <br/>

<br />
<br />
Start the Windows 10 VM which will serve as our End-User:  <br/>

<br />
<br />
Complete the User Config:  <br/>

<br />
<br />
Add the Domain Controller: <br/>

<br />
<br />
Address Leases:  <br/>

<br />
<br />
Active Directory Users and Computers:  <br/>

<br />
<br />
</p>

<!--
 ```diff
- text in red
+ text in green
! text in orange
# text in gray
@@ text in purple (and bold)@@
```
--!>
