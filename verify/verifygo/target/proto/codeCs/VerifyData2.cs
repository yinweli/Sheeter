// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: verifyData2.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Sheeter {

  /// <summary>Holder for reflection information generated from verifyData2.proto</summary>
  public static partial class VerifyData2Reflection {

    #region Descriptor
    /// <summary>File descriptor for verifyData2.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static VerifyData2Reflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChF2ZXJpZnlEYXRhMi5wcm90bxIHc2hlZXRlchoMcmV3YXJkLnByb3RvIpQB",
            "CgtWZXJpZnlEYXRhMhIkCgZSZXdhcmQYASABKAsyDy5zaGVldGVyLlJld2Fy",
            "ZEgAiAEBEhMKBkVuYWJsZRgCIAEoCEgBiAEBEhAKA0tleRgDIAEoA0gCiAEB",
            "EhEKBE5hbWUYBCABKAlIA4gBAUIJCgdfUmV3YXJkQgkKB19FbmFibGVCBgoE",
            "X0tleUIHCgVfTmFtZSKNAQoRVmVyaWZ5RGF0YTJTdG9yZXISNAoFRGF0YXMY",
            "ASADKAsyJS5zaGVldGVyLlZlcmlmeURhdGEyU3RvcmVyLkRhdGFzRW50cnka",
            "QgoKRGF0YXNFbnRyeRILCgNrZXkYASABKAMSIwoFdmFsdWUYAiABKAsyFC5z",
            "aGVldGVyLlZlcmlmeURhdGEyOgI4AUILWgkuO3NoZWV0ZXJiBnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Sheeter.RewardReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Sheeter.VerifyData2), global::Sheeter.VerifyData2.Parser, new[]{ "Reward", "Enable", "Key", "Name" }, new[]{ "Reward", "Enable", "Key", "Name" }, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Sheeter.VerifyData2Storer), global::Sheeter.VerifyData2Storer.Parser, new[]{ "Datas" }, null, null, null, new pbr::GeneratedClrTypeInfo[] { null, })
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class VerifyData2 : pb::IMessage<VerifyData2>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<VerifyData2> _parser = new pb::MessageParser<VerifyData2>(() => new VerifyData2());
    private pb::UnknownFieldSet _unknownFields;
    private int _hasBits0;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<VerifyData2> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Sheeter.VerifyData2Reflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public VerifyData2() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public VerifyData2(VerifyData2 other) : this() {
      _hasBits0 = other._hasBits0;
      reward_ = other.reward_ != null ? other.reward_.Clone() : null;
      enable_ = other.enable_;
      key_ = other.key_;
      name_ = other.name_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public VerifyData2 Clone() {
      return new VerifyData2(this);
    }

    /// <summary>Field number for the "Reward" field.</summary>
    public const int RewardFieldNumber = 1;
    private global::Sheeter.Reward reward_;
    /// <summary>
    /// 
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Sheeter.Reward Reward {
      get { return reward_; }
      set {
        reward_ = value;
      }
    }

    /// <summary>Field number for the "Enable" field.</summary>
    public const int EnableFieldNumber = 2;
    private bool enable_;
    /// <summary>
    /// 是否啟用
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Enable {
      get { if ((_hasBits0 & 1) != 0) { return enable_; } else { return false; } }
      set {
        _hasBits0 |= 1;
        enable_ = value;
      }
    }
    /// <summary>Gets whether the "Enable" field is set</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool HasEnable {
      get { return (_hasBits0 & 1) != 0; }
    }
    /// <summary>Clears the value of the "Enable" field</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void ClearEnable() {
      _hasBits0 &= ~1;
    }

    /// <summary>Field number for the "Key" field.</summary>
    public const int KeyFieldNumber = 3;
    private long key_;
    /// <summary>
    /// 索引
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public long Key {
      get { if ((_hasBits0 & 2) != 0) { return key_; } else { return 0L; } }
      set {
        _hasBits0 |= 2;
        key_ = value;
      }
    }
    /// <summary>Gets whether the "Key" field is set</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool HasKey {
      get { return (_hasBits0 & 2) != 0; }
    }
    /// <summary>Clears the value of the "Key" field</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void ClearKey() {
      _hasBits0 &= ~2;
    }

    /// <summary>Field number for the "Name" field.</summary>
    public const int NameFieldNumber = 4;
    private string name_;
    /// <summary>
    /// 名稱
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string Name {
      get { return name_ ?? ""; }
      set {
        name_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }
    /// <summary>Gets whether the "Name" field is set</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool HasName {
      get { return name_ != null; }
    }
    /// <summary>Clears the value of the "Name" field</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void ClearName() {
      name_ = null;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as VerifyData2);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(VerifyData2 other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(Reward, other.Reward)) return false;
      if (Enable != other.Enable) return false;
      if (Key != other.Key) return false;
      if (Name != other.Name) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (reward_ != null) hash ^= Reward.GetHashCode();
      if (HasEnable) hash ^= Enable.GetHashCode();
      if (HasKey) hash ^= Key.GetHashCode();
      if (HasName) hash ^= Name.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (reward_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(Reward);
      }
      if (HasEnable) {
        output.WriteRawTag(16);
        output.WriteBool(Enable);
      }
      if (HasKey) {
        output.WriteRawTag(24);
        output.WriteInt64(Key);
      }
      if (HasName) {
        output.WriteRawTag(34);
        output.WriteString(Name);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (reward_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(Reward);
      }
      if (HasEnable) {
        output.WriteRawTag(16);
        output.WriteBool(Enable);
      }
      if (HasKey) {
        output.WriteRawTag(24);
        output.WriteInt64(Key);
      }
      if (HasName) {
        output.WriteRawTag(34);
        output.WriteString(Name);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      if (reward_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Reward);
      }
      if (HasEnable) {
        size += 1 + 1;
      }
      if (HasKey) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(Key);
      }
      if (HasName) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Name);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(VerifyData2 other) {
      if (other == null) {
        return;
      }
      if (other.reward_ != null) {
        if (reward_ == null) {
          Reward = new global::Sheeter.Reward();
        }
        Reward.MergeFrom(other.Reward);
      }
      if (other.HasEnable) {
        Enable = other.Enable;
      }
      if (other.HasKey) {
        Key = other.Key;
      }
      if (other.HasName) {
        Name = other.Name;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            if (reward_ == null) {
              Reward = new global::Sheeter.Reward();
            }
            input.ReadMessage(Reward);
            break;
          }
          case 16: {
            Enable = input.ReadBool();
            break;
          }
          case 24: {
            Key = input.ReadInt64();
            break;
          }
          case 34: {
            Name = input.ReadString();
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10: {
            if (reward_ == null) {
              Reward = new global::Sheeter.Reward();
            }
            input.ReadMessage(Reward);
            break;
          }
          case 16: {
            Enable = input.ReadBool();
            break;
          }
          case 24: {
            Key = input.ReadInt64();
            break;
          }
          case 34: {
            Name = input.ReadString();
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class VerifyData2Storer : pb::IMessage<VerifyData2Storer>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<VerifyData2Storer> _parser = new pb::MessageParser<VerifyData2Storer>(() => new VerifyData2Storer());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<VerifyData2Storer> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Sheeter.VerifyData2Reflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public VerifyData2Storer() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public VerifyData2Storer(VerifyData2Storer other) : this() {
      datas_ = other.datas_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public VerifyData2Storer Clone() {
      return new VerifyData2Storer(this);
    }

    /// <summary>Field number for the "Datas" field.</summary>
    public const int DatasFieldNumber = 1;
    private static readonly pbc::MapField<long, global::Sheeter.VerifyData2>.Codec _map_datas_codec
        = new pbc::MapField<long, global::Sheeter.VerifyData2>.Codec(pb::FieldCodec.ForInt64(8, 0L), pb::FieldCodec.ForMessage(18, global::Sheeter.VerifyData2.Parser), 10);
    private readonly pbc::MapField<long, global::Sheeter.VerifyData2> datas_ = new pbc::MapField<long, global::Sheeter.VerifyData2>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::MapField<long, global::Sheeter.VerifyData2> Datas {
      get { return datas_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as VerifyData2Storer);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(VerifyData2Storer other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!Datas.Equals(other.Datas)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      hash ^= Datas.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      datas_.WriteTo(output, _map_datas_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      datas_.WriteTo(ref output, _map_datas_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      size += datas_.CalculateSize(_map_datas_codec);
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(VerifyData2Storer other) {
      if (other == null) {
        return;
      }
      datas_.Add(other.datas_);
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            datas_.AddEntriesFrom(input, _map_datas_codec);
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10: {
            datas_.AddEntriesFrom(ref input, _map_datas_codec);
            break;
          }
        }
      }
    }
    #endif

  }

  #endregion

}

#endregion Designer generated code
