using SheeterJson;
using System;
using System.IO;

namespace verifycs {
    public static class VerifyJson {
        public static void Test() {
            var path = Path.Combine("json", "data"); // 工作目錄在target

            verifyJsonFrom1(path);
            verifyJsonFrom2(path);
            verifyJsonMerge1(path);
            verifyJsonMerge2(path);
        }

        private static string read(string path, string name) {
            return File.ReadAllText(Path.Combine(path, name));
        }

        private static void assert(bool condition) {
            if (condition == false)
                throw new Exception("verify json: verify failed");
        }

        private static void verifyJsonFrom1(string path) {
            var reader = new VerifyData1Reader();

            if (string.IsNullOrEmpty(reader.FromData(read(path, reader.DataFile()))) == false) {
                throw new Exception("verify json: read failed");
            } // if

            VerifyData1 actual;

            assert(reader.TryGetValue(1, out actual));
            assert(actual != null);
            assert(actual.Key == 1);
            assert(actual.Hide == false);
            assert(actual.Enable == true);
            assert(actual.Name == "名稱1");
            assert(actual.Reward.Desc == "獎勵說明1");
            assert(actual.Reward.Gold == 100);
            assert(actual.Reward.Diamond == 10);
            assert(actual.Reward.Crystal == 199);
            assert(actual.Reward.FelIron == 5);
            assert(actual.Reward.Atium == 1);
            assert(actual.Reward.Item.Length == 3);
            assert(actual.Reward.Item[0].ItemID == 10001);
            assert(actual.Reward.Item[0].Type == 1);
            assert(actual.Reward.Item[0].Count == 1);
            assert(actual.Reward.Item[1].ItemID == 10002);
            assert(actual.Reward.Item[1].Type == 2);
            assert(actual.Reward.Item[1].Count == 2);
            assert(actual.Reward.Item[2].ItemID == 0);
            assert(actual.Reward.Item[2].Type == 0);
            assert(actual.Reward.Item[2].Count == 0);

            assert(reader.TryGetValue(2, out actual));
            assert(actual != null);
            assert(actual.Key == 2);
            assert(actual.Hide == false);
            assert(actual.Enable == false);
            assert(actual.Name == "名稱2");
            assert(actual.Reward.Desc == "獎勵說明2");
            assert(actual.Reward.Gold == 200);
            assert(actual.Reward.Diamond == 20);
            assert(actual.Reward.Crystal == 299);
            assert(actual.Reward.FelIron == 10);
            assert(actual.Reward.Atium == 2);
            assert(actual.Reward.Item.Length == 3);
            assert(actual.Reward.Item[0].ItemID == 10001);
            assert(actual.Reward.Item[0].Type == 1);
            assert(actual.Reward.Item[0].Count == 1);
            assert(actual.Reward.Item[1].ItemID == 10002);
            assert(actual.Reward.Item[1].Type == 2);
            assert(actual.Reward.Item[1].Count == 2);
            assert(actual.Reward.Item[2].ItemID == 10003);
            assert(actual.Reward.Item[2].Type == 3);
            assert(actual.Reward.Item[2].Count == 3);

            assert(reader.TryGetValue(3, out actual) == false);
            assert(actual == null);

            Console.WriteLine("verify json from 1: success");
        }

        private static void verifyJsonFrom2(string path) {
            var reader = new VerifyData2Reader();

            if (string.IsNullOrEmpty(reader.FromData(read(path, reader.DataFile()))) == false) {
                throw new Exception("verify json: read failed");
            } // if

            VerifyData2 actual;

            assert(reader.TryGetValue(1, out actual));
            assert(actual != null);
            assert(actual.Key == 1);
            assert(actual.Hide == false);
            assert(actual.Enable == true);
            assert(actual.Name == "名稱1");
            assert(actual.Reward.Desc == "獎勵說明1");
            assert(actual.Reward.Gold == 100);
            assert(actual.Reward.Diamond == 10);
            assert(actual.Reward.Crystal == 0);
            assert(actual.Reward.FelIron == 0);
            assert(actual.Reward.Atium == 0);
            assert(actual.Reward.Item.Length == 3);
            assert(actual.Reward.Item[0].ItemID == 10001);
            assert(actual.Reward.Item[0].Type == 1);
            assert(actual.Reward.Item[0].Count == 1);
            assert(actual.Reward.Item[1].ItemID == 10002);
            assert(actual.Reward.Item[1].Type == 2);
            assert(actual.Reward.Item[1].Count == 2);
            assert(actual.Reward.Item[2].ItemID == 0);
            assert(actual.Reward.Item[2].Type == 0);
            assert(actual.Reward.Item[2].Count == 0);

            assert(reader.TryGetValue(2, out actual));
            assert(actual != null);
            assert(actual.Key == 2);
            assert(actual.Hide == false);
            assert(actual.Enable == false);
            assert(actual.Name == "名稱2");
            assert(actual.Reward.Desc == "獎勵說明2");
            assert(actual.Reward.Gold == 200);
            assert(actual.Reward.Diamond == 20);
            assert(actual.Reward.Crystal == 0);
            assert(actual.Reward.FelIron == 0);
            assert(actual.Reward.Atium == 0);
            assert(actual.Reward.Item.Length == 3);
            assert(actual.Reward.Item[0].ItemID == 10001);
            assert(actual.Reward.Item[0].Type == 1);
            assert(actual.Reward.Item[0].Count == 1);
            assert(actual.Reward.Item[1].ItemID == 10002);
            assert(actual.Reward.Item[1].Type == 2);
            assert(actual.Reward.Item[1].Count == 2);
            assert(actual.Reward.Item[2].ItemID == 10003);
            assert(actual.Reward.Item[2].Type == 3);
            assert(actual.Reward.Item[2].Count == 3);

            assert(reader.TryGetValue(3, out actual) == false);
            assert(actual == null);

            Console.WriteLine("verify json from 2: success");
        }

