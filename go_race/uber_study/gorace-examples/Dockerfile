FROM ubuntu:20.04
ENV TZ=America/Los_Angeles
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt-get update
RUN apt-get install -y build-essential git-core vim cmake g++ golang
CMD [mkdir /datarace]
WORKDIR /datarace
COPY . .
