- whichever folder your `dockerfile` is the Image is going to be built on that level of Folder and anything down 

- `docker build -t rpfeb23/docker-whalesay-fortunes .` Execute this to build your own image named as `rpfeb23/docker-whalesay-fortunes` you can give anyname. This command will grab Dockerfile and get `docker/whalesay` image first then get the `fortunes` app and in the CMD feed output of `fortunes` to `COWSAY`

- Now run `docker images` and you should see the image named `rpfeb23/docker-whalesay-fortunes` created. Also notice `docker/whalesay` image is also downloaded

- Run the New Image `docker run rpfeb23/docker-whalesay-fortunes`

