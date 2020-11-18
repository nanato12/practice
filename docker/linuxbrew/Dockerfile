FROM ubuntu:18.04

# Environment variable
ENV DEBIAN_FRONTEND noninteractive
ENV USER nanato12
ENV HOME /home/${USER}
ENV SHELL /bin/bash
ENV PW password

# Update and install pakckages
RUN apt-get update
RUN apt-get install -y --no-install-recommends sudo \
    ca-certificates \
    build-essential \
    locales \
    curl \
    file \
    git

#　localの設定
RUN locale-gen en_US.UTF-8

# Add normal User
RUN useradd -m ${USER}
# Add sudo
RUN gpasswd -a ${USER} sudo
# Set password
RUN echo "${USER}:${PW}" | chpasswd
# Login shell
RUN sed -i.bak -r s#${HOME}:\(.+\)#${HOME}:${SHELL}# /etc/passwd

# Set user and home dirctory
USER ${USER}
WORKDIR ${HOME}

# Set Linuxbrew PATH
RUN echo 'export PATH=/home/linuxbrew/.linuxbrew/bin:$PATH' >> ${HOME}/.bashrc
