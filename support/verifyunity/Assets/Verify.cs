using UnityEngine;

namespace verifycs {
    public class Verify : MonoBehaviour {
        void Start() {
            VerifyJson.verifyJsonFrom();
            VerifyJson.verifyJsonMerge();
            VerifyProto.verifyProtoFrom();
            VerifyProto.verifyProtoMerge();
            VerifyEnum.verifyEnum();
        }
    }
}