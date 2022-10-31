// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: path3Data.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace SheeterProto {

  /// <summary>Holder for reflection information generated from path3Data.proto</summary>
  public static partial class Path3DataReflection {

    #region Descriptor
    /// <summary>File descriptor for path3Data.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static Path3DataReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Cg9wYXRoM0RhdGEucHJvdG8SDHNoZWV0ZXJQcm90bxoMcmV3YXJkLnByb3Rv",
            "IrcBCglQYXRoM0RhdGESKQoGUmV3YXJkGAEgASgLMhQuc2hlZXRlclByb3Rv",
            "LlJld2FyZEgAiAEBEhMKBkVuYWJsZRgCIAEoCEgBiAEBEhMKBklnbm9yZRgD",
            "IAEoA0gCiAEBEhAKA0tleRgEIAEoA0gDiAEBEhEKBE5hbWUYBSABKAlIBIgB",
            "AUIJCgdfUmV3YXJkQgkKB19FbmFibGVCCQoHX0lnbm9yZUIGCgRfS2V5QgcK",
            "BV9OYW1lIpEBCg9QYXRoM0RhdGFTdG9yZXISNwoFRGF0YXMYASADKAsyKC5z",
            "aGVldGVyUHJvdG8uUGF0aDNEYXRhU3RvcmVyLkRhdGFzRW50cnkaRQoKRGF0",
            "YXNFbnRyeRILCgNrZXkYASABKAMSJgoFdmFsdWUYAiABKAsyFy5zaGVldGVy",
            "UHJvdG8uUGF0aDNEYXRhOgI4AUIQWg4uO3NoZWV0ZXJQcm90b2IGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::SheeterProto.RewardReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::SheeterProto.Path3Data), global::SheeterProto.Path3Data.Parser, new[]{ "Reward", "Enable", "Ignore", "Key", "Name" }, new[]{ "Reward", "Enable", "Ignore", "Key", "Name" }, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::SheeterProto.Path3DataStorer), global::SheeterProto.Path3DataStorer.Parser, new[]{ "Datas" }, null, null, null, new pbr::GeneratedClrTypeInfo[] { null, })
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class Path3Data : pb::IMessage<Path3Data>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<Path3Data> _parser = new pb::MessageParser<Path3Data>(() => new Path3Data());
    private pb::UnknownFieldSet _unknownFields;
    private int _hasBits0;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<Path3Data> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::SheeterProto.Path3DataReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Path3Data() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Path3Data(Path3Data other) : this() {
      _hasBits0 = other._hasBits0;
      reward_ = other.reward_ != null ? other.reward_.Clone() : null;
      enable_ = other.enable_;
      ignore_ = other.ignore_;
      key_ = other.key_;
      name_ = other.name_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Path3Data Clone() {
      return new Path3Data(this);
    }

    /// <summary>Field number for the "Reward" field.</summary>
    public const int RewardFieldNumber = 1;
    private global::SheeterProto.Reward reward_;
    /// <summary>
    /// 
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::SheeterProto.Reward Reward {
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

    /// <summary>Field number for the "Ignore" field.</summary>
    public const int IgnoreFieldNumber = 3;
    private long ignore_;
    /// <summary>
    /// 忽略
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public long Ignore {
      get { if ((_hasBits0 & 2) != 0) { return ignore_; } else { return 0L; } }
      set {
        _hasBits0 |= 2;
        ignore_ = value;
      }
    }
    /// <summary>Gets whether the "Ignore" field is set</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool HasIgnore {
      get { return (_hasBits0 & 2) != 0; }
    }
    /// <summary>Clears the value of the "Ignore" field</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void ClearIgnore() {
      _hasBits0 &= ~2;
    }

    /// <summary>Field number for the "Key" field.</summary>
    public const int KeyFieldNumber = 4;
    private long key_;
    /// <summary>
    /// 索引
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public long Key {
      get { if ((_hasBits0 & 4) != 0) { return key_; } else { return 0L; } }
      set {
        _hasBits0 |= 4;
        key_ = value;
      }
    }
    /// <summary>Gets whether the "Key" field is set</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool HasKey {
      get { return (_hasBits0 & 4) != 0; }
    }
    /// <summary>Clears the value of the "Key" field</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void ClearKey() {
      _hasBits0 &= ~4;
    }

    /// <summary>Field number for the "Name" field.</summary>
    public const int NameFieldNumber = 5;
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
      return Equals(other as Path3Data);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(Path3Data other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(Reward, other.Reward)) return false;
      if (Enable != other.Enable) return false;
      if (Ignore != other.Ignore) return false;
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
      if (HasIgnore) hash ^= Ignore.GetHashCode();
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
      if (HasIgnore) {
        output.WriteRawTag(24);
        output.WriteInt64(Ignore);
      }
      if (HasKey) {
        output.WriteRawTag(32);
        output.WriteInt64(Key);
      }
      if (HasName) {
        output.WriteRawTag(42);
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
      if (HasIgnore) {
        output.WriteRawTag(24);
        output.WriteInt64(Ignore);
      }
      if (HasKey) {
        output.WriteRawTag(32);
        output.WriteInt64(Key);
      }
      if (HasName) {
        output.WriteRawTag(42);
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
      if (HasIgnore) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(Ignore);
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
    public void MergeFrom(Path3Data other) {
      if (other == null) {
        return;
      }
      if (other.reward_ != null) {
        if (reward_ == null) {
          Reward = new global::SheeterProto.Reward();
        }
        Reward.MergeFrom(other.Reward);
      }
      if (other.HasEnable) {
        Enable = other.Enable;
      }
      if (other.HasIgnore) {
        Ignore = other.Ignore;
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
              Reward = new global::SheeterProto.Reward();
            }
            input.ReadMessage(Reward);
            break;
          }
          case 16: {
            Enable = input.ReadBool();
            break;
          }
          case 24: {
            Ignore = input.ReadInt64();
            break;
          }
          case 32: {
            Key = input.ReadInt64();
            break;
          }
          case 42: {
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
              Reward = new global::SheeterProto.Reward();
            }
            input.ReadMessage(Reward);
            break;
          }
          case 16: {
            Enable = input.ReadBool();
            break;
          }
          case 24: {
            Ignore = input.ReadInt64();
            break;
          }
          case 32: {
            Key = input.ReadInt64();
            break;
          }
          case 42: {
            Name = input.ReadString();
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class Path3DataStorer : pb::IMessage<Path3DataStorer>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<Path3DataStorer> _parser = new pb::MessageParser<Path3DataStorer>(() => new Path3DataStorer());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<Path3DataStorer> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::SheeterProto.Path3DataReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Path3DataStorer() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Path3DataStorer(Path3DataStorer other) : this() {
      datas_ = other.datas_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Path3DataStorer Clone() {
      return new Path3DataStorer(this);
    }

    /// <summary>Field number for the "Datas" field.</summary>
    public const int DatasFieldNumber = 1;
    private static readonly pbc::MapField<long, global::SheeterProto.Path3Data>.Codec _map_datas_codec
        = new pbc::MapField<long, global::SheeterProto.Path3Data>.Codec(pb::FieldCodec.ForInt64(8, 0L), pb::FieldCodec.ForMessage(18, global::SheeterProto.Path3Data.Parser), 10);
    private readonly pbc::MapField<long, global::SheeterProto.Path3Data> datas_ = new pbc::MapField<long, global::SheeterProto.Path3Data>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::MapField<long, global::SheeterProto.Path3Data> Datas {
      get { return datas_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as Path3DataStorer);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(Path3DataStorer other) {
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
    public void MergeFrom(Path3DataStorer other) {
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
