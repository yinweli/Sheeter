﻿using Newtonsoft.Json;
using System;
using verifydata;

namespace verifycs {
    public class Verifycs {
        public static void Main()
        {
            var reader = Reader.FromJsonFile(Reader.Json);

            if (reader == null) {
                throw new Exception("verify cs: read failed");
            } // if

            var expects = new Struct[] {
                new Struct {
                    Reward = new Reward {
                        Item = new Item[] {
                            new Item{ Count = 10, ItemId = 10001, Type = 1, },
                            new Item{ Count = 0, ItemId = 0, Type = 0, },
                            new Item{ Count = 0, ItemId = 0, Type = 0, },
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
                new Struct {
                    Reward = new Reward {
                        Item = new Item[] {
                            new Item{ Count = 10, ItemId = 10001, Type = 1, },
                            new Item{ Count = 5, ItemId = 10002, Type = 1, },
                            new Item{ Count = 0, ItemId = 0, Type = 0, },
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
                new Struct {
                    Reward = new Reward {
                        Item = new Item[] {
                            new Item{ Count = 10, ItemId = 10001, Type = 1, },
                            new Item{ Count = 5, ItemId = 10002, Type = 1, },
                            new Item{ Count = 2, ItemId = 10003, Type = 1, },
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
                new Struct {
                    Reward = new Reward {
                        Item = new Item[] {
                            new Item{ Count = 10, ItemId = 10001, Type = 1, },
                            new Item{ Count = 5, ItemId = 10002, Type = 1, },
                            new Item{ Count = 3, ItemId = 10003, Type = 1, },
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

            foreach (var itor in expects) {
                if (reader.TryGetValue(itor.Key, out var actual) == false || JsonConvert.SerializeObject(itor) != JsonConvert.SerializeObject(actual)) {
                    throw new Exception("verify cs: compare failed");
                } // if
            } // for

            Console.WriteLine("verify cs: success");
        }
    }
}