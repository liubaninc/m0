// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: m0/storage/share_encrypt_storage.proto

package liubaninc.m0.storage;

public final class ShareEncryptStorageOuterClass {
  private ShareEncryptStorageOuterClass() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface ShareEncryptStorageOrBuilder extends
      // @@protoc_insertion_point(interface_extends:liubaninc.m0.storage.ShareEncryptStorage)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <code>string creator = 1;</code>
     */
    java.lang.String getCreator();
    /**
     * <code>string creator = 1;</code>
     */
    com.google.protobuf.ByteString
        getCreatorBytes();

    /**
     * <code>string index = 2;</code>
     */
    java.lang.String getIndex();
    /**
     * <code>string index = 2;</code>
     */
    com.google.protobuf.ByteString
        getIndexBytes();

    /**
     * <code>string shareIndex = 3;</code>
     */
    java.lang.String getShareIndex();
    /**
     * <code>string shareIndex = 3;</code>
     */
    com.google.protobuf.ByteString
        getShareIndexBytes();

    /**
     * <code>string envelope = 4;</code>
     */
    java.lang.String getEnvelope();
    /**
     * <code>string envelope = 4;</code>
     */
    com.google.protobuf.ByteString
        getEnvelopeBytes();

    /**
     * <code>string sharer = 5;</code>
     */
    java.lang.String getSharer();
    /**
     * <code>string sharer = 5;</code>
     */
    com.google.protobuf.ByteString
        getSharerBytes();
  }
  /**
   * Protobuf type {@code liubaninc.m0.storage.ShareEncryptStorage}
   */
  public  static final class ShareEncryptStorage extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:liubaninc.m0.storage.ShareEncryptStorage)
      ShareEncryptStorageOrBuilder {
  private static final long serialVersionUID = 0L;
    // Use ShareEncryptStorage.newBuilder() to construct.
    private ShareEncryptStorage(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private ShareEncryptStorage() {
      creator_ = "";
      index_ = "";
      shareIndex_ = "";
      envelope_ = "";
      sharer_ = "";
    }

    @java.lang.Override
    public final com.google.protobuf.UnknownFieldSet
    getUnknownFields() {
      return this.unknownFields;
    }
    private ShareEncryptStorage(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      this();
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      int mutable_bitField0_ = 0;
      com.google.protobuf.UnknownFieldSet.Builder unknownFields =
          com.google.protobuf.UnknownFieldSet.newBuilder();
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 10: {
              java.lang.String s = input.readStringRequireUtf8();

              creator_ = s;
              break;
            }
            case 18: {
              java.lang.String s = input.readStringRequireUtf8();

              index_ = s;
              break;
            }
            case 26: {
              java.lang.String s = input.readStringRequireUtf8();

              shareIndex_ = s;
              break;
            }
            case 34: {
              java.lang.String s = input.readStringRequireUtf8();

              envelope_ = s;
              break;
            }
            case 42: {
              java.lang.String s = input.readStringRequireUtf8();

              sharer_ = s;
              break;
            }
            default: {
              if (!parseUnknownField(
                  input, unknownFields, extensionRegistry, tag)) {
                done = true;
              }
              break;
            }
          }
        }
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(this);
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(
            e).setUnfinishedMessage(this);
      } finally {
        this.unknownFields = unknownFields.build();
        makeExtensionsImmutable();
      }
    }
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return liubaninc.m0.storage.ShareEncryptStorageOuterClass.internal_static_liubaninc_m0_storage_ShareEncryptStorage_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return liubaninc.m0.storage.ShareEncryptStorageOuterClass.internal_static_liubaninc_m0_storage_ShareEncryptStorage_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage.class, liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage.Builder.class);
    }

    public static final int CREATOR_FIELD_NUMBER = 1;
    private volatile java.lang.Object creator_;
    /**
     * <code>string creator = 1;</code>
     */
    public java.lang.String getCreator() {
      java.lang.Object ref = creator_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        creator_ = s;
        return s;
      }
    }
    /**
     * <code>string creator = 1;</code>
     */
    public com.google.protobuf.ByteString
        getCreatorBytes() {
      java.lang.Object ref = creator_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        creator_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int INDEX_FIELD_NUMBER = 2;
    private volatile java.lang.Object index_;
    /**
     * <code>string index = 2;</code>
     */
    public java.lang.String getIndex() {
      java.lang.Object ref = index_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        index_ = s;
        return s;
      }
    }
    /**
     * <code>string index = 2;</code>
     */
    public com.google.protobuf.ByteString
        getIndexBytes() {
      java.lang.Object ref = index_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        index_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int SHAREINDEX_FIELD_NUMBER = 3;
    private volatile java.lang.Object shareIndex_;
    /**
     * <code>string shareIndex = 3;</code>
     */
    public java.lang.String getShareIndex() {
      java.lang.Object ref = shareIndex_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        shareIndex_ = s;
        return s;
      }
    }
    /**
     * <code>string shareIndex = 3;</code>
     */
    public com.google.protobuf.ByteString
        getShareIndexBytes() {
      java.lang.Object ref = shareIndex_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        shareIndex_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int ENVELOPE_FIELD_NUMBER = 4;
    private volatile java.lang.Object envelope_;
    /**
     * <code>string envelope = 4;</code>
     */
    public java.lang.String getEnvelope() {
      java.lang.Object ref = envelope_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        envelope_ = s;
        return s;
      }
    }
    /**
     * <code>string envelope = 4;</code>
     */
    public com.google.protobuf.ByteString
        getEnvelopeBytes() {
      java.lang.Object ref = envelope_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        envelope_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int SHARER_FIELD_NUMBER = 5;
    private volatile java.lang.Object sharer_;
    /**
     * <code>string sharer = 5;</code>
     */
    public java.lang.String getSharer() {
      java.lang.Object ref = sharer_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        sharer_ = s;
        return s;
      }
    }
    /**
     * <code>string sharer = 5;</code>
     */
    public com.google.protobuf.ByteString
        getSharerBytes() {
      java.lang.Object ref = sharer_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        sharer_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    private byte memoizedIsInitialized = -1;
    @java.lang.Override
    public final boolean isInitialized() {
      byte isInitialized = memoizedIsInitialized;
      if (isInitialized == 1) return true;
      if (isInitialized == 0) return false;

      memoizedIsInitialized = 1;
      return true;
    }

    @java.lang.Override
    public void writeTo(com.google.protobuf.CodedOutputStream output)
                        throws java.io.IOException {
      if (!getCreatorBytes().isEmpty()) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 1, creator_);
      }
      if (!getIndexBytes().isEmpty()) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 2, index_);
      }
      if (!getShareIndexBytes().isEmpty()) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 3, shareIndex_);
      }
      if (!getEnvelopeBytes().isEmpty()) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 4, envelope_);
      }
      if (!getSharerBytes().isEmpty()) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 5, sharer_);
      }
      unknownFields.writeTo(output);
    }

    @java.lang.Override
    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      if (!getCreatorBytes().isEmpty()) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, creator_);
      }
      if (!getIndexBytes().isEmpty()) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(2, index_);
      }
      if (!getShareIndexBytes().isEmpty()) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(3, shareIndex_);
      }
      if (!getEnvelopeBytes().isEmpty()) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(4, envelope_);
      }
      if (!getSharerBytes().isEmpty()) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(5, sharer_);
      }
      size += unknownFields.getSerializedSize();
      memoizedSize = size;
      return size;
    }

    @java.lang.Override
    public boolean equals(final java.lang.Object obj) {
      if (obj == this) {
       return true;
      }
      if (!(obj instanceof liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage)) {
        return super.equals(obj);
      }
      liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage other = (liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage) obj;

      if (!getCreator()
          .equals(other.getCreator())) return false;
      if (!getIndex()
          .equals(other.getIndex())) return false;
      if (!getShareIndex()
          .equals(other.getShareIndex())) return false;
      if (!getEnvelope()
          .equals(other.getEnvelope())) return false;
      if (!getSharer()
          .equals(other.getSharer())) return false;
      if (!unknownFields.equals(other.unknownFields)) return false;
      return true;
    }

    @java.lang.Override
    public int hashCode() {
      if (memoizedHashCode != 0) {
        return memoizedHashCode;
      }
      int hash = 41;
      hash = (19 * hash) + getDescriptor().hashCode();
      hash = (37 * hash) + CREATOR_FIELD_NUMBER;
      hash = (53 * hash) + getCreator().hashCode();
      hash = (37 * hash) + INDEX_FIELD_NUMBER;
      hash = (53 * hash) + getIndex().hashCode();
      hash = (37 * hash) + SHAREINDEX_FIELD_NUMBER;
      hash = (53 * hash) + getShareIndex().hashCode();
      hash = (37 * hash) + ENVELOPE_FIELD_NUMBER;
      hash = (53 * hash) + getEnvelope().hashCode();
      hash = (37 * hash) + SHARER_FIELD_NUMBER;
      hash = (53 * hash) + getSharer().hashCode();
      hash = (29 * hash) + unknownFields.hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }
    public static liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }
    public static liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage parseFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    @java.lang.Override
    public Builder newBuilderForType() { return newBuilder(); }
    public static Builder newBuilder() {
      return DEFAULT_INSTANCE.toBuilder();
    }
    public static Builder newBuilder(liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage prototype) {
      return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
    }
    @java.lang.Override
    public Builder toBuilder() {
      return this == DEFAULT_INSTANCE
          ? new Builder() : new Builder().mergeFrom(this);
    }

    @java.lang.Override
    protected Builder newBuilderForType(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      Builder builder = new Builder(parent);
      return builder;
    }
    /**
     * Protobuf type {@code liubaninc.m0.storage.ShareEncryptStorage}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:liubaninc.m0.storage.ShareEncryptStorage)
        liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorageOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return liubaninc.m0.storage.ShareEncryptStorageOuterClass.internal_static_liubaninc_m0_storage_ShareEncryptStorage_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return liubaninc.m0.storage.ShareEncryptStorageOuterClass.internal_static_liubaninc_m0_storage_ShareEncryptStorage_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage.class, liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage.Builder.class);
      }

      // Construct using liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage.newBuilder()
      private Builder() {
        maybeForceBuilderInitialization();
      }

      private Builder(
          com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
        super(parent);
        maybeForceBuilderInitialization();
      }
      private void maybeForceBuilderInitialization() {
        if (com.google.protobuf.GeneratedMessageV3
                .alwaysUseFieldBuilders) {
        }
      }
      @java.lang.Override
      public Builder clear() {
        super.clear();
        creator_ = "";

        index_ = "";

        shareIndex_ = "";

        envelope_ = "";

        sharer_ = "";

        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return liubaninc.m0.storage.ShareEncryptStorageOuterClass.internal_static_liubaninc_m0_storage_ShareEncryptStorage_descriptor;
      }

      @java.lang.Override
      public liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage getDefaultInstanceForType() {
        return liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage.getDefaultInstance();
      }

      @java.lang.Override
      public liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage build() {
        liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage buildPartial() {
        liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage result = new liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage(this);
        result.creator_ = creator_;
        result.index_ = index_;
        result.shareIndex_ = shareIndex_;
        result.envelope_ = envelope_;
        result.sharer_ = sharer_;
        onBuilt();
        return result;
      }

      @java.lang.Override
      public Builder clone() {
        return super.clone();
      }
      @java.lang.Override
      public Builder setField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          java.lang.Object value) {
        return super.setField(field, value);
      }
      @java.lang.Override
      public Builder clearField(
          com.google.protobuf.Descriptors.FieldDescriptor field) {
        return super.clearField(field);
      }
      @java.lang.Override
      public Builder clearOneof(
          com.google.protobuf.Descriptors.OneofDescriptor oneof) {
        return super.clearOneof(oneof);
      }
      @java.lang.Override
      public Builder setRepeatedField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          int index, java.lang.Object value) {
        return super.setRepeatedField(field, index, value);
      }
      @java.lang.Override
      public Builder addRepeatedField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          java.lang.Object value) {
        return super.addRepeatedField(field, value);
      }
      @java.lang.Override
      public Builder mergeFrom(com.google.protobuf.Message other) {
        if (other instanceof liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage) {
          return mergeFrom((liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage other) {
        if (other == liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage.getDefaultInstance()) return this;
        if (!other.getCreator().isEmpty()) {
          creator_ = other.creator_;
          onChanged();
        }
        if (!other.getIndex().isEmpty()) {
          index_ = other.index_;
          onChanged();
        }
        if (!other.getShareIndex().isEmpty()) {
          shareIndex_ = other.shareIndex_;
          onChanged();
        }
        if (!other.getEnvelope().isEmpty()) {
          envelope_ = other.envelope_;
          onChanged();
        }
        if (!other.getSharer().isEmpty()) {
          sharer_ = other.sharer_;
          onChanged();
        }
        this.mergeUnknownFields(other.unknownFields);
        onChanged();
        return this;
      }

      @java.lang.Override
      public final boolean isInitialized() {
        return true;
      }

      @java.lang.Override
      public Builder mergeFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws java.io.IOException {
        liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage parsedMessage = null;
        try {
          parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          parsedMessage = (liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage) e.getUnfinishedMessage();
          throw e.unwrapIOException();
        } finally {
          if (parsedMessage != null) {
            mergeFrom(parsedMessage);
          }
        }
        return this;
      }

      private java.lang.Object creator_ = "";
      /**
       * <code>string creator = 1;</code>
       */
      public java.lang.String getCreator() {
        java.lang.Object ref = creator_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          creator_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string creator = 1;</code>
       */
      public com.google.protobuf.ByteString
          getCreatorBytes() {
        java.lang.Object ref = creator_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          creator_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string creator = 1;</code>
       */
      public Builder setCreator(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        creator_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>string creator = 1;</code>
       */
      public Builder clearCreator() {
        
        creator_ = getDefaultInstance().getCreator();
        onChanged();
        return this;
      }
      /**
       * <code>string creator = 1;</code>
       */
      public Builder setCreatorBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        creator_ = value;
        onChanged();
        return this;
      }

      private java.lang.Object index_ = "";
      /**
       * <code>string index = 2;</code>
       */
      public java.lang.String getIndex() {
        java.lang.Object ref = index_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          index_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string index = 2;</code>
       */
      public com.google.protobuf.ByteString
          getIndexBytes() {
        java.lang.Object ref = index_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          index_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string index = 2;</code>
       */
      public Builder setIndex(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        index_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>string index = 2;</code>
       */
      public Builder clearIndex() {
        
        index_ = getDefaultInstance().getIndex();
        onChanged();
        return this;
      }
      /**
       * <code>string index = 2;</code>
       */
      public Builder setIndexBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        index_ = value;
        onChanged();
        return this;
      }

      private java.lang.Object shareIndex_ = "";
      /**
       * <code>string shareIndex = 3;</code>
       */
      public java.lang.String getShareIndex() {
        java.lang.Object ref = shareIndex_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          shareIndex_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string shareIndex = 3;</code>
       */
      public com.google.protobuf.ByteString
          getShareIndexBytes() {
        java.lang.Object ref = shareIndex_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          shareIndex_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string shareIndex = 3;</code>
       */
      public Builder setShareIndex(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        shareIndex_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>string shareIndex = 3;</code>
       */
      public Builder clearShareIndex() {
        
        shareIndex_ = getDefaultInstance().getShareIndex();
        onChanged();
        return this;
      }
      /**
       * <code>string shareIndex = 3;</code>
       */
      public Builder setShareIndexBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        shareIndex_ = value;
        onChanged();
        return this;
      }

      private java.lang.Object envelope_ = "";
      /**
       * <code>string envelope = 4;</code>
       */
      public java.lang.String getEnvelope() {
        java.lang.Object ref = envelope_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          envelope_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string envelope = 4;</code>
       */
      public com.google.protobuf.ByteString
          getEnvelopeBytes() {
        java.lang.Object ref = envelope_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          envelope_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string envelope = 4;</code>
       */
      public Builder setEnvelope(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        envelope_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>string envelope = 4;</code>
       */
      public Builder clearEnvelope() {
        
        envelope_ = getDefaultInstance().getEnvelope();
        onChanged();
        return this;
      }
      /**
       * <code>string envelope = 4;</code>
       */
      public Builder setEnvelopeBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        envelope_ = value;
        onChanged();
        return this;
      }

      private java.lang.Object sharer_ = "";
      /**
       * <code>string sharer = 5;</code>
       */
      public java.lang.String getSharer() {
        java.lang.Object ref = sharer_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          sharer_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string sharer = 5;</code>
       */
      public com.google.protobuf.ByteString
          getSharerBytes() {
        java.lang.Object ref = sharer_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          sharer_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string sharer = 5;</code>
       */
      public Builder setSharer(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        sharer_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>string sharer = 5;</code>
       */
      public Builder clearSharer() {
        
        sharer_ = getDefaultInstance().getSharer();
        onChanged();
        return this;
      }
      /**
       * <code>string sharer = 5;</code>
       */
      public Builder setSharerBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        sharer_ = value;
        onChanged();
        return this;
      }
      @java.lang.Override
      public final Builder setUnknownFields(
          final com.google.protobuf.UnknownFieldSet unknownFields) {
        return super.setUnknownFields(unknownFields);
      }

      @java.lang.Override
      public final Builder mergeUnknownFields(
          final com.google.protobuf.UnknownFieldSet unknownFields) {
        return super.mergeUnknownFields(unknownFields);
      }


      // @@protoc_insertion_point(builder_scope:liubaninc.m0.storage.ShareEncryptStorage)
    }

    // @@protoc_insertion_point(class_scope:liubaninc.m0.storage.ShareEncryptStorage)
    private static final liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage();
    }

    public static liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<ShareEncryptStorage>
        PARSER = new com.google.protobuf.AbstractParser<ShareEncryptStorage>() {
      @java.lang.Override
      public ShareEncryptStorage parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        return new ShareEncryptStorage(input, extensionRegistry);
      }
    };

    public static com.google.protobuf.Parser<ShareEncryptStorage> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<ShareEncryptStorage> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public liubaninc.m0.storage.ShareEncryptStorageOuterClass.ShareEncryptStorage getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_liubaninc_m0_storage_ShareEncryptStorage_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_liubaninc_m0_storage_ShareEncryptStorage_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n&m0/storage/share_encrypt_storage.proto" +
      "\022\024liubaninc.m0.storage\032\024gogoproto/gogo.p" +
      "roto\"k\n\023ShareEncryptStorage\022\017\n\007creator\030\001" +
      " \001(\t\022\r\n\005index\030\002 \001(\t\022\022\n\nshareIndex\030\003 \001(\t\022" +
      "\020\n\010envelope\030\004 \001(\t\022\016\n\006sharer\030\005 \001(\tB)Z\'git" +
      "hub.com/liubaninc/m0/x/storage/typesb\006pr" +
      "oto3"
    };
    com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner assigner =
        new com.google.protobuf.Descriptors.FileDescriptor.    InternalDescriptorAssigner() {
          public com.google.protobuf.ExtensionRegistry assignDescriptors(
              com.google.protobuf.Descriptors.FileDescriptor root) {
            descriptor = root;
            return null;
          }
        };
    com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.protobuf.GoGoProtos.getDescriptor(),
        }, assigner);
    internal_static_liubaninc_m0_storage_ShareEncryptStorage_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_liubaninc_m0_storage_ShareEncryptStorage_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_liubaninc_m0_storage_ShareEncryptStorage_descriptor,
        new java.lang.String[] { "Creator", "Index", "ShareIndex", "Envelope", "Sharer", });
    com.google.protobuf.GoGoProtos.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}