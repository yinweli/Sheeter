using Sheeter;

namespace verify
{
    public static class Verify
    {
        public static void Main()
        {
            var sheet = new Sheeter.Sheeter(new FileLoader());

            Assert(sheet.FromData().GetAwaiter().GetResult());
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
                        Assert(actual.Name1 == 1);
                        Assert(actual.Name2 == 10);
                        Assert(actual.Name3 == 11);
                        Assert(actual.Name4 == 12);
                        Assert(actual.Name5 == 13);
                        Assert(actual.Float == 0.01f);
                        Assert(actual.Ratio.ParseAs<Ratio>().Float32() == 0.01f);
                        Assert(actual.Duration.ParseAs<Duration>().Interval() == TimeSpan.Parse("1.02:03:04.005"));

                        Assert(sheet.VerifyData.TryGetValue(2, out actual));
                        Assert(actual.Name1 == 2);
                        Assert(actual.Name2 == 20);
                        Assert(actual.Name3 == 21);
                        Assert(actual.Name4 == 22);
                        Assert(actual.Name5 == 23);
                        Assert(actual.Float == 0.0001f);
                        Assert(actual.Ratio.ParseAs<Ratio>().Float32() == 0.0001f);
                        Assert(actual.Duration.ParseAs<Duration>().Interval() == TimeSpan.Parse("1.02:03:04.005"));

                        Assert(sheet.VerifyData.TryGetValue(3, out actual) == false);

                        Assert(sheet.VerifyData.TryGetValue(4, out actual));
                        Assert(actual.Name1 == 4);
                        Assert(actual.Name2 == 40);
                        Assert(actual.Name3 == 41);
                        Assert(actual.Name4 == 42);
                        Assert(actual.Name5 == 43);

                        Assert(sheet.VerifyData.TryGetValue(5, out actual));
                        Assert(actual.Name1 == 5);
                        Assert(actual.Name2 == 50);
                        Assert(actual.Name3 == 51);
                        Assert(actual.Name4 == 52);
                        Assert(actual.Name5 == 53);

                        Assert(sheet.MergeData.TryGetValue(1, out actual));
                        Assert(actual.Name1 == 1);
                        Assert(actual.Name2 == 10);
                        Assert(actual.Name3 == 11);
                        Assert(actual.Name4 == 12);
                        Assert(actual.Name5 == 13);

                        Assert(sheet.MergeData.TryGetValue(2, out actual));
                        Assert(actual.Name1 == 2);
                        Assert(actual.Name2 == 20);
                        Assert(actual.Name3 == 21);
                        Assert(actual.Name4 == 22);
                        Assert(actual.Name5 == 23);

                        Assert(sheet.MergeData.TryGetValue(3, out actual) == false);

                        Assert(sheet.MergeData.TryGetValue(4, out actual));
                        Assert(actual.Name1 == 4);
                        Assert(actual.Name2 == 40);
                        Assert(actual.Name3 == 41);
                        Assert(actual.Name4 == 42);
                        Assert(actual.Name5 == 43);

                        Assert(sheet.MergeData.TryGetValue(5, out actual));
                        Assert(actual.Name1 == 5);
                        Assert(actual.Name2 == 50);
                        Assert(actual.Name3 == 51);
                        Assert(actual.Name4 == 52);
                        Assert(actual.Name5 == 53);
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
