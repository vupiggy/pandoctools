all: image container

image:
	podman build -f ./Dockerfile.`grep "^NAME=" /etc/os-release | cut -d '=' -f2` \
                     -t toolbox-doc --rm --force-rm --layers=false

container: image
	toolbox create --container toolbox-doc --image toolbox-doc --hostname toolbox-doc
