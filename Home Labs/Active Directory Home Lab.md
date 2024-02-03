<h1>Active Directory Home Lab</h1>

 ### youtube link

<h2>Description</h2>
<b>This project is creating a simulated large scale Active Directory environment. We accomplish this by running two Windows virtual machines using Oracle Virtual Box. One of the machines will act as the domain controller and the other acts as an individual client accessing the network through the internal network.  In order to simulate the large environment, we will be using a Powershell script to create 1,000 users. This exercise includes Domain /AD DS, RAS/NAT, DHCP (1 Scope), Domain Controller setup, VMWare Network, and client access of the internet through domain controller on an AD environment.
<br/>


<h2>Languages and Utilities Used</h2>

- <b>Active Directory</b>
- <b>Oracle Virtual Box</b>
- <b>RAS/NAT</b>
- <b>PowerShell</b> 
- <b>DHCP</b>
- <b>IAM</b>


<h2>Environments Used </h2>

- <b>Windows 10</b>
- <b>Windows 19</b>
<h2>Program walk-through:</h2>

<p align="center">
Download and Spin up the Windows 10 and Windows 2019 VMs: <br/>
 ![windows 19 download 1](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/d3479784-4d60-452a-9945-3857275e6faf)
 ![windows 10 download 21](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/f9d43c5f-01ea-423d-af35-e377f18cb53d)

<br />
<br />
Set up IP Addressing on Domain Controller (Windows 2019 vm). Two NICs are required. One will be connecting to the internet, and the other wil be our internal network. The internet IP will be dynamically assigned, and we will manually set up our internal network IP:  <br/>
 ![DC Internet Settings 2](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/b3de260e-ce9e-4460-99ce-46c4c0690d4e)
 ![DC Internet Identified Home IP Address 3](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/cd489856-2595-496e-8b80-2da05f4fabf7)
 ![DC Internet Identified Internal Network 4](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/7d921b10-8ca0-44bf-923e-efb99347a81e)

<br />
<br />
Rename PC through Start Menu > System > Rename PC: <br/>
![Renaming PC to DC1 6](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/818e6b04-2a21-49c1-9ac3-4f38ca3a9eba)

<br />
<br />
Set IP Address for Domain Controller Internal Network: <br/>
![Set IP Address for Domain Controller Internal Network 5](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/a6587e4b-5116-49b3-a8d4-c7f44de3fa72)

<br />
<br />
Install Active Directory Domain Services:  <br/>
![Downloading AD and Selecting Domain 7](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/940fcc44-b0bf-4dce-8340-988335a5062d)

<br />
<br />
Configure the Post Deployment Settings:  <br/>
![Post Deployment Configuration 8](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/e4c1dfc5-c991-4a7e-9aad-aef9493e9c41)
![Post Deployment Configuration 2 9](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/d568f5b4-6a7b-4d89-93b9-bafaf0a304ff)

<br />
<br />
Create an Organization Unit Named ADMINS, Create a New User and Make it the Domain Admin Account:  <br/>
![Organizational Unit and New Admin User Creation 10](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/9d1afc4f-29cf-4b3d-b63d-30645a6660b4)

<br />
<br />

Logging in as New Admin Account: <br/>
![Logging in as New Admin User 11](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/8e62d8b0-2b35-403c-a378-e4ad72389fcd)

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
