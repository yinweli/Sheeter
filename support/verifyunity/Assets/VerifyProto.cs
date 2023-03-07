using SheeterProto;
using System;
using System.Collections.Generic;
using System.IO;
using System.Threading;
using UnityEngine;

namespace verifycs
{
    public static class VerifyProto
    {
        public static void verifyProtoFrom(int threads)
        {
            var loader = new ProtoFileLoader();
            var depot = new Depot() { Loader = loader };

            assert(depot.FromData());
            verify(depot, threads);

            Debug.Log("verify proto from: success");
        }

        public static void verifyProtoMerge(int threads)
        {
            var loader = new ProtoFileLoader();
            var depot = new Depot() { Loader = loader };

            assert(depot.MergeData());
            verify(depot, threads);

            Debug.Log("verify proto merge: success");
        }

        private static void verify(Depot depot, int threads)
        {
            var threadList = new List<Thread>();

            for (var t = 0; t < threads; t++)
            {
                threadList.Add(
                    new Thread(() =>
                    {
                        VerifyData1 actual1;
                        VerifyData2 actual2;

                        for (var i = 1; i <= 100; i++)
                        {
                            assert(depot.VerifyData1.TryGetValue(i, out actual1));
                            assert(actual1 != null);
                            assert(actual1.Key == i);
                            assert(actual1.Hide == false);
                            assert(actual1.Enable == (i % 2 == 1));
                            assert(actual1.Name == "名稱" + i);
                            assert(actual1.Reward.Desc == "獎勵說明" + i);
                            assert(actual1.Reward.Gold == i * 2);
                            assert(actual1.Reward.Diamond == i * 3);
                            assert(actual1.Reward.Crystal == i * 4);
                            assert(actual1.Reward.FelIron == i * 5);
                            assert(actual1.Reward.Atium == i * 6);
                            assert(actual1.Reward.Item.Count == 3);
                            assert(actual1.Reward.Item[0].ItemID == 1000 + i);
                            assert(actual1.Reward.Item[0].Type == 0);
                            assert(actual1.Reward.Item[0].Count == i);
                            assert(actual1.Reward.Item[1].ItemID == 10000 + i);
                            assert(actual1.Reward.Item[1].Type == 1);
                            assert(actual1.Reward.Item[1].Count == i);
                            assert(actual1.Reward.Item[2].ItemID == 100000 + i);
                            assert(actual1.Reward.Item[2].Type == 2);
                            assert(actual1.Reward.Item[2].Count == i);

                            assert(depot.VerifyData2.TryGetValue(i.ToString(), out actual2));
                            assert(actual2 != null);
                            assert(actual2.Key == i.ToString());
                            assert(actual2.Hide == false);
                            assert(actual2.Enable == (i % 2 == 1));
                            assert(actual2.Name == "名稱" + i);
                            assert(actual2.Reward.Desc == "獎勵說明" + i);
                            assert(actual2.Reward.Gold == i * 2);
                            assert(actual2.Reward.Diamond == i * 3);
                            assert(actual2.Reward.Crystal == 0);
                            assert(actual2.Reward.FelIron == 0);
                            assert(actual2.Reward.Atium == 0);
                            assert(actual2.Reward.Item.Count == 3);
                            assert(actual2.Reward.Item[0].ItemID == 1000 + i);
                            assert(actual2.Reward.Item[0].Type == 0);
                            assert(actual2.Reward.Item[0].Count == i);
                            assert(actual2.Reward.Item[1].ItemID == 10000 + i);
                            assert(actual2.Reward.Item[1].Type == 1);
                            assert(actual2.Reward.Item[1].Count == i);
                            assert(actual2.Reward.Item[2].ItemID == 100000 + i);
                            assert(actual2.Reward.Item[2].Type == 2);
                            assert(actual2.Reward.Item[2].Count == i);
                        } // for

                        assert(depot.VerifyData1.TryGetValue(101, out actual1) == false);
                        assert(actual1 == null);

                        assert(depot.VerifyData2.TryGetValue(101.ToString(), out actual2) == false);
                        assert(actual2 == null);
                    })
                );
            } // for

            foreach (var itor in threadList)
                itor.Start();

            foreach (var itor in threadList)
                itor.Join();
        }

        private static void assert(bool condition)
        {
            if (condition == false)
                throw new Exception("verify proto: verify failed");
        }
    }

    class ProtoFileLoader : Loader
    {
        public void Error(string name, string message)
        {
            Debug.Log(name + ": proto file load failed: " + message);
        }

        public byte[] Load(FileName filename)
        {
            return File.ReadAllBytes(Path.Combine("Assets", "target", "proto", "data", filename.File));
        }
    }
}
