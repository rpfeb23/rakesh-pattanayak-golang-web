- whichever folder your `Dockerfile` is the Image is going to be built on that level of Folder and anything down 

- `docker build -t rpfeb23/docker-curl .` Execute this to build your own image named as `rpfeb23/docker-curl` you can give anyname. 

- Now run `docker images` and you should see the image named `rpfeb23/docker-curl` created. 

- Run the New Image Interactively  `docker run -it dockerimagename`

    -   `curl --head www.google.com`