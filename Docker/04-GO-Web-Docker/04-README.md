- whichever folder your `dockerfile` is the Image is going to be built on that level of Folder and anything down 

- `docker build -t rpfeb23/first-golang-docker .` Execute this to build your own image named as `rpfeb23/first-golang-docker` you can give anyname. This command will grab Dockerfile and get `docker/whalesay` image first then get the `fortunes` app and in the CMD feed output of `fortunes` to `COWSAY`

- Now run `docker images` and you should see the image named `rpfeb23/first-golang-docker` created. Also notice `docker/whalesay` image is also downloaded

- Run the New Image `docker run -d -p 80:80 rpfeb23/first-golang-docker`

  -d is running as its own daemon
  
  -p is port mapping physical machine:container port (80:80)
  
  ---------
  `docker login`
  `docker push rpfeb23/first-golang-docker`

