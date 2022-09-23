using Newtonsoft.Json;
using sheeter;
using System;
using System.IO;

namespace verifycs
{
    public class VerifyJsonCs
    {
        public static void Test()
        {
            var path = Path.Combine("json", ".data", VerifyData1Reader.Json); // 工作目錄在target
            var reader = VerifyData1Reader.FromJsonFile(path);

            if (reader == null)
            {
                throw new Exception("verify json cs: read failed");
            } // if

            var expects = new VerifyData1[] {
                new VerifyData1 {
                    Reward = new Reward {
                        Item = new Item[] {
                            new Item{ Count = 10, ItemID = 10001, Type = 1, },
                            new Item{ Count = 0, ItemID = 0, Type = 0, },
                            new Item{ Count = 0, ItemID = 0, Type = 0, },
                        },
                        Atium = 2,
                        Crystal = 120,
                        Diamond = 10,
                        FelIron = 6,
                        Gold = 500,
                    },
                    Enable = true,
                    Key = 1,
                    Name = "名稱1",
                },
                new VerifyData1 {
                    Reward = new Reward {
                        Item = new Item[] {
                            new Item{ Count = 10, ItemID = 10001, Type = 1, },
                            new Item{ Count = 5, ItemID = 10002, Type = 1, },
                            new Item{ Count = 0, ItemID = 0, Type = 0, },
                        },
                        Atium = 2,
                        Crystal = 135,
                        Diamond = 12,
                        FelIron = 8,
                        Gold = 550,
                    },
                    Enable = true,
                    Key = 2,
                    Name = "名稱2",
                },
                new VerifyData1 {
                    Reward = new Reward {
                        Item = new Item[] {
                            new Item{ Count = 10, ItemID = 10001, Type = 1, },
                            new Item{ Count = 5, ItemID = 10002, Type = 1, },
                            new Item{ Count = 2, ItemID = 10003, Type = 1, },
                        },
                        Atium = 3,
                        Crystal = 150,
                        Diamond = 14,
                        FelIron = 10,
                        Gold = 600,
                    },
                    Enable = false,
                    Key = 3,
                    Name = "名稱3",
                },
                new VerifyData1 {
                    Reward = new Reward {
                        Item = new Item[] {
                            new Item{ Count = 10, ItemID = 10001, Type = 1, },
                            new Item{ Count = 5, ItemID = 10002, Type = 1, },
                            new Item{ Count = 3, ItemID = 10003, Type = 1, },
                        },
                        Atium = 3,
                        Crystal = 165,
                        Diamond = 16,
                        FelIron = 12,
                        Gold = 650,
                    },
                    Enable = false,
                    Key = 4,
                    Name = "名稱4",
                },
            };

            foreach (var itor in expects)
            {
                if (reader.TryGetValue(itor.Key, out var actual) == false || JsonConvert.SerializeObject(itor) != JsonConvert.SerializeObject(actual))
                {
                    throw new Exception("verify json cs: compare failed");
                } // if
            } // for

            Console.WriteLine("verify json cs: success");
        }
    }
}
