FROM iron/go:dev

RUN apk update && \
    apk add openssh gcc libc-dev cyrus-sasl-dev linux-headers python-dev automake autoconf inotify-tools py2-pip openssl-dev

RUN mkdir -p ~/.ssh && ssh-keyscan -H github.com >> ~/.ssh/known_hosts && chmod 644 ~/.ssh/known_hosts

RUN git clone https://github.com/facebook/watchman.git && \
    cd watchman && \
    git checkout tags/v4.7.0 && \
    ./autogen.sh && \
    ./configure && \
    make && \
    make install && \
    cd .. && \
    rm -rf watchman && \
    pip install --upgrade setuptools && \
    pip install pywatchman

RUN go get -u github.com/kardianos/govendor

WORKDIR /go/src/github.com/alecholmez/http-server
COPY . /go/src/github.com/alecholmez/http-server
# 
# CMD echo "Vendoring..." && \
#     git config --global url."git@github.com:".insteadOf "https://github.com/" && \
#     /go/bin/govendor sync && \
#     /go/bin/govendor generate +local
#    ./autobuild.sh mongo:27017

EXPOSE 6060
EXPOSE 9000
