namespace verifycs {
    public class Verifycs {
        public static void Main() {
            const int threads = 1000;

            VerifyJson.verifyJsonFrom(threads);
            VerifyJson.verifyJsonMerge(threads);
            VerifyProto.verifyProtoFrom(threads);
            VerifyProto.verifyProtoMerge(threads);
            VerifyEnum.verifyEnum();
        }
    }
}
