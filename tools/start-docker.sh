#! /bin/bash

local_project_dir="${HOME}/Projects"
docker_project_dir="/root/projects"
#local_ssh_dir="${HOME}/.ssh"
#docker_ssh_dir="/root/.ssh"
image="doctorwechat-dev:v1.1"

docker run -ti --privileged \
		   -v ${local_project_dir}:${docker_project_dir} \
		   #-v ${local_ssh_dir}:${docker_ssh_dir} \
		   ${image} /bin/bash