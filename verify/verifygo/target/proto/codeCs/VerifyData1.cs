// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: verifyData1.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace SheeterProto {

  /// <summary>Holder for reflection information generated from verifyData1.proto</summary>
  public static partial class VerifyData1Reflection {

    #region Descriptor
    /// <summary>File descriptor for verifyData1.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static VerifyData1Reflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChF2ZXJpZnlEYXRhMS5wcm90bxIMc2hlZXRlclByb3RvGgxyZXdhcmQucHJv",
            "dG8itQEKC1ZlcmlmeURhdGExEikKBlJld2FyZBgBIAEoCzIULnNoZWV0ZXJQ",
            "cm90by5SZXdhcmRIAIgBARITCgZFbmFibGUYAiABKAhIAYgBARIRCgRIaWRl",
            "GAMgASgISAKIAQESEAoDS2V5GAQgASgDSAOIAQESEQoETmFtZRgFIAEoCUgE",
            "iAEBQgkKB19SZXdhcmRCCQoHX0VuYWJsZUIHCgVfSGlkZUIGCgRfS2V5QgcK",
            "BV9OYW1lIpcBChFWZXJpZnlEYXRhMVN0b3JlchI5CgVEYXRhcxgBIAMoCzIq",
            "LnNoZWV0ZXJQcm90by5WZXJpZnlEYXRhMVN0b3Jlci5EYXRhc0VudHJ5GkcK",
            "CkRhdGFzRW50cnkSCwoDa2V5GAEgASgDEigKBXZhbHVlGAIgASgLMhkuc2hl",
            "ZXRlclByb3RvLlZlcmlmeURhdGExOgI4AUIQWg4uO3NoZWV0ZXJQcm90b2IG",
            "cHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::SheeterProto.RewardReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::SheeterProto.VerifyData1), global::SheeterProto.VerifyData1.Parser, new[]{ "Reward", "Enable", "Hide", "Key", "Name" }, new[]{ "Reward", "Enable", "Hide", "Key", "Name" }, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::SheeterProto.VerifyData1Storer), global::SheeterProto.VerifyData1Storer.Parser, new[]{ "Datas" }, null, null, null, new pbr::GeneratedClrTypeInfo[] { null, })
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class VerifyData1 : pb::IMessage<VerifyData1>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<VerifyData1> _parser = new pb::MessageParser<VerifyData1>(() => new VerifyData1());
    private pb::UnknownFieldSet _unknownFields;
    private int _hasBits0;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<VerifyData1> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::SheeterProto.VerifyData1Reflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public VerifyData1() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public VerifyData1(VerifyData1 other) : this() {
      _hasBits0 = other._hasBits0;
      reward_ = other.reward_ != null ? other.reward_.Clone() : null;
      enable_ = other.enable_;
      hide_ = other.hide_;
      key_ = other.key_;
      name_ = other.name_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public VerifyData1 Clone() {
      return new VerifyData1(this);
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

    /// <summary>Field number for the "Hide" field.</summary>
    public const int HideFieldNumber = 3;
    private bool hide_;
    /// <summary>
    /// 隱藏
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Hide {
      get { if ((_hasBits0 & 2) != 0) { return hide_; } else { return false; } }
      set {
        _hasBits0 |= 2;
        hide_ = value;
      }
    }
    /// <summary>Gets whether the "Hide" field is set</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool HasHide {
      get { return (_hasBits0 & 2) != 0; }
    }
    /// <summary>Clears the value of the "Hide" field</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void ClearHide() {
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
      return Equals(other as VerifyData1);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(VerifyData1 other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(Reward, other.Reward)) return false;
      if (Enable != other.Enable) return false;
      if (Hide != other.Hide) return false;
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
      if (HasHide) hash ^= Hide.GetHashCode();
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
      if (HasHide) {
        output.WriteRawTag(24);
        output.WriteBool(Hide);
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
      if (HasHide) {
        output.WriteRawTag(24);
        output.WriteBool(Hide);
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
      if (HasHide) {
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
    public void MergeFrom(VerifyData1 other) {
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
      if (other.HasHide) {
        Hide = other.Hide;
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
            Hide = input.ReadBool();
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
            Hide = input.ReadBool();
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

  public sealed partial class VerifyData1Storer : pb::IMessage<VerifyData1Storer>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<VerifyData1Storer> _parser = new pb::MessageParser<VerifyData1Storer>(() => new VerifyData1Storer());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<VerifyData1Storer> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::SheeterProto.VerifyData1Reflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public VerifyData1Storer() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public VerifyData1Storer(VerifyData1Storer other) : this() {
      datas_ = other.datas_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public VerifyData1Storer Clone() {
      return new VerifyData1Storer(this);
    }

    /// <summary>Field number for the "Datas" field.</summary>
    public const int DatasFieldNumber = 1;
    private static readonly pbc::MapField<long, global::SheeterProto.VerifyData1>.Codec _map_datas_codec
        = new pbc::MapField<long, global::SheeterProto.VerifyData1>.Codec(pb::FieldCodec.ForInt64(8, 0L), pb::FieldCodec.ForMessage(18, global::SheeterProto.VerifyData1.Parser), 10);
    private readonly pbc::MapField<long, global::SheeterProto.VerifyData1> datas_ = new pbc::MapField<long, global::SheeterProto.VerifyData1>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::MapField<long, global::SheeterProto.VerifyData1> Datas {
      get { return datas_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as VerifyData1Storer);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(VerifyData1Storer other) {
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
    public void MergeFrom(VerifyData1Storer other) {
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
