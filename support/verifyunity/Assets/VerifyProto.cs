using SheeterProto;
using System;
using System.IO;
using UnityEngine;

namespace verifycs {
    public static class VerifyProto {
        public static void verifyProtoFrom() {
            var loader = new ProtoFileLoader();
            var depot = new Depot() { Loader = loader };

            assert(depot.FromData());

            VerifyData1 actual1;

            assert(depot.VerifyData1.TryGetValue(1, out actual1));
            assert(actual1 != null);
            assert(actual1.Key == 1);
            assert(actual1.Hide == false);
            assert(actual1.Enable == true);
            assert(actual1.Name == "名稱1");
            assert(actual1.Reward.Desc == "獎勵說明1");
            assert(actual1.Reward.Gold == 100);
            assert(actual1.Reward.Diamond == 10);
            assert(actual1.Reward.Crystal == 199);
            assert(actual1.Reward.FelIron == 5);
            assert(actual1.Reward.Atium == 1);
            assert(actual1.Reward.Item.Count == 3);
            assert(actual1.Reward.Item[0].ItemID == 10001);
            assert(actual1.Reward.Item[0].Type == 1);
            assert(actual1.Reward.Item[0].Count == 1);
            assert(actual1.Reward.Item[1].ItemID == 10002);
            assert(actual1.Reward.Item[1].Type == 2);
            assert(actual1.Reward.Item[1].Count == 2);
            assert(actual1.Reward.Item[2].ItemID == 0);
            assert(actual1.Reward.Item[2].Type == 0);
            assert(actual1.Reward.Item[2].Count == 0);

            assert(depot.VerifyData1.TryGetValue(2, out actual1));
            assert(actual1 != null);
            assert(actual1.Key == 2);
            assert(actual1.Hide == false);
            assert(actual1.Enable == false);
            assert(actual1.Name == "名稱2");
            assert(actual1.Reward.Desc == "獎勵說明2");
            assert(actual1.Reward.Gold == 200);
            assert(actual1.Reward.Diamond == 20);
            assert(actual1.Reward.Crystal == 299);
            assert(actual1.Reward.FelIron == 10);
            assert(actual1.Reward.Atium == 2);
            assert(actual1.Reward.Item.Count == 3);
            assert(actual1.Reward.Item[0].ItemID == 10001);
            assert(actual1.Reward.Item[0].Type == 1);
            assert(actual1.Reward.Item[0].Count == 1);
            assert(actual1.Reward.Item[1].ItemID == 10002);
            assert(actual1.Reward.Item[1].Type == 2);
            assert(actual1.Reward.Item[1].Count == 2);
            assert(actual1.Reward.Item[2].ItemID == 10003);
            assert(actual1.Reward.Item[2].Type == 3);
            assert(actual1.Reward.Item[2].Count == 3);

            assert(depot.VerifyData1.TryGetValue(3, out actual1) == false);
            assert(actual1 == null);

            VerifyData2 actual2;

            assert(depot.VerifyData2.TryGetValue(1, out actual2));
            assert(actual2 != null);
            assert(actual2.Key == 1);
            assert(actual2.Hide == false);
            assert(actual2.Enable == true);
            assert(actual2.Name == "名稱1");
            assert(actual2.Reward.Desc == "獎勵說明1");
            assert(actual2.Reward.Gold == 100);
            assert(actual2.Reward.Diamond == 10);
            assert(actual2.Reward.Crystal == 0);
            assert(actual2.Reward.FelIron == 0);
            assert(actual2.Reward.Atium == 0);
            assert(actual2.Reward.Item.Count == 3);
            assert(actual2.Reward.Item[0].ItemID == 10001);
            assert(actual2.Reward.Item[0].Type == 1);
            assert(actual2.Reward.Item[0].Count == 1);
            assert(actual2.Reward.Item[1].ItemID == 10002);
            assert(actual2.Reward.Item[1].Type == 2);
            assert(actual2.Reward.Item[1].Count == 2);
            assert(actual2.Reward.Item[2].ItemID == 0);
            assert(actual2.Reward.Item[2].Type == 0);
            assert(actual2.Reward.Item[2].Count == 0);

            assert(depot.VerifyData2.TryGetValue(2, out actual2));
            assert(actual2 != null);
            assert(actual2.Key == 2);
            assert(actual2.Hide == false);
            assert(actual2.Enable == false);
            assert(actual2.Name == "名稱2");
            assert(actual2.Reward.Desc == "獎勵說明2");
            assert(actual2.Reward.Gold == 200);
            assert(actual2.Reward.Diamond == 20);
            assert(actual2.Reward.Crystal == 0);
            assert(actual2.Reward.FelIron == 0);
            assert(actual2.Reward.Atium == 0);
            assert(actual2.Reward.Item.Count == 3);
            assert(actual2.Reward.Item[0].ItemID == 10001);
            assert(actual2.Reward.Item[0].Type == 1);
            assert(actual2.Reward.Item[0].Count == 1);
            assert(actual2.Reward.Item[1].ItemID == 10002);
            assert(actual2.Reward.Item[1].Type == 2);
            assert(actual2.Reward.Item[1].Count == 2);
            assert(actual2.Reward.Item[2].ItemID == 10003);
            assert(actual2.Reward.Item[2].Type == 3);
            assert(actual2.Reward.Item[2].Count == 3);

            assert(depot.VerifyData2.TryGetValue(3, out actual2) == false);
            assert(actual2 == null);

            Debug.Log("verify proto from: success");
        }

