// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: cosmos/genutil/v1beta1/genesis.proto

package cosmos.genutil.v1beta1;

public final class Genesis {
  private Genesis() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface GenesisStateOrBuilder extends
      // @@protoc_insertion_point(interface_extends:cosmos.genutil.v1beta1.GenesisState)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <pre>
     * gen_txs defines the genesis transactions.
     * </pre>
     *
     * <code>repeated bytes gen_txs = 1 [(.gogoproto.jsontag) = "gentxs", (.gogoproto.moretags) = "yaml:&#92;"gentxs&#92;"", (.gogoproto.casttype) = "encoding/json.RawMessage"];</code>
     */
    java.util.List<com.google.protobuf.ByteString> getGenTxsList();
    /**
     * <pre>
     * gen_txs defines the genesis transactions.
     * </pre>
     *
     * <code>repeated bytes gen_txs = 1 [(.gogoproto.jsontag) = "gentxs", (.gogoproto.moretags) = "yaml:&#92;"gentxs&#92;"", (.gogoproto.casttype) = "encoding/json.RawMessage"];</code>
     */
    int getGenTxsCount();
    /**
     * <pre>
     * gen_txs defines the genesis transactions.
     * </pre>
     *
     * <code>repeated bytes gen_txs = 1 [(.gogoproto.jsontag) = "gentxs", (.gogoproto.moretags) = "yaml:&#92;"gentxs&#92;"", (.gogoproto.casttype) = "encoding/json.RawMessage"];</code>
     */
    com.google.protobuf.ByteString getGenTxs(int index);
  }
  /**
   * <pre>
   * GenesisState defines the raw genesis transaction in JSON.
   * </pre>
   *
   * Protobuf type {@code cosmos.genutil.v1beta1.GenesisState}
   */
  public  static final class GenesisState extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:cosmos.genutil.v1beta1.GenesisState)
      GenesisStateOrBuilder {
  private static final long serialVersionUID = 0L;
    // Use GenesisState.newBuilder() to construct.
    private GenesisState(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private GenesisState() {
      genTxs_ = java.util.Collections.emptyList();
    }

    @java.lang.Override
    public final com.google.protobuf.UnknownFieldSet
    getUnknownFields() {
      return this.unknownFields;
    }
    private GenesisState(
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
              if (!((mutable_bitField0_ & 0x00000001) != 0)) {
                genTxs_ = new java.util.ArrayList<com.google.protobuf.ByteString>();
                mutable_bitField0_ |= 0x00000001;
              }
              genTxs_.add(input.readBytes());
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
        if (((mutable_bitField0_ & 0x00000001) != 0)) {
          genTxs_ = java.util.Collections.unmodifiableList(genTxs_); // C
        }
        this.unknownFields = unknownFields.build();
        makeExtensionsImmutable();
      }
    }
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return cosmos.genutil.v1beta1.Genesis.internal_static_cosmos_genutil_v1beta1_GenesisState_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return cosmos.genutil.v1beta1.Genesis.internal_static_cosmos_genutil_v1beta1_GenesisState_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              cosmos.genutil.v1beta1.Genesis.GenesisState.class, cosmos.genutil.v1beta1.Genesis.GenesisState.Builder.class);
    }

    public static final int GEN_TXS_FIELD_NUMBER = 1;
    private java.util.List<com.google.protobuf.ByteString> genTxs_;
    /**
     * <pre>
     * gen_txs defines the genesis transactions.
     * </pre>
     *
     * <code>repeated bytes gen_txs = 1 [(.gogoproto.jsontag) = "gentxs", (.gogoproto.moretags) = "yaml:&#92;"gentxs&#92;"", (.gogoproto.casttype) = "encoding/json.RawMessage"];</code>
     */
    public java.util.List<com.google.protobuf.ByteString>
        getGenTxsList() {
      return genTxs_;
    }
    /**
     * <pre>
     * gen_txs defines the genesis transactions.
     * </pre>
     *
     * <code>repeated bytes gen_txs = 1 [(.gogoproto.jsontag) = "gentxs", (.gogoproto.moretags) = "yaml:&#92;"gentxs&#92;"", (.gogoproto.casttype) = "encoding/json.RawMessage"];</code>
     */
    public int getGenTxsCount() {
      return genTxs_.size();
    }
    /**
     * <pre>
     * gen_txs defines the genesis transactions.
     * </pre>
     *
     * <code>repeated bytes gen_txs = 1 [(.gogoproto.jsontag) = "gentxs", (.gogoproto.moretags) = "yaml:&#92;"gentxs&#92;"", (.gogoproto.casttype) = "encoding/json.RawMessage"];</code>
     */
    public com.google.protobuf.ByteString getGenTxs(int index) {
      return genTxs_.get(index);
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
      for (int i = 0; i < genTxs_.size(); i++) {
        output.writeBytes(1, genTxs_.get(i));
      }
      unknownFields.writeTo(output);
    }

    @java.lang.Override
    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      {
        int dataSize = 0;
        for (int i = 0; i < genTxs_.size(); i++) {
          dataSize += com.google.protobuf.CodedOutputStream
            .computeBytesSizeNoTag(genTxs_.get(i));
        }
        size += dataSize;
        size += 1 * getGenTxsList().size();
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
      if (!(obj instanceof cosmos.genutil.v1beta1.Genesis.GenesisState)) {
        return super.equals(obj);
      }
      cosmos.genutil.v1beta1.Genesis.GenesisState other = (cosmos.genutil.v1beta1.Genesis.GenesisState) obj;

      if (!getGenTxsList()
          .equals(other.getGenTxsList())) return false;
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
      if (getGenTxsCount() > 0) {
        hash = (37 * hash) + GEN_TXS_FIELD_NUMBER;
        hash = (53 * hash) + getGenTxsList().hashCode();
      }
      hash = (29 * hash) + unknownFields.hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static cosmos.genutil.v1beta1.Genesis.GenesisState parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static cosmos.genutil.v1beta1.Genesis.GenesisState parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static cosmos.genutil.v1beta1.Genesis.GenesisState parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static cosmos.genutil.v1beta1.Genesis.GenesisState parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static cosmos.genutil.v1beta1.Genesis.GenesisState parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static cosmos.genutil.v1beta1.Genesis.GenesisState parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static cosmos.genutil.v1beta1.Genesis.GenesisState parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static cosmos.genutil.v1beta1.Genesis.GenesisState parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }
    public static cosmos.genutil.v1beta1.Genesis.GenesisState parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }
    public static cosmos.genutil.v1beta1.Genesis.GenesisState parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static cosmos.genutil.v1beta1.Genesis.GenesisState parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static cosmos.genutil.v1beta1.Genesis.GenesisState parseFrom(
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
    public static Builder newBuilder(cosmos.genutil.v1beta1.Genesis.GenesisState prototype) {
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
     * GenesisState defines the raw genesis transaction in JSON.
     * </pre>
     *
     * Protobuf type {@code cosmos.genutil.v1beta1.GenesisState}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:cosmos.genutil.v1beta1.GenesisState)
        cosmos.genutil.v1beta1.Genesis.GenesisStateOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return cosmos.genutil.v1beta1.Genesis.internal_static_cosmos_genutil_v1beta1_GenesisState_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return cosmos.genutil.v1beta1.Genesis.internal_static_cosmos_genutil_v1beta1_GenesisState_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                cosmos.genutil.v1beta1.Genesis.GenesisState.class, cosmos.genutil.v1beta1.Genesis.GenesisState.Builder.class);
      }

      // Construct using cosmos.genutil.v1beta1.Genesis.GenesisState.newBuilder()
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
        genTxs_ = java.util.Collections.emptyList();
        bitField0_ = (bitField0_ & ~0x00000001);
        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return cosmos.genutil.v1beta1.Genesis.internal_static_cosmos_genutil_v1beta1_GenesisState_descriptor;
      }

      @java.lang.Override
      public cosmos.genutil.v1beta1.Genesis.GenesisState getDefaultInstanceForType() {
        return cosmos.genutil.v1beta1.Genesis.GenesisState.getDefaultInstance();
      }

      @java.lang.Override
      public cosmos.genutil.v1beta1.Genesis.GenesisState build() {
        cosmos.genutil.v1beta1.Genesis.GenesisState result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public cosmos.genutil.v1beta1.Genesis.GenesisState buildPartial() {
        cosmos.genutil.v1beta1.Genesis.GenesisState result = new cosmos.genutil.v1beta1.Genesis.GenesisState(this);
        int from_bitField0_ = bitField0_;
        if (((bitField0_ & 0x00000001) != 0)) {
          genTxs_ = java.util.Collections.unmodifiableList(genTxs_);
          bitField0_ = (bitField0_ & ~0x00000001);
        }
        result.genTxs_ = genTxs_;
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
        if (other instanceof cosmos.genutil.v1beta1.Genesis.GenesisState) {
          return mergeFrom((cosmos.genutil.v1beta1.Genesis.GenesisState)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(cosmos.genutil.v1beta1.Genesis.GenesisState other) {
        if (other == cosmos.genutil.v1beta1.Genesis.GenesisState.getDefaultInstance()) return this;
        if (!other.genTxs_.isEmpty()) {
          if (genTxs_.isEmpty()) {
            genTxs_ = other.genTxs_;
            bitField0_ = (bitField0_ & ~0x00000001);
          } else {
            ensureGenTxsIsMutable();
            genTxs_.addAll(other.genTxs_);
          }
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
        cosmos.genutil.v1beta1.Genesis.GenesisState parsedMessage = null;
        try {
          parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          parsedMessage = (cosmos.genutil.v1beta1.Genesis.GenesisState) e.getUnfinishedMessage();
          throw e.unwrapIOException();
        } finally {
          if (parsedMessage != null) {
            mergeFrom(parsedMessage);
          }
        }
        return this;
      }
      private int bitField0_;

      private java.util.List<com.google.protobuf.ByteString> genTxs_ = java.util.Collections.emptyList();
      private void ensureGenTxsIsMutable() {
        if (!((bitField0_ & 0x00000001) != 0)) {
          genTxs_ = new java.util.ArrayList<com.google.protobuf.ByteString>(genTxs_);
          bitField0_ |= 0x00000001;
         }
      }
      /**
       * <pre>
       * gen_txs defines the genesis transactions.
       * </pre>
       *
       * <code>repeated bytes gen_txs = 1 [(.gogoproto.jsontag) = "gentxs", (.gogoproto.moretags) = "yaml:&#92;"gentxs&#92;"", (.gogoproto.casttype) = "encoding/json.RawMessage"];</code>
       */
      public java.util.List<com.google.protobuf.ByteString>
          getGenTxsList() {
        return ((bitField0_ & 0x00000001) != 0) ?
                 java.util.Collections.unmodifiableList(genTxs_) : genTxs_;
      }
      /**
       * <pre>
       * gen_txs defines the genesis transactions.
       * </pre>
       *
       * <code>repeated bytes gen_txs = 1 [(.gogoproto.jsontag) = "gentxs", (.gogoproto.moretags) = "yaml:&#92;"gentxs&#92;"", (.gogoproto.casttype) = "encoding/json.RawMessage"];</code>
       */
      public int getGenTxsCount() {
        return genTxs_.size();
      }
      /**
       * <pre>
       * gen_txs defines the genesis transactions.
       * </pre>
       *
       * <code>repeated bytes gen_txs = 1 [(.gogoproto.jsontag) = "gentxs", (.gogoproto.moretags) = "yaml:&#92;"gentxs&#92;"", (.gogoproto.casttype) = "encoding/json.RawMessage"];</code>
       */
      public com.google.protobuf.ByteString getGenTxs(int index) {
        return genTxs_.get(index);
      }
      /**
       * <pre>
       * gen_txs defines the genesis transactions.
       * </pre>
       *
       * <code>repeated bytes gen_txs = 1 [(.gogoproto.jsontag) = "gentxs", (.gogoproto.moretags) = "yaml:&#92;"gentxs&#92;"", (.gogoproto.casttype) = "encoding/json.RawMessage"];</code>
       */
      public Builder setGenTxs(
          int index, com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  ensureGenTxsIsMutable();
        genTxs_.set(index, value);
        onChanged();
        return this;
      }
      /**
       * <pre>
       * gen_txs defines the genesis transactions.
       * </pre>
       *
       * <code>repeated bytes gen_txs = 1 [(.gogoproto.jsontag) = "gentxs", (.gogoproto.moretags) = "yaml:&#92;"gentxs&#92;"", (.gogoproto.casttype) = "encoding/json.RawMessage"];</code>
       */
      public Builder addGenTxs(com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  ensureGenTxsIsMutable();
        genTxs_.add(value);
        onChanged();
        return this;
      }
      /**
       * <pre>
       * gen_txs defines the genesis transactions.
       * </pre>
       *
       * <code>repeated bytes gen_txs = 1 [(.gogoproto.jsontag) = "gentxs", (.gogoproto.moretags) = "yaml:&#92;"gentxs&#92;"", (.gogoproto.casttype) = "encoding/json.RawMessage"];</code>
       */
      public Builder addAllGenTxs(
          java.lang.Iterable<? extends com.google.protobuf.ByteString> values) {
        ensureGenTxsIsMutable();
        com.google.protobuf.AbstractMessageLite.Builder.addAll(
            values, genTxs_);
        onChanged();
        return this;
      }
      /**
       * <pre>
       * gen_txs defines the genesis transactions.
       * </pre>
       *
       * <code>repeated bytes gen_txs = 1 [(.gogoproto.jsontag) = "gentxs", (.gogoproto.moretags) = "yaml:&#92;"gentxs&#92;"", (.gogoproto.casttype) = "encoding/json.RawMessage"];</code>
       */
      public Builder clearGenTxs() {
        genTxs_ = java.util.Collections.emptyList();
        bitField0_ = (bitField0_ & ~0x00000001);
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


      // @@protoc_insertion_point(builder_scope:cosmos.genutil.v1beta1.GenesisState)
    }

    // @@protoc_insertion_point(class_scope:cosmos.genutil.v1beta1.GenesisState)
    private static final cosmos.genutil.v1beta1.Genesis.GenesisState DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new cosmos.genutil.v1beta1.Genesis.GenesisState();
    }

    public static cosmos.genutil.v1beta1.Genesis.GenesisState getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<GenesisState>
        PARSER = new com.google.protobuf.AbstractParser<GenesisState>() {
      @java.lang.Override
      public GenesisState parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        return new GenesisState(input, extensionRegistry);
      }
    };

    public static com.google.protobuf.Parser<GenesisState> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<GenesisState> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public cosmos.genutil.v1beta1.Genesis.GenesisState getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_cosmos_genutil_v1beta1_GenesisState_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_cosmos_genutil_v1beta1_GenesisState_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n$cosmos/genutil/v1beta1/genesis.proto\022\026" +
      "cosmos.genutil.v1beta1\032\024gogoproto/gogo.p" +
      "roto\"X\n\014GenesisState\022H\n\007gen_txs\030\001 \003(\014B7\372" +
      "\336\037\030encoding/json.RawMessage\352\336\037\006gentxs\362\336\037" +
      "\ryaml:\"gentxs\"B.Z,github.com/cosmos/cosm" +
      "os-sdk/x/genutil/typesb\006proto3"
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
    internal_static_cosmos_genutil_v1beta1_GenesisState_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_cosmos_genutil_v1beta1_GenesisState_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_cosmos_genutil_v1beta1_GenesisState_descriptor,
        new java.lang.String[] { "GenTxs", });
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(com.google.protobuf.GoGoProtos.casttype);
    registry.add(com.google.protobuf.GoGoProtos.jsontag);
    registry.add(com.google.protobuf.GoGoProtos.moretags);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
    com.google.protobuf.GoGoProtos.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}