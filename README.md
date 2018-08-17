
# terraform-provider-sampleclassis

This project is providing the code for talks and a blog post on custom providers. 

## Building with drone as long as you have the cli

```sh
drone exec --repo-name terraform-provider-sampleclassis
```

##Dockerfile for including the pluging

Once the binaries are built simply run 

```sh
docker build -t baniol/terraform-sampleclassis .
docker run -i -t -v $(pwd):/app/ -w /app/ baniol/terraform-sampleclassis init -plugin-dir=/usr/local/terraform-plugins 
docker run -i -t -v $(pwd):/app/ -w /app/ baniol/terraform-sampleclassis plan
```

This should create a docker image that includes your custom plugin for other folks to use :)

