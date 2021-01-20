// GENERATED CODE -- DO NOT EDIT!

// package: hentai_cafe
// file: hentai-cafe.proto

import * as hentai_cafe_pb from "./hentai-cafe_pb";
import * as grpc from "grpc";

interface IAPIService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  gallery: grpc.MethodDefinition<hentai_cafe_pb.GalleryRequest, hentai_cafe_pb.GalleryResponse>;
}

export const APIService: IAPIService;

export class APIClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  gallery(argument: hentai_cafe_pb.GalleryRequest, callback: grpc.requestCallback<hentai_cafe_pb.GalleryResponse>): grpc.ClientUnaryCall;
  gallery(argument: hentai_cafe_pb.GalleryRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<hentai_cafe_pb.GalleryResponse>): grpc.ClientUnaryCall;
  gallery(argument: hentai_cafe_pb.GalleryRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<hentai_cafe_pb.GalleryResponse>): grpc.ClientUnaryCall;
}
