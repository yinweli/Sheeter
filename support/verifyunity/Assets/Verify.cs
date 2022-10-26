using UnityEngine;

namespace verifycs {
    public class Verify : MonoBehaviour {
        void Start() {
            const int threads = 1000;

            VerifyJson.verifyJsonFrom(threads);
            VerifyJson.verifyJsonMerge(threads);
            VerifyProto.verifyProtoFrom(threads);
            VerifyProto.verifyProtoMerge(threads);
            VerifyEnum.verifyEnum();
        }
    }
}
