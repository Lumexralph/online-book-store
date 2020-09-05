import * as jspb from "google-protobuf"

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';

export class Product extends jspb.Message {
  getId(): number;
  setId(value: number): Product;

  getUuid(): string;
  setUuid(value: string): Product;

  getName(): string;
  setName(value: string): Product;

  getDescription(): string;
  setDescription(value: string): Product;

  getPrice(): number;
  setPrice(value: number): Product;

  getSlug(): string;
  setSlug(value: string): Product;

  getInactive(): boolean;
  setInactive(value: boolean): Product;

  getQuantity(): number;
  setQuantity(value: number): Product;

  getImageUrl(): string;
  setImageUrl(value: string): Product;

  getCategoriesList(): Array<Category>;
  setCategoriesList(value: Array<Category>): Product;
  clearCategoriesList(): Product;
  addCategories(value?: Category, index?: number): Category;

  getCreatedAt(): Timestamp | undefined;
  setCreatedAt(value?: Timestamp): Product;
  hasCreatedAt(): boolean;
  clearCreatedAt(): Product;

  getUpdatedAt(): Timestamp | undefined;
  setUpdatedAt(value?: Timestamp): Product;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): Product;

  getDeletedAt(): Timestamp | undefined;
  setDeletedAt(value?: Timestamp): Product;
  hasDeletedAt(): boolean;
  clearDeletedAt(): Product;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Product.AsObject;
  static toObject(includeInstance: boolean, msg: Product): Product.AsObject;
  static serializeBinaryToWriter(message: Product, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Product;
  static deserializeBinaryFromReader(message: Product, reader: jspb.BinaryReader): Product;
}

export namespace Product {
  export type AsObject = {
    id: number,
    uuid: string,
    name: string,
    description: string,
    price: number,
    slug: string,
    inactive: boolean,
    quantity: number,
    imageUrl: string,
    categoriesList: Array<Category.AsObject>,
    createdAt?: Timestamp.AsObject,
    updatedAt?: Timestamp.AsObject,
    deletedAt?: Timestamp.AsObject,
  }
}

export class Category extends jspb.Message {
  getId(): number;
  setId(value: number): Category;

  getName(): string;
  setName(value: string): Category;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Category.AsObject;
  static toObject(includeInstance: boolean, msg: Category): Category.AsObject;
  static serializeBinaryToWriter(message: Category, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Category;
  static deserializeBinaryFromReader(message: Category, reader: jspb.BinaryReader): Category;
}

export namespace Category {
  export type AsObject = {
    id: number,
    name: string,
  }
}

export class Timestamp extends jspb.Message {
  getTimestamp(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTimestamp(value?: google_protobuf_timestamp_pb.Timestamp): Timestamp;
  hasTimestamp(): boolean;
  clearTimestamp(): Timestamp;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Timestamp.AsObject;
  static toObject(includeInstance: boolean, msg: Timestamp): Timestamp.AsObject;
  static serializeBinaryToWriter(message: Timestamp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Timestamp;
  static deserializeBinaryFromReader(message: Timestamp, reader: jspb.BinaryReader): Timestamp;
}

export namespace Timestamp {
  export type AsObject = {
    timestamp?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class AddProductRequest extends jspb.Message {
  getProduct(): Product | undefined;
  setProduct(value?: Product): AddProductRequest;
  hasProduct(): boolean;
  clearProduct(): AddProductRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddProductRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddProductRequest): AddProductRequest.AsObject;
  static serializeBinaryToWriter(message: AddProductRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddProductRequest;
  static deserializeBinaryFromReader(message: AddProductRequest, reader: jspb.BinaryReader): AddProductRequest;
}

export namespace AddProductRequest {
  export type AsObject = {
    product?: Product.AsObject,
  }
}

export class AddProductResponse extends jspb.Message {
  getCreatedProduct(): Product | undefined;
  setCreatedProduct(value?: Product): AddProductResponse;
  hasCreatedProduct(): boolean;
  clearCreatedProduct(): AddProductResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddProductResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddProductResponse): AddProductResponse.AsObject;
  static serializeBinaryToWriter(message: AddProductResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddProductResponse;
  static deserializeBinaryFromReader(message: AddProductResponse, reader: jspb.BinaryReader): AddProductResponse;
}

export namespace AddProductResponse {
  export type AsObject = {
    createdProduct?: Product.AsObject,
  }
}

