namespace verifycs {
    public class Verifycs {
        public static void Main() {
            VerifyJson.verifyJsonFrom();
            VerifyJson.verifyJsonMerge();
            VerifyProto.verifyProtoFrom();
            VerifyProto.verifyProtoMerge();
        }
    }
}