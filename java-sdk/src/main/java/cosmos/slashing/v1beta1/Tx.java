// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: cosmos/slashing/v1beta1/tx.proto

package cosmos.slashing.v1beta1;

public final class Tx {
  private Tx() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface MsgUnjailOrBuilder extends
      // @@protoc_insertion_point(interface_extends:cosmos.slashing.v1beta1.MsgUnjail)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <code>string validator_addr = 1 [(.gogoproto.jsontag) = "address", (.gogoproto.moretags) = "yaml:&#92;"address&#92;""];</code>
     */
    java.lang.String getValidatorAddr();
    /**
     * <code>string validator_addr = 1 [(.gogoproto.jsontag) = "address", (.gogoproto.moretags) = "yaml:&#92;"address&#92;""];</code>
     */
    com.google.protobuf.ByteString
        getValidatorAddrBytes();
  }
  /**
   * <pre>
   * MsgUnjail defines the Msg/Unjail request type
   * </pre>
   *
   * Protobuf type {@code cosmos.slashing.v1beta1.MsgUnjail}
   */
  public  static final class MsgUnjail extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:cosmos.slashing.v1beta1.MsgUnjail)
      MsgUnjailOrBuilder {
  private static final long serialVersionUID = 0L;
    // Use MsgUnjail.newBuilder() to construct.
    private MsgUnjail(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private MsgUnjail() {
      validatorAddr_ = "";
    }

    @java.lang.Override
    public final com.google.protobuf.UnknownFieldSet
    getUnknownFields() {
      return this.unknownFields;
    }
    private MsgUnjail(
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

              validatorAddr_ = s;
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
      return cosmos.slashing.v1beta1.Tx.internal_static_cosmos_slashing_v1beta1_MsgUnjail_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return cosmos.slashing.v1beta1.Tx.internal_static_cosmos_slashing_v1beta1_MsgUnjail_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              cosmos.slashing.v1beta1.Tx.MsgUnjail.class, cosmos.slashing.v1beta1.Tx.MsgUnjail.Builder.class);
    }

    public static final int VALIDATOR_ADDR_FIELD_NUMBER = 1;
    private volatile java.lang.Object validatorAddr_;
    /**
     * <code>string validator_addr = 1 [(.gogoproto.jsontag) = "address", (.gogoproto.moretags) = "yaml:&#92;"address&#92;""];</code>
     */
    public java.lang.String getValidatorAddr() {
      java.lang.Object ref = validatorAddr_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        validatorAddr_ = s;
        return s;
      }
    }
    /**
     * <code>string validator_addr = 1 [(.gogoproto.jsontag) = "address", (.gogoproto.moretags) = "yaml:&#92;"address&#92;""];</code>
     */
    public com.google.protobuf.ByteString
        getValidatorAddrBytes() {
      java.lang.Object ref = validatorAddr_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        validatorAddr_ = b;
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
      if (!getValidatorAddrBytes().isEmpty()) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 1, validatorAddr_);
      }
      unknownFields.writeTo(output);
    }

    @java.lang.Override
    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      if (!getValidatorAddrBytes().isEmpty()) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, validatorAddr_);
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
      if (!(obj instanceof cosmos.slashing.v1beta1.Tx.MsgUnjail)) {
        return super.equals(obj);
      }
      cosmos.slashing.v1beta1.Tx.MsgUnjail other = (cosmos.slashing.v1beta1.Tx.MsgUnjail) obj;

      if (!getValidatorAddr()
          .equals(other.getValidatorAddr())) return false;
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
      hash = (37 * hash) + VALIDATOR_ADDR_FIELD_NUMBER;
      hash = (53 * hash) + getValidatorAddr().hashCode();
      hash = (29 * hash) + unknownFields.hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static cosmos.slashing.v1beta1.Tx.MsgUnjail parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjail parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjail parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjail parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjail parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjail parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjail parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjail parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjail parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjail parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjail parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjail parseFrom(
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
    public static Builder newBuilder(cosmos.slashing.v1beta1.Tx.MsgUnjail prototype) {
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
     * <pre>
     * MsgUnjail defines the Msg/Unjail request type
     * </pre>
     *
     * Protobuf type {@code cosmos.slashing.v1beta1.MsgUnjail}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:cosmos.slashing.v1beta1.MsgUnjail)
        cosmos.slashing.v1beta1.Tx.MsgUnjailOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return cosmos.slashing.v1beta1.Tx.internal_static_cosmos_slashing_v1beta1_MsgUnjail_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return cosmos.slashing.v1beta1.Tx.internal_static_cosmos_slashing_v1beta1_MsgUnjail_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                cosmos.slashing.v1beta1.Tx.MsgUnjail.class, cosmos.slashing.v1beta1.Tx.MsgUnjail.Builder.class);
      }

      // Construct using cosmos.slashing.v1beta1.Tx.MsgUnjail.newBuilder()
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
        validatorAddr_ = "";

        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return cosmos.slashing.v1beta1.Tx.internal_static_cosmos_slashing_v1beta1_MsgUnjail_descriptor;
      }

      @java.lang.Override
      public cosmos.slashing.v1beta1.Tx.MsgUnjail getDefaultInstanceForType() {
        return cosmos.slashing.v1beta1.Tx.MsgUnjail.getDefaultInstance();
      }

      @java.lang.Override
      public cosmos.slashing.v1beta1.Tx.MsgUnjail build() {
        cosmos.slashing.v1beta1.Tx.MsgUnjail result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public cosmos.slashing.v1beta1.Tx.MsgUnjail buildPartial() {
        cosmos.slashing.v1beta1.Tx.MsgUnjail result = new cosmos.slashing.v1beta1.Tx.MsgUnjail(this);
        result.validatorAddr_ = validatorAddr_;
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
        if (other instanceof cosmos.slashing.v1beta1.Tx.MsgUnjail) {
          return mergeFrom((cosmos.slashing.v1beta1.Tx.MsgUnjail)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(cosmos.slashing.v1beta1.Tx.MsgUnjail other) {
        if (other == cosmos.slashing.v1beta1.Tx.MsgUnjail.getDefaultInstance()) return this;
        if (!other.getValidatorAddr().isEmpty()) {
          validatorAddr_ = other.validatorAddr_;
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
        cosmos.slashing.v1beta1.Tx.MsgUnjail parsedMessage = null;
        try {
          parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          parsedMessage = (cosmos.slashing.v1beta1.Tx.MsgUnjail) e.getUnfinishedMessage();
          throw e.unwrapIOException();
        } finally {
          if (parsedMessage != null) {
            mergeFrom(parsedMessage);
          }
        }
        return this;
      }

      private java.lang.Object validatorAddr_ = "";
      /**
       * <code>string validator_addr = 1 [(.gogoproto.jsontag) = "address", (.gogoproto.moretags) = "yaml:&#92;"address&#92;""];</code>
       */
      public java.lang.String getValidatorAddr() {
        java.lang.Object ref = validatorAddr_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          validatorAddr_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string validator_addr = 1 [(.gogoproto.jsontag) = "address", (.gogoproto.moretags) = "yaml:&#92;"address&#92;""];</code>
       */
      public com.google.protobuf.ByteString
          getValidatorAddrBytes() {
        java.lang.Object ref = validatorAddr_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          validatorAddr_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string validator_addr = 1 [(.gogoproto.jsontag) = "address", (.gogoproto.moretags) = "yaml:&#92;"address&#92;""];</code>
       */
      public Builder setValidatorAddr(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        validatorAddr_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>string validator_addr = 1 [(.gogoproto.jsontag) = "address", (.gogoproto.moretags) = "yaml:&#92;"address&#92;""];</code>
       */
      public Builder clearValidatorAddr() {
        
        validatorAddr_ = getDefaultInstance().getValidatorAddr();
        onChanged();
        return this;
      }
      /**
       * <code>string validator_addr = 1 [(.gogoproto.jsontag) = "address", (.gogoproto.moretags) = "yaml:&#92;"address&#92;""];</code>
       */
      public Builder setValidatorAddrBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        validatorAddr_ = value;
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


      // @@protoc_insertion_point(builder_scope:cosmos.slashing.v1beta1.MsgUnjail)
    }

    // @@protoc_insertion_point(class_scope:cosmos.slashing.v1beta1.MsgUnjail)
    private static final cosmos.slashing.v1beta1.Tx.MsgUnjail DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new cosmos.slashing.v1beta1.Tx.MsgUnjail();
    }

    public static cosmos.slashing.v1beta1.Tx.MsgUnjail getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<MsgUnjail>
        PARSER = new com.google.protobuf.AbstractParser<MsgUnjail>() {
      @java.lang.Override
      public MsgUnjail parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        return new MsgUnjail(input, extensionRegistry);
      }
    };

    public static com.google.protobuf.Parser<MsgUnjail> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<MsgUnjail> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public cosmos.slashing.v1beta1.Tx.MsgUnjail getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  public interface MsgUnjailResponseOrBuilder extends
      // @@protoc_insertion_point(interface_extends:cosmos.slashing.v1beta1.MsgUnjailResponse)
      com.google.protobuf.MessageOrBuilder {
  }
  /**
   * <pre>
   * MsgUnjailResponse defines the Msg/Unjail response type
   * </pre>
   *
   * Protobuf type {@code cosmos.slashing.v1beta1.MsgUnjailResponse}
   */
  public  static final class MsgUnjailResponse extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:cosmos.slashing.v1beta1.MsgUnjailResponse)
      MsgUnjailResponseOrBuilder {
  private static final long serialVersionUID = 0L;
    // Use MsgUnjailResponse.newBuilder() to construct.
    private MsgUnjailResponse(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private MsgUnjailResponse() {
    }

    @java.lang.Override
    public final com.google.protobuf.UnknownFieldSet
    getUnknownFields() {
      return this.unknownFields;
    }
    private MsgUnjailResponse(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      this();
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
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
      return cosmos.slashing.v1beta1.Tx.internal_static_cosmos_slashing_v1beta1_MsgUnjailResponse_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return cosmos.slashing.v1beta1.Tx.internal_static_cosmos_slashing_v1beta1_MsgUnjailResponse_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              cosmos.slashing.v1beta1.Tx.MsgUnjailResponse.class, cosmos.slashing.v1beta1.Tx.MsgUnjailResponse.Builder.class);
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
      unknownFields.writeTo(output);
    }

    @java.lang.Override
    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      size += unknownFields.getSerializedSize();
      memoizedSize = size;
      return size;
    }

    @java.lang.Override
    public boolean equals(final java.lang.Object obj) {
      if (obj == this) {
       return true;
      }
      if (!(obj instanceof cosmos.slashing.v1beta1.Tx.MsgUnjailResponse)) {
        return super.equals(obj);
      }
      cosmos.slashing.v1beta1.Tx.MsgUnjailResponse other = (cosmos.slashing.v1beta1.Tx.MsgUnjailResponse) obj;

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
      hash = (29 * hash) + unknownFields.hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static cosmos.slashing.v1beta1.Tx.MsgUnjailResponse parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjailResponse parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjailResponse parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjailResponse parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjailResponse parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjailResponse parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjailResponse parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjailResponse parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjailResponse parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjailResponse parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjailResponse parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static cosmos.slashing.v1beta1.Tx.MsgUnjailResponse parseFrom(
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
    public static Builder newBuilder(cosmos.slashing.v1beta1.Tx.MsgUnjailResponse prototype) {
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
     * <pre>
     * MsgUnjailResponse defines the Msg/Unjail response type
     * </pre>
     *
     * Protobuf type {@code cosmos.slashing.v1beta1.MsgUnjailResponse}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:cosmos.slashing.v1beta1.MsgUnjailResponse)
        cosmos.slashing.v1beta1.Tx.MsgUnjailResponseOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return cosmos.slashing.v1beta1.Tx.internal_static_cosmos_slashing_v1beta1_MsgUnjailResponse_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return cosmos.slashing.v1beta1.Tx.internal_static_cosmos_slashing_v1beta1_MsgUnjailResponse_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                cosmos.slashing.v1beta1.Tx.MsgUnjailResponse.class, cosmos.slashing.v1beta1.Tx.MsgUnjailResponse.Builder.class);
      }

      // Construct using cosmos.slashing.v1beta1.Tx.MsgUnjailResponse.newBuilder()
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
        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return cosmos.slashing.v1beta1.Tx.internal_static_cosmos_slashing_v1beta1_MsgUnjailResponse_descriptor;
      }

      @java.lang.Override
      public cosmos.slashing.v1beta1.Tx.MsgUnjailResponse getDefaultInstanceForType() {
        return cosmos.slashing.v1beta1.Tx.MsgUnjailResponse.getDefaultInstance();
      }

      @java.lang.Override
      public cosmos.slashing.v1beta1.Tx.MsgUnjailResponse build() {
        cosmos.slashing.v1beta1.Tx.MsgUnjailResponse result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public cosmos.slashing.v1beta1.Tx.MsgUnjailResponse buildPartial() {
        cosmos.slashing.v1beta1.Tx.MsgUnjailResponse result = new cosmos.slashing.v1beta1.Tx.MsgUnjailResponse(this);
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
        if (other instanceof cosmos.slashing.v1beta1.Tx.MsgUnjailResponse) {
          return mergeFrom((cosmos.slashing.v1beta1.Tx.MsgUnjailResponse)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(cosmos.slashing.v1beta1.Tx.MsgUnjailResponse other) {
        if (other == cosmos.slashing.v1beta1.Tx.MsgUnjailResponse.getDefaultInstance()) return this;
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
        cosmos.slashing.v1beta1.Tx.MsgUnjailResponse parsedMessage = null;
        try {
          parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          parsedMessage = (cosmos.slashing.v1beta1.Tx.MsgUnjailResponse) e.getUnfinishedMessage();
          throw e.unwrapIOException();
        } finally {
          if (parsedMessage != null) {
            mergeFrom(parsedMessage);
          }
        }
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


      // @@protoc_insertion_point(builder_scope:cosmos.slashing.v1beta1.MsgUnjailResponse)
    }

    // @@protoc_insertion_point(class_scope:cosmos.slashing.v1beta1.MsgUnjailResponse)
    private static final cosmos.slashing.v1beta1.Tx.MsgUnjailResponse DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new cosmos.slashing.v1beta1.Tx.MsgUnjailResponse();
    }

    public static cosmos.slashing.v1beta1.Tx.MsgUnjailResponse getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<MsgUnjailResponse>
        PARSER = new com.google.protobuf.AbstractParser<MsgUnjailResponse>() {
      @java.lang.Override
      public MsgUnjailResponse parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        return new MsgUnjailResponse(input, extensionRegistry);
      }
    };

    public static com.google.protobuf.Parser<MsgUnjailResponse> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<MsgUnjailResponse> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public cosmos.slashing.v1beta1.Tx.MsgUnjailResponse getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_cosmos_slashing_v1beta1_MsgUnjail_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_cosmos_slashing_v1beta1_MsgUnjail_fieldAccessorTable;
  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_cosmos_slashing_v1beta1_MsgUnjailResponse_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_cosmos_slashing_v1beta1_MsgUnjailResponse_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n cosmos/slashing/v1beta1/tx.proto\022\027cosm" +
      "os.slashing.v1beta1\032\024gogoproto/gogo.prot" +
      "o\"L\n\tMsgUnjail\0225\n\016validator_addr\030\001 \001(\tB\035" +
      "\362\336\037\016yaml:\"address\"\352\336\037\007address:\010\210\240\037\000\230\240\037\001\"" +
      "\023\n\021MsgUnjailResponse2_\n\003Msg\022X\n\006Unjail\022\"." +
      "cosmos.slashing.v1beta1.MsgUnjail\032*.cosm" +
      "os.slashing.v1beta1.MsgUnjailResponseB3Z" +
      "-github.com/cosmos/cosmos-sdk/x/slashing" +
      "/types\250\342\036\001b\006proto3"
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
    internal_static_cosmos_slashing_v1beta1_MsgUnjail_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_cosmos_slashing_v1beta1_MsgUnjail_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_cosmos_slashing_v1beta1_MsgUnjail_descriptor,
        new java.lang.String[] { "ValidatorAddr", });
    internal_static_cosmos_slashing_v1beta1_MsgUnjailResponse_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_cosmos_slashing_v1beta1_MsgUnjailResponse_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_cosmos_slashing_v1beta1_MsgUnjailResponse_descriptor,
        new java.lang.String[] { });
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(com.google.protobuf.GoGoProtos.equalAll);
    registry.add(com.google.protobuf.GoGoProtos.goprotoGetters);
    registry.add(com.google.protobuf.GoGoProtos.goprotoStringer);
    registry.add(com.google.protobuf.GoGoProtos.jsontag);
    registry.add(com.google.protobuf.GoGoProtos.moretags);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
    com.google.protobuf.GoGoProtos.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
