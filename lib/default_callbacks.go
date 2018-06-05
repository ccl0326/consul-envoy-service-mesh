package lib

import (
	"github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/golang/glog"
)

type DefaultCallbacks struct{}

func (c DefaultCallbacks) OnStreamOpen(id int64, url string) {
	glog.V(10).Infof("on stream open, stream id: %d, type url: %s", id, url)
}

func (c DefaultCallbacks) OnStreamClosed(id int64) {
	glog.V(10).Infof("on stream closed, stream id: %d", id)
}

func (c DefaultCallbacks) OnStreamRequest(id int64, request *v2.DiscoveryRequest) {
	glog.V(10).Infof("on stream request, stream id: %d, request: %v", id, *request)
}

func (c DefaultCallbacks) OnStreamResponse(id int64, request *v2.DiscoveryRequest, response *v2.DiscoveryResponse) {
	glog.V(10).Infof("on stream response, stream id: %d, request: %v, response: %v", id, *request, *response)
}

func (c DefaultCallbacks) OnFetchRequest(request *v2.DiscoveryRequest) {
	glog.V(10).Infof("on fetch request, request: %v", *request)
}

func (c DefaultCallbacks) OnFetchResponse(request *v2.DiscoveryRequest, response *v2.DiscoveryResponse) {
	glog.V(10).Infof("on fetch response, request: %v, response: %v", *request, *response)
}
