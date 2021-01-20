// package: hentai_cafe
// file: hentai-cafe.proto

import * as jspb from "google-protobuf";

export class GalleryRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GalleryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GalleryRequest): GalleryRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GalleryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GalleryRequest;
  static deserializeBinaryFromReader(message: GalleryRequest, reader: jspb.BinaryReader): GalleryRequest;
}

export namespace GalleryRequest {
  export type AsObject = {
    id: number,
  }
}

export class GalleryResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GalleryResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GalleryResponse): GalleryResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GalleryResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GalleryResponse;
  static deserializeBinaryFromReader(message: GalleryResponse, reader: jspb.BinaryReader): GalleryResponse;
}

export namespace GalleryResponse {
  export type AsObject = {
  }
}

