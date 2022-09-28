// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: reward.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Sheeter {

  /// <summary>Holder for reflection information generated from reward.proto</summary>
  public static partial class RewardReflection {

    #region Descriptor
    /// <summary>File descriptor for reward.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static RewardReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CgxyZXdhcmQucHJvdG8SB3NoZWV0ZXIaCml0ZW0ucHJvdG8ixQEKBlJld2Fy",
            "ZBIbCgRJdGVtGAEgAygLMg0uc2hlZXRlci5JdGVtEhIKBUF0aXVtGAIgASgD",
            "SACIAQESFAoHQ3J5c3RhbBgDIAEoA0gBiAEBEhQKB0RpYW1vbmQYBCABKANI",
            "AogBARIUCgdGZWxJcm9uGAUgASgDSAOIAQESEQoER29sZBgGIAEoA0gEiAEB",
            "QggKBl9BdGl1bUIKCghfQ3J5c3RhbEIKCghfRGlhbW9uZEIKCghfRmVsSXJv",
            "bkIHCgVfR29sZCJ+CgxSZXdhcmRTdG9yZXISLwoFRGF0YXMYASADKAsyIC5z",
            "aGVldGVyLlJld2FyZFN0b3Jlci5EYXRhc0VudHJ5Gj0KCkRhdGFzRW50cnkS",
            "CwoDa2V5GAEgASgDEh4KBXZhbHVlGAIgASgLMg8uc2hlZXRlci5SZXdhcmQ6",
            "AjgBQgtaCS47c2hlZXRlcmIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Sheeter.ItemReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Sheeter.Reward), global::Sheeter.Reward.Parser, new[]{ "Item", "Atium", "Crystal", "Diamond", "FelIron", "Gold" }, new[]{ "Atium", "Crystal", "Diamond", "FelIron", "Gold" }, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Sheeter.RewardStorer), global::Sheeter.RewardStorer.Parser, new[]{ "Datas" }, null, null, null, new pbr::GeneratedClrTypeInfo[] { null, })
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class Reward : pb::IMessage<Reward>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<Reward> _parser = new pb::MessageParser<Reward>(() => new Reward());
    private pb::UnknownFieldSet _unknownFields;
    private int _hasBits0;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<Reward> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Sheeter.RewardReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Reward() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Reward(Reward other) : this() {
      _hasBits0 = other._hasBits0;
      item_ = other.item_.Clone();
      atium_ = other.atium_;
      crystal_ = other.crystal_;
      diamond_ = other.diamond_;
      felIron_ = other.felIron_;
      gold_ = other.gold_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Reward Clone() {
      return new Reward(this);
    }

    /// <summary>Field number for the "Item" field.</summary>
    public const int ItemFieldNumber = 1;
    private static readonly pb::FieldCodec<global::Sheeter.Item> _repeated_item_codec
        = pb::FieldCodec.ForMessage(10, global::Sheeter.Item.Parser);
    private readonly pbc::RepeatedField<global::Sheeter.Item> item_ = new pbc::RepeatedField<global::Sheeter.Item>();
    /// <summary>
    /// 
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::RepeatedField<global::Sheeter.Item> Item {
      get { return item_; }
    }

    /// <summary>Field number for the "Atium" field.</summary>
    public const int AtiumFieldNumber = 2;
    private long atium_;
    /// <summary>
    /// 天金
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public long Atium {
      get { if ((_hasBits0 & 1) != 0) { return atium_; } else { return 0L; } }
      set {
        _hasBits0 |= 1;
        atium_ = value;
      }
    }
    /// <summary>Gets whether the "Atium" field is set</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool HasAtium {
      get { return (_hasBits0 & 1) != 0; }
    }
    /// <summary>Clears the value of the "Atium" field</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void ClearAtium() {
      _hasBits0 &= ~1;
    }

    /// <summary>Field number for the "Crystal" field.</summary>
    public const int CrystalFieldNumber = 3;
    private long crystal_;
    /// <summary>
    /// 魔晶
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public long Crystal {
      get { if ((_hasBits0 & 2) != 0) { return crystal_; } else { return 0L; } }
      set {
        _hasBits0 |= 2;
        crystal_ = value;
      }
    }
    /// <summary>Gets whether the "Crystal" field is set</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool HasCrystal {
      get { return (_hasBits0 & 2) != 0; }
    }
    /// <summary>Clears the value of the "Crystal" field</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void ClearCrystal() {
      _hasBits0 &= ~2;
    }

    /// <summary>Field number for the "Diamond" field.</summary>
    public const int DiamondFieldNumber = 4;
    private long diamond_;
    /// <summary>
    /// 鑽石
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public long Diamond {
      get { if ((_hasBits0 & 4) != 0) { return diamond_; } else { return 0L; } }
      set {
        _hasBits0 |= 4;
        diamond_ = value;
      }
    }
    /// <summary>Gets whether the "Diamond" field is set</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool HasDiamond {
      get { return (_hasBits0 & 4) != 0; }
    }
    /// <summary>Clears the value of the "Diamond" field</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void ClearDiamond() {
      _hasBits0 &= ~4;
    }

    /// <summary>Field number for the "FelIron" field.</summary>
    public const int FelIronFieldNumber = 5;
    private long felIron_;
    /// <summary>
    /// 精鐵
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public long FelIron {
      get { if ((_hasBits0 & 8) != 0) { return felIron_; } else { return 0L; } }
      set {
        _hasBits0 |= 8;
        felIron_ = value;
      }
    }
    /// <summary>Gets whether the "FelIron" field is set</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool HasFelIron {
      get { return (_hasBits0 & 8) != 0; }
    }
    /// <summary>Clears the value of the "FelIron" field</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void ClearFelIron() {
      _hasBits0 &= ~8;
    }

    /// <summary>Field number for the "Gold" field.</summary>
    public const int GoldFieldNumber = 6;
    private long gold_;
    /// <summary>
    /// 金幣
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public long Gold {
      get { if ((_hasBits0 & 16) != 0) { return gold_; } else { return 0L; } }
      set {
        _hasBits0 |= 16;
        gold_ = value;
      }
    }
    /// <summary>Gets whether the "Gold" field is set</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool HasGold {
      get { return (_hasBits0 & 16) != 0; }
    }
    /// <summary>Clears the value of the "Gold" field</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void ClearGold() {
      _hasBits0 &= ~16;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as Reward);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(Reward other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if(!item_.Equals(other.item_)) return false;
      if (Atium != other.Atium) return false;
      if (Crystal != other.Crystal) return false;
      if (Diamond != other.Diamond) return false;
      if (FelIron != other.FelIron) return false;
      if (Gold != other.Gold) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      hash ^= item_.GetHashCode();
      if (HasAtium) hash ^= Atium.GetHashCode();
      if (HasCrystal) hash ^= Crystal.GetHashCode();
      if (HasDiamond) hash ^= Diamond.GetHashCode();
      if (HasFelIron) hash ^= FelIron.GetHashCode();
      if (HasGold) hash ^= Gold.GetHashCode();
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
      item_.WriteTo(output, _repeated_item_codec);
      if (HasAtium) {
        output.WriteRawTag(16);
        output.WriteInt64(Atium);
      }
      if (HasCrystal) {
        output.WriteRawTag(24);
        output.WriteInt64(Crystal);
      }
      if (HasDiamond) {
        output.WriteRawTag(32);
        output.WriteInt64(Diamond);
      }
      if (HasFelIron) {
        output.WriteRawTag(40);
        output.WriteInt64(FelIron);
      }
      if (HasGold) {
        output.WriteRawTag(48);
        output.WriteInt64(Gold);
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
      item_.WriteTo(ref output, _repeated_item_codec);
      if (HasAtium) {
        output.WriteRawTag(16);
        output.WriteInt64(Atium);
      }
      if (HasCrystal) {
        output.WriteRawTag(24);
        output.WriteInt64(Crystal);
      }
      if (HasDiamond) {
        output.WriteRawTag(32);
        output.WriteInt64(Diamond);
      }
      if (HasFelIron) {
        output.WriteRawTag(40);
        output.WriteInt64(FelIron);
      }
      if (HasGold) {
        output.WriteRawTag(48);
        output.WriteInt64(Gold);
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
      size += item_.CalculateSize(_repeated_item_codec);
      if (HasAtium) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(Atium);
      }
      if (HasCrystal) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(Crystal);
      }
      if (HasDiamond) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(Diamond);
      }
      if (HasFelIron) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(FelIron);
      }
      if (HasGold) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(Gold);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(Reward other) {
      if (other == null) {
        return;
      }
      item_.Add(other.item_);
      if (other.HasAtium) {
        Atium = other.Atium;
      }
      if (other.HasCrystal) {
        Crystal = other.Crystal;
      }
      if (other.HasDiamond) {
        Diamond = other.Diamond;
      }
      if (other.HasFelIron) {
        FelIron = other.FelIron;
      }
      if (other.HasGold) {
        Gold = other.Gold;
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
            item_.AddEntriesFrom(input, _repeated_item_codec);
            break;
          }
          case 16: {
            Atium = input.ReadInt64();
            break;
          }
          case 24: {
            Crystal = input.ReadInt64();
            break;
          }
          case 32: {
            Diamond = input.ReadInt64();
            break;
          }
          case 40: {
            FelIron = input.ReadInt64();
            break;
          }
          case 48: {
            Gold = input.ReadInt64();
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
            item_.AddEntriesFrom(ref input, _repeated_item_codec);
            break;
          }
          case 16: {
            Atium = input.ReadInt64();
            break;
          }
          case 24: {
            Crystal = input.ReadInt64();
            break;
          }
          case 32: {
            Diamond = input.ReadInt64();
            break;
          }
          case 40: {
            FelIron = input.ReadInt64();
            break;
          }
          case 48: {
            Gold = input.ReadInt64();
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class RewardStorer : pb::IMessage<RewardStorer>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<RewardStorer> _parser = new pb::MessageParser<RewardStorer>(() => new RewardStorer());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<RewardStorer> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Sheeter.RewardReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public RewardStorer() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public RewardStorer(RewardStorer other) : this() {
      datas_ = other.datas_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public RewardStorer Clone() {
      return new RewardStorer(this);
    }

    /// <summary>Field number for the "Datas" field.</summary>
    public const int DatasFieldNumber = 1;
    private static readonly pbc::MapField<long, global::Sheeter.Reward>.Codec _map_datas_codec
        = new pbc::MapField<long, global::Sheeter.Reward>.Codec(pb::FieldCodec.ForInt64(8, 0L), pb::FieldCodec.ForMessage(18, global::Sheeter.Reward.Parser), 10);
    private readonly pbc::MapField<long, global::Sheeter.Reward> datas_ = new pbc::MapField<long, global::Sheeter.Reward>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::MapField<long, global::Sheeter.Reward> Datas {
      get { return datas_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as RewardStorer);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(RewardStorer other) {
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
    public void MergeFrom(RewardStorer other) {
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
