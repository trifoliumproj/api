// package: nhentai
// file: nhentai.proto

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
  getId(): number;
  setId(value: number): void;

  getMediaId(): number;
  setMediaId(value: number): void;

  hasTitle(): boolean;
  clearTitle(): void;
  getTitle(): GalleryResponse.Title | undefined;
  setTitle(value?: GalleryResponse.Title): void;

  hasImages(): boolean;
  clearImages(): void;
  getImages(): GalleryResponse.Images | undefined;
  setImages(value?: GalleryResponse.Images): void;

  getScanlator(): string;
  setScanlator(value: string): void;

  getUploadDate(): string;
  setUploadDate(value: string): void;

  clearTagsList(): void;
  getTagsList(): Array<GalleryResponse.Tag>;
  setTagsList(value: Array<GalleryResponse.Tag>): void;
  addTags(value?: GalleryResponse.Tag, index?: number): GalleryResponse.Tag;

  getNumPages(): number;
  setNumPages(value: number): void;

  getNumFavorites(): number;
  setNumFavorites(value: number): void;

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
    id: number,
    mediaId: number,
    title?: GalleryResponse.Title.AsObject,
    images?: GalleryResponse.Images.AsObject,
    scanlator: string,
    uploadDate: string,
    tagsList: Array<GalleryResponse.Tag.AsObject>,
    numPages: number,
    numFavorites: number,
  }

  export class Title extends jspb.Message {
    getJapanese(): string;
    setJapanese(value: string): void;

    getEnglish(): string;
    setEnglish(value: string): void;

    getPretty(): string;
    setPretty(value: string): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Title.AsObject;
    static toObject(includeInstance: boolean, msg: Title): Title.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Title, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Title;
    static deserializeBinaryFromReader(message: Title, reader: jspb.BinaryReader): Title;
  }

  export namespace Title {
    export type AsObject = {
      japanese: string,
      english: string,
      pretty: string,
    }
  }

  export class Images extends jspb.Message {
    clearPagesList(): void;
    getPagesList(): Array<GalleryResponse.Images.Image>;
    setPagesList(value: Array<GalleryResponse.Images.Image>): void;
    addPages(value?: GalleryResponse.Images.Image, index?: number): GalleryResponse.Images.Image;

    hasCover(): boolean;
    clearCover(): void;
    getCover(): GalleryResponse.Images.Image | undefined;
    setCover(value?: GalleryResponse.Images.Image): void;

    hasThumbnail(): boolean;
    clearThumbnail(): void;
    getThumbnail(): GalleryResponse.Images.Image | undefined;
    setThumbnail(value?: GalleryResponse.Images.Image): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Images.AsObject;
    static toObject(includeInstance: boolean, msg: Images): Images.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Images, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Images;
    static deserializeBinaryFromReader(message: Images, reader: jspb.BinaryReader): Images;
  }

  export namespace Images {
    export type AsObject = {
      pagesList: Array<GalleryResponse.Images.Image.AsObject>,
      cover?: GalleryResponse.Images.Image.AsObject,
      thumbnail?: GalleryResponse.Images.Image.AsObject,
    }

    export class Image extends jspb.Message {
      getT(): string;
      setT(value: string): void;

      getW(): number;
      setW(value: number): void;

      getH(): number;
      setH(value: number): void;

      serializeBinary(): Uint8Array;
      toObject(includeInstance?: boolean): Image.AsObject;
      static toObject(includeInstance: boolean, msg: Image): Image.AsObject;
      static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
      static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
      static serializeBinaryToWriter(message: Image, writer: jspb.BinaryWriter): void;
      static deserializeBinary(bytes: Uint8Array): Image;
      static deserializeBinaryFromReader(message: Image, reader: jspb.BinaryReader): Image;
    }

    export namespace Image {
      export type AsObject = {
        t: string,
        w: number,
        h: number,
      }
    }
  }

  export class Tag extends jspb.Message {
    getId(): number;
    setId(value: number): void;

    getType(): string;
    setType(value: string): void;

    getName(): string;
    setName(value: string): void;

    getUrl(): string;
    setUrl(value: string): void;

    getCount(): number;
    setCount(value: number): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Tag.AsObject;
    static toObject(includeInstance: boolean, msg: Tag): Tag.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Tag, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Tag;
    static deserializeBinaryFromReader(message: Tag, reader: jspb.BinaryReader): Tag;
  }

  export namespace Tag {
    export type AsObject = {
      id: number,
      type: string,
      name: string,
      url: string,
      count: number,
    }
  }
}

