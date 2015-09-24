FROM golang:1.5

ENV APPROOT /go/src/github.com/zakuro9715/AOJRESTAPI

WORKDIR $APPROOT
RUN mkdir -p $APPROOT
COPY . $APPROOT

RUN go build && go install
CMD ["AOJRESTAPI"]
