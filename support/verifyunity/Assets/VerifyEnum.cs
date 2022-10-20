using System;
using UnityEngine;

namespace verifycs {
    public static class VerifyEnum {
        public static void verifyEnum() {
            assert((int)SheeterEnum.VerifyEnum.Name0 == 0);
            assert((int)SheeterEnum.VerifyEnum.Name1 == 1);
            assert((int)SheeterEnum.VerifyEnum.Name2 == 2);
            Debug.Log("verify enum from: success");
        }

        private static void assert(bool condition) {
            if (condition == false)
                throw new Exception("verify enum: verify failed");
        }
    }
}