        private static void verifyJsonMerge1(string path) {
            var reader = new VerifyData1Reader();

            if (string.IsNullOrEmpty(reader.MergeData(read(path, reader.DataFile()))) == false) {
                throw new Exception("verify json: read failed");
            } // if

            VerifyData1 actual;

            assert(reader.TryGetValue(1, out actual));
            assert(actual != null);
            assert(actual.Key == 1);
            assert(actual.Hide == false);
            assert(actual.Enable == true);
            assert(actual.Name == "名稱1");
            assert(actual.Reward.Desc == "獎勵說明1");
            assert(actual.Reward.Gold == 100);
            assert(actual.Reward.Diamond == 10);
            assert(actual.Reward.Crystal == 199);
            assert(actual.Reward.FelIron == 5);
            assert(actual.Reward.Atium == 1);
            assert(actual.Reward.Item.Length == 3);
            assert(actual.Reward.Item[0].ItemID == 10001);
            assert(actual.Reward.Item[0].Type == 1);
            assert(actual.Reward.Item[0].Count == 1);
            assert(actual.Reward.Item[1].ItemID == 10002);
            assert(actual.Reward.Item[1].Type == 2);
            assert(actual.Reward.Item[1].Count == 2);
            assert(actual.Reward.Item[2].ItemID == 0);
            assert(actual.Reward.Item[2].Type == 0);
            assert(actual.Reward.Item[2].Count == 0);

            assert(reader.TryGetValue(2, out actual));
            assert(actual != null);
            assert(actual.Key == 2);
            assert(actual.Hide == false);
            assert(actual.Enable == false);
            assert(actual.Name == "名稱2");
            assert(actual.Reward.Desc == "獎勵說明2");
            assert(actual.Reward.Gold == 200);
            assert(actual.Reward.Diamond == 20);
            assert(actual.Reward.Crystal == 299);
            assert(actual.Reward.FelIron == 10);
            assert(actual.Reward.Atium == 2);
            assert(actual.Reward.Item.Length == 3);
            assert(actual.Reward.Item[0].ItemID == 10001);
            assert(actual.Reward.Item[0].Type == 1);
            assert(actual.Reward.Item[0].Count == 1);
            assert(actual.Reward.Item[1].ItemID == 10002);
            assert(actual.Reward.Item[1].Type == 2);
            assert(actual.Reward.Item[1].Count == 2);
            assert(actual.Reward.Item[2].ItemID == 10003);
            assert(actual.Reward.Item[2].Type == 3);
            assert(actual.Reward.Item[2].Count == 3);

            assert(reader.TryGetValue(3, out actual) == false);
            assert(actual == null);

            Console.WriteLine("verify json merge 1: success");
        }

        private static void verifyJsonMerge2(string path) {
            var reader = new VerifyData2Reader();

            if (string.IsNullOrEmpty(reader.MergeData(read(path, reader.DataFile()))) == false) {
                throw new Exception("verify json: read failed");
            } // if

            VerifyData2 actual;

            assert(reader.TryGetValue(1, out actual));
            assert(actual != null);
            assert(actual.Key == 1);
            assert(actual.Hide == false);
            assert(actual.Enable == true);
            assert(actual.Name == "名稱1");
            assert(actual.Reward.Desc == "獎勵說明1");
            assert(actual.Reward.Gold == 100);
            assert(actual.Reward.Diamond == 10);
            assert(actual.Reward.Crystal == 0);
            assert(actual.Reward.FelIron == 0);
            assert(actual.Reward.Atium == 0);
            assert(actual.Reward.Item.Length == 3);
            assert(actual.Reward.Item[0].ItemID == 10001);
            assert(actual.Reward.Item[0].Type == 1);
            assert(actual.Reward.Item[0].Count == 1);
            assert(actual.Reward.Item[1].ItemID == 10002);
            assert(actual.Reward.Item[1].Type == 2);
            assert(actual.Reward.Item[1].Count == 2);
            assert(actual.Reward.Item[2].ItemID == 0);
            assert(actual.Reward.Item[2].Type == 0);
            assert(actual.Reward.Item[2].Count == 0);

            assert(reader.TryGetValue(2, out actual));
            assert(actual != null);
            assert(actual.Key == 2);
            assert(actual.Hide == false);
            assert(actual.Enable == false);
            assert(actual.Name == "名稱2");
            assert(actual.Reward.Desc == "獎勵說明2");
            assert(actual.Reward.Gold == 200);
            assert(actual.Reward.Diamond == 20);
            assert(actual.Reward.Crystal == 0);
            assert(actual.Reward.FelIron == 0);
            assert(actual.Reward.Atium == 0);
            assert(actual.Reward.Item.Length == 3);
            assert(actual.Reward.Item[0].ItemID == 10001);
            assert(actual.Reward.Item[0].Type == 1);
            assert(actual.Reward.Item[0].Count == 1);
            assert(actual.Reward.Item[1].ItemID == 10002);
            assert(actual.Reward.Item[1].Type == 2);
            assert(actual.Reward.Item[1].Count == 2);
            assert(actual.Reward.Item[2].ItemID == 10003);
            assert(actual.Reward.Item[2].Type == 3);
            assert(actual.Reward.Item[2].Count == 3);

            assert(reader.TryGetValue(3, out actual) == false);
            assert(actual == null);

            Console.WriteLine("verify json merge 2: success");
        }
    }
}