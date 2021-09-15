// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: m0/mibc/itx.proto

package liubaninc.m0.mibc;

public final class ItxOuterClass {
  private ItxOuterClass() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface ItxOrBuilder extends
      // @@protoc_insertion_point(interface_extends:liubaninc.m0.mibc.Itx)
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
     * <code>uint64 id = 2;</code>
     */
    long getId();

    /**
     * <code>string sourceHash = 3;</code>
     */
    java.lang.String getSourceHash();
    /**
     * <code>string sourceHash = 3;</code>
     */
    com.google.protobuf.ByteString
        getSourceHashBytes();

    /**
     * <code>string destinationHash = 4;</code>
     */
    java.lang.String getDestinationHash();
    /**
     * <code>string destinationHash = 4;</code>
     */
    com.google.protobuf.ByteString
        getDestinationHashBytes();

    /**
     * <code>string log = 6;</code>
     */
    java.lang.String getLog();
    /**
     * <code>string log = 6;</code>
     */
    com.google.protobuf.ByteString
        getLogBytes();
  }
  /**
   * Protobuf type {@code liubaninc.m0.mibc.Itx}
   */
  public  static final class Itx extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:liubaninc.m0.mibc.Itx)
      ItxOrBuilder {
  private static final long serialVersionUID = 0L;
    // Use Itx.newBuilder() to construct.
    private Itx(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private Itx() {
      creator_ = "";
      sourceHash_ = "";
      destinationHash_ = "";
      log_ = "";
    }

    @java.lang.Override
    public final com.google.protobuf.UnknownFieldSet
    getUnknownFields() {
      return this.unknownFields;
    }
    private Itx(
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
            case 16: {

              id_ = input.readUInt64();
              break;
            }
            case 26: {
              java.lang.String s = input.readStringRequireUtf8();

              sourceHash_ = s;
              break;
            }
            case 34: {
              java.lang.String s = input.readStringRequireUtf8();

              destinationHash_ = s;
              break;
            }
            case 50: {
              java.lang.String s = input.readStringRequireUtf8();

              log_ = s;
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
      return liubaninc.m0.mibc.ItxOuterClass.internal_static_liubaninc_m0_mibc_Itx_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return liubaninc.m0.mibc.ItxOuterClass.internal_static_liubaninc_m0_mibc_Itx_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              liubaninc.m0.mibc.ItxOuterClass.Itx.class, liubaninc.m0.mibc.ItxOuterClass.Itx.Builder.class);
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

    public static final int ID_FIELD_NUMBER = 2;
    private long id_;
    /**
     * <code>uint64 id = 2;</code>
     */
    public long getId() {
      return id_;
    }

    public static final int SOURCEHASH_FIELD_NUMBER = 3;
    private volatile java.lang.Object sourceHash_;
    /**
     * <code>string sourceHash = 3;</code>
     */
    public java.lang.String getSourceHash() {
      java.lang.Object ref = sourceHash_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        sourceHash_ = s;
        return s;
      }
    }
    /**
     * <code>string sourceHash = 3;</code>
     */
    public com.google.protobuf.ByteString
        getSourceHashBytes() {
      java.lang.Object ref = sourceHash_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        sourceHash_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int DESTINATIONHASH_FIELD_NUMBER = 4;
    private volatile java.lang.Object destinationHash_;
    /**
     * <code>string destinationHash = 4;</code>
     */
    public java.lang.String getDestinationHash() {
      java.lang.Object ref = destinationHash_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        destinationHash_ = s;
        return s;
      }
    }
    /**
     * <code>string destinationHash = 4;</code>
     */
    public com.google.protobuf.ByteString
        getDestinationHashBytes() {
      java.lang.Object ref = destinationHash_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        destinationHash_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int LOG_FIELD_NUMBER = 6;
    private volatile java.lang.Object log_;
    /**
     * <code>string log = 6;</code>
     */
    public java.lang.String getLog() {
      java.lang.Object ref = log_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        log_ = s;
        return s;
      }
    }
    /**
     * <code>string log = 6;</code>
     */
    public com.google.protobuf.ByteString
        getLogBytes() {
      java.lang.Object ref = log_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        log_ = b;
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
      if (id_ != 0L) {
        output.writeUInt64(2, id_);
      }
      if (!getSourceHashBytes().isEmpty()) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 3, sourceHash_);
      }
      if (!getDestinationHashBytes().isEmpty()) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 4, destinationHash_);
      }
      if (!getLogBytes().isEmpty()) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 6, log_);
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
      if (id_ != 0L) {
        size += com.google.protobuf.CodedOutputStream
          .computeUInt64Size(2, id_);
      }
      if (!getSourceHashBytes().isEmpty()) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(3, sourceHash_);
      }
      if (!getDestinationHashBytes().isEmpty()) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(4, destinationHash_);
      }
      if (!getLogBytes().isEmpty()) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(6, log_);
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
      if (!(obj instanceof liubaninc.m0.mibc.ItxOuterClass.Itx)) {
        return super.equals(obj);
      }
      liubaninc.m0.mibc.ItxOuterClass.Itx other = (liubaninc.m0.mibc.ItxOuterClass.Itx) obj;

      if (!getCreator()
          .equals(other.getCreator())) return false;
      if (getId()
          != other.getId()) return false;
      if (!getSourceHash()
          .equals(other.getSourceHash())) return false;
      if (!getDestinationHash()
          .equals(other.getDestinationHash())) return false;
      if (!getLog()
          .equals(other.getLog())) return false;
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
      hash = (37 * hash) + ID_FIELD_NUMBER;
      hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
          getId());
      hash = (37 * hash) + SOURCEHASH_FIELD_NUMBER;
      hash = (53 * hash) + getSourceHash().hashCode();
      hash = (37 * hash) + DESTINATIONHASH_FIELD_NUMBER;
      hash = (53 * hash) + getDestinationHash().hashCode();
      hash = (37 * hash) + LOG_FIELD_NUMBER;
      hash = (53 * hash) + getLog().hashCode();
      hash = (29 * hash) + unknownFields.hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static liubaninc.m0.mibc.ItxOuterClass.Itx parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static liubaninc.m0.mibc.ItxOuterClass.Itx parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static liubaninc.m0.mibc.ItxOuterClass.Itx parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static liubaninc.m0.mibc.ItxOuterClass.Itx parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static liubaninc.m0.mibc.ItxOuterClass.Itx parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static liubaninc.m0.mibc.ItxOuterClass.Itx parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static liubaninc.m0.mibc.ItxOuterClass.Itx parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static liubaninc.m0.mibc.ItxOuterClass.Itx parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }
    public static liubaninc.m0.mibc.ItxOuterClass.Itx parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }
    public static liubaninc.m0.mibc.ItxOuterClass.Itx parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static liubaninc.m0.mibc.ItxOuterClass.Itx parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static liubaninc.m0.mibc.ItxOuterClass.Itx parseFrom(
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
    public static Builder newBuilder(liubaninc.m0.mibc.ItxOuterClass.Itx prototype) {
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
     * Protobuf type {@code liubaninc.m0.mibc.Itx}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:liubaninc.m0.mibc.Itx)
        liubaninc.m0.mibc.ItxOuterClass.ItxOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return liubaninc.m0.mibc.ItxOuterClass.internal_static_liubaninc_m0_mibc_Itx_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return liubaninc.m0.mibc.ItxOuterClass.internal_static_liubaninc_m0_mibc_Itx_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                liubaninc.m0.mibc.ItxOuterClass.Itx.class, liubaninc.m0.mibc.ItxOuterClass.Itx.Builder.class);
      }

      // Construct using liubaninc.m0.mibc.ItxOuterClass.Itx.newBuilder()
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

        id_ = 0L;

        sourceHash_ = "";

        destinationHash_ = "";

        log_ = "";

        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return liubaninc.m0.mibc.ItxOuterClass.internal_static_liubaninc_m0_mibc_Itx_descriptor;
      }

      @java.lang.Override
      public liubaninc.m0.mibc.ItxOuterClass.Itx getDefaultInstanceForType() {
        return liubaninc.m0.mibc.ItxOuterClass.Itx.getDefaultInstance();
      }

      @java.lang.Override
      public liubaninc.m0.mibc.ItxOuterClass.Itx build() {
        liubaninc.m0.mibc.ItxOuterClass.Itx result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public liubaninc.m0.mibc.ItxOuterClass.Itx buildPartial() {
        liubaninc.m0.mibc.ItxOuterClass.Itx result = new liubaninc.m0.mibc.ItxOuterClass.Itx(this);
        result.creator_ = creator_;
        result.id_ = id_;
        result.sourceHash_ = sourceHash_;
        result.destinationHash_ = destinationHash_;
        result.log_ = log_;
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
        if (other instanceof liubaninc.m0.mibc.ItxOuterClass.Itx) {
          return mergeFrom((liubaninc.m0.mibc.ItxOuterClass.Itx)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(liubaninc.m0.mibc.ItxOuterClass.Itx other) {
        if (other == liubaninc.m0.mibc.ItxOuterClass.Itx.getDefaultInstance()) return this;
        if (!other.getCreator().isEmpty()) {
          creator_ = other.creator_;
          onChanged();
        }
        if (other.getId() != 0L) {
          setId(other.getId());
        }
        if (!other.getSourceHash().isEmpty()) {
          sourceHash_ = other.sourceHash_;
          onChanged();
        }
        if (!other.getDestinationHash().isEmpty()) {
          destinationHash_ = other.destinationHash_;
          onChanged();
        }
        if (!other.getLog().isEmpty()) {
          log_ = other.log_;
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
        liubaninc.m0.mibc.ItxOuterClass.Itx parsedMessage = null;
        try {
          parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          parsedMessage = (liubaninc.m0.mibc.ItxOuterClass.Itx) e.getUnfinishedMessage();
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

      private long id_ ;
      /**
       * <code>uint64 id = 2;</code>
       */
      public long getId() {
        return id_;
      }
      /**
       * <code>uint64 id = 2;</code>
       */
      public Builder setId(long value) {
        
        id_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>uint64 id = 2;</code>
       */
      public Builder clearId() {
        
        id_ = 0L;
        onChanged();
        return this;
      }

      private java.lang.Object sourceHash_ = "";
      /**
       * <code>string sourceHash = 3;</code>
       */
      public java.lang.String getSourceHash() {
        java.lang.Object ref = sourceHash_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          sourceHash_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string sourceHash = 3;</code>
       */
      public com.google.protobuf.ByteString
          getSourceHashBytes() {
        java.lang.Object ref = sourceHash_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          sourceHash_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string sourceHash = 3;</code>
       */
      public Builder setSourceHash(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        sourceHash_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>string sourceHash = 3;</code>
       */
      public Builder clearSourceHash() {
        
        sourceHash_ = getDefaultInstance().getSourceHash();
        onChanged();
        return this;
      }
      /**
       * <code>string sourceHash = 3;</code>
       */
      public Builder setSourceHashBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        sourceHash_ = value;
        onChanged();
        return this;
      }

      private java.lang.Object destinationHash_ = "";
      /**
       * <code>string destinationHash = 4;</code>
       */
      public java.lang.String getDestinationHash() {
        java.lang.Object ref = destinationHash_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          destinationHash_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string destinationHash = 4;</code>
       */
      public com.google.protobuf.ByteString
          getDestinationHashBytes() {
        java.lang.Object ref = destinationHash_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          destinationHash_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string destinationHash = 4;</code>
       */
      public Builder setDestinationHash(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        destinationHash_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>string destinationHash = 4;</code>
       */
      public Builder clearDestinationHash() {
        
        destinationHash_ = getDefaultInstance().getDestinationHash();
        onChanged();
        return this;
      }
      /**
       * <code>string destinationHash = 4;</code>
       */
      public Builder setDestinationHashBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        destinationHash_ = value;
        onChanged();
        return this;
      }

      private java.lang.Object log_ = "";
      /**
       * <code>string log = 6;</code>
       */
      public java.lang.String getLog() {
        java.lang.Object ref = log_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          log_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string log = 6;</code>
       */
      public com.google.protobuf.ByteString
          getLogBytes() {
        java.lang.Object ref = log_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          log_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string log = 6;</code>
       */
      public Builder setLog(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        log_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>string log = 6;</code>
       */
      public Builder clearLog() {
        
        log_ = getDefaultInstance().getLog();
        onChanged();
        return this;
      }
      /**
       * <code>string log = 6;</code>
       */
      public Builder setLogBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        log_ = value;
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


      // @@protoc_insertion_point(builder_scope:liubaninc.m0.mibc.Itx)
    }

    // @@protoc_insertion_point(class_scope:liubaninc.m0.mibc.Itx)
    private static final liubaninc.m0.mibc.ItxOuterClass.Itx DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new liubaninc.m0.mibc.ItxOuterClass.Itx();
    }

    public static liubaninc.m0.mibc.ItxOuterClass.Itx getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<Itx>
        PARSER = new com.google.protobuf.AbstractParser<Itx>() {
      @java.lang.Override
      public Itx parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        return new Itx(input, extensionRegistry);
      }
    };

    public static com.google.protobuf.Parser<Itx> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<Itx> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public liubaninc.m0.mibc.ItxOuterClass.Itx getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_liubaninc_m0_mibc_Itx_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_liubaninc_m0_mibc_Itx_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\021m0/mibc/itx.proto\022\021liubaninc.m0.mibc\"\\" +
      "\n\003Itx\022\017\n\007creator\030\001 \001(\t\022\n\n\002id\030\002 \001(\004\022\022\n\nso" +
      "urceHash\030\003 \001(\t\022\027\n\017destinationHash\030\004 \001(\t\022" +
      "\013\n\003log\030\006 \001(\tB&Z$github.com/liubaninc/m0/" +
      "x/mibc/typesb\006proto3"
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
        }, assigner);
    internal_static_liubaninc_m0_mibc_Itx_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_liubaninc_m0_mibc_Itx_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_liubaninc_m0_mibc_Itx_descriptor,
        new java.lang.String[] { "Creator", "Id", "SourceHash", "DestinationHash", "Log", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}