        public static void verifyProtoMerge() {
            var loader = new ProtoFileLoader();
            var depot = new Depot() { Loader = loader };

            assert(depot.MergeData());

            VerifyData1 actual1;

            assert(depot.VerifyData1.TryGetValue(1, out actual1));
            assert(actual1 != null);
            assert(actual1.Key == 1);
            assert(actual1.Hide == false);
            assert(actual1.Enable == true);
            assert(actual1.Name == "名稱1");
            assert(actual1.Reward.Desc == "獎勵說明1");
            assert(actual1.Reward.Gold == 100);
            assert(actual1.Reward.Diamond == 10);
            assert(actual1.Reward.Crystal == 199);
            assert(actual1.Reward.FelIron == 5);
            assert(actual1.Reward.Atium == 1);
            assert(actual1.Reward.Item.Count == 3);
            assert(actual1.Reward.Item[0].ItemID == 10001);
            assert(actual1.Reward.Item[0].Type == 1);
            assert(actual1.Reward.Item[0].Count == 1);
            assert(actual1.Reward.Item[1].ItemID == 10002);
            assert(actual1.Reward.Item[1].Type == 2);
            assert(actual1.Reward.Item[1].Count == 2);
            assert(actual1.Reward.Item[2].ItemID == 0);
            assert(actual1.Reward.Item[2].Type == 0);
            assert(actual1.Reward.Item[2].Count == 0);

            assert(depot.VerifyData1.TryGetValue(2, out actual1));
            assert(actual1 != null);
            assert(actual1.Key == 2);
            assert(actual1.Hide == false);
            assert(actual1.Enable == false);
            assert(actual1.Name == "名稱2");
            assert(actual1.Reward.Desc == "獎勵說明2");
            assert(actual1.Reward.Gold == 200);
            assert(actual1.Reward.Diamond == 20);
            assert(actual1.Reward.Crystal == 299);
            assert(actual1.Reward.FelIron == 10);
            assert(actual1.Reward.Atium == 2);
            assert(actual1.Reward.Item.Count == 3);
            assert(actual1.Reward.Item[0].ItemID == 10001);
            assert(actual1.Reward.Item[0].Type == 1);
            assert(actual1.Reward.Item[0].Count == 1);
            assert(actual1.Reward.Item[1].ItemID == 10002);
            assert(actual1.Reward.Item[1].Type == 2);
            assert(actual1.Reward.Item[1].Count == 2);
            assert(actual1.Reward.Item[2].ItemID == 10003);
            assert(actual1.Reward.Item[2].Type == 3);
            assert(actual1.Reward.Item[2].Count == 3);

            assert(depot.VerifyData1.TryGetValue(3, out actual1) == false);
            assert(actual1 == null);

            VerifyData2 actual2;

            assert(depot.VerifyData2.TryGetValue(1, out actual2));
            assert(actual2 != null);
            assert(actual2.Key == 1);
            assert(actual2.Hide == false);
            assert(actual2.Enable == true);
            assert(actual2.Name == "名稱1");
            assert(actual2.Reward.Desc == "獎勵說明1");
            assert(actual2.Reward.Gold == 100);
            assert(actual2.Reward.Diamond == 10);
            assert(actual2.Reward.Crystal == 0);
            assert(actual2.Reward.FelIron == 0);
            assert(actual2.Reward.Atium == 0);
            assert(actual2.Reward.Item.Count == 3);
            assert(actual2.Reward.Item[0].ItemID == 10001);
            assert(actual2.Reward.Item[0].Type == 1);
            assert(actual2.Reward.Item[0].Count == 1);
            assert(actual2.Reward.Item[1].ItemID == 10002);
            assert(actual2.Reward.Item[1].Type == 2);
            assert(actual2.Reward.Item[1].Count == 2);
            assert(actual2.Reward.Item[2].ItemID == 0);
            assert(actual2.Reward.Item[2].Type == 0);
            assert(actual2.Reward.Item[2].Count == 0);

            assert(depot.VerifyData2.TryGetValue(2, out actual2));
            assert(actual2 != null);
            assert(actual2.Key == 2);
            assert(actual2.Hide == false);
            assert(actual2.Enable == false);
            assert(actual2.Name == "名稱2");
            assert(actual2.Reward.Desc == "獎勵說明2");
            assert(actual2.Reward.Gold == 200);
            assert(actual2.Reward.Diamond == 20);
            assert(actual2.Reward.Crystal == 0);
            assert(actual2.Reward.FelIron == 0);
            assert(actual2.Reward.Atium == 0);
            assert(actual2.Reward.Item.Count == 3);
            assert(actual2.Reward.Item[0].ItemID == 10001);
            assert(actual2.Reward.Item[0].Type == 1);
            assert(actual2.Reward.Item[0].Count == 1);
            assert(actual2.Reward.Item[1].ItemID == 10002);
            assert(actual2.Reward.Item[1].Type == 2);
            assert(actual2.Reward.Item[1].Count == 2);
            assert(actual2.Reward.Item[2].ItemID == 10003);
            assert(actual2.Reward.Item[2].Type == 3);
            assert(actual2.Reward.Item[2].Count == 3);

            assert(depot.VerifyData2.TryGetValue(3, out actual2) == false);
            assert(actual2 == null);

            Debug.Log("verify proto merge: success");
        }

        private static void assert(bool condition) {
            if (condition == false)
                throw new Exception("verify proto: verify failed");
        }
    }

    class ProtoFileLoader : Loader {
        public void Error(string name, string message) {
            Debug.Log(name + ": proto file load failed: " + message);
        }

        public byte[] Load(string name, string ext, string fullname) {
            return File.ReadAllBytes(Path.Combine("Assets", "target", "proto", "data", fullname));
        }
    }
}