// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var hentai$cafe_pb = require('./hentai-cafe_pb.js');

function serialize_hentai_cafe_GalleryRequest(arg) {
  if (!(arg instanceof hentai$cafe_pb.GalleryRequest)) {
    throw new Error('Expected argument of type hentai_cafe.GalleryRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_hentai_cafe_GalleryRequest(buffer_arg) {
  return hentai$cafe_pb.GalleryRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_hentai_cafe_GalleryResponse(arg) {
  if (!(arg instanceof hentai$cafe_pb.GalleryResponse)) {
    throw new Error('Expected argument of type hentai_cafe.GalleryResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_hentai_cafe_GalleryResponse(buffer_arg) {
  return hentai$cafe_pb.GalleryResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var APIService = exports.APIService = {
  gallery: {
    path: '/hentai_cafe.API/Gallery',
    requestStream: false,
    responseStream: false,
    requestType: hentai$cafe_pb.GalleryRequest,
    responseType: hentai$cafe_pb.GalleryResponse,
    requestSerialize: serialize_hentai_cafe_GalleryRequest,
    requestDeserialize: deserialize_hentai_cafe_GalleryRequest,
    responseSerialize: serialize_hentai_cafe_GalleryResponse,
    responseDeserialize: deserialize_hentai_cafe_GalleryResponse,
  },
};

exports.APIClient = grpc.makeGenericClientConstructor(APIService);
