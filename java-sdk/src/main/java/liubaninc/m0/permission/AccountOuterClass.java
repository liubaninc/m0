// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: m0/permission/account.proto

package liubaninc.m0.permission;

public final class AccountOuterClass {
  private AccountOuterClass() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface AccountOrBuilder extends
      // @@protoc_insertion_point(interface_extends:liubaninc.m0.permission.Account)
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
     * <code>string address = 2;</code>
     */
    java.lang.String getAddress();
    /**
     * <code>string address = 2;</code>
     */
    com.google.protobuf.ByteString
        getAddressBytes();

    /**
     * <code>repeated string perms = 3;</code>
     */
    java.util.List<java.lang.String>
        getPermsList();
    /**
     * <code>repeated string perms = 3;</code>
     */
    int getPermsCount();
    /**
     * <code>repeated string perms = 3;</code>
     */
    java.lang.String getPerms(int index);
    /**
     * <code>repeated string perms = 3;</code>
     */
    com.google.protobuf.ByteString
        getPermsBytes(int index);
  }
  /**
   * Protobuf type {@code liubaninc.m0.permission.Account}
   */
  public  static final class Account extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:liubaninc.m0.permission.Account)
      AccountOrBuilder {
  private static final long serialVersionUID = 0L;
    // Use Account.newBuilder() to construct.
    private Account(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private Account() {
      creator_ = "";
      address_ = "";
      perms_ = com.google.protobuf.LazyStringArrayList.EMPTY;
    }

    @java.lang.Override
    public final com.google.protobuf.UnknownFieldSet
    getUnknownFields() {
      return this.unknownFields;
    }
    private Account(
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

              address_ = s;
              break;
            }
            case 26: {
              java.lang.String s = input.readStringRequireUtf8();
              if (!((mutable_bitField0_ & 0x00000004) != 0)) {
                perms_ = new com.google.protobuf.LazyStringArrayList();
                mutable_bitField0_ |= 0x00000004;
              }
              perms_.add(s);
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
        if (((mutable_bitField0_ & 0x00000004) != 0)) {
          perms_ = perms_.getUnmodifiableView();
        }
        this.unknownFields = unknownFields.build();
        makeExtensionsImmutable();
      }
    }
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return liubaninc.m0.permission.AccountOuterClass.internal_static_liubaninc_m0_permission_Account_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return liubaninc.m0.permission.AccountOuterClass.internal_static_liubaninc_m0_permission_Account_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              liubaninc.m0.permission.AccountOuterClass.Account.class, liubaninc.m0.permission.AccountOuterClass.Account.Builder.class);
    }

    private int bitField0_;
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

    public static final int ADDRESS_FIELD_NUMBER = 2;
    private volatile java.lang.Object address_;
    /**
     * <code>string address = 2;</code>
     */
    public java.lang.String getAddress() {
      java.lang.Object ref = address_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        address_ = s;
        return s;
      }
    }
    /**
     * <code>string address = 2;</code>
     */
    public com.google.protobuf.ByteString
        getAddressBytes() {
      java.lang.Object ref = address_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        address_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int PERMS_FIELD_NUMBER = 3;
    private com.google.protobuf.LazyStringList perms_;
    /**
     * <code>repeated string perms = 3;</code>
     */
    public com.google.protobuf.ProtocolStringList
        getPermsList() {
      return perms_;
    }
    /**
     * <code>repeated string perms = 3;</code>
     */
    public int getPermsCount() {
      return perms_.size();
    }
    /**
     * <code>repeated string perms = 3;</code>
     */
    public java.lang.String getPerms(int index) {
      return perms_.get(index);
    }
    /**
     * <code>repeated string perms = 3;</code>
     */
    public com.google.protobuf.ByteString
        getPermsBytes(int index) {
      return perms_.getByteString(index);
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
      if (!getAddressBytes().isEmpty()) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 2, address_);
      }
      for (int i = 0; i < perms_.size(); i++) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 3, perms_.getRaw(i));
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
      if (!getAddressBytes().isEmpty()) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(2, address_);
      }
      {
        int dataSize = 0;
        for (int i = 0; i < perms_.size(); i++) {
          dataSize += computeStringSizeNoTag(perms_.getRaw(i));
        }
        size += dataSize;
        size += 1 * getPermsList().size();
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
      if (!(obj instanceof liubaninc.m0.permission.AccountOuterClass.Account)) {
        return super.equals(obj);
      }
      liubaninc.m0.permission.AccountOuterClass.Account other = (liubaninc.m0.permission.AccountOuterClass.Account) obj;

      if (!getCreator()
          .equals(other.getCreator())) return false;
      if (!getAddress()
          .equals(other.getAddress())) return false;
      if (!getPermsList()
          .equals(other.getPermsList())) return false;
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
      hash = (37 * hash) + ADDRESS_FIELD_NUMBER;
      hash = (53 * hash) + getAddress().hashCode();
      if (getPermsCount() > 0) {
        hash = (37 * hash) + PERMS_FIELD_NUMBER;
        hash = (53 * hash) + getPermsList().hashCode();
      }
      hash = (29 * hash) + unknownFields.hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static liubaninc.m0.permission.AccountOuterClass.Account parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static liubaninc.m0.permission.AccountOuterClass.Account parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static liubaninc.m0.permission.AccountOuterClass.Account parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static liubaninc.m0.permission.AccountOuterClass.Account parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static liubaninc.m0.permission.AccountOuterClass.Account parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static liubaninc.m0.permission.AccountOuterClass.Account parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static liubaninc.m0.permission.AccountOuterClass.Account parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static liubaninc.m0.permission.AccountOuterClass.Account parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }
    public static liubaninc.m0.permission.AccountOuterClass.Account parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }
    public static liubaninc.m0.permission.AccountOuterClass.Account parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static liubaninc.m0.permission.AccountOuterClass.Account parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static liubaninc.m0.permission.AccountOuterClass.Account parseFrom(
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
    public static Builder newBuilder(liubaninc.m0.permission.AccountOuterClass.Account prototype) {
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
     * Protobuf type {@code liubaninc.m0.permission.Account}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:liubaninc.m0.permission.Account)
        liubaninc.m0.permission.AccountOuterClass.AccountOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return liubaninc.m0.permission.AccountOuterClass.internal_static_liubaninc_m0_permission_Account_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return liubaninc.m0.permission.AccountOuterClass.internal_static_liubaninc_m0_permission_Account_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                liubaninc.m0.permission.AccountOuterClass.Account.class, liubaninc.m0.permission.AccountOuterClass.Account.Builder.class);
      }

      // Construct using liubaninc.m0.permission.AccountOuterClass.Account.newBuilder()
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

        address_ = "";

        perms_ = com.google.protobuf.LazyStringArrayList.EMPTY;
        bitField0_ = (bitField0_ & ~0x00000004);
        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return liubaninc.m0.permission.AccountOuterClass.internal_static_liubaninc_m0_permission_Account_descriptor;
      }

      @java.lang.Override
      public liubaninc.m0.permission.AccountOuterClass.Account getDefaultInstanceForType() {
        return liubaninc.m0.permission.AccountOuterClass.Account.getDefaultInstance();
      }

      @java.lang.Override
      public liubaninc.m0.permission.AccountOuterClass.Account build() {
        liubaninc.m0.permission.AccountOuterClass.Account result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public liubaninc.m0.permission.AccountOuterClass.Account buildPartial() {
        liubaninc.m0.permission.AccountOuterClass.Account result = new liubaninc.m0.permission.AccountOuterClass.Account(this);
        int from_bitField0_ = bitField0_;
        int to_bitField0_ = 0;
        result.creator_ = creator_;
        result.address_ = address_;
        if (((bitField0_ & 0x00000004) != 0)) {
          perms_ = perms_.getUnmodifiableView();
          bitField0_ = (bitField0_ & ~0x00000004);
        }
        result.perms_ = perms_;
        result.bitField0_ = to_bitField0_;
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
        if (other instanceof liubaninc.m0.permission.AccountOuterClass.Account) {
          return mergeFrom((liubaninc.m0.permission.AccountOuterClass.Account)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(liubaninc.m0.permission.AccountOuterClass.Account other) {
        if (other == liubaninc.m0.permission.AccountOuterClass.Account.getDefaultInstance()) return this;
        if (!other.getCreator().isEmpty()) {
          creator_ = other.creator_;
          onChanged();
        }
        if (!other.getAddress().isEmpty()) {
          address_ = other.address_;
          onChanged();
        }
        if (!other.perms_.isEmpty()) {
          if (perms_.isEmpty()) {
            perms_ = other.perms_;
            bitField0_ = (bitField0_ & ~0x00000004);
          } else {
            ensurePermsIsMutable();
            perms_.addAll(other.perms_);
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
        liubaninc.m0.permission.AccountOuterClass.Account parsedMessage = null;
        try {
          parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          parsedMessage = (liubaninc.m0.permission.AccountOuterClass.Account) e.getUnfinishedMessage();
          throw e.unwrapIOException();
        } finally {
          if (parsedMessage != null) {
            mergeFrom(parsedMessage);
          }
        }
        return this;
      }
      private int bitField0_;

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

      private java.lang.Object address_ = "";
      /**
       * <code>string address = 2;</code>
       */
      public java.lang.String getAddress() {
        java.lang.Object ref = address_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          address_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string address = 2;</code>
       */
      public com.google.protobuf.ByteString
          getAddressBytes() {
        java.lang.Object ref = address_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          address_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string address = 2;</code>
       */
      public Builder setAddress(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        address_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>string address = 2;</code>
       */
      public Builder clearAddress() {
        
        address_ = getDefaultInstance().getAddress();
        onChanged();
        return this;
      }
      /**
       * <code>string address = 2;</code>
       */
      public Builder setAddressBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        address_ = value;
        onChanged();
        return this;
      }

      private com.google.protobuf.LazyStringList perms_ = com.google.protobuf.LazyStringArrayList.EMPTY;
      private void ensurePermsIsMutable() {
        if (!((bitField0_ & 0x00000004) != 0)) {
          perms_ = new com.google.protobuf.LazyStringArrayList(perms_);
          bitField0_ |= 0x00000004;
         }
      }
      /**
       * <code>repeated string perms = 3;</code>
       */
      public com.google.protobuf.ProtocolStringList
          getPermsList() {
        return perms_.getUnmodifiableView();
      }
      /**
       * <code>repeated string perms = 3;</code>
       */
      public int getPermsCount() {
        return perms_.size();
      }
      /**
       * <code>repeated string perms = 3;</code>
       */
      public java.lang.String getPerms(int index) {
        return perms_.get(index);
      }
      /**
       * <code>repeated string perms = 3;</code>
       */
      public com.google.protobuf.ByteString
          getPermsBytes(int index) {
        return perms_.getByteString(index);
      }
      /**
       * <code>repeated string perms = 3;</code>
       */
      public Builder setPerms(
          int index, java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  ensurePermsIsMutable();
        perms_.set(index, value);
        onChanged();
        return this;
      }
      /**
       * <code>repeated string perms = 3;</code>
       */
      public Builder addPerms(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  ensurePermsIsMutable();
        perms_.add(value);
        onChanged();
        return this;
      }
      /**
       * <code>repeated string perms = 3;</code>
       */
      public Builder addAllPerms(
          java.lang.Iterable<java.lang.String> values) {
        ensurePermsIsMutable();
        com.google.protobuf.AbstractMessageLite.Builder.addAll(
            values, perms_);
        onChanged();
        return this;
      }
      /**
       * <code>repeated string perms = 3;</code>
       */
      public Builder clearPerms() {
        perms_ = com.google.protobuf.LazyStringArrayList.EMPTY;
        bitField0_ = (bitField0_ & ~0x00000004);
        onChanged();
        return this;
      }
      /**
       * <code>repeated string perms = 3;</code>
       */
      public Builder addPermsBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        ensurePermsIsMutable();
        perms_.add(value);
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


      // @@protoc_insertion_point(builder_scope:liubaninc.m0.permission.Account)
    }

    // @@protoc_insertion_point(class_scope:liubaninc.m0.permission.Account)
    private static final liubaninc.m0.permission.AccountOuterClass.Account DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new liubaninc.m0.permission.AccountOuterClass.Account();
    }

    public static liubaninc.m0.permission.AccountOuterClass.Account getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<Account>
        PARSER = new com.google.protobuf.AbstractParser<Account>() {
      @java.lang.Override
      public Account parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        return new Account(input, extensionRegistry);
      }
    };

    public static com.google.protobuf.Parser<Account> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<Account> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public liubaninc.m0.permission.AccountOuterClass.Account getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_liubaninc_m0_permission_Account_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_liubaninc_m0_permission_Account_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\033m0/permission/account.proto\022\027liubaninc" +
      ".m0.permission\032\024gogoproto/gogo.proto\":\n\007" +
      "Account\022\017\n\007creator\030\001 \001(\t\022\017\n\007address\030\002 \001(" +
      "\t\022\r\n\005perms\030\003 \003(\tB,Z*github.com/liubaninc" +
      "/m0/x/permission/typesb\006proto3"
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
    internal_static_liubaninc_m0_permission_Account_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_liubaninc_m0_permission_Account_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_liubaninc_m0_permission_Account_descriptor,
        new java.lang.String[] { "Creator", "Address", "Perms", });
    com.google.protobuf.GoGoProtos.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
