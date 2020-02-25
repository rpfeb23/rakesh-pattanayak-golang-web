- Launch an EC2 Instance
- SSH to the EC2 Instance
- `sudo yum update -y`

### Install Docker on EC2 Virtual Machine

- `sudo yum install -y docker`

### Start Docker

- `sudo service docker start`

### Add ec2-user to docker group so you can execute docker commands without using sudo

-   `sudo usermod -a -G docker ec2-user`

### Logout and Logbackin to pick the new permisssion

-   `exit`
-   ssh

### Verify ec2-user can now run docker commands without sudo

-   `docker info`
- `docker version`

### Now run docker image in a container

`docker run -d -p 80:80 rpfeb23/first-golang-docker`

It will pull from dockerhub and run the container

-   'docker ps' will give you the containerID

### Go to your Public DNS of EC2 or Public IP in the browser and validate code is running