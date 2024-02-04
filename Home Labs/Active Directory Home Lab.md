<h1>Active Directory Home Lab</h1>

<h2>Description</h2>
<b>This project is creating a simulated large scale corporate Active Directory environment. This is accomplished by running two Windows virtual machines using Oracle VirtualBox. One of the machines will act as the domain controller and the other acts as an individual client/end-user accessing the internet through the internal network.  In order to simulate the large environment, we will be using a Powershell script to create 1,000 users. This exercise includes AD DS, RAS/NAT, DHCP (1 Scope), Domain Controller setup, VM Network, and more.
<br/>
 
![AD Home Lab Architecture Graphic](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/14325a6a-36d2-465a-8e99-9866658ca53a)


<h2>Languages and Utilities Used</h2>

- <b>Active Directory</b>
- <b>Oracle VirtualBox</b>
- <b>RAS/NAT</b>
- <b>PowerShell</b> 
- <b>DHCP</b>
- <b>IAM</b>


<h2>Environments Used </h2>

- <b>Windows 10</b>
- <b>Windows 19</b>
<h2>Corporate AD Simulation Walk-Through:</h2>

<p align="center">
Download and Spin up a Windows 2019 VM to Serve as our Domain Controller Named DC1: <br/>
 
![windows 19 download 1](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/d3479784-4d60-452a-9945-3857275e6faf)

<br />
<br />
Set up IP Addressing on Domain Controller (Windows 2019 vm). Two NICs are Required. One Will be Connecting to the Internet, and the Other wil be our Internal Network. The Internet IP will be Dynamically Assigned, and We will Manually sSet Up our Internal Network IP:  <br/>

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

![Adding Routing and Remote Acces 12](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/c4067562-98b0-44a5-a1ec-07b0dd55ae2c)

<br />
<br />
Install NAT:  <br/>

![Successfully Configured NAT 13](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/cf6f1f1b-fbdd-49e0-b979-eeb05a590d75)

<br />
<br />
Set Up a DHCP Server on the Domain Controller: <br/>

![Installing DHCP 14](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/d13b700f-717d-42cb-be22-861b544eaf05)

<br />
<br />
Adding a Scope to IPv4, including Lease Duration, DHCP Options, and adding an IP Addres for the Router:  <br/>

![Adding Scope to IPv4 15](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/1df77e45-9406-4f54-b46c-f42c9020ba39)

<br />
<br />
Coding Powershell Script to Generate 1,000+ Users:  <br/>

![Full Powershell Script 18](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/ce5941bf-5939-4dcb-9b0b-03eba4fa37b8)

<br />
<br />
Name List Randomly Generated in .names for the Script to Pull from: <br/>

![Name list for Powershell Script 16](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/af0fff7a-9ab8-4803-bbb4-6ffa6af90954)

<br />
<br />
Set Powershell ISE to Unrestricted: <br/>

![Powershell ISE Set Execution to Unrestricted 17](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/8f986bee-4552-48d4-8cd5-d2ee22b579aa)

<br />
<br />
Powershell User Creation Script in Process: <br/>

![Powershell Script Creating Users 19](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/3e429430-08cd-4c83-bf99-f2ae61826a63)

<br />
<br />
Users Successfully Added to _USERS OU: <br/>

![Users Successfully Added to _USERS OU in AD 20](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/0cb3a5a1-9270-405e-93ee-6583e5cf73cc)

<br />
<br />
Download and Spin Up a Windows 10 VM which will serve as our End-User Named CLIENT1:  <br/>

![windows 10 download 21](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/9fd4f67a-299d-4f06-ab0a-c5f924bc2110)

<br />
<br />
Confirming Internet Access on CLIENT1 VM through our Internal Network:  <br/>

![22 Confirming Internet Access through the DC Default Gateway on CLIENT1 VM](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/26ad7728-5b20-4753-ba79-e5929e33bc03)

<br />
<br />
Renaming Client PC and Joining the Domain. <br/>

![23 renaming client VM and joining domain](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/d98992da-5d11-4533-b407-dd1db3fc6b40)

<br />
<br />
Add the Domain Controller Using the Admin Credentials we made. In this case a-CodyRochester: <br/>

![24 Joining Domain Using a-CodyRochester admin account](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/6126ca4c-a2fe-456c-8876-b598881d2911)

<br />
<br />
Select Random User to Test Login. Antonio Bargo. Username "abargo" Password "Password1" Due to the Script Setting the Same Password for All Users it Created :  <br/>

![25 Select Random User to Log in from CLIENT1](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/a30910a8-c86c-47b0-8900-fa18d8bdf128)

<br />
<br />
Logging in as "abargo": <br/>

![26 Logging in as abargo User](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/ddc6fe2b-fa7e-47c4-92b3-132628fe7c70)

<br />
<br />
CLIENT1 Using the DHCP through the Domain Controller to Access External Internet. This is also Visible on the "Computers" Tab in the AD Users and Computers Menu:  <br/>

![27 CLIENT1 Using the DHCP through the DC to Access the Internet](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/9aba46bd-2dc5-4bc7-9475-5f6abc2906d1)

<br />
<br />

User "abargo" Accessing the Internet through NAT/DHCP/mydomain.com: <br/>

![28 Accessing External Internet through DHCP Connection on CLIENT1](https://github.com/Cody-Rochester/Cybersecurity/assets/107632714/900f0319-e50a-4b12-b97c-2e4085b2b22d)

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
