// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: m0/validator/params.proto

package liubaninc.m0.validator;

public final class ParamsOuterClass {
  private ParamsOuterClass() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface ParamsOrBuilder extends
      // @@protoc_insertion_point(interface_extends:liubaninc.m0.validator.Params)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <pre>
     * unbonding_time is the time duration of unbonding.
     * </pre>
     *
     * <code>.google.protobuf.Duration unbonding_time = 1 [(.gogoproto.nullable) = false, (.gogoproto.moretags) = "yaml:&#92;"unbonding_time&#92;"", (.gogoproto.stdduration) = true];</code>
     */
    boolean hasUnbondingTime();
    /**
     * <pre>
     * unbonding_time is the time duration of unbonding.
     * </pre>
     *
     * <code>.google.protobuf.Duration unbonding_time = 1 [(.gogoproto.nullable) = false, (.gogoproto.moretags) = "yaml:&#92;"unbonding_time&#92;"", (.gogoproto.stdduration) = true];</code>
     */
    com.google.protobuf.Duration getUnbondingTime();
    /**
     * <pre>
     * unbonding_time is the time duration of unbonding.
     * </pre>
     *
     * <code>.google.protobuf.Duration unbonding_time = 1 [(.gogoproto.nullable) = false, (.gogoproto.moretags) = "yaml:&#92;"unbonding_time&#92;"", (.gogoproto.stdduration) = true];</code>
     */
    com.google.protobuf.DurationOrBuilder getUnbondingTimeOrBuilder();

    /**
     * <pre>
     * max_validators is the maximum number of validators.
     * </pre>
     *
     * <code>uint32 max_validators = 2 [(.gogoproto.moretags) = "yaml:&#92;"max_validators&#92;""];</code>
     */
    int getMaxValidators();

    /**
     * <pre>
     * min_validators is the minimum number of validators.
     * </pre>
     *
     * <code>uint32 min_validators = 3 [(.gogoproto.moretags) = "yaml:&#92;"min_validators&#92;""];</code>
     */
    int getMinValidators();

    /**
     * <pre>
     * historical_entries is the number of historical entries to persist.
     * </pre>
     *
     * <code>uint32 historical_entries = 4 [(.gogoproto.moretags) = "yaml:&#92;"historical_entries&#92;""];</code>
     */
    int getHistoricalEntries();
  }
  /**
   * <pre>
   * Params defines the parameters for the staking module.
   * </pre>
   *
   * Protobuf type {@code liubaninc.m0.validator.Params}
   */
  public  static final class Params extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:liubaninc.m0.validator.Params)
      ParamsOrBuilder {
  private static final long serialVersionUID = 0L;
    // Use Params.newBuilder() to construct.
    private Params(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private Params() {
    }

    @java.lang.Override
    public final com.google.protobuf.UnknownFieldSet
    getUnknownFields() {
      return this.unknownFields;
    }
    private Params(
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
              com.google.protobuf.Duration.Builder subBuilder = null;
              if (unbondingTime_ != null) {
                subBuilder = unbondingTime_.toBuilder();
              }
              unbondingTime_ = input.readMessage(com.google.protobuf.Duration.parser(), extensionRegistry);
              if (subBuilder != null) {
                subBuilder.mergeFrom(unbondingTime_);
                unbondingTime_ = subBuilder.buildPartial();
              }

              break;
            }
            case 16: {

              maxValidators_ = input.readUInt32();
              break;
            }
            case 24: {

              minValidators_ = input.readUInt32();
              break;
            }
            case 32: {

              historicalEntries_ = input.readUInt32();
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
      return liubaninc.m0.validator.ParamsOuterClass.internal_static_liubaninc_m0_validator_Params_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return liubaninc.m0.validator.ParamsOuterClass.internal_static_liubaninc_m0_validator_Params_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              liubaninc.m0.validator.ParamsOuterClass.Params.class, liubaninc.m0.validator.ParamsOuterClass.Params.Builder.class);
    }

    public static final int UNBONDING_TIME_FIELD_NUMBER = 1;
    private com.google.protobuf.Duration unbondingTime_;
    /**
     * <pre>
     * unbonding_time is the time duration of unbonding.
     * </pre>
     *
     * <code>.google.protobuf.Duration unbonding_time = 1 [(.gogoproto.nullable) = false, (.gogoproto.moretags) = "yaml:&#92;"unbonding_time&#92;"", (.gogoproto.stdduration) = true];</code>
     */
    public boolean hasUnbondingTime() {
      return unbondingTime_ != null;
    }
    /**
     * <pre>
     * unbonding_time is the time duration of unbonding.
     * </pre>
     *
     * <code>.google.protobuf.Duration unbonding_time = 1 [(.gogoproto.nullable) = false, (.gogoproto.moretags) = "yaml:&#92;"unbonding_time&#92;"", (.gogoproto.stdduration) = true];</code>
     */
    public com.google.protobuf.Duration getUnbondingTime() {
      return unbondingTime_ == null ? com.google.protobuf.Duration.getDefaultInstance() : unbondingTime_;
    }
    /**
     * <pre>
     * unbonding_time is the time duration of unbonding.
     * </pre>
     *
     * <code>.google.protobuf.Duration unbonding_time = 1 [(.gogoproto.nullable) = false, (.gogoproto.moretags) = "yaml:&#92;"unbonding_time&#92;"", (.gogoproto.stdduration) = true];</code>
     */
    public com.google.protobuf.DurationOrBuilder getUnbondingTimeOrBuilder() {
      return getUnbondingTime();
    }

    public static final int MAX_VALIDATORS_FIELD_NUMBER = 2;
    private int maxValidators_;
    /**
     * <pre>
     * max_validators is the maximum number of validators.
     * </pre>
     *
     * <code>uint32 max_validators = 2 [(.gogoproto.moretags) = "yaml:&#92;"max_validators&#92;""];</code>
     */
    public int getMaxValidators() {
      return maxValidators_;
    }

    public static final int MIN_VALIDATORS_FIELD_NUMBER = 3;
    private int minValidators_;
    /**
     * <pre>
     * min_validators is the minimum number of validators.
     * </pre>
     *
     * <code>uint32 min_validators = 3 [(.gogoproto.moretags) = "yaml:&#92;"min_validators&#92;""];</code>
     */
    public int getMinValidators() {
      return minValidators_;
    }

    public static final int HISTORICAL_ENTRIES_FIELD_NUMBER = 4;
    private int historicalEntries_;
    /**
     * <pre>
     * historical_entries is the number of historical entries to persist.
     * </pre>
     *
     * <code>uint32 historical_entries = 4 [(.gogoproto.moretags) = "yaml:&#92;"historical_entries&#92;""];</code>
     */
    public int getHistoricalEntries() {
      return historicalEntries_;
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
      if (unbondingTime_ != null) {
        output.writeMessage(1, getUnbondingTime());
      }
      if (maxValidators_ != 0) {
        output.writeUInt32(2, maxValidators_);
      }
      if (minValidators_ != 0) {
        output.writeUInt32(3, minValidators_);
      }
      if (historicalEntries_ != 0) {
        output.writeUInt32(4, historicalEntries_);
      }
      unknownFields.writeTo(output);
    }

    @java.lang.Override
    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      if (unbondingTime_ != null) {
        size += com.google.protobuf.CodedOutputStream
          .computeMessageSize(1, getUnbondingTime());
      }
      if (maxValidators_ != 0) {
        size += com.google.protobuf.CodedOutputStream
          .computeUInt32Size(2, maxValidators_);
      }
      if (minValidators_ != 0) {
        size += com.google.protobuf.CodedOutputStream
          .computeUInt32Size(3, minValidators_);
      }
      if (historicalEntries_ != 0) {
        size += com.google.protobuf.CodedOutputStream
          .computeUInt32Size(4, historicalEntries_);
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
      if (!(obj instanceof liubaninc.m0.validator.ParamsOuterClass.Params)) {
        return super.equals(obj);
      }
      liubaninc.m0.validator.ParamsOuterClass.Params other = (liubaninc.m0.validator.ParamsOuterClass.Params) obj;

      if (hasUnbondingTime() != other.hasUnbondingTime()) return false;
      if (hasUnbondingTime()) {
        if (!getUnbondingTime()
            .equals(other.getUnbondingTime())) return false;
      }
      if (getMaxValidators()
          != other.getMaxValidators()) return false;
      if (getMinValidators()
          != other.getMinValidators()) return false;
      if (getHistoricalEntries()
          != other.getHistoricalEntries()) return false;
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
      if (hasUnbondingTime()) {
        hash = (37 * hash) + UNBONDING_TIME_FIELD_NUMBER;
        hash = (53 * hash) + getUnbondingTime().hashCode();
      }
      hash = (37 * hash) + MAX_VALIDATORS_FIELD_NUMBER;
      hash = (53 * hash) + getMaxValidators();
      hash = (37 * hash) + MIN_VALIDATORS_FIELD_NUMBER;
      hash = (53 * hash) + getMinValidators();
      hash = (37 * hash) + HISTORICAL_ENTRIES_FIELD_NUMBER;
      hash = (53 * hash) + getHistoricalEntries();
      hash = (29 * hash) + unknownFields.hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static liubaninc.m0.validator.ParamsOuterClass.Params parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static liubaninc.m0.validator.ParamsOuterClass.Params parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static liubaninc.m0.validator.ParamsOuterClass.Params parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static liubaninc.m0.validator.ParamsOuterClass.Params parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static liubaninc.m0.validator.ParamsOuterClass.Params parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static liubaninc.m0.validator.ParamsOuterClass.Params parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static liubaninc.m0.validator.ParamsOuterClass.Params parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static liubaninc.m0.validator.ParamsOuterClass.Params parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }
    public static liubaninc.m0.validator.ParamsOuterClass.Params parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }
    public static liubaninc.m0.validator.ParamsOuterClass.Params parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static liubaninc.m0.validator.ParamsOuterClass.Params parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static liubaninc.m0.validator.ParamsOuterClass.Params parseFrom(
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
    public static Builder newBuilder(liubaninc.m0.validator.ParamsOuterClass.Params prototype) {
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
     * Params defines the parameters for the staking module.
     * </pre>
     *
     * Protobuf type {@code liubaninc.m0.validator.Params}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:liubaninc.m0.validator.Params)
        liubaninc.m0.validator.ParamsOuterClass.ParamsOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return liubaninc.m0.validator.ParamsOuterClass.internal_static_liubaninc_m0_validator_Params_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return liubaninc.m0.validator.ParamsOuterClass.internal_static_liubaninc_m0_validator_Params_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                liubaninc.m0.validator.ParamsOuterClass.Params.class, liubaninc.m0.validator.ParamsOuterClass.Params.Builder.class);
      }

      // Construct using liubaninc.m0.validator.ParamsOuterClass.Params.newBuilder()
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
        if (unbondingTimeBuilder_ == null) {
          unbondingTime_ = null;
        } else {
          unbondingTime_ = null;
          unbondingTimeBuilder_ = null;
        }
        maxValidators_ = 0;

        minValidators_ = 0;

        historicalEntries_ = 0;

        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return liubaninc.m0.validator.ParamsOuterClass.internal_static_liubaninc_m0_validator_Params_descriptor;
      }

      @java.lang.Override
      public liubaninc.m0.validator.ParamsOuterClass.Params getDefaultInstanceForType() {
        return liubaninc.m0.validator.ParamsOuterClass.Params.getDefaultInstance();
      }

      @java.lang.Override
      public liubaninc.m0.validator.ParamsOuterClass.Params build() {
        liubaninc.m0.validator.ParamsOuterClass.Params result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public liubaninc.m0.validator.ParamsOuterClass.Params buildPartial() {
        liubaninc.m0.validator.ParamsOuterClass.Params result = new liubaninc.m0.validator.ParamsOuterClass.Params(this);
        if (unbondingTimeBuilder_ == null) {
          result.unbondingTime_ = unbondingTime_;
        } else {
          result.unbondingTime_ = unbondingTimeBuilder_.build();
        }
        result.maxValidators_ = maxValidators_;
        result.minValidators_ = minValidators_;
        result.historicalEntries_ = historicalEntries_;
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
        if (other instanceof liubaninc.m0.validator.ParamsOuterClass.Params) {
          return mergeFrom((liubaninc.m0.validator.ParamsOuterClass.Params)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(liubaninc.m0.validator.ParamsOuterClass.Params other) {
        if (other == liubaninc.m0.validator.ParamsOuterClass.Params.getDefaultInstance()) return this;
        if (other.hasUnbondingTime()) {
          mergeUnbondingTime(other.getUnbondingTime());
        }
        if (other.getMaxValidators() != 0) {
          setMaxValidators(other.getMaxValidators());
        }
        if (other.getMinValidators() != 0) {
          setMinValidators(other.getMinValidators());
        }
        if (other.getHistoricalEntries() != 0) {
          setHistoricalEntries(other.getHistoricalEntries());
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
        liubaninc.m0.validator.ParamsOuterClass.Params parsedMessage = null;
        try {
          parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          parsedMessage = (liubaninc.m0.validator.ParamsOuterClass.Params) e.getUnfinishedMessage();
          throw e.unwrapIOException();
        } finally {
          if (parsedMessage != null) {
            mergeFrom(parsedMessage);
          }
        }
        return this;
      }

      private com.google.protobuf.Duration unbondingTime_;
      private com.google.protobuf.SingleFieldBuilderV3<
          com.google.protobuf.Duration, com.google.protobuf.Duration.Builder, com.google.protobuf.DurationOrBuilder> unbondingTimeBuilder_;
      /**
       * <pre>
       * unbonding_time is the time duration of unbonding.
       * </pre>
       *
       * <code>.google.protobuf.Duration unbonding_time = 1 [(.gogoproto.nullable) = false, (.gogoproto.moretags) = "yaml:&#92;"unbonding_time&#92;"", (.gogoproto.stdduration) = true];</code>
       */
      public boolean hasUnbondingTime() {
        return unbondingTimeBuilder_ != null || unbondingTime_ != null;
      }
      /**
       * <pre>
       * unbonding_time is the time duration of unbonding.
       * </pre>
       *
       * <code>.google.protobuf.Duration unbonding_time = 1 [(.gogoproto.nullable) = false, (.gogoproto.moretags) = "yaml:&#92;"unbonding_time&#92;"", (.gogoproto.stdduration) = true];</code>
       */
      public com.google.protobuf.Duration getUnbondingTime() {
        if (unbondingTimeBuilder_ == null) {
          return unbondingTime_ == null ? com.google.protobuf.Duration.getDefaultInstance() : unbondingTime_;
        } else {
          return unbondingTimeBuilder_.getMessage();
        }
      }
      /**
       * <pre>
       * unbonding_time is the time duration of unbonding.
       * </pre>
       *
       * <code>.google.protobuf.Duration unbonding_time = 1 [(.gogoproto.nullable) = false, (.gogoproto.moretags) = "yaml:&#92;"unbonding_time&#92;"", (.gogoproto.stdduration) = true];</code>
       */
      public Builder setUnbondingTime(com.google.protobuf.Duration value) {
        if (unbondingTimeBuilder_ == null) {
          if (value == null) {
            throw new NullPointerException();
          }
          unbondingTime_ = value;
          onChanged();
        } else {
          unbondingTimeBuilder_.setMessage(value);
        }

        return this;
      }
      /**
       * <pre>
       * unbonding_time is the time duration of unbonding.
       * </pre>
       *
       * <code>.google.protobuf.Duration unbonding_time = 1 [(.gogoproto.nullable) = false, (.gogoproto.moretags) = "yaml:&#92;"unbonding_time&#92;"", (.gogoproto.stdduration) = true];</code>
       */
      public Builder setUnbondingTime(
          com.google.protobuf.Duration.Builder builderForValue) {
        if (unbondingTimeBuilder_ == null) {
          unbondingTime_ = builderForValue.build();
          onChanged();
        } else {
          unbondingTimeBuilder_.setMessage(builderForValue.build());
        }

        return this;
      }
      /**
       * <pre>
       * unbonding_time is the time duration of unbonding.
       * </pre>
       *
       * <code>.google.protobuf.Duration unbonding_time = 1 [(.gogoproto.nullable) = false, (.gogoproto.moretags) = "yaml:&#92;"unbonding_time&#92;"", (.gogoproto.stdduration) = true];</code>
       */
      public Builder mergeUnbondingTime(com.google.protobuf.Duration value) {
        if (unbondingTimeBuilder_ == null) {
          if (unbondingTime_ != null) {
            unbondingTime_ =
              com.google.protobuf.Duration.newBuilder(unbondingTime_).mergeFrom(value).buildPartial();
          } else {
            unbondingTime_ = value;
          }
          onChanged();
        } else {
          unbondingTimeBuilder_.mergeFrom(value);
        }

        return this;
      }
      /**
       * <pre>
       * unbonding_time is the time duration of unbonding.
       * </pre>
       *
       * <code>.google.protobuf.Duration unbonding_time = 1 [(.gogoproto.nullable) = false, (.gogoproto.moretags) = "yaml:&#92;"unbonding_time&#92;"", (.gogoproto.stdduration) = true];</code>
       */
      public Builder clearUnbondingTime() {
        if (unbondingTimeBuilder_ == null) {
          unbondingTime_ = null;
          onChanged();
        } else {
          unbondingTime_ = null;
          unbondingTimeBuilder_ = null;
        }

        return this;
      }
      /**
       * <pre>
       * unbonding_time is the time duration of unbonding.
       * </pre>
       *
       * <code>.google.protobuf.Duration unbonding_time = 1 [(.gogoproto.nullable) = false, (.gogoproto.moretags) = "yaml:&#92;"unbonding_time&#92;"", (.gogoproto.stdduration) = true];</code>
       */
      public com.google.protobuf.Duration.Builder getUnbondingTimeBuilder() {
        
        onChanged();
        return getUnbondingTimeFieldBuilder().getBuilder();
      }
      /**
       * <pre>
       * unbonding_time is the time duration of unbonding.
       * </pre>
       *
       * <code>.google.protobuf.Duration unbonding_time = 1 [(.gogoproto.nullable) = false, (.gogoproto.moretags) = "yaml:&#92;"unbonding_time&#92;"", (.gogoproto.stdduration) = true];</code>
       */
      public com.google.protobuf.DurationOrBuilder getUnbondingTimeOrBuilder() {
        if (unbondingTimeBuilder_ != null) {
          return unbondingTimeBuilder_.getMessageOrBuilder();
        } else {
          return unbondingTime_ == null ?
              com.google.protobuf.Duration.getDefaultInstance() : unbondingTime_;
        }
      }
      /**
       * <pre>
       * unbonding_time is the time duration of unbonding.
       * </pre>
       *
       * <code>.google.protobuf.Duration unbonding_time = 1 [(.gogoproto.nullable) = false, (.gogoproto.moretags) = "yaml:&#92;"unbonding_time&#92;"", (.gogoproto.stdduration) = true];</code>
       */
      private com.google.protobuf.SingleFieldBuilderV3<
          com.google.protobuf.Duration, com.google.protobuf.Duration.Builder, com.google.protobuf.DurationOrBuilder> 
          getUnbondingTimeFieldBuilder() {
        if (unbondingTimeBuilder_ == null) {
          unbondingTimeBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
              com.google.protobuf.Duration, com.google.protobuf.Duration.Builder, com.google.protobuf.DurationOrBuilder>(
                  getUnbondingTime(),
                  getParentForChildren(),
                  isClean());
          unbondingTime_ = null;
        }
        return unbondingTimeBuilder_;
      }

      private int maxValidators_ ;
      /**
       * <pre>
       * max_validators is the maximum number of validators.
       * </pre>
       *
       * <code>uint32 max_validators = 2 [(.gogoproto.moretags) = "yaml:&#92;"max_validators&#92;""];</code>
       */
      public int getMaxValidators() {
        return maxValidators_;
      }
      /**
       * <pre>
       * max_validators is the maximum number of validators.
       * </pre>
       *
       * <code>uint32 max_validators = 2 [(.gogoproto.moretags) = "yaml:&#92;"max_validators&#92;""];</code>
       */
      public Builder setMaxValidators(int value) {
        
        maxValidators_ = value;
        onChanged();
        return this;
      }
      /**
       * <pre>
       * max_validators is the maximum number of validators.
       * </pre>
       *
       * <code>uint32 max_validators = 2 [(.gogoproto.moretags) = "yaml:&#92;"max_validators&#92;""];</code>
       */
      public Builder clearMaxValidators() {
        
        maxValidators_ = 0;
        onChanged();
        return this;
      }

      private int minValidators_ ;
      /**
       * <pre>
       * min_validators is the minimum number of validators.
       * </pre>
       *
       * <code>uint32 min_validators = 3 [(.gogoproto.moretags) = "yaml:&#92;"min_validators&#92;""];</code>
       */
      public int getMinValidators() {
        return minValidators_;
      }
      /**
       * <pre>
       * min_validators is the minimum number of validators.
       * </pre>
       *
       * <code>uint32 min_validators = 3 [(.gogoproto.moretags) = "yaml:&#92;"min_validators&#92;""];</code>
       */
      public Builder setMinValidators(int value) {
        
        minValidators_ = value;
        onChanged();
        return this;
      }
      /**
       * <pre>
       * min_validators is the minimum number of validators.
       * </pre>
       *
       * <code>uint32 min_validators = 3 [(.gogoproto.moretags) = "yaml:&#92;"min_validators&#92;""];</code>
       */
      public Builder clearMinValidators() {
        
        minValidators_ = 0;
        onChanged();
        return this;
      }

      private int historicalEntries_ ;
      /**
       * <pre>
       * historical_entries is the number of historical entries to persist.
       * </pre>
       *
       * <code>uint32 historical_entries = 4 [(.gogoproto.moretags) = "yaml:&#92;"historical_entries&#92;""];</code>
       */
      public int getHistoricalEntries() {
        return historicalEntries_;
      }
      /**
       * <pre>
       * historical_entries is the number of historical entries to persist.
       * </pre>
       *
       * <code>uint32 historical_entries = 4 [(.gogoproto.moretags) = "yaml:&#92;"historical_entries&#92;""];</code>
       */
      public Builder setHistoricalEntries(int value) {
        
        historicalEntries_ = value;
        onChanged();
        return this;
      }
      /**
       * <pre>
       * historical_entries is the number of historical entries to persist.
       * </pre>
       *
       * <code>uint32 historical_entries = 4 [(.gogoproto.moretags) = "yaml:&#92;"historical_entries&#92;""];</code>
       */
      public Builder clearHistoricalEntries() {
        
        historicalEntries_ = 0;
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


      // @@protoc_insertion_point(builder_scope:liubaninc.m0.validator.Params)
    }

    // @@protoc_insertion_point(class_scope:liubaninc.m0.validator.Params)
    private static final liubaninc.m0.validator.ParamsOuterClass.Params DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new liubaninc.m0.validator.ParamsOuterClass.Params();
    }

    public static liubaninc.m0.validator.ParamsOuterClass.Params getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<Params>
        PARSER = new com.google.protobuf.AbstractParser<Params>() {
      @java.lang.Override
      public Params parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        return new Params(input, extensionRegistry);
      }
    };

    public static com.google.protobuf.Parser<Params> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<Params> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public liubaninc.m0.validator.ParamsOuterClass.Params getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_liubaninc_m0_validator_Params_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_liubaninc_m0_validator_Params_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\031m0/validator/params.proto\022\026liubaninc.m" +
      "0.validator\032\024gogoproto/gogo.proto\032\036googl" +
      "e/protobuf/duration.proto\"\211\002\n\006Params\022T\n\016" +
      "unbonding_time\030\001 \001(\0132\031.google.protobuf.D" +
      "urationB!\310\336\037\000\230\337\037\001\362\336\037\025yaml:\"unbonding_tim" +
      "e\"\0221\n\016max_validators\030\002 \001(\rB\031\362\336\037\025yaml:\"ma" +
      "x_validators\"\0221\n\016min_validators\030\003 \001(\rB\031\362" +
      "\336\037\025yaml:\"min_validators\"\0229\n\022historical_e" +
      "ntries\030\004 \001(\rB\035\362\336\037\031yaml:\"historical_entri" +
      "es\":\010\350\240\037\001\230\240\037\000B+Z)github.com/liubaninc/m0" +
      "/x/validator/typesb\006proto3"
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
          com.google.protobuf.DurationProto.getDescriptor(),
        }, assigner);
    internal_static_liubaninc_m0_validator_Params_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_liubaninc_m0_validator_Params_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_liubaninc_m0_validator_Params_descriptor,
        new java.lang.String[] { "UnbondingTime", "MaxValidators", "MinValidators", "HistoricalEntries", });
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(com.google.protobuf.GoGoProtos.equal);
    registry.add(com.google.protobuf.GoGoProtos.goprotoStringer);
    registry.add(com.google.protobuf.GoGoProtos.moretags);
    registry.add(com.google.protobuf.GoGoProtos.nullable);
    registry.add(com.google.protobuf.GoGoProtos.stdduration);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
    com.google.protobuf.GoGoProtos.getDescriptor();
    com.google.protobuf.DurationProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}