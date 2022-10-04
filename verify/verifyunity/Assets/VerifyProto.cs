using SheeterProto;
using System;
using System.IO;
using UnityEngine;

public static class VerifyProto
{
    public static void Test()
    {
        var path = Path.Combine("Assets", "target", "proto", "data");

        Test1(path);
        Test2(path);
    }

    private static void Test1(string path)
    {
        var reader = new VerifyData1Reader();

        if (reader.FromPath(path) == false) // 工作目錄在target
        {
            throw new Exception("verify proto: read failed");
        } // if

        VerifyData1 actual;

        Assert(reader.Data.TryGetValue(1, out actual));
        Assert(actual != null);
        Assert(actual.Key == 1);
        Assert(actual.Hide == false);
        Assert(actual.Enable == true);
        Assert(actual.Name == "名稱1");
        Assert(actual.Reward.Desc == "獎勵說明1");
        Assert(actual.Reward.Gold == 100);
        Assert(actual.Reward.Diamond == 10);
        Assert(actual.Reward.Crystal == 199);
        Assert(actual.Reward.FelIron == 5);
        Assert(actual.Reward.Atium == 1);
        Assert(actual.Reward.Item.Count == 3);
        Assert(actual.Reward.Item[0].ItemID == 10001);
        Assert(actual.Reward.Item[0].Type == 1);
        Assert(actual.Reward.Item[0].Count == 1);
        Assert(actual.Reward.Item[1].ItemID == 10002);
        Assert(actual.Reward.Item[1].Type == 2);
        Assert(actual.Reward.Item[1].Count == 2);
        Assert(actual.Reward.Item[2].ItemID == 0);
        Assert(actual.Reward.Item[2].Type == 0);
        Assert(actual.Reward.Item[2].Count == 0);

        Assert(reader.Data.TryGetValue(2, out actual));
        Assert(actual != null);
        Assert(actual.Key == 2);
        Assert(actual.Hide == false);
        Assert(actual.Enable == false);
        Assert(actual.Name == "名稱2");
        Assert(actual.Reward.Desc == "獎勵說明2");
        Assert(actual.Reward.Gold == 200);
        Assert(actual.Reward.Diamond == 20);
        Assert(actual.Reward.Crystal == 299);
        Assert(actual.Reward.FelIron == 10);
        Assert(actual.Reward.Atium == 2);
        Assert(actual.Reward.Item.Count == 3);
        Assert(actual.Reward.Item[0].ItemID == 10001);
        Assert(actual.Reward.Item[0].Type == 1);
        Assert(actual.Reward.Item[0].Count == 1);
        Assert(actual.Reward.Item[1].ItemID == 10002);
        Assert(actual.Reward.Item[1].Type == 2);
        Assert(actual.Reward.Item[1].Count == 2);
        Assert(actual.Reward.Item[2].ItemID == 10003);
        Assert(actual.Reward.Item[2].Type == 3);
        Assert(actual.Reward.Item[2].Count == 3);

        Assert(reader.Data.TryGetValue(3, out actual) == false);
        Assert(actual == null);

        Debug.Log("verify proto: success");
    }

    private static void Test2(string path)
    {
        var reader = new VerifyData2Reader();

        if (reader.FromPath(path) == false) // 工作目錄在target
        {
            throw new Exception("verify proto: read failed");
        } // if

        VerifyData2 actual;

        Assert(reader.Data.TryGetValue(1, out actual));
        Assert(actual != null);
        Assert(actual.Key == 1);
        Assert(actual.Hide == false);
        Assert(actual.Enable == true);
        Assert(actual.Name == "名稱1");
        Assert(actual.Reward.Desc == "獎勵說明1");
        Assert(actual.Reward.Gold == 100);
        Assert(actual.Reward.Diamond == 10);
        Assert(actual.Reward.Crystal == 0);
        Assert(actual.Reward.FelIron == 0);
        Assert(actual.Reward.Atium == 0);
        Assert(actual.Reward.Item.Count == 3);
        Assert(actual.Reward.Item[0].ItemID == 10001);
        Assert(actual.Reward.Item[0].Type == 1);
        Assert(actual.Reward.Item[0].Count == 1);
        Assert(actual.Reward.Item[1].ItemID == 10002);
        Assert(actual.Reward.Item[1].Type == 2);
        Assert(actual.Reward.Item[1].Count == 2);
        Assert(actual.Reward.Item[2].ItemID == 0);
        Assert(actual.Reward.Item[2].Type == 0);
        Assert(actual.Reward.Item[2].Count == 0);

        Assert(reader.Data.TryGetValue(2, out actual));
        Assert(actual != null);
        Assert(actual.Key == 2);
        Assert(actual.Hide == false);
        Assert(actual.Enable == false);
        Assert(actual.Name == "名稱2");
        Assert(actual.Reward.Desc == "獎勵說明2");
        Assert(actual.Reward.Gold == 200);
        Assert(actual.Reward.Diamond == 20);
        Assert(actual.Reward.Crystal == 0);
        Assert(actual.Reward.FelIron == 0);
        Assert(actual.Reward.Atium == 0);
        Assert(actual.Reward.Item.Count == 3);
        Assert(actual.Reward.Item[0].ItemID == 10001);
        Assert(actual.Reward.Item[0].Type == 1);
        Assert(actual.Reward.Item[0].Count == 1);
        Assert(actual.Reward.Item[1].ItemID == 10002);
        Assert(actual.Reward.Item[1].Type == 2);
        Assert(actual.Reward.Item[1].Count == 2);
        Assert(actual.Reward.Item[2].ItemID == 10003);
        Assert(actual.Reward.Item[2].Type == 3);
        Assert(actual.Reward.Item[2].Count == 3);

        Assert(reader.Data.TryGetValue(3, out actual) == false);
        Assert(actual == null);

        Debug.Log("verify proto: success");
    }

    private static void Assert(bool condition)
    {
        if (condition == false)
            throw new Exception("verify proto: verify failed");
    }
}
