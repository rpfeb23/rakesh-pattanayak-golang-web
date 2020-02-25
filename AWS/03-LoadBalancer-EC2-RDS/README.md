GOOS=linux GOARCH=amd64 go build -o ALBEC2RDS

ssh -i ~/Documents/AWS/AWS-Solutions-Architect-Associate/AWS-Certified-Architect\ Asscoaite/myNorthvirginia1stkeypair.pem ec2-user@ec2-54-174-44-45.compute-1.amazonaws.com

sudo yum update

mkdir LoadBalancerEC2RDS

scp -i ~/Documents/AWS/AWS-Solutions-Architect-Associate/AWS-Certified-Architect\ Asscoaite/myNorthvirginia1stkeypair.pem -r Templates ec2-user@ec2-3-90-187-40.compute-1.amazonaws.com:LoadBalancerEC2RDS

scp -i ~/Documents/AWS/AWS-Solutions-Architect-Associate/AWS-Certified-Architect\ Asscoaite/myNorthvirginia1stkeypair.pem ALBEC2RDS ec2-user@ec2-54-174-44-45.compute-1.amazonaws.com:LoadBalancerEC2RDS

chmod 700 ALBEC2RDS 

cd /etc/systemd/system/

sudo nano rakesh.service



2nd Instance
AWS Terminal
ssh -i ~/Documents/AWS/AWS-Solutions-Architect-Associate/AWS-Certified-Architect\ Asscoaite/myNorthvirginia1stkeypair.pem ec2-user@ec2-54-167-127-231.compute-1.amazonaws.com

sudo yum update

mkdir LoadBalancerEC2RDS

cd LoadBalancerEC2RDS/

GO Terminal

scp -i ~/Documents/AWS/AWS-Solutions-Architect-Associate/AWS-Certified-Architect\ Asscoaite/myNorthvirginia1stkeypair.pem ALBEC2RDS ec2-user@ec2-54-167-127-231.compute-1.amazonaws.com:LoadBalancerEC2RDS

scp -i ~/Documents/AWS/AWS-Solutions-Architect-Associate/AWS-Certified-Architect\ Asscoaite/myNorthvirginia1stkeypair.pem -r Templates ec2-user@ec2-3-80-200-20.compute-1.amazonaws.com:LoadBalancerEC2RDS



chmod 700 ALBEC2RDS 

cd /etc/systemd/system/

sudo nano rakesh.service

`[Unit]
    Description=Go Server
    
    [Service]
    ExecStart=/home/ec2-user/LoadBalancerEC2RDS/ALBEC2RDS
     WorkingDirectory=/home/ec2-user/LoadBalancerEC2RDS
 
    User=root
    Group=root
    Restart=always
    
    [Install]
    WantedBy=multi-user.target
`

---------

sudo systemctl enable rakesh.service

sudo systemctl start rakesh.service

sudo systemctl status rakesh.service

-----
Run in Local

1. Change ListenandServe to :8080
2. Hardcode EC2instanceid = "Dummy"
3. Comment the EC2instanceinfo func call
4. RDS Database: Make Public Access YES
5. RDS Database : Add Security Group WebDMZ

----------
Before Build and Putting in AWS

1. Change ListenandServe to :80
2. Comment the Hardcode EC2instanceid = "Dummy"
3. UnComment the EC2instanceinfo func call
4. RDS Database: Make Public Access NO
5. RDS Database : REMOVE Security Group WebDMZ

-----
Two Key Security groups

sg1 : Allow Access from HTTP 80 attach to the ALB

sg2 : 
    - Add HTTP 80 and Source as sg1
    - Add SSH from myIP/ Bastion Host
    - Add MYSQL/Auroral 3306 and Source as sg2
    - Attach all EC2 and RDS 
    
-----------
For re-build
1. SSH to EC2 Instances
2. Remove the Binary
3. COpy the Binary from Local to EC2. Copy Templates if changed
4. Reboot the EC2 Instances
(Above works for me)

you can try
a. remove Binary in EC2
b. sudo systemctl stop rakesh.service
c. copy the Binary from local to EC2
3. chmod 700 Binary
4. sudo systemctl daemon-reload
5. sudo systemctl start rakesh.service
6. sudo systemctl status rakesh.service



