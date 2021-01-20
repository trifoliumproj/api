// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var nhentai_pb = require('./nhentai_pb.js');

function serialize_nhentai_GalleryRequest(arg) {
  if (!(arg instanceof nhentai_pb.GalleryRequest)) {
    throw new Error('Expected argument of type nhentai.GalleryRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_nhentai_GalleryRequest(buffer_arg) {
  return nhentai_pb.GalleryRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_nhentai_GalleryResponse(arg) {
  if (!(arg instanceof nhentai_pb.GalleryResponse)) {
    throw new Error('Expected argument of type nhentai.GalleryResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_nhentai_GalleryResponse(buffer_arg) {
  return nhentai_pb.GalleryResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var APIService = exports.APIService = {
  gallery: {
    path: '/nhentai.API/Gallery',
    requestStream: false,
    responseStream: false,
    requestType: nhentai_pb.GalleryRequest,
    responseType: nhentai_pb.GalleryResponse,
    requestSerialize: serialize_nhentai_GalleryRequest,
    requestDeserialize: deserialize_nhentai_GalleryRequest,
    responseSerialize: serialize_nhentai_GalleryResponse,
    responseDeserialize: deserialize_nhentai_GalleryResponse,
  },
};

exports.APIClient = grpc.makeGenericClientConstructor(APIService);
