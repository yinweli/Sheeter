using UnityEngine;

namespace verifycs {
    public class Verify : MonoBehaviour {
        void Start() {
            VerifyJson.Test();
            VerifyProto.Test();
        }
    }
}