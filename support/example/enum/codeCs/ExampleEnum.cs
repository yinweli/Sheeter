// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: exampleEnum.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace SheeterEnum {

  /// <summary>Holder for reflection information generated from exampleEnum.proto</summary>
  public static partial class ExampleEnumReflection {

    #region Descriptor
    /// <summary>File descriptor for exampleEnum.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static ExampleEnumReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChFleGFtcGxlRW51bS5wcm90bxILc2hlZXRlckVudW0qLgoLRXhhbXBsZUVu",
            "dW0SCQoFTmFtZTAQABIJCgVOYW1lMRABEgkKBU5hbWUyEAJCD1oNLjtzaGVl",
            "dGVyRW51bWIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::SheeterEnum.ExampleEnum), }, null, null));
    }
    #endregion

  }
  #region Enums
  public enum ExampleEnum {
    /// <summary>
    /// 第0個列舉
    /// </summary>
    [pbr::OriginalName("Name0")] Name0 = 0,
    /// <summary>
    /// 第1個列舉
    /// </summary>
    [pbr::OriginalName("Name1")] Name1 = 1,
    /// <summary>
    /// 第2個列舉
    /// </summary>
    [pbr::OriginalName("Name2")] Name2 = 2,
  }

  #endregion

}

#endregion Designer generated code