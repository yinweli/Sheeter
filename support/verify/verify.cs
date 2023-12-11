using Sheeter;

namespace verify
{
    public static class Verify
    {
        public static void Main()
        {
            var sheet = new Sheeter.Sheeter(new FileLoader());

            Assert(sheet.FromData());
            Check(sheet, 1000);
            Console.WriteLine("verify success");
        }

        private static void Check(Sheeter.Sheeter sheet, int threads)
        {
            var threadList = new List<Thread>();

            for (var t = 0; t < threads; t++)
            {
                threadList.Add(
                    new Thread(() =>
                    {
                        Assert(sheet.VerifyData.TryGetValue(1, out var actual));
                        Assert(actual.Pkey == 1);
                        Assert(actual.Name1 == 10);
                        Assert(actual.Name2 == 11);
                        Assert(actual.Name3 == 12);
                        Assert(actual.Name4 == 13);

                        Assert(sheet.VerifyData.TryGetValue(2, out actual));
                        Assert(actual.Pkey == 2);
                        Assert(actual.Name1 == 20);
                        Assert(actual.Name2 == 21);
                        Assert(actual.Name3 == 22);
                        Assert(actual.Name4 == 23);

                        Assert(sheet.VerifyData.TryGetValue(3, out actual) == false);

                        Assert(sheet.VerifyData.TryGetValue(4, out actual));
                        Assert(actual.Pkey == 4);
                        Assert(actual.Name1 == 40);
                        Assert(actual.Name2 == 41);
                        Assert(actual.Name3 == 42);
                        Assert(actual.Name4 == 43);

                        Assert(sheet.VerifyData.TryGetValue(5, out actual));
                        Assert(actual.Pkey == 5);
                        Assert(actual.Name1 == 50);
                        Assert(actual.Name2 == 51);
                        Assert(actual.Name3 == 52);
                        Assert(actual.Name4 == 53);

                        Assert(sheet.MergeData.TryGetValue(1, out actual));
                        Assert(actual.Pkey == 1);
                        Assert(actual.Name1 == 10);
                        Assert(actual.Name2 == 11);
                        Assert(actual.Name3 == 12);
                        Assert(actual.Name4 == 13);

                        Assert(sheet.MergeData.TryGetValue(2, out actual));
                        Assert(actual.Pkey == 2);
                        Assert(actual.Name1 == 20);
                        Assert(actual.Name2 == 21);
                        Assert(actual.Name3 == 22);
                        Assert(actual.Name4 == 23);

                        Assert(sheet.MergeData.TryGetValue(3, out actual) == false);

                        Assert(sheet.MergeData.TryGetValue(4, out actual));
                        Assert(actual.Pkey == 4);
                        Assert(actual.Name1 == 40);
                        Assert(actual.Name2 == 41);
                        Assert(actual.Name3 == 42);
                        Assert(actual.Name4 == 43);

                        Assert(sheet.MergeData.TryGetValue(5, out actual));
                        Assert(actual.Pkey == 5);
                        Assert(actual.Name1 == 50);
                        Assert(actual.Name2 == 51);
                        Assert(actual.Name3 == 52);
                        Assert(actual.Name4 == 53);
                    })
                );
            } // for

            foreach (var itor in threadList)
                itor.Start();

            foreach (var itor in threadList)
                itor.Join();
        }

        private static void Assert(bool condition)
        {
            if (condition == false)
                throw new Exception("verify failed");
        }
    }

    public class FileLoader : Loader
    {
        public string Load(FileName filename)
        {
            return File.ReadAllText(Path.Combine("json", filename.File));
        }

        public void Error(string name, string message)
        {
            Console.WriteLine("file load failed: " + name + ": " + message);
        }
    }
}
