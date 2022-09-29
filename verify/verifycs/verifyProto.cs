using SheeterProto;
using System;
using System.IO;

namespace verifycs
{
    public class VerifyProto
    {
        public static void Test()
        {
            var reader = new VerifyData1Reader();

            if (reader.FromPathHalf(Path.Combine("proto", "data")) == false) // 工作目錄在target
            {
                throw new Exception("verify proto: read failed");
            } // if

            VerifyData1 actual;

            Assert(reader.Data.TryGetValue(1, out actual));
            Assert(actual != null);
            Assert(actual.Key == 1);
            Assert(actual.Name == "名稱1");
            Assert(actual.Enable == true);
            Assert(actual.Reward.Atium == 2);
            Assert(actual.Reward.Crystal == 120);
            Assert(actual.Reward.Diamond == 10);
            Assert(actual.Reward.FelIron == 6);
            Assert(actual.Reward.Gold == 500);
            Assert(actual.Reward.Item.Count == 3);
            Assert(actual.Reward.Item[0].ItemID == 10001);
            Assert(actual.Reward.Item[0].Type == 1);
            Assert(actual.Reward.Item[0].Count == 10);
            Assert(actual.Reward.Item[1].ItemID == 0);
            Assert(actual.Reward.Item[1].Type == 0);
            Assert(actual.Reward.Item[1].Count == 0);
            Assert(actual.Reward.Item[2].ItemID == 0);
            Assert(actual.Reward.Item[2].Type == 0);
            Assert(actual.Reward.Item[2].Count == 0);

            Assert(reader.Data.TryGetValue(2, out actual));
            Assert(actual != null);
            Assert(actual.Key == 2);
            Assert(actual.Name == "名稱2");
            Assert(actual.Enable == true);
            Assert(actual.Reward.Atium == 2);
            Assert(actual.Reward.Crystal == 135);
            Assert(actual.Reward.Diamond == 12);
            Assert(actual.Reward.FelIron == 8);
            Assert(actual.Reward.Gold == 550);
            Assert(actual.Reward.Item.Count == 3);
            Assert(actual.Reward.Item[0].ItemID == 10001);
            Assert(actual.Reward.Item[0].Type == 1);
            Assert(actual.Reward.Item[0].Count == 10);
            Assert(actual.Reward.Item[1].ItemID == 10002);
            Assert(actual.Reward.Item[1].Type == 1);
            Assert(actual.Reward.Item[1].Count == 5);
            Assert(actual.Reward.Item[2].ItemID == 0);
            Assert(actual.Reward.Item[2].Type == 0);
            Assert(actual.Reward.Item[2].Count == 0);

            Assert(reader.Data.TryGetValue(3, out actual));
            Assert(actual != null);
            Assert(actual.Key == 3);
            Assert(actual.Name == "名稱3");
            Assert(actual.Enable == false);
            Assert(actual.Reward.Atium == 3);
            Assert(actual.Reward.Crystal == 150);
            Assert(actual.Reward.Diamond == 14);
            Assert(actual.Reward.FelIron == 10);
            Assert(actual.Reward.Gold == 600);
            Assert(actual.Reward.Item.Count == 3);
            Assert(actual.Reward.Item[0].ItemID == 10001);
            Assert(actual.Reward.Item[0].Type == 1);
            Assert(actual.Reward.Item[0].Count == 10);
            Assert(actual.Reward.Item[1].ItemID == 10002);
            Assert(actual.Reward.Item[1].Type == 1);
            Assert(actual.Reward.Item[1].Count == 5);
            Assert(actual.Reward.Item[2].ItemID == 10003);
            Assert(actual.Reward.Item[2].Type == 1);
            Assert(actual.Reward.Item[2].Count == 2);

            Assert(reader.Data.TryGetValue(4, out actual));
            Assert(actual != null);
            Assert(actual.Key == 4);
            Assert(actual.Name == "名稱4");
            Assert(actual.Enable == false);
            Assert(actual.Reward.Atium == 3);
            Assert(actual.Reward.Crystal == 165);
            Assert(actual.Reward.Diamond == 16);
            Assert(actual.Reward.FelIron == 12);
            Assert(actual.Reward.Gold == 650);
            Assert(actual.Reward.Item.Count == 3);
            Assert(actual.Reward.Item[0].ItemID == 10001);
            Assert(actual.Reward.Item[0].Type == 1);
            Assert(actual.Reward.Item[0].Count == 10);
            Assert(actual.Reward.Item[1].ItemID == 10002);
            Assert(actual.Reward.Item[1].Type == 1);
            Assert(actual.Reward.Item[1].Count == 5);
            Assert(actual.Reward.Item[2].ItemID == 10003);
            Assert(actual.Reward.Item[2].Type == 1);
            Assert(actual.Reward.Item[2].Count == 3);

            Assert(reader.Data.TryGetValue(5, out actual) == false);
            Assert(actual == null);

            Console.WriteLine("verify proto: success");
        }

        private static void Assert(bool condition)
        {
            if (condition == false)
                throw new Exception("verify proto: verify failed");
        }
    }
}
